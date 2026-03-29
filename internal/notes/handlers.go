package notes

import (
	"log"
	"net/http"

	"github.com/Davidmuthee12/notes-rest-API/internal/json"
)

type handler struct {
	service Service
}

type createNoteRequest struct {
	Content string `json:"content"`
}

func NewHandler(service Service) *handler {
	return &handler {
		service: service,
	}
}

func (h handler) ListNotes(w http.ResponseWriter, r *http.Request) {
	// 1. Call the Service --> ListNotes
	//  2. Return JSON in a HTTP Response
	notes, err := h.service.ListNotes(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to list notes", http.StatusInternalServerError)
		return 
	}

	json.Write(w, http.StatusOK, notes)
}

func (h handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req createNoteRequest

	// Decode request body
	err := json.Read(r.Body, &req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.Content == "" {
		http.Error(w, "Content is required", http.StatusBadRequest)
		return
	}

	// Call service with content
	note, err := h.service.CreateNote(r.Context(), req.Content)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, note)
}