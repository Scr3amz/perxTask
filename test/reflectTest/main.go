package main

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/Scr3amz/perxTask/pkg/models"
)

func main(){
	queue := make(map[int]models.Task, 0)

	queue[0] = *models.AddTask(5,2,3,4,5)
	queue[1] = *models.AddTask(6,2,3,4,5)
	queue[2] = *models.AddTask(7,2,3,4,5)

	// fmt.Println(queue)

	elems := reflect.ValueOf(queue).MapKeys()
	keySlice := make([]int, len(elems))
	for i:=0; i < len(elems); i++ {
		keySlice[i] = elems[i].Interface().(int)
	}
	sort.Ints(keySlice)
	
	for num, key := range keySlice {
		fmt.Printf("Key is: %d\n", key)
		fmt.Printf("Task %d: ", num+1)
		task := queue[key]
		fmt.Println(task)
	}

}