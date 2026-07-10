package main

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id string
	tls bool
}

type Server struct {
	Opts
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id: "default",
		tls: false,
	}
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	} 
	return &Server{
		Opts: o,
	}
}

func withTls() OptFunc {
	return func(o *Opts) {
		o.tls = true
	}
}

func main() {
	s := newServer(withTls())
	fmt.Printf("%+v", s)
}