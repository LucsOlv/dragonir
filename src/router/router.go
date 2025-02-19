package router

import (
	"dragonir/src/handlers"
	"dragonir/src/views"

	"github.com/gorilla/mux"
)

func NewRouter() (*mux.Router, error) {
	// Initialize template
	template, err := views.NewTemplate()
	if err != nil {
		return nil, err
	}

	// Initialize handlers
	taskHandler := handlers.NewTaskHandler(template)

	// Create router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", taskHandler.Index).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")

	return r, nil
}
