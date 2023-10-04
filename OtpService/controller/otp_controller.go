package controller

import (
	"OtpService/dtos"
	"OtpService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type userController struct {
	service service.OtpService
}

func NewUserController(userService service.OtpService, router *mux.Router) *userController {

	controller := userController{service: userService}
	router.HandleFunc("/otp/validate", controller.ValidateOtp).Methods("POST")
	router.HandleFunc("/otp/generate", controller.GenerateOtp).Methods("POST")

	return &controller
}

func (c userController) ValidateOtp(w http.ResponseWriter, r *http.Request) {
	var request dtos.ValidateOtpRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	response, err := c.service.ValidateOtp(&request)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func (c userController) GenerateOtp(w http.ResponseWriter, r *http.Request) {
	var request dtos.GenerateOtpRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	response, err := c.service.RequestOtp(&request)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
