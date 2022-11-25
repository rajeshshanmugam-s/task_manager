package main

import (
	"database/sql"
	"net/http"

	"fmt"

	// "github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "sample_db"
)

type task struct {
	Task_Name string `json:"task_name"`
	Sub_Task  string `json:"sub_task"`
	Manager   string `json:"manager"`
	Date      string `json:"date"`
	Username  string `json:"username"`
}

var tasks = []task{}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successful")

	router := gin.Default()
	router.POST("/add_task", addTask)

	router.Run("localhost:8080")
}

func addTask(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, tasks)
}
