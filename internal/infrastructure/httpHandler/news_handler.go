package httpHandler

import (
	"encoding/json"
	"net/http"
	"news-back-go/internal/app/core"
	"news-back-go/internal/app/ports"

	"github.com/gorilla/mux"
)

type NewsHandler struct {
	service ports.NewsRepository
}

func NewNewsHandler(service ports.NewsRepository) *NewsHandler {
	return &NewsHandler{service: service}
}

func (h *NewsHandler) CreateNews(w http.ResponseWriter, r *http.Request) {
	var news core.News
	json.NewDecoder(r.Body).Decode(&news)
	err := h.service.Create(&news)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *NewsHandler) GetByIdNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, exists := vars["id"]
	if !exists {
		http.Error(w, "El ID es requerido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para obtener la noticia
	news, err := h.service.GetById(id)
	if err != nil {
		http.Error(w, "No se encontró la noticia", http.StatusNotFound)
		return
	}

	// Enviar la noticia como respuesta en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func (h *NewsHandler) GetAllNews(w http.ResponseWriter, r *http.Request) {
	// Llamar al servicio para obtener la noticia
	news, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "No se encontró noticias", http.StatusNotFound)
		return
	}

	// Enviar la noticia como respuesta en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func (h *NewsHandler) UpdateNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var news core.News

	id, exists := vars["id"]
	if !exists {
		http.Error(w, "El ID es requerido", http.StatusBadRequest)
		return
	}

	var updatedNews core.News
	err := json.NewDecoder(r.Body).Decode(&updatedNews)
	if err != nil {
		http.Error(w, "Formato de solicitud inválido", http.StatusBadRequest)
		return
	}

	updatedNews.ID = id

	// Llamar al servicio para obtener la noticia
	err = h.service.Update(&updatedNews)
	if err != nil {
		http.Error(w, "No se pudo editar la noticia", http.StatusNotFound)
		return
	}

	// Enviar la noticia como respuesta en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func (h *NewsHandler) DeleteNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, exists := vars["id"]
	if !exists {
		http.Error(w, "El ID es requerido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para obtener la noticia
	err := h.service.Delete(id)
	if err != nil {
		http.Error(w, "No se pudo eliminar la noticia", http.StatusNotFound)
		return
	}

	// Enviar la noticia como respuesta en formato JSON
	w.WriteHeader(http.StatusAccepted)
}
