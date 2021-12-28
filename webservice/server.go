package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Done        bool   `json: "Done"`
}

var tasks = []Task{{
	ID:          "1192",
	Title:       "ReoderTest",
	Description: "The name is good",
	Priority:    1,
	Done:        false,
},
	{
		ID:          "1193",
		Title:       "Name",
		Description: "The name is good",
		Priority:    1,
		Done:        false,
	},
}

func postTask(c *gin.Context) {
	var task Task

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&task); err != nil {
		return
	}

	// Add the new album to the slice.
	tasks = append(tasks, task)
	c.IndentedJSON(http.StatusCreated, task)
}

// getTaskByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range tasks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
func deleteTaskByID(c *gin.Context) {
	id := c.Param("id")
	removeIndex := 0
	found := false
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for k, a := range tasks {
		if a.ID == id {
			removeIndex = k
			found = true
			break
		}
	}
	if !found {
		tasks = append(tasks[:removeIndex], tasks[removeIndex+1:]...)
		c.IndentedJSON(http.StatusOK, tasks)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
func getTasks(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, tasks)
}
func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.GET("/task/:id", getTaskByID)
	router.POST("/task", postTask)
	router.DELETE("/task/:id", deleteTaskByID)
	router.Run("localhost:8080")
}
