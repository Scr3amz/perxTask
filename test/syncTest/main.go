package main

import (
	"fmt"
	"sync"

	"github.com/Scr3amz/perxTask/pkg/models"
	//"time"
	//"github.com/Scr3amz/perxTask/pkg/models"
)

func main(){
	// queue := make(map[int]models.Task, 0)

	// queue[0] = *models.AddTask(5,2,3,4,5)
	// // queue[1] = *models.AddTask(6,2,3,4,5)
	// // queue[2] = *models.AddTask(7,2,3,4,5)

	// fmt.Println(queue[0])
	// task := queue[0]
	// go func(task models.Task) {
	// 	task.N = -1
	// 	queue[0] = task

	// }(task)
	// time.Sleep( time.Second * 3)
	// fmt.Println(queue[0])


	slice := make([]models.Task, 0)
	for i := 0; i < 5; i++ {
		slice = append(slice, *models.AddTask(i,0,0,0,0))
	}

	for _, el := range slice {
		fmt.Print(el, ", ")
	}
	fmt.Print("\n")

	var wg sync.WaitGroup

	go func(task *models.Task){
		//time.Sleep( time.Millisecond * 500)
		task.TTL = -1
		wg.Done()
	}(&slice[2])
	wg.Add(1)

	go func(){
		slice = slice[1:]
		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()

	for _, el := range slice {
		fmt.Print(el, ", ")
	}
	fmt.Print("\n")
	
}