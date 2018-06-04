package utils

import (
	"net/http"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/opera443399/cicd-backend/log"
)

//Daemon graphql service
type Daemon struct {
	port        int
	extHandlers map[string]http.Handler
	*SvcHandler
}

//AddExtHandler extend handler to graphql service
func (d *Daemon) AddExtHandler(url string, handler http.Handler) {
	d.extHandlers[url] = handler
}

//Listen graphql service
func (d *Daemon) Listen() {
	http.Handle("/graphql", d)
	for url, handler := range d.extHandlers {
		http.Handle(url, handler)
	}
	
	log.Info("Service [%s] is listening on port *:%d\n", d.Service.String(), d.port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", d.port), nil)
}

//NewGQLDaemon graphql service
func NewGQLDaemon(service ServicesCode, schema graphql.Schema) *Daemon {
	return &Daemon{
		port:        PortDefault + int(service),
		extHandlers: make(map[string]http.Handler),
		SvcHandler: NewGQLHandler(service, schema),
	}
}
