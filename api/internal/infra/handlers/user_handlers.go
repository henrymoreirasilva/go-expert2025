package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/henrymoreirasilva/go-api/internal/dto"
	"github.com/henrymoreirasilva/go-api/internal/entity"
	"github.com/henrymoreirasilva/go-api/internal/infra/database"
)

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, JwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       db,
		Jwt:          jwt,
		JwtExpiresIn: JwtExpiresIn,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.JwtInput
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	if !u.ValidatePassword(user.Password) {
		http.Error(w, "password incorrect", http.StatusUnauthorized)
		return
	}

	mapClaims := map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	}
	_, tokenString, _ := h.Jwt.Encode(mapClaims)
	accessToken := map[string]string{"AccessToken": tokenString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": accessToken["AccessToken"]})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		http.Error(w, "create user fail", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
