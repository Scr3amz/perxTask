package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Scr3amz/perxTask/pkg/models"
)

const (
	port = ":8080"
)

/*
Queue - Очередь задач
QueueRunning - Задачи в процессе
QueueDone - Выполненные задачи
N - параллельно выполняющиеся задачи
*/
type App struct {
	Queue []models.Task
	QueueRunning []models.Task
	QueueDone []models.Task
	N int
	Wg sync.WaitGroup
}

func main() {
	app := App{
        //Queue: make([]models.Task, 0),
		Queue: make([]models.Task, 0),
        QueueRunning: make([]models.Task, 0),
        QueueDone: make([]models.Task, 0),
		N : 0,
		Wg: sync.WaitGroup{},
    }
	_,err := fmt.Scan(&app.N)
	if err!= nil {
        log.Fatal(err)
    }
	
	for i:=0 ; i < 6; i++ {
		app.Queue = append(app.Queue, *models.AddTask(i,1,1,2,1))
	}
	


    http.HandleFunc("/tasks", app.GetTasks)
    http.HandleFunc("/tasks/add", app.AddTask)

	fmt.Printf("Server is running on http://127.0.0.1%s/tasks\n", port)

	err = http.ListenAndServe(port, nil)
	if err!= nil {
        log.Fatal("ListenAndServe", err)
    }
	
}
