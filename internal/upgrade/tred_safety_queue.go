package upgrade

import (
	"fmt"
	"sync"
)

type Queue struct {
	mutex sync.Mutex
	queue []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.queue = append(q.queue, item)
}

func (q *Queue) Dequeue() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.queue) == 0 {
		return nil
	}

	item := q.queue[0]
	q.queue = q.queue[1:]

	return item
}

func (q *Queue) Len() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return len(q.queue)
}

func SafetyQueue() {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	for queue.Len() > 0 {
		item := queue.Dequeue()
		fmt.Println(queue.queue, item.(int))
	}
}
