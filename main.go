package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type OkvedItem struct {
	GlobalID int64 `json:"global_id"`
}

func main() {
	// Прямая ссылка на файл в репозитории
	url := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"

	// Выполняем HTTP GET запрос
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка при запросе к GitHub: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка HTTP: %s", resp.Status)
	}

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка чтения тела ответа: %v", err)
	}

	// Декодируем JSON
	var items []OkvedItem
	if err := json.Unmarshal(body, &items); err != nil {
		log.Fatalf("Ошибка декодирования JSON: %v", err)
	}

	// Вычисляем сумму global_id
	var sum int64
	for _, item := range items {
		sum += item.GlobalID
	}

	fmt.Println(sum)
}
