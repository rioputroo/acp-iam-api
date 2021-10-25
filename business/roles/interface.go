package roles

type Service interface {
	//GetAllRoles get list of all available roles
	GetAllRoles() ([]Roles, error)

	//GetRoles get detail of a role
	GetRoles(id uint) (*Roles, error)

	//AddRoles add new role to the system, return error if failed
	AddRoles(name string, isActive bool) error

	//UpdateRoles update a role in the system, return error if failed
	UpdateRoles(id uint, name string, isActive bool) error

	//DeleteRoles delete a role from the system, return error if failed
	DeleteRoles(id uint) error
}

type Repository interface {
	//GetAllRoles get list of all available roles
	GetAllRoles() ([]Roles, error)

	//GetRoles get detail of a role
	GetRoles(id uint) (*Roles, error)

	//AddRoles add new role to the system, return error if failed
	AddRoles(name string, isActive bool) error

	//UpdateRoles update a role in the system, return error if failed
	UpdateRoles(id uint, roles Roles) error

	//DeleteRoles delete a role from the system, return error if failed
	DeleteRoles(id uint) error
}
