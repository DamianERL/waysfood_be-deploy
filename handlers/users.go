package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	dto "waysfood/dto/result"
	usersdto "waysfood/dto/users"
	"waysfood/models"
	"waysfood/repositories"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

// connection to repo ->db
type handlerUser struct {
	UserRepository repositories.UserRepository
}

// connection to routing
func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//look for every users
	//with method finduser
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	// for i, p := range users {
	// 	users[i].Image = p.Image
	// }


	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: users}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerUser) FindPartners (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users,err:=h.UserRepository.FindPartners("patner")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status:"success", Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//take id from params
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//get the specific user method from repo
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// user.Image = path_file + user.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	id := int(userInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile") // add this code
	filepath := dataContex.(string)             // add this code

		
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waysfood"})

	if err != nil {
		fmt.Println(err.Error())
	}


	
	request := usersdto.UpdateUserRequest{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Location: r.FormValue("location"),
		Phone:    r.FormValue("phone"),
		Image: resp.SecureURL,
	}


	user, _ := h.UserRepository.GetUser(int(id))

	//if they have changed
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Phone != "" {
		user.Phone = request.Phone
	}
	if request.Location != "" {
		user.Location = request.Location
	}
	if request.Image != "false" {
		user.Image = request.Image
	}

	//update method from repo
	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//take id from params
	//and find the Specific user
	//check if user exist
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//delete method from repo
	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//take request from client and  use dto for what data we take
	request := new(usersdto.UpdateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//use validation for check data (and bcz dto have validation to)
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//check again use models for what data will push to db
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Location: request.Location,
	}

	//create method from repo
	data, err := h.UserRepository.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}
