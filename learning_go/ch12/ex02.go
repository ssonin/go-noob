package main

import (
	"fmt"
	"math/rand"
)

type Message struct {
	id  string
	val int
}

func (m Message) String() string {
	return fmt.Sprintf("id:%s, val:%d", m.id, m.val)
}

func ch12ex02() {
	ch1 := make(chan Message)
	ch2 := make(chan Message)
	go func(name string) {
		for i := 0; i < 10; i++ {
			ch1 <- Message{id: name, val: rand.Int()}
		}
		close(ch1)
	}("go routine 1")
	go func(name string) {
		for i := 0; i < 10; i++ {
			ch2 <- Message{id: name, val: rand.Int()}
		}
		close(ch2)
	}("go routine 2")
	for ch1 != nil || ch2 != nil {
		select {
		case msg, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println(msg)
		case msg, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println(msg)
		}
	}
}
