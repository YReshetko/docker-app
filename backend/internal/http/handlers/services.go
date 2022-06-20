package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/YReshetko/docker-app/backend/internal/model"
)

//Services struct for handling services requests
//@Rest(path = "/api/v1/services")
type Services struct {
	model *model.Model
}

func NewServices(model *model.Model) *Services {
	return &Services{
		model: model,
	}
}

// @Rest(method = "GET")
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
