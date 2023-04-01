package main

type RingBuf struct {
	buffer []int
}

func (r *RingBuf) bufFill() {
	for i := range r.buffer {
		r.buffer[i] = i
	}
}

func newRingBuf(size int) *RingBuf {
	return &RingBuf{
		buffer: make([]int, size),
	}
}

func main() {

}
