package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/YReshetko/docker-app/backend/internal/model"
	"log"
	"net/http"
)

type Services struct {
	model *model.Model
}

func NewServices(model *model.Model) *Services {
	return &Services{
		model: model,
	}
}

func (s *Services) getServices(w http.ResponseWriter, r *http.Request) {
	s.model.Services()

	body, err := json.Marshal(s.model.Services())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		handleErrorsOnWrite(fmt.Fprint(w, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	handleErrorsOnWrite(w.Write(body))
}

func handleErrorsOnWrite(_ int, err error) {
	if err != nil {
		log.Println(err)
	}
}

func (s *Services) Handlers() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		http.MethodGet: s.getServices,
	}

}

func (s *Services) Path() string {
	return "/api/v1/services"
}
