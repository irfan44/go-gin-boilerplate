package example_handler

func (h *exampleHandler) MapRoutes() {
	h.r.Group("/accounts").
		GET("", h.GetExamples).
		GET("/:id", h.GetExampleById).
		POST("", h.CreateExample).
		PUT("/:id", h.UpdateExample).
		DELETE("/:id", h.DeleteExample)
}
