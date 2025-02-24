package handler

import (
	"fmt"
	"net/http"
)

func (h *exampleHandler) MapRoutes() {
	h.mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, "/examples"),
		h.GetExamples(),
	)
	h.mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodGet, "/examples/{id}"),
		h.GetExampleById(),
	)
	h.mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodPost, "/examples"),
		h.CreateExample(),
	)
	h.mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodPut, "/examples/{id}"),
		h.UpdateExample(),
	)
	h.mux.HandleFunc(
		fmt.Sprintf("%s %s", http.MethodDelete, "/examples/{id}"),
		h.DeleteExample(),
	)
}
