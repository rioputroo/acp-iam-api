package roles

import (
	"acp-iam-api/business"
)

type service struct {
	repository Repository
}

//NewService Construct roles service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

//GetAllRoles get all available roles, return nil if empty
func (s service) GetAllRoles() ([]Roles, error) {
	roles, err := s.repository.GetAllRoles()

	if err != nil {
		return []Roles{}, err
	}

	return roles, err
}

//AddRoles add new role to the system
func (s service) AddRoles(name string, isActive bool) error {
	return s.repository.AddRoles(name, isActive)
}

//GetRoles get single roles from the system
func (s service) GetRoles(id uint) (*Roles, error) {
	role, err := s.repository.GetRoles(id)

	if err != nil {
		return nil, business.ErrNotFound
	}

	return role, nil
}

//UpdateRoles update a role in the system
func (s service) UpdateRoles(id uint, name string, isActive bool) error {
	rolesData, err := s.repository.GetRoles(id)

	if err != nil {
		return err
	} else if rolesData == nil {
		return business.ErrNotFound
	}

	updatedRoles := rolesData.ModifyRoles(name, isActive)

	return s.repository.UpdateRoles(id, updatedRoles)
}

func (s service) DeleteRoles(id uint) error {
	return s.repository.DeleteRoles(id)
}
