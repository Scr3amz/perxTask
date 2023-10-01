package main

import (
	"fmt"
	"log"
	"net/http"

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
	Queue map[int]models.Task
	QueueRunning map[int]models.Task
	QueueDone map[int]models.Task
	N int
	TaskIdx int
}

func main() {
	app := App{
        //Queue: make([]models.Task, 0),
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
		app.Queue[i] = *models.AddTask(5+i, 0,0,3,3)
		app.TaskIdx = i
		//fmt.Println("Init idx: ",app.TaskIdx)
	}

    http.HandleFunc("/tasks", app.GetTasks)
    http.HandleFunc("/tasks/add", app.AddTask)

	fmt.Printf("Server is running on http://127.0.0.1%s/tasks\n", port)

	err = http.ListenAndServe(port, nil)
	if err!= nil {
        log.Fatal("ListenAndServe", err)
    }
	
}
