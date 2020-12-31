package msgq

import "fmt"

var Queue *msgq

func init() {
	Queue = &msgq{
		q: map[string]chan []byte{},
	}
}

type msgq struct{
	q map[string]chan []byte
}

func (m *msgq) Poll(id string) <-chan []byte {
	fmt.Printf("poll: %s",id)
	if q, ok := m.q[id]; ok {
		return q
	} else {
		m.q[id] = make(chan []byte, 10000)
		return m.q[id]
	}
}

func (m *msgq) Enqueue(id string, msg []byte) {
	fmt.Printf("enqueue: %s",id)
	if q, ok := m.q[id]; ok {
		q <- msg
	} else {
		m.q[id] = make(chan []byte)
		m.q[id] <- msg
	}
}

func (m *msgq) Close(id string) {
	close(m.q[id])
}
