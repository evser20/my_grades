package accepted

import (
	"fmt"
	"math/rand"
	"time"
)

// реализация структуры
type Gen struct {
	Min int
	Max int
}

type Iterator interface {
	Next() bool
	Value() int
}

func (g *Gen) Next() bool {
	ok := false
	if g.Max != 0 {
		g.Min = g.Min + 1
		if g.Min == g.Max {
			ok = true
		}
	}
	return ok
}

func (g *Gen) Value() int {
	return g.Max
}

func StructRealization() {
	rand.Seed(time.Now().UnixNano())
	g := &Gen{
		Min: 0,
		Max: rand.Intn(15),
	}
	for i := 1; i < 16; i++ {
		if g.Next() {
			fmt.Printf("%d. Max is equal: %d\n", i, g.Value())
			break
		}
		fmt.Printf("%d. Continue to search max value...\n", i)
		time.Sleep(time.Second * 1)
	}
}
