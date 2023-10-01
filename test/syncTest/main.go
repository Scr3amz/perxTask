package main

import (
	"fmt"
	"time"

	"github.com/Scr3amz/perxTask/pkg/models"
)

func main(){
	queue := make(map[int]models.Task, 0)

	queue[0] = *models.AddTask(5,2,3,4,5)
	// queue[1] = *models.AddTask(6,2,3,4,5)
	// queue[2] = *models.AddTask(7,2,3,4,5)

	fmt.Println(queue[0])
	task := queue[0]
	go func(task models.Task) {
		task.N = -1
		queue[0] = task

	}(task)
	time.Sleep( time.Second * 3)
	fmt.Println(queue[0])
}