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
		fmt.Println("Started transition")
		app.transitionTask()
    }

	w.Write([]byte("----Queue of tasks:----\n\n"))
	writeQueue(w, app.Queue)

	w.Write([]byte("\n----Queue of running tasks:----\n\n"))
	writeQueue(w, app.QueueRunning)

	w.Write([]byte("\n----Queue of done tasks:----\n\n"))
	writeQueue(w, app.QueueDone)
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
	idx := app.TaskIdx + 1
	fmt.Println("Idx then add: ", idx)
	app.TaskIdx++
	app.Queue[idx] = task
}

func (app *App) transitionTask() {
	key := app.TaskIdx - len(app.Queue) + 1
	task := app.Queue[key]
	fmt.Println(key, task)
	task.StartTime = time.Now().Local().Format(time.ANSIC)
	app.QueueRunning[key] = task
	delete(app.Queue, key)
	go app.runTask(key)
}

func writeQueue(w http.ResponseWriter, q map[int]models.Task) {
	for num, t := range q {
		js, err := json.Marshal(t)
		//js, err := json.MarshalIndent(t,"","\t")
		if err!= nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
			return
        }
		taskText := []byte(fmt.Sprintf("Task %d: %s\n", num, js))
		w.Write(taskText)
	}
}

func (app *App) runTask(key int) {
	fmt.Println("Start task ")
	task := app.QueueRunning[key]
	res := task.N1
	for n:=0; n < task.N; n++ {
		res += task.D
		task.Iteration = n
		fmt.Println(task.Iteration)
		app.QueueRunning[key] = task
		time.Sleep(time.Duration(task.I) * time.Second)
	}
	task.EndTime = time.Now().Local().Format(time.ANSIC)
	app.QueueDone[key] = task
	fmt.Println("task complited")
	delete(app.QueueRunning, key)
	go func(key int) {
		time.Sleep(time.Duration(task.TTL) * time.Second)
		delete(app.QueueDone, key)
		fmt.Println("task deleted")
	}(key)
}