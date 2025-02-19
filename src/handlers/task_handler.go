package handlers

import (
	"encoding/json"
	"net/http"

	"dragonir/src/models"
	"dragonir/src/views"
)

type TaskHandler struct {
	template *views.Template
	tasks    []models.Task // In-memory storage for demonstration
}

func NewTaskHandler(template *views.Template) *TaskHandler {
	return &TaskHandler{
		template: template,
		tasks:    make([]models.Task, 0),
	}
}

func (h *TaskHandler) Index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
		Tasks []models.Task
	}{
		Title: "Home",
		Tasks: h.tasks,
	}

	err := h.template.Render(w, "layout.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add task to in-memory storage
	h.tasks = append(h.tasks, task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
