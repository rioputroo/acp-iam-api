package user

import "time"

//User define base user model for admin/customer
type User struct {
	Id         int
	Email      string
	Password   string
	Name       string
	Roles      int
	IsActive   bool
	CreatedAt  time.Time
	ModifiedAt time.Time
}

//Roles define all roles available for auth
type Roles struct {
	id         int
	name       string
	isActive   bool
	createdAt  time.Time
	modifiedAt time.Time
}
