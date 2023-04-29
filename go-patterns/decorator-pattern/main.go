package main

import "fmt"

type DB interface {
	Store(string) error
}

type Storage struct{}

func (s *Storage) Store(val string) error {
	fmt.Println("storing into db", val)
	return nil
}

func myExecuteFunc(db DB) ExecuteFn {
	return func(s string) {
		fmt.Println("my Execute func", s)
		db.Store(s)
	}
}

func main() {
	s := &Storage{}
	Execute(myExecuteFunc(s))
}

// Third party.
type ExecuteFn func(string)

func Execute(fn ExecuteFn) {
	fn("hi, everyone")
}

/*

func Execute(fn ExecuteFn){
			(fn ExecuteFn)-> fn("hi everyone") -> myExecuteFunc("hi everyone") -> fmt.Println(s)
			}

*/
