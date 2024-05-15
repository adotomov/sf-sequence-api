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

type SequenceService interface {
	GetSequences(context.Context) ([]*models.Sequence, error)
	CreateSequence(context.Context, types.CreateSequenceRequest) (*models.Sequence, error)
	GetSequenceByID(context.Context, string) (*models.Sequence, error)
	UpdateSequence(context.Context, types.UpdateSequenceRequest) (*models.Sequence, error)
}

type SequenceHandler struct {
	s SequenceService
}

func NewSequenceHandler(svc SequenceService) *SequenceHandler {
	return &SequenceHandler{
		s: svc,
	}
}

func (h *SequenceHandler) GetSequences(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	res, err := h.s.GetSequences(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonRes, err := json.Marshal(res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

func (h *SequenceHandler) CreateSequence(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var rb types.CreateSequenceRequest

	if err := util.DecodeJSONBody(w, r, &rb); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.s.CreateSequence(ctx, rb)
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

func (h *SequenceHandler) GetSequenceByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	id := r.URL.Query().Get("sequenceId")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.s.GetSequenceByID(ctx, id)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonRes, err := json.Marshal(res)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

func (h *SequenceHandler) UpdateSequence(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var rb types.UpdateSequenceRequest

	if err := util.DecodeJSONBody(w, r, &rb); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.s.UpdateSequence(ctx, rb)
	if err != nil {
		log.Print(err)
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
