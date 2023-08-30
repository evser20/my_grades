package accepted

import (
	"fmt"
	"sync"
)

// потокобезопасные горутины
func TredSafetyRoutines() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	go func() {
		defer wg.Done()

		// Блокируем мьютекс, чтобы гарантировать, что первая горутина завершится первой
		mu.Lock()
		defer mu.Unlock()

		fmt.Println("Горутина 1 выполняется")
	}()

	go func() {
		defer wg.Done()

		// Блокируем мьютекс, чтобы гарантировать, что вторая горутина начнет выполнение после первой
		mu.Lock()
		defer mu.Unlock()

		fmt.Println("Горутина 2 выполняется")
	}()

	wg.Wait()

	fmt.Println("Обе горутины завершились")
	// допилить
}
