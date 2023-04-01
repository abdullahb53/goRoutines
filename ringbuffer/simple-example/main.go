package main

import "fmt"

type human struct {
	name string
	age  int
}

type RingBuffer struct {
	buffer  []*human
	readId  int
	writeId int
}

func (r *RingBuffer) Add(item *human) {
	r.buffer[r.writeId] = item
	if ((r.writeId + 1) % len(r.buffer)) == r.readId {
		r.readId = (r.readId + 1) % len(r.buffer)
	}
	r.writeId = (r.writeId + 1) % len(r.buffer)
}

func (r *RingBuffer) Get() *human {
	r.readId = (r.readId + 1) % len(r.buffer)
	if r.readId == 0 {
		return r.buffer[len(r.buffer)-1]
	}
	return r.buffer[r.readId-1]
}

func (r *RingBuffer) orderShow() {
	buff := r.buffer
	writeIdx := r.writeId
	tmpWriteIdx := writeIdx
	size := len(r.buffer)

	fmt.Println("\n --- get all  ---")
	defer fmt.Println("\n--- get all  ---")

	for {
		if buff[writeIdx] != nil {
			fmt.Print("[", buff[writeIdx].name, "] ")
		}
		writeIdx = (writeIdx + 1) % size

		if writeIdx == tmpWriteIdx {
			break
		}
	}
}

func (r *RingBuffer) unorderShow() {
	// copy buffer
	buf := r.buffer

	fmt.Println("\n--- unorderShowed ---")
	defer fmt.Println("\n--- unorderShowed ---")
	for i := range buf {
		if buf[i] != nil {
			fmt.Print("[", buf[i].name, "] ")
		}
	}
}

func newHuman(name string, age int) *human {
	return &human{
		name: name,
		age:  age,
	}
}

func newRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer:  make([]*human, size),
		readId:  0,
		writeId: 0,
	}
}

func main() {

	ringBuff := newRingBuffer(4)

	ringBuff.Add(newHuman("1", 27))
	ringBuff.Add(newHuman("2", 27))
	ringBuff.Add(newHuman("3", 27))
	ringBuff.Add(newHuman("4", 27))
	ringBuff.Add(newHuman("5", 27))
	ringBuff.Add(newHuman("6", 27))
	ringBuff.Add(newHuman("-1", 27))
	ringBuff.Add(newHuman("-2", 27))
	ringBuff.Add(newHuman("5", 27))

	// 5634
	ringBuff.unorderShow()

	ringBuff.orderShow()

}
