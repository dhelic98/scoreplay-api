package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhelic98/scoreplay-api/application/dto"
	"github.com/dhelic98/scoreplay-api/application/service"
	"github.com/dhelic98/scoreplay-api/interface/validator"
	"github.com/google/uuid"
)

type TagHandler struct {
	Service service.ITagService
}

func NewTagHandler(tagService service.ITagService) *TagHandler {
	return &TagHandler{Service: tagService}
}

func (handler *TagHandler) CreateTagHandler(w http.ResponseWriter, r *http.Request) {
	var createTagdto dto.CreateTagDTO

	if err := json.NewDecoder(r.Body).Decode(&createTagdto); err != nil {
		http.Error(w, "Malformed JSON input", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := validator.GetValidatorInstance().Struct(createTagdto); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	if err := handler.Service.CreateTag(r.Context(), &createTagdto); err != nil {
		http.Error(w, "Failed to create tag", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *TagHandler) GetAllTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := handler.Service.GetAllTags(r.Context())
	if err != nil {
		http.Error(w, "Tags not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)

}

func (handler *TagHandler) GetTagByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	tag, err := handler.Service.GetTagById(r.Context(), id)
	if err != nil {
		http.Error(w, "Tag not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tag)
}
