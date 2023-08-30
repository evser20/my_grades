package accepted

import (
	"fmt"
	"sync"
	"time"
)

// реализация кэша
type Cache struct {
	data sync.Map
}

func (c *Cache) Set(key string, value interface{}) {
	c.data.Store(key, value)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	value, ok := c.data.Load(key)
	return value, ok
}

func CasheRealization() {
	cache := &Cache{}

	// Добавляем данные в кэш
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	// Получаем данные из кэша
	value1, ok1 := cache.Get("key1")
	fmt.Println("Значение для key1:", value1, "Найдено:", ok1)

	value2, ok2 := cache.Get("key2")
	fmt.Println("Значение для key2:", value2, "Найдено:", ok2)

	// Ждем некоторое время
	time.Sleep(2 * time.Second)

	// Получаем данные из кэша после истечения некоторого времени
	value1, ok1 = cache.Get("key1")
	fmt.Println("Значение для key1 после истечения времени жизни:", value1, "Найдено:", ok1)

	// Очищаем кэш
	cache.data.Range(func(key, value interface{}) bool {
		cache.data.Delete(key)
		return true
	})

	// Получаем данные из кэша после очистки
	value2, ok2 = cache.Get("key2")
	fmt.Println("Значение для key2 после очистки:", value2, "Найдено:", ok2)
}
