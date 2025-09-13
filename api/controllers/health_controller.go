package controllers

import "net/http"

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (c *HealthController) Answer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
