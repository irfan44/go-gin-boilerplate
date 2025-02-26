package user_handler

func (h *userHandler) MapRoutes() {
	h.r.Group("/users", h.m.AdminAuthorization()).
		GET("", h.GetUsers).
		GET("/:id", h.GetUserById).
		POST("", h.CreateUser).
		PUT("/:id", h.UpdateUser).
		DELETE("/:id", h.DeleteUser)
}
