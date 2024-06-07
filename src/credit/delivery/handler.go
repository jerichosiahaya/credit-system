package delivery

import (
	"context"
	"encoding/json"
	"kredit-plus/src/common"
	"kredit-plus/src/domain"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

type creditDelivery struct {
	ctx context.Context
	uc  domain.Usecase
}

func New(router *chi.Mux, ctx context.Context, uc domain.Usecase) *chi.Mux {
	handler := &creditDelivery{
		uc: uc,
		ctx: ctx,
	}
	corsMiddleware := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Api-Key"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300,
    })

	router.Use(middleware.Heartbeat("/ping"), corsMiddleware.Handler)
	router.Post("/create/customer", handler.CreateCustomer)
	return router
}

func (d *creditDelivery) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	errorResponse := domain.GetResponse{}

	data := domain.Customer{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		errorResponse.Status = string(common.Error)
		errorResponse.Message = err.Error()
		response, _ := json.Marshal(errorResponse)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(response)
		return
	}
	
	data, err = d.uc.CreateCustomer(d.ctx, data)
	if err != nil {
		errorResponse.Status = string(common.Error)
		errorResponse.Message = err.Error()
		response, _ := json.Marshal(errorResponse)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(response)
		return
	}

	res := domain.GetResponse{
		Status:  string(common.Success),
		Message: string(common.SuccessPost),
		Data:    data,
	}
	response, err := json.Marshal(res)
	if err != nil {
		errorResponse.Status = string(common.Error)
		errorResponse.Message = err.Error()
		response, _ := json.Marshal(errorResponse)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(response)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (d *creditDelivery) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	errorResponse := domain.GetResponse{}

	data := domain.Transaction{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		errorResponse.Status = string(common.Error)
		errorResponse.Message = err.Error()
		response, _ := json.Marshal(errorResponse)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(response)
		return
	}
	
	data, err = d.uc.CreateTransaction(d.ctx, data)
	if err != nil {
		errorResponse.Status = string(common.Error)
		errorResponse.Message = err.Error()
		response, _ := json.Marshal(errorResponse)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(response)
		return
	}

	res := domain.GetResponse{
		Status:  string(common.Success),
		Message: string(common.SuccessPost),
		Data:    data,
	}
	response, err := json.Marshal(res)
	if err != nil {
		errorResponse.Status = string(common.Error)
		errorResponse.Message = err.Error()
		response, _ := json.Marshal(errorResponse)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(response)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}