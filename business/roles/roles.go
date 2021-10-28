package roles

import (
	"gorm.io/gorm"
)

//Roles core data type for roles
type Roles struct {
	gorm.Model
	Name     string
	IsActive bool
}

func (oldData *Roles) Error() string {
	panic("implement me")
}

type RolesUserResponse struct {
	id        uint
	name      string
	is_active bool
}

//NewRoles create new roles
func NewRoles(name string, isActive bool) Roles {
	return Roles{
		Name:     name,
		IsActive: isActive,
	}
}

//ModifyRoles update existing Roles data
func (oldData *Roles) ModifyRoles(newName string, newIsActive bool) Roles {
	return Roles{
		Name:     newName,
		IsActive: newIsActive,
	}
}
