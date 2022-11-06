package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	cartdto "waysfood/dto/cart"
	dto "waysfood/dto/result"
	"waysfood/models"
	"waysfood/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart  {
	return &handlerCart{CartRepository}
}

func (h *handlerCart ) FindCart(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	carts,err:=h.CartRepository.FindCart()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: carts}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) GEtCart (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	transID := int(userInfo["id"].(float64))

	cart, err := h.CartRepository.GEtCart(transID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

w.WriteHeader(http.StatusOK)
response := dto.SuccessResult{Status:"success", Data: cart}
json.NewEncoder(w).Encode(response)
}


func (h *handlerCart) CreateCart (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	userInfo :=r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	request:= new (cartdto.CreateCart)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}	

	
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	time:=time.Now()
	miliTime := time.Unix()

	cart:=models.Cart{
		ID: int(miliTime),
		UserID: userId,
		Status:"pending",
	}

	statusCheck,_:=h.CartRepository.FindbyIDCart(userId,"pending")
	if statusCheck.Status == "pending" {
		response := dto.SuccessResult{Status:"success pending", Data: cart}
		json.NewEncoder(w).Encode(response)
	} else {
		data, _ := h.CartRepository.CreateCart(cart)
		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Status: "success transaction ", Data: data}
		json.NewEncoder(w).Encode(response)
	}
}

func (h *handlerCart) DeleteCart (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	id,_:= strconv.Atoi(mux.Vars(r)["id"])
	cart,_:=h.CartRepository.GEtCart(id)

	data,_:=h.CartRepository.DeleteCart(cart)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}


func (h *handlerCart) FindbyIDCart (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	
	userInfo :=r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	cart,_:=h.CartRepository.FindbyIDCart(userId,"pending")

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: cart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) UpdateCart(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	request:= new(cartdto.UpdateCart)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	cart, err := h.CartRepository.FindbyIDCart(userId, "pending")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	if request.SubTotal !=0{
		cart.UserID= request.SubTotal
	}
	
	if request.Total !=0{
		cart.Total= request.Total
	}
	if request.QTY !=0{
		cart.QTY= request.QTY
	}
	// if request.Shipping !=0{
	// 	cart.Shipping= request.Shipping
	// }

	if request.Status !="pending"{
		cart.Status= request.Status
	}
	
	data,_:=h.CartRepository.UpdateCart(cart)
	
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

