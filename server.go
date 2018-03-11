package httpware

import (
	"net/http"
	"strconv"
)

//midware add process before and after the privided http handler.
type midware func(http.Handler) http.Handler

//Server is the a server dealing with http request.
type Server struct {
	router http.Handler
	server http.Server
	middle []midware
}

//NewServer create a Server with privided port.
func NewServer(port int) *Server {
	srv := new(Server)
	srv.server.Addr = ":" + strconv.Itoa(port)
	srv.middle = make([]midware, 0)
	return srv
}

//Run start the service.
func (this *Server) Run() {
	if len(this.server.Addr) == 0 {
		this.server.Addr = ":8081"
	}
	go this.server.ListenAndServe()
}

//SetRouter set a router to Server and should not be called after Run.
func (this *Server) SetRouter(handler http.Handler) {
	this.router = handler
	for i := len(this.middle) - 1; i >= 0; i-- {
		handler = this.middle[i](handler)
	}
	this.server.Handler = handler
}

//AddMidware add midware to Server and should not be called after Run.
func (this *Server) AddMidware(wares ...midware) {
	if this.middle == nil {
		this.middle = make([]midware, 0)
	}
	this.middle = append(this.middle, wares...)
	handler := this.router
	for i := len(this.middle) - 1; i >= 0; i-- {
		handler = this.middle[i](handler)
	}
	this.server.Handler = handler
}
