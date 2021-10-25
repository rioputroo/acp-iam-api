package request

type UpdateUsersRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	RolesId  uint   `json:"roles_id"`
	IsActive bool   `json:"is_active"`
}
