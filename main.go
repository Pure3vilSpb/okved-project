package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Student struct {
	Rating []int `json:"Rating"`
}

type Group struct {
	Students []Student `json:"Students"`
}

type Result struct {
	Average float64 `json:"Average"`
}

func main() {
	// Читаем данные из stdin
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	// Декодируем JSON
	var group Group
	if err := json.Unmarshal(data, &group); err != nil {
		return
	}

	// Вычисляем среднее количество оценок
	var totalRatings int
	studentCount := len(group.Students)

	for _, student := range group.Students {
		totalRatings += len(student.Rating)
	}

	average := 0.0
	if studentCount > 0 {
		average = float64(totalRatings) / float64(studentCount)
	}

	// Формируем результат
	result := Result{
		Average: average,
	}

	// Кодируем результат в JSON с отступами
	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return
	}

	// Выводим результат
	os.Stdout.Write(output)
}
