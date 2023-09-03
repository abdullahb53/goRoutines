package main

type Pool struct {
	work chan func()
	cls  chan struct{}
}

func NewPool(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		cls:  make(chan struct{}, size),
	}
}

func (p *Pool) Schedule(task func()) {
	select {
	case p.work <- task:
	case p.cls <- struct{}{}:
		go p.Worker(task)
	}
}

func (p *Pool) Worker(task func()) {
	defer func() { <-p.cls }()
	for {
		task()
		task = <-p.work
	}
}

func main() {
	size := 32
	pool := NewPool(size)
	for i := 0; i < size; i++ {
		pool.Schedule(func() {})
	}
}
