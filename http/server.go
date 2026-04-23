package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	handlers *HTTPHandlers
}

func NewHTTPServer(httpHand *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		handlers: httpHand,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()
	router.Path("/tasks").Methods("POST").HandlerFunc(s.handlers.HandleAddTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.handlers.HandleGetTask)
	router.Path("/tasks").Methods("GET").Queries("completed", "false").HandlerFunc(s.handlers.HandleGetAllUncompletedTask)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.handlers.HandleGetAllTask)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.handlers.HandleCompleteTask)
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.handlers.HandleDeleteTask)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil

}
