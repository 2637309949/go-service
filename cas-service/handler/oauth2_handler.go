package handler

import (
	"comm/auth"
	"comm/logger"
	"comm/mark"
	whttp "comm/service/web/http"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-session/session/v3"
)

func (h *Handler) OAuth2Authorize(rw http.ResponseWriter, r *http.Request) {
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "OAuth2Authorize")()

	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do OAuth2Authorize", acc.Name)
	}

	store, err := session.Start(r.Context(), rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	r.Form = form

	store.Delete("ReturnUri")
	store.Save()

	timemark.Mark("HandleAuthorizeRequest")
	err = h.OAuthServer.HandleAuthorizeRequest(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
}

func (h *Handler) OAuth2Token(rw http.ResponseWriter, r *http.Request) {
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "OAuth2Token")()

	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do OAuth2Token", acc.Name)
	}

	timemark.Mark("HandleTokenRequest")
	err := h.OAuthServer.HandleTokenRequest(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) OAuth2Affirm(rw http.ResponseWriter, r *http.Request) {
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "OAuth2Affirm")()

	store, err := session.Start(r.Context(), rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		rw.Header().Set("Location", "/login")
		rw.WriteHeader(http.StatusFound)
		return
	}

	whttp.OutputHTML(rw, r, "static/affirm.html")
}

func (h *Handler) UserAuthorizeHandler(rw http.ResponseWriter, r *http.Request) (userID string, err error) {
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "UserAuthorizeHandler")()

	store, err := session.Start(r.Context(), rw, r)
	if err != nil {
		return
	}

	err = r.ParseForm()
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		store.Set("ReturnUri", r.Form)
		store.Save()

		rw.Header().Set("Location", "/cas/user/login")
		rw.WriteHeader(http.StatusFound)
		return
	}
	userID = fmt.Sprintf("%v", uid)
	store.Save()
	return
}
