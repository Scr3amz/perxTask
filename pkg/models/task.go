package models

import (
	"time"
)


/*
n - количество элементов (целочисленное)
d - дельта между элементами последовательности (вещественное)
n1 - Стартовое значение (вещественное)
I - интервал в секундах между итерациями (вещественное)
TTL - время хранения результата в секундах (вещественное)
Iteration int - Текущая итерация
StatementTime - Время постановки задачи
StartTime - Время старта задачи
EndTime - Время окончания задачи (в случае если задача завершена)
*/
type Task struct {
	N int `json:"n"`
	D float64 `json:"d"`
	N1 float64 `json:"n1"`
	I float64 `json:"i"`
	TTL float64 `json:"ttl"`
	Iteration int `json:"iteration,omitempty"`
	StatementTime string `json:"statement_time,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime string `json:"end_time,omitempty"`
}

func AddTask(n int,  d, n1, i, ttl float64) *Task {

	task := &Task{N: n, D: d, N1: n1, I: i, TTL: ttl, StatementTime: time.Now().Local().Format(time.ANSIC)}
	return task
}