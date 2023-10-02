package utils

import(
	"fmt"
	"reflect"
	"sort"
	"encoding/json"
	"net/http"

	"github.com/Scr3amz/perxTask/pkg/models"
)

/*
Функция сортирует задачи в очереди по времени добавления и выводит 
*/
func WriteQueue(w http.ResponseWriter, q map[int]models.Task) {
	elems:= reflect.ValueOf(q).MapKeys()
	keySlice := make([]int, len(elems))
	for i:=0; i < len(elems); i++ {
		keySlice[i] = elems[i].Interface().(int)
	}
	sort.Ints(keySlice)
	for num, key := range keySlice {
		task := q[key]
		js, err := json.MarshalIndent(task, "", "\t")
		if err!= nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
			return
        }
		taskText := []byte(fmt.Sprintf("Task %d: %s\n", num+1, js))
		w.Write(taskText)
	}
}
