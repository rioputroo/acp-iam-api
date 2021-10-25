package request

type InsertUsersRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RolesId  uint   `json:"roles_id"`
	IsActive bool   `json:"is_active"`
}
