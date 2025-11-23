package router

import (
    "net/http"
    "spotify/controller"
)

func NewWithMux(mux *http.ServeMux) *http.ServeMux {
    mux.HandleFunc("/damso", controller.Damso)
    return mux
}
