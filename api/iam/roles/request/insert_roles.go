package request

type InsertRolesRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
