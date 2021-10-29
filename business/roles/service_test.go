package roles_test

import (
	_ "acp-iam-api/api/iam/roles"
	"acp-iam-api/api/iam/roles/request"
	"acp-iam-api/business"
	"acp-iam-api/business/roles"
	rolesMock "acp-iam-api/business/roles/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

const id uint = 0

const (
	name      = "admin"
	is_active = true
)

var (
	rolesService    roles.Service
	rolesRepository rolesMock.Repository

	rolesData       roles.Roles
	rolesArray      []roles.Roles
	insertRolesData request.InsertRolesRequest
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {

	rolesData = roles.NewRoles(
		name,
		is_active,
	)

	rolesArray = append(rolesArray, rolesData)

	insertRolesData = request.InsertRolesRequest{
		Name:     name,
		IsActive: is_active,
	}

	rolesService = roles.NewService(&rolesRepository)
}

func TestServiceGetRoles(t *testing.T) {
	t.Run("Expect found the roles", func(t *testing.T) {
		rolesRepository.On("GetRoles", mock.AnythingOfType("uint")).Return(&rolesData, nil).Once()

		role, err := rolesService.GetRoles(id)

		assert.Nil(t, err)

		assert.NotNil(t, role)

		assert.Equal(t, id, role.Model.ID)
		assert.Equal(t, name, role.Name)
		assert.Equal(t, is_active, role.IsActive)
	})

	t.Run("Expect roles not found", func(t *testing.T) {
		rolesRepository.On("GetRoles", mock.AnythingOfType("uint")).Return(nil, business.ErrNotFound).Once()

		role, err := rolesService.GetRoles(id)

		assert.NotNil(t, err)

		assert.Nil(t, role)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestServiceGetAllRoles(t *testing.T) {
	t.Run("Expect found all list roles", func(t *testing.T) {
		rolesRepository.On("GetAllRoles", mock.Anything).Return(rolesArray, nil).Once()

		roles, err := rolesService.GetAllRoles()

		assert.Nil(t, err)

		assert.Equal(t, 1, len(roles))

		assert.NotNil(t, roles)
	})

	t.Run("Expect failed found all list roles", func(t *testing.T) {
		rolesRepository.On("GetAllRoles", mock.Anything).Return(nil, business.ErrNotFound).Once()

		roles, err := rolesService.GetAllRoles()

		assert.NotNil(t, err)

		assert.Empty(t, roles)
	})
}

func TestServiceAddRoles(t *testing.T) {
	t.Run("Expect add role success", func(t *testing.T) {
		rolesRepository.On("AddRoles", mock.AnythingOfType("string"), mock.AnythingOfType("bool")).Return(nil).Once()

		err := rolesService.AddRoles(name, is_active)

		assert.Nil(t, err)
	})

	t.Run("Expect add roles not found", func(t *testing.T) {
		rolesRepository.On("AddRoles", mock.AnythingOfType("string"), mock.AnythingOfType("bool")).Return(business.ErrInternalServerError).Once()

		err := rolesService.AddRoles(name, is_active)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})
}

func TestServiceUpdateRoles(t *testing.T) {
	t.Run("Expect update role success", func(t *testing.T) {
		rolesRepository.On("GetRoles", mock.AnythingOfType("uint")).Return(&rolesData, nil).Once()
		rolesRepository.On("UpdateRoles", mock.AnythingOfType("uint"), mock.AnythingOfType("roles.Roles")).Return(nil).Once()

		err := rolesService.UpdateRoles(id, name, is_active)

		assert.Nil(t, err)

	})

	t.Run("Expect update roles failed", func(t *testing.T) {
		rolesRepository.On("GetRoles", mock.AnythingOfType("uint")).Return(nil, business.ErrNotFound).Once()
		rolesRepository.On("UpdateRoles", mock.AnythingOfType("uint"), mock.AnythingOfType("roles.Roles")).Return(business.ErrNotFound).Once()

		err2 := rolesService.UpdateRoles(id, name, is_active)

		assert.NotNil(t, err2)

		assert.Equal(t, err2, business.ErrNotFound)
	})
}

func TestServiceDeleteRoles(t *testing.T) {
	t.Run("Expect delete role success", func(t *testing.T) {
		rolesRepository.On("GetRoles", mock.AnythingOfType("uint")).Return(&rolesData, nil).Once()
		rolesRepository.On("DeleteRoles", mock.AnythingOfType("uint")).Return(nil).Once()

		err := rolesService.DeleteRoles(id)

		assert.Nil(t, err)
	})

	t.Run("Expect delete role failed", func(t *testing.T) {
		rolesRepository.On("GetRoles", mock.AnythingOfType("uint")).Return(&rolesData, nil).Once()
		rolesRepository.On("DeleteRoles", mock.AnythingOfType("uint")).Return(business.ErrInternalServerError).Once()

		err := rolesService.DeleteRoles(id)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})
}
