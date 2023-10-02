package main

import (
	"fmt"
	"log"
	"net/http"

	perxtasks "github.com/Scr3amz/perxTask/internal/perxtasks"
	"github.com/Scr3amz/perxTask/pkg/models"
)

const (
	port = ":8080"
)

func main() {
	app := perxtasks.App{
		Queue: make(map[int]models.Task, 0),
        QueueRunning: make(map[int]models.Task, 0),
        QueueDone: make(map[int]models.Task, 0),
		N : 0,
		TaskIdx: 0,
    }
	_,err := fmt.Scan(&app.N)
	if err!= nil {
        log.Fatal(err)
    }
	
	for i := 1; i < 7; i++ {
		app.Queue[i] = *models.AddTask(5+i, 0,0,2,10)
		app.TaskIdx = i
	}

    http.HandleFunc("/tasks", app.GetTasks)
    http.HandleFunc("/tasks/add", app.AddTask)

	fmt.Printf("Server is running on http://127.0.0.1%s/tasks\n", port)

	err = http.ListenAndServe(port, nil)
	if err!= nil {
        log.Fatal("ListenAndServe", err)
    }
	
}
