package users

import (
	"acp-iam-api/api/iam/users/request"
	roles2 "acp-iam-api/business/roles"
	"acp-iam-api/business/users"
	"acp-iam-api/repository/roles"
	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

type UserTable struct {
	gorm.Model
	ID       uint   `gorm:"id;primaryKey;autoIncrement"`
	RolesID  uint   `gorm:"foreignKey:RolesID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
	Name     string `gorm:"name"`
	IsActive bool   `gorm:"is_active"`
}

func newUserTable(users2 users.Users) *UserTable {
	return &UserTable{
		RolesID:  users2.RolesId,
		Email:    users2.Email,
		Name:     users2.Name,
		IsActive: users2.IsActive,
	}
}

//NewGormDBRepository Generate Gorm DB users repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (col *UserTable) ToUser() users.Users {
	var users users.Users

	users.ID = col.ID
	users.Name = col.Name
	users.Email = col.Email
	users.IsActive = col.IsActive
	users.RolesId = col.RolesID

	return users
}

func (repo *GormRepository) GetAllUsers() ([]users.Users, error) {
	var userTable []UserTable

	err := repo.DB.Find(&userTable).Error

	if err != nil {
		return nil, err
	}

	var result []users.Users

	for _, value := range userTable {
		result = append(result, value.ToUser())
	}

	return result, nil
}

func (repo *GormRepository) GetUsers(id uint) (*users.Users, *roles2.Roles, error) {
	var usersTable UserTable
	var rolesTable roles.RolesTable

	err := repo.DB.Where("id = ?", id).First(&usersTable).Error

	if err != nil {
		return nil, nil, err
	}

	result := usersTable.ToUser()

	err2 := repo.DB.Where("id = ?", result.RolesId).First(&rolesTable).Error

	if err2 != nil {
		return nil, nil, err2
	}

	resultUserRoles := rolesTable.ToRoles()

	return &result, &resultUserRoles, nil
}

func (repo *GormRepository) AddUsers(request *request.InsertUsersRequest) error {
	err := repo.DB.Create(&UserTable{
		RolesID:  request.RolesId,
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
		IsActive: request.IsActive,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) UpdateUsers(id uint, users users.Users) error {
	usersData := newUserTable(users)

	err := repo.DB.Model(&usersData).Where("id = ? ", id).Updates(map[string]interface{}{"name": usersData.Name, "email": usersData.Email, "is_active": usersData.IsActive, "roles_id": usersData.RolesID}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) DeleteUsers(id uint) error {
	var usersTable UserTable

	err := repo.DB.Model(&usersTable).Where("id = ?", id).Updates(map[string]interface{}{"is_active": false}).Delete(&usersTable).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) Login(email string, password string) (*users.Users, error) {
	var userData UserTable

	err := repo.DB.Where("email = ?", email).Where("password = ?", password).First(&userData).Error
	if err != nil {
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}
