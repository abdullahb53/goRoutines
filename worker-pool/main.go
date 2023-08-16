package main

const WORKERsize = 256

type Pool struct {
	Work chan func()
	Cls  chan struct{}
}

func NewPool(size int) *Pool {
	return &Pool{
		Work: make(chan func()),
		Cls:  make(chan struct{}, size),
	}
}

func (p *Pool) worker(task func()) {
	defer func() {
		<-p.Cls
	}()
	for {
		task()
		task = <-p.Work
	}
}

func (p *Pool) Schedule(task func()) {
	select {
	case p.Work <- task:
	case p.Cls <- struct{}{}:
		go p.worker(task)
	}
}

func main() {
	_ = NewPool(WORKERsize)

}
