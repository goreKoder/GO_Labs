package main //			go run GO_6Laba/zad_6/main.go

import (
	"fmt"
	"sync"
)

// Задача, которую будет выполнять воркер
type Task struct {
	ID int
}

// Функция, выполняемая воркером
func worker(id int, tasks chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Воркер %d выполняет задачу %d\n", id, task.ID)
	}
}

func main() {
	var numWorkers int
	fmt.Println("Введите число воркеров: ")
	fmt.Scan(&numWorkers)
	tasks := make(chan Task)
	var wg sync.WaitGroup

	// Запускаем воркеры
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Добавляем задачи в очередь
	for i := 1; i <= 100; i++ {
		tasks <- Task{ID: i}
	}

	close(tasks) // Закрываем канал, чтобы воркеры знали, что больше задач не будет
	wg.Wait()    // Ждем завершения всех воркеров
	// time.Sleep(8 * time.Second)
}

//			go run GO_6Laba/zad_6/main.go
