package main

import (
	"encoding/json"
	"time"
	"fmt"
	"net/http"

	"github.com/Scr3amz/perxTask/pkg/models"

)

func (app *App) GetTasks(w http.ResponseWriter, r *http.Request) {
	for len(app.Queue) > 0 && len(app.QueueRunning) < app.N {
		app.TransitionTask()
		// fmt.Println("Длинна очереди:", len(app.Queue))
		// fmt.Println("Длинна процессов:", len(app.QueueRunning), app.N)
    }

	w.Write([]byte("Queue of tasks:\n\n"))
	WriteQueue(w, app.Queue)

	w.Write([]byte("Queue of running tasks:\n\n"))
	WriteQueue(w, app.QueueRunning)

	w.Write([]byte("Queue of done tasks:\n\n"))
	WriteQueue(w, app.QueueDone)
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

func (app *App) TransitionTask() {
	task := app.Queue[0]
	task.StartTime = time.Now().Local().Format(time.ANSIC)
	app.QueueRunning = append(app.QueueRunning, task)
	go app.startTask(&task)
	app.Queue = app.Queue[:len(app.Queue)-1]
}

func WriteQueue(w http.ResponseWriter, q []models.Task) {
	for num, t := range q {
		js, err := json.Marshal(t)
		//js, err := json.MarshalIndent(t,"","\t")
		if err!= nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
			return
        }
		taskText := []byte(fmt.Sprintf("Task %d: %s\n", num+1, js))
		w.Write(taskText)
	}
}

func (app *App) startTask(task *models.Task) {
	fmt.Println("Start task ", task)
	res := task.N1
	for n:=0; n < task.N; n++ {
		res += task.D
		task.Iteration = n
		task.N1 = 0
		fmt.Println(task.Iteration)
		time.Sleep(time.Duration(task.I) * time.Second)
	}
}