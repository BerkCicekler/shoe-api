package image

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ImageServiceHandler struct{}

func (h *ImageServiceHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/image/{filename}", h.handleImage).Methods("GET")
}

func (s *ImageServiceHandler) handleImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	imagePath := "images/" + filename

	w.Header().Set("Content-Type", "image/svg+xml")

	http.ServeFile(w, r, imagePath)
}