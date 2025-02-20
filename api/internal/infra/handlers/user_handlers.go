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

type Error struct {
	Message string `json:"message"`
}

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

// Get JWT token
// @Summary			Get token
// @Description		Get JWT Token authorization
// @Tags			Users
// @Accept			json
// @Produce			json
// @Param 			request	body	dto.JwtInput	true	"User credentials"
// @Success 		201	{object}	dto.GetJWTOutput
// @BadRequest		400
// @Unauthorized	401
// @NotFound 		404
// @Failure 		500	{object}	Error
// @Router			/users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.JwtInput
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: "user not found"}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: "password incorrect"}
		json.NewEncoder(w).Encode(error)
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

// Create user godoc
// @Summary			Create user
// @Description		Create user
// @Tags			Users
// @Accept			json
// @Produce			json
// @Param 			request	body	dto.CreateUserInput	true	"User request"
// @Success 		201
// @Failure 		500	{object}	Error
// @Router			/users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: "create user fail"}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
