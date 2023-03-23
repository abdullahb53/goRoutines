package main

type pidStorage struct {
	pidMap map[int]int
	msg    chan any
	getMsg chan any
}

func newPidStorage() *pidStorage {
	g := &pidStorage{
		pidMap: make(map[int]int),
		msg:    make(chan any, 10),
		getMsg: make(chan any, 10),
	}
	go g.loop()
	go g.getloop()

	return g
}

func (ps *pidStorage) FillMap() {
	for i := 0; i < 900000; i++ {
		ps.pidMap[i] = i
	}
}

func (ps *pidStorage) handleNewMessage(message int) {
	ps.receive(message)
}

func (ps *pidStorage) receive(msg any) {
	ps.msg <- msg
}

func (ps *pidStorage) loop() {
	for val := range ps.msg {
		ps.handleMessage(val)
	}
}

func (ps *pidStorage) handleMessage(message any) {
	switch v := message.(type) {
	case int:
		ps.SetValue(v)
	}
}

func (ps *pidStorage) SetValue(val int) {
	ps.pidMap[val] = 9999
}

func main() {

}

func (ps *pidStorage) handleGetMessage(msg int) {
	ps.receiveGet(msg)
}

func (ps *pidStorage) receiveGet(msg any) {
	ps.getMsg <- msg
}

func (ps *pidStorage) getloop() {
	for val := range ps.msg {
		ps.GetValue(val)
	}
}

func (ps *pidStorage) GetValue(message any) int {
	switch v := message.(type) {
	case int:
		return ps.GetValue(v)
	}
	return 0
}
