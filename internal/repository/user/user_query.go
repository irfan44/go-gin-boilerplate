package user_repository

const (
	GET_USERS        = "SELECT id, username, role, created_at, updated_at FROM users"
	GET_USER_BY_ID   = "SELECT id, username, role, created_at, updated_at FROM users WHERE id = $1"
	GET_USER_BY_NAME = "SELECT id, username, password, role, created_at, updated_at FROM users WHERE username = $1"
	CREATE_USER      = "INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id, username, role, created_at, updated_at"
	UPDATE_USER      = "UPDATE users SET username = $1, role = $2, updated_at = $3 WHERE id = $4 RETURNING  id, username, role, created_at, updated_at"
	DELETE_USER      = "DELETE FROM users WHERE id = $1"
)
