package main

import "fmt"

type Opts func(*Config)

type Config struct {
	maxConnection int
	id            string
	tls           bool
}

type Server struct {
	Config
}

func defaultSetup() Config {
	return Config{
		maxConnection: 100,
		id:            "DefaultID",
		tls:           false,
	}
}

func withTLS(cfg *Config) {
	cfg.tls = true
}

func withMaxConnection(number int) Opts {
	return func(cfg *Config) {
		cfg.maxConnection = number
	}
}

func newServer(ops ...Opts) *Server {
	o := defaultSetup()
	for _, fn := range ops {
		fn(&o)
	}
	return &Server{
		Config: o,
	}
}

func main() {
	s := newServer(withMaxConnection(60), withTLS)
	fmt.Printf("%+v\n", s)

}
