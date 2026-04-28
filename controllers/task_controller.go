package controllers

import (
	"GO-1/config"
	"GO-1/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetAllTask(ctx *gin.Context) {

	var tasks []models.Task

	// 1. Ambil data dari context
	val, exists := ctx.Get("currentUser")

	// 2. Cek apakah "currentUser" beneran ada?
	if !exists || val == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Sesi tidak ditemukan, silakan login ulang"})
		return
	}

	// 3. Pastikan tipe datanya string sebelum diconvert
	userIDStr, ok := val.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Format UserID salah"})
		return
	}

	// 4. Baru convert ke ObjectID
	objID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// 5. Query dengan filter
	filter := bson.M{"user_id": objID}
	cursor, err := config.TaskCollection.Find(context.TODO(), filter)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &tasks); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, tasks)

}

func AddTask(ctx *gin.Context) {
	var newTask models.Task

	val, _ := ctx.Get("currentUser")
	userIDStr := val.(string)

	objID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan di context"})
		return
	}

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return // Penting: stop eksekusi di sini!
	}

	newTask.UserId = objID

	if newTask.ID.IsZero() {
		newTask.ID = primitive.NewObjectID()
	}

	_, err = config.TaskCollection.InsertOne(context.TODO(), newTask)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Data berhasil ditambah", "data": newTask})
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Format ID salah"})
		return
	}

	_, err = config.TaskCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil menghapus data"})
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	update := bson.M{"$set": bson.M{"isdone": true}}
	_, err = config.TaskCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil update status task"})

}
