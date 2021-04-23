package httphandler

import (
	"net/http"

	"github.com/bartam1/mauth/internal/application/domain"
	"github.com/bartam1/mauth/internal/application/service"
	"github.com/bartam1/mauth/pkg/errors/httperror"
	"github.com/go-chi/render"
)

type Entity struct {
	service service.Interface
}

func New(service service.Interface) Entity {
	return Entity{service}
}

func (h Entity) LoginPage(w http.ResponseWriter, r *http.Request) {
	page, err := h.service.LoginPage()
	if err != nil {
		httperror.InternalError("Can't get login page!", err, w, r)
		return
	}

	render.HTML(w, r, page)
}
func (h Entity) UserAuth(w http.ResponseWriter, r *http.Request) {
	ua := domain.UserAuth{}
	if err := render.Decode(r, &ua); err != nil {
		httperror.BadRequest("invalid-request", err, w, r)
		return
	}
	err := h.service.UserAuth(r.Context(), ua)
	if err != nil {
		httperror.Unauthorised("Auth failed!", err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}
