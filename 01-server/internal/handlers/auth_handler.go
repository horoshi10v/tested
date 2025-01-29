package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"01-server/internal/models"
	"01-server/internal/repository"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/swaggo/http-swagger"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewAuthHandler(userRepo repository.UserRepository, secret string) *AuthHandler {
	return &AuthHandler{
		userRepo:  userRepo,
		jwtSecret: secret,
	}
}

// Register godoc
// @Summary      Register new user
// @Description  Creates a new user with hashed password and sets JWT token in cookie
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.UserRequest  true  "User Data"
// @Success      201   {object}  models.User
// @Failure      400   {string}  string "Invalid JSON"
// @Failure      500   {string}  string "Internal Server Error"
// @Router       /register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Cannot hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashed)

	createdID, err := h.userRepo.CreateUser(user)
	if err != nil {
		http.Error(w, "Cannot register user", http.StatusInternalServerError)
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": createdID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		http.Error(w, "JWT error", http.StatusInternalServerError)
		return
	}

	// Set JWT in cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    tokenStr,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 2),
		SameSite: http.SameSiteStrictMode,
	})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created and logged in",
	})
}

// Login godoc
// @Summary      Login
// @Description  Authenticates user and sets JWT token in cookie
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.UserRequest  true  "Username/Password"
// @Success      200          {string}  string "Login successful"
// @Failure      401          {string}  string "Unauthorized"
// @Failure      500          {string}  string "Internal Server Error"
// @Router       /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var cred models.User
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	user, err := h.userRepo.GetByUsername(cred.Username)
	if err != nil || user.ID == 0 {
		http.Error(w, "Invalid username", http.StatusUnauthorized)
		return
	}
	// Compare hash
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password)) != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		http.Error(w, "JWT error", http.StatusInternalServerError)
		return
	}

	// Set JWT in cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    tokenStr,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 2),
		SameSite: http.SameSiteStrictMode,
	})

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
	})
}

// Logout godoc
// @Summary      Logout
// @Description  Clears JWT token from cookie
// @Tags         Auth
// @Success      200  {string}  string "Logged out successfully"
// @Router       /logout [post]
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		SameSite: http.SameSiteStrictMode,
	})

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged out successfully",
	})
}
