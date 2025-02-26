package auth_handler

func (h *authHandler) MapRoutes() {
	h.r.Group("/auth").
		POST("/login", h.Login).
		POST("/register", h.Register)
}
