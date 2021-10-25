package roles

import (
	"acp-iam-api/business/roles"
	"gorm.io/gorm"
)

//GormRepository The implementation of roles.Repository object
type GormRepository struct {
	DB *gorm.DB
}

type RolesTable struct {
	gorm.Model
	ID       uint   `gorm:"id;primaryKey;autoIncrement"`
	Name     string `gorm:"name"`
	IsActive bool   `gorm:"is_active"`
}

func newRolesTable(roles2 roles.Roles) *RolesTable {

	return &RolesTable{
		roles2.Model,
		roles2.ID,
		roles2.Name,
		roles2.IsActive,
	}

}

func (col *RolesTable) ToRoles() roles.Roles {
	var roles roles.Roles

	roles.Model.ID = col.ID
	roles.Name = col.Name
	roles.IsActive = col.IsActive

	return roles
}

//NewGormDBRepository Generate Gorm DB roles repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (repo *GormRepository) GetAllRoles() ([]roles.Roles, error) {
	var rolesTable []RolesTable

	err := repo.DB.Find(&rolesTable).Error

	if err != nil {
		return nil, err
	}

	var result []roles.Roles

	for _, value := range rolesTable {
		result = append(result, value.ToRoles())
	}

	return result, nil
}

func (repo *GormRepository) GetRoles(id uint) (*roles.Roles, error) {
	var rolesTable RolesTable

	err := repo.DB.Where("id = ?", id).First(&rolesTable).Error

	if err != nil {
		return nil, err
	}

	result := rolesTable.ToRoles()

	return &result, nil
}

func (repo *GormRepository) AddRoles(name string, isActive bool) error {
	err := repo.DB.Create(&RolesTable{ID: 0, Name: name, IsActive: isActive}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) UpdateRoles(id uint, rolesParam roles.Roles) error {

	rolesData := newRolesTable(rolesParam)

	err := repo.DB.Model(&rolesData).Where("id = ? ", id).Updates(map[string]interface{}{"name": rolesData.Name, "is_active": rolesData.IsActive}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) DeleteRoles(id uint) error {
	var rolesTable RolesTable

	err := repo.DB.Model(&rolesTable).Where("id = ?", id).Updates(map[string]interface{}{"is_active": false}).Delete(&rolesTable).Error

	if err != nil {
		return err
	}

	return nil
}
