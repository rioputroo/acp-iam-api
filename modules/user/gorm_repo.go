package user

import (
	"acp-iam-api/business/user"
	"gorm.io/gorm"
	"time"
)

type GormRepository struct {
	DB *gorm.DB
}

func (repo *GormRepository) Login(email string, password string) (string, error) {
	panic("implement me")
}

type UserTable struct {
	ID         int       `gorm:"id;primaryKey;autoIncrement"`
	Name       string    `gorm:"name"`
	Email      string    `gorm:"email"`
	Password   string    `gorm:"password"`
	Roles      int       `gorm:"roles"`
	IsActive   bool      `gorm:"is_active"`
	CreatedAt  time.Time `gorm:"created_at"`
	ModifiedAt time.Time `gorm:"modified_at"`
}

//NewGormDBRepository Generate Gorm DB auth repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (repo *GormRepository) Register(email string, password string) (string, error) {
	userData := newUserTable(email, password)

	err := repo.DB.Create(userData).Error
	if err != nil {
		return "", err
	}

	return "", nil

}

func newUserTable(email string, password string) *UserTable {
	return &UserTable{
		1,
		"Rio",
		email,
		password,
		1,
		true,
		time.Now(),
		time.Now(),
	}
}

func (repo *GormRepository) FindUserByEmailAndPassword(email string, password string) (*user.User, error) {
	var userData UserTable

	err := repo.DB.Where("email = ?", email).Where("password = ?", password).First(&userData).Error
	if err != nil {
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}

func (col *UserTable) ToUser() user.User {
	var user user.User

	user.Id = col.ID
	user.Name = col.Name
	user.Password = col.Password
	user.Roles = col.Roles
	user.IsActive = col.IsActive
	user.CreatedAt = col.CreatedAt
	user.ModifiedAt = col.ModifiedAt

	return user
}
