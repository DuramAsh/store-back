package http

import (
	"net/http"
	"store-back/internal/domain/product"
	"store-back/internal/service/general"
	"store-back/pkg/server/response"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GeneralHandler struct {
	GeneralService *general.Service
}

func NewGeneralHandler(generalService *general.Service) *GeneralHandler {
	return &GeneralHandler{GeneralService: generalService}
}

func (h *GeneralHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/products", h.selectProducts)

	r.Post("/", h.createOrder)
	r.Get("/{id}", h.getOrderByID)
	r.Get("/orders/{email}", h.selectOrdersByClient)

	r.Post("/login", h.login)
	r.Post("/register", h.register)

	return r
}

func (h *GeneralHandler) selectProducts(w http.ResponseWriter, r *http.Request) {
	res, err := h.GeneralService.SelectProducts(r.Context())
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

func (h *GeneralHandler) selectOrdersByClient(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	res, err := h.GeneralService.SelectOrdersByClient(r.Context(), email)
	println("email", email)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

func (h *GeneralHandler) getOrderByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.GeneralService.GetOrderByID(r.Context(), id)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

func (h *GeneralHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	req := product.OrderRequest{}

	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, r.Body)
		return
	}

	res, err := h.GeneralService.CreateOrder(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

func (h *GeneralHandler) login(w http.ResponseWriter, r *http.Request) {
	req := product.LoginRequest{}

	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, r.Body)
		return
	}

	res, err := h.GeneralService.Login(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

func (h *GeneralHandler) register(w http.ResponseWriter, r *http.Request) {
	req := product.LoginRequest{}

	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, r.Body)
		return
	}

	res, err := h.GeneralService.Register(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}
