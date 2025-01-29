package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"01-server/internal/models"
	"01-server/internal/repository"
	_ "github.com/swaggo/http-swagger"
)

type SellerHandler struct {
	repo repository.SellerRepository
}

func NewSellerHandler(repo repository.SellerRepository) *SellerHandler {
	return &SellerHandler{
		repo: repo,
	}
}

func (h *SellerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/sellers" {
		switch r.Method {
		case http.MethodGet:
			h.getAll(w, r)
		case http.MethodPost:
			h.create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		h.getByID(w, id)
	case http.MethodPut:
		h.update(w, r, id)
	case http.MethodDelete:
		h.delete(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getAll godoc
// @Summary      Get all sellers
// @Description  Returns array of sellers
// @Tags         Sellers
// @Produce      json
// @Success      200  {array}   models.Seller
// @Router       /sellers [get]
func (h *SellerHandler) getAll(w http.ResponseWriter, r *http.Request) {
	sellers, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sellers)
}

// create godoc
// @Summary      Create seller
// @Description  Creates a new seller
// @Tags         Sellers
// @Accept       json
// @Produce      json
// @Param        seller  body      models.SellerRequest  true  "Seller Data"
// @Success      201     {object}  models.Seller
// @Failure      400     {string}  string "Invalid JSON"
// @Router       /sellers [post]
// @Security     BearerAuth
func (h *SellerHandler) create(w http.ResponseWriter, r *http.Request) {
	var s models.Seller
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	newID, err := h.repo.Create(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.ID = newID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

// getByID godoc
// @Summary      Get seller by ID
// @Description  Returns a single seller by its ID
// @Tags         Sellers
// @Produce      json
// @Param        id   path      int  true  "Seller ID"
// @Success      200  {object}  models.Seller
// @Failure      404  {string}  string "Not found"
// @Failure      400  {string}  string "Invalid ID"
// @Router       /sellers/{id} [get]
func (h *SellerHandler) getByID(w http.ResponseWriter, id int) {
	s, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if s.ID == 0 {
		http.NotFound(w, nil)
		return
	}
	json.NewEncoder(w).Encode(s)
}

// update godoc
// @Summary      Update seller
// @Description  Updates existing seller by ID
// @Tags         Sellers
// @Accept       json
// @Produce      json
// @Param        id      path      int           true "Seller ID"
// @Param        seller  body      models.SellerRequest true "Seller Data"
// @Success      200     {object}  models.Seller
// @Failure      400     {string}  string "Invalid JSON / ID"
// @Router       /sellers/{id} [put]
// @Security     BearerAuth
func (h *SellerHandler) update(w http.ResponseWriter, r *http.Request, id int) {
	var s models.Seller
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	s.ID = id
	if err := h.repo.Update(s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(s)
}

// delete godoc
// @Summary      Delete seller
// @Description  Removes a seller by ID
// @Tags         Sellers
// @Param        id   path      int  true  "Seller ID"
// @Success      204  {string}  string "No Content"
// @Failure      400  {string}  string "Invalid ID"
// @Router       /sellers/{id} [delete]
// @Security     BearerAuth
func (h *SellerHandler) delete(w http.ResponseWriter, id int) {
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
