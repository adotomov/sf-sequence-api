package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/adotomov/sf-sequence-api/internal/models"
	"github.com/adotomov/sf-sequence-api/internal/types"
	util "github.com/adotomov/sf-sequence-api/internal/utils"
)

type StepService interface {
	CreateStep(context.Context, types.CreateStepRequest) (*models.Step, error)
	UpdateStep(context.Context, types.UpdateStepRequest) (*models.Step, error)
	DeleteStep(context.Context, string) error
}

type StepHandler struct {
	s StepService
}

func NewStepHandler(svc StepService) *StepHandler {
	return &StepHandler{
		s: svc,
	}
}

func (h *StepHandler) CreateStep(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var rb types.CreateStepRequest

	if err := util.DecodeJSONBody(w, r, &rb); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.s.CreateStep(ctx, rb)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonRes, err := json.Marshal(res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonRes)
}

func (h *StepHandler) UpdateStep(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var rb types.UpdateStepRequest

	if err := util.DecodeJSONBody(w, r, &rb); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.s.UpdateStep(ctx, rb)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonRes, err := json.Marshal(res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonRes)
}

func (h *StepHandler) DeleteStep(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	id := r.URL.Query().Get("stepId")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.s.DeleteStep(ctx, id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
