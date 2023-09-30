package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Scr3amz/perxTask/pkg/models"
)

const (
	port = ":8080"
)

type App struct {
	Queue []models.Task
	QueueRunning []models.Task
	QueueDone []models.Task
}

func main() {
	app := App{
        //Queue: make([]models.Task, 0),
		Queue: make([]models.Task, 0),
        QueueRunning: make([]models.Task, 0),
        QueueDone: make([]models.Task, 0),
    }
	
	app.Queue = append(app.Queue, *models.AddTask(1,2,3,4,5))
	// app.Queue = append(app.Queue, models.Task{})
	app.Queue = append(app.Queue, *models.AddTask(2,3,4,5,1))

    http.HandleFunc("/tasks", app.GetTasks)
    http.HandleFunc("/tasks/add", app.AddTask)

	fmt.Printf("Server is running on http://127.0.0.1%s/tasks", port)

	err := http.ListenAndServe(port, nil)
	if err!= nil {
        log.Fatal("ListenAndServe", err)
    }
	
}

func (app *App) GetTasks(w http.ResponseWriter, r *http.Request) {
	
	for _, t := range app.Queue {
		js, err := json.MarshalIndent(t,"","\t")
		if err!= nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
			return
        }
		w.Write(js)
		w.Write([]byte("\n"))
	}
}

func (app *App) AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	reqBody := r.Body
	defer reqBody.Close()
	var task models.Task
	err := json.NewDecoder(reqBody).Decode(&task)
	if err!= nil {
		fmt.Printf( "%s\n", err.Error())
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }
	task.StatementTime = time.Now().Local().Format(time.ANSIC)
	task.StartTime = ""
	task.EndTime = ""
	app.Queue = append(app.Queue, task)
	
}