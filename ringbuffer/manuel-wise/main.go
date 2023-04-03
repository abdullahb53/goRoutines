package main

const sizeStack int = 8

type RingBuff struct {
	buffer   []int
	writeIdx int
	readIdx  int
	mask     int
}

func (r *RingBuff) bufFill() {
	print("\n")
	for i := range r.buffer {
		r.buffer[i] = i
	}
}

func (r *RingBuff) show() {
	print("\n")
	for i := range r.buffer {
		print("[", r.buffer[i], "], ")
	}
}

func (r *RingBuff) put(outBuf []int) {
	i := len(outBuf) - 1
	for {
		if (r.writeIdx+1)%(r.mask+1) == r.readIdx {
			continue
		}
		r.buffer[r.writeIdx&r.mask] = outBuf[i]
		r.writeIdx = (r.writeIdx + 1) % (r.mask + 1)
		i--
		// oBuff is completely moved.
		if i == -1 {
			break
		}
	}
}

func (r *RingBuff) read(size int) []int {
	res := make([]int, 0, size)
	for {
		if ((r.readIdx + 1) % (r.mask + 1)) == ((r.writeIdx) % (r.mask + 1)) {
			continue
		}
		r.readIdx = (r.readIdx + 1) % (r.mask + 1)
		res = append(res, r.buffer[r.readIdx])
		size--
		if size == 1 {
			r.readIdx = (r.readIdx + 1) % (r.mask + 1)
			res = append(res, r.buffer[r.readIdx])
			r.readIdx = r.writeIdx - 1
			return res
		}

	}
}

func newRingBuff() *RingBuff {
	stackBuffer := [sizeStack]int{}
	return &RingBuff{
		buffer:   stackBuffer[:],
		writeIdx: 1,
		readIdx:  0,
		mask:     sizeStack - 1,
	}
}

func fibonacciFill(buf *[]int) {
	size := len((*buf))
	(*buf)[0] = 1
	(*buf)[1] = 2
	for i := 2; i < size; i++ {
		(*buf)[i] = (*buf)[i-2] + (*buf)[i-1]
	}

	for i := range *buf {
		print("[", (*buf)[i], "], ")
	}

}

func main() {

	incomingBuff := make([]int, 11)
	fibonacciFill(&incomingBuff)

	ringBuff := newRingBuff()

	go ringBuff.put(incomingBuff)
	res := ringBuff.read(len(incomingBuff))

	println("\n")
	for i := range res {
		println(i, "->", res[i])
	}
	ringBuff.show()
}
