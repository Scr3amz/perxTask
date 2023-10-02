package perxtasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Scr3amz/perxTask/internal/utils"
	"github.com/Scr3amz/perxTask/pkg/models"
)

/*
Queue - Очередь задач
QueueRunning - Задачи в процессе
QueueDone - Выполненные задачи
N - параллельно выполняющиеся задачи
TaskIdx - индекс последней задачи
*/
type App struct {
	Queue map[int]models.Task
	QueueRunning map[int]models.Task
	QueueDone map[int]models.Task
	N int
	TaskIdx int
}

/*
Обработчик, который выводит отсортированный список всех задач
*/
func (app *App) GetTasks(w http.ResponseWriter, r *http.Request) {
	for len(app.Queue) > 0 && len(app.QueueRunning) < app.N {
		fmt.Println("Started transition")
		app.TransitionTask()
    }

	w.Write([]byte("----Queue of tasks:----\n\n"))
	utils.WriteQueue(w, app.Queue)

	w.Write([]byte("\n----Queue of running tasks:----\n\n"))
	utils.WriteQueue(w, app.QueueRunning)

	w.Write([]byte("\n----Queue of done tasks:----\n\n"))
	utils.WriteQueue(w, app.QueueDone)
}

/*
Обработчик, который добавляет задачу в очередь
*/
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

/*
Функция, которая перемещает задачу из очереди в очередь выполняющихся
*/
func (app *App) TransitionTask() {
	key := app.TaskIdx - len(app.Queue) + 1
	task := app.Queue[key]
	fmt.Println(key, task)
	task.StartTime = time.Now().Local().Format(time.ANSIC)
	app.QueueRunning[key] = task
	delete(app.Queue, key)
	go app.runTask(key)
}

/*
Функция, которая запускает арифметическую прогрессию, и после выполнения
перемещает задачу в очередь завершённых
*/
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




