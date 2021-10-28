package request

type UpdateRolesRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
