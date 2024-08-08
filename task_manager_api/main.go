package main

import (
	// "fmt"
	"net/http"
	"task_manager_api/models"

	"github.com/gin-gonic/gin"
)

func getTasks(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"tasks": models.Tasks,})
	}
func getTaskById(ctx *gin.Context){
		id := ctx.Param("id")

		for _, task := range models.Tasks{
			if task.ID == id {
				ctx.JSON(http.StatusOK, task)
				return
			}
		}

		ctx.JSON(http.StatusNotFound, gin.H{"error:":"Task not found"})
	}

func addTask(ctx *gin.Context){
	var task models.Task

	if err := ctx.ShouldBindJSON(&task); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Tasks = append(models.Tasks, task)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func updateTask(ctx *gin.Context){
	id := ctx.Param("id")

	var updatedTask models.Task

	if err := ctx.ShouldBindBodyWithJSON(&updatedTask); err!= nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range models.Tasks{
		if task.ID == id {
			if updatedTask.Title != ""{
				models.Tasks[i].Title = updatedTask.Title
			}
            if updatedTask.Description != "" {
                models.Tasks[i].Description = updatedTask.Description
            }
			if updatedTask.DueDate != models.Tasks[i].DueDate  {
                models.Tasks[i].DueDate = updatedTask.DueDate
            }	
			if updatedTask.Status != models.Tasks[i].Status  {
                models.Tasks[i].Status = updatedTask.Status
            }
			ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})				
		}

		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
}

func deleteTask(ctx *gin.Context){
	id := ctx.Param("id")

	for i, task := range models.Tasks{

		if task.ID == id{
			models.Tasks = append(models.Tasks[:i],models.Tasks[i:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}
func main(){
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskById)
	router.POST("/tasks", addTask)
	router.PUT("tasks/:id",updateTask)
	router.DELETE("tasks/:id", deleteTask)

	router.Run()
}

