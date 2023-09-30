package main

import (
	"encoding/json"
	"fmt"

	"github.com/Scr3amz/perxTask/pkg/models"
)



func main(){
	
	n := 0 
	var d, n1, i, ttl float64
	queue := make([]models.Task, 0)
	queueRunning := make([]models.Task, 0)


	for n != -1 {
		fmt.Scan(&n, &d, &n1, &i, &ttl)
		task := *models.AddTask(n, d, n1, i, ttl)
		queue = append(queue, task)

		fmt.Println(jsconConvert(task))

		if len(queueRunning) < 2 {
			transition(&queue, &queueRunning)
		}
		printData(queue,queueRunning)
	}
}

func transition(queue, queueRunning *[]models.Task) {
	q := *queue
	if len(q) == 0 { return }
	task := q[len(q)-1]
	*queue = append(q[:len(q)-1])
	*queueRunning = append(*queueRunning, task)
}

func printData(queue, queueRunning []models.Task) {
	fmt.Println("Tasks in queue: ")
	for _, t := range queue {
		fmt.Println(t)
	}
	fmt.Println("Tasks is running: ")
	for _, t := range queueRunning {
		fmt.Println(t)
	}
}

func jsconConvert(task models.Task) string {
	js, _ := json.MarshalIndent(task, "", "\t")
    return string(js)
}