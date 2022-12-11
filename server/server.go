package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	message "github.com/leeswindell/dchat/messages"
)

type User struct {
	DisplayName string
}

type server struct {
	mu       sync.Mutex
	users    []User
	messages []message.Message
	mux      mux.Router
}

func newServer() {
	r := mux.NewRouter()
	r.HandleFunc("", HomeHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

}
