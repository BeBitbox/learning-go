package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total      int
	updateTime time.Time
}

type SuperCounter Counter

type FrenchEvenNumbers int

func (c *Counter) AddOne() {
	if c == nil {
		fmt.Printf("You try to add one on a nil value!\n")
	} else {
		c.total++
		c.updateTime = time.Now()
	}
}

func (c Counter) PrintCopy() string {
	return fmt.Sprintf("Counter total:%d and update time:%v", c.total, c.updateTime)
}

func Types() {
	initializedCounter := Counter{
		total:      41,
		updateTime: time.Time{},
	}
	var nilPointer *Counter

	fmt.Printf("Counter: %v\n", initializedCounter.PrintCopy())
	initializedCounter.AddOne()
	fmt.Printf("Counter after AddOne: %v\n", initializedCounter.PrintCopy())
	nilPointer.AddOne()

	super := SuperCounter{
		total:      80086,
		updateTime: time.Time{},
	}
	fmt.Printf("Type of SuperCounter is not inheritance of Counter! '%T', methods of counter not available, just the fields %d\n", super, super.total)

	const (
		Zero FrenchEvenNumbers = iota * 2
		Deux
		Quatre
		Six
	)
	fmt.Printf("French numbers (%d, %d, %d, %d,...)\n", Zero, Deux, Quatre, Six)
}
