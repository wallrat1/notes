package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
	"zametki/database"
	"zametki/models"
)

func resetCache(val string) {
	database.RedisClient.Del(val)
}
func redisToCash(notes []models.Note, authorId uint) {
	notesJson, err := json.Marshal(notes)
	if err != nil {
		log.Printf("Ошибка при сериализации заметок: %v", err)
		return
	}
	err = database.RedisClient.Set(fmt.Sprintf("notes/%d", authorId), string(notesJson), 1488*time.Minute).Err()
	if err != nil {
		log.Println(fmt.Sprintf("Ошибка сохранения в редис %v", err))
		return
	}
}
func takeCashe(val string, ctx *gin.Context) {
	log.Println("кэш найден")
	notes := make([]models.Note, 0)
	json.Unmarshal([]byte(val), &notes)
	ctx.JSON(http.StatusOK, notes)
}
func GetNoteHandler(ctx *gin.Context) {
	authorId, err := ExtractUserID(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "нет доступа"})
	}
	id := ctx.Param("id")
	var note models.Note
	filter := bson.M{"id": id}
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", authorId))
	errFind := collection.FindOne(ctx, filter).Decode(&note)
	if errFind != nil {
		ctx.JSON(http.StatusInternalServerError, "Не найдено")
	} else {
		ctx.JSON(http.StatusOK, &note)
	}
}
func GetNotesHandler(ctx *gin.Context) {
	authorId, err := ExtractUserID(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "нет доступа"})
	}
	var notes []models.Note
	val, err := database.RedisClient.Get(fmt.Sprintf("notes/%d", authorId)).Result()
	if err == redis.Nil {
		log.Println("кэш не найден, загрузка из бд")
		collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", authorId))
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var note models.Note
			err := cursor.Decode(&note)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			notes = append(notes, note)
		}
		if err := cursor.Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		if len(notes) == 0 {
			ctx.JSON(http.StatusNotFound, "заметок не найдено")
		} else {
			ctx.JSON(http.StatusOK, notes)
			redisToCash(notes, authorId)
		}
	} else {
		takeCashe(val, ctx)
	}

}
func CreateNoteHandler(ctx *gin.Context) {
	var note models.Note
	err := ctx.ShouldBindJSON(&note)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}
	note.Id = uuid.New().String()
	note.AuthorID, err = ExtractUserID(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "нет доступа"})
	}
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", note.AuthorID))
	_, err = collection.InsertOne(ctx, note)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	resetCache(fmt.Sprintf("notes/%d", note.AuthorID))
	ctx.JSON(http.StatusOK, gin.H{
		"note":    note,
		"message": "Заметка успешно создана"})
}
func UpdateNoteHandler(ctx *gin.Context) {
	authorId, err := ExtractUserID(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "нет доступа"})
	}
	id := ctx.Param("id")
	var note models.Note
	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при распаковке " + err.Error()})
		return
	}
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", authorId))
	updateFields := bson.M{}
	if note.Name != nil {
		updateFields["name"] = note.Name
	}
	if note.Content != nil {
		updateFields["content"] = note.Content
	}
	update := bson.M{"$set": updateFields}
	filter := bson.M{"id": id}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result.MatchedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
	} else {
		resetCache(fmt.Sprintf("notes/%d", authorId))
		ctx.JSON(http.StatusOK, "Обновлено")
	}
}
func DeleteNoteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	authorId, err := ExtractUserID(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "нет доступа"})
	}
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", authorId))
	filter := bson.M{"id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
	} else {
		resetCache(fmt.Sprintf("notes/%d", authorId))
		ctx.JSON(http.StatusOK, "Удалено")
	}
}
