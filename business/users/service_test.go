package users_test

import (
	"acp-iam-api/api/iam/users/request"
	"acp-iam-api/business"
	"acp-iam-api/business/roles"
	"acp-iam-api/business/users"
	usersMock "acp-iam-api/business/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

const id uint = 0

const (
	email     = "rio@gmail.com"
	password  = "rioputro"
	name      = "Rio Trilaksono Putro"
	is_active = true
	roles_id  = 2
)

var (
	usersService    users.Service
	usersRepository usersMock.Repository

	usersData  users.Users
	rolesData  roles.Roles
	usersArray []users.Users

	insertUsersRequest request.InsertUsersRequest
	updateUsersRequest request.UpdateUsersRequest
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {

	insertUsersRequest = request.InsertUsersRequest{
		Name:     name,
		Email:    email,
		Password: password,
		RolesId:  roles_id,
		IsActive: is_active,
	}

	usersData = users.NewUsers(
		insertUsersRequest,
	)

	usersArray = append(usersArray, usersData)

	usersService = users.NewService(&usersRepository)
}

func TestService_GetAllUsers(t *testing.T) {
	t.Run("Expect found all list users", func(t *testing.T) {
		usersRepository.On("GetAllUsers", mock.Anything).Return(usersArray, nil).Once()

		users, err := usersService.GetAllUsers()

		assert.Nil(t, err)

		assert.Equal(t, 1, len(users))

		assert.NotNil(t, users)
	})
}

func TestService_GetUsers(t *testing.T) {
	t.Run("Expect found the user", func(t *testing.T) {
		usersRepository.On("GetUsers", mock.AnythingOfType("uint")).Return(&usersData, &rolesData, nil).Once()

		user, _, err := usersService.GetUsers(id)

		assert.Nil(t, err)

		assert.NotNil(t, user)

		assert.Equal(t, id, user.ID)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, roles_id, int(user.RolesId))
		assert.Equal(t, is_active, user.IsActive)
	})

	t.Run("Expect user not found", func(t *testing.T) {
		usersRepository.On("GetUsers", mock.AnythingOfType("uint")).Return(nil, nil, business.ErrNotFound).Once()

		user, _, err := usersService.GetUsers(id)

		assert.NotNil(t, err)

		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestService_AddUsers(t *testing.T) {
	t.Run("Expect add user success", func(t *testing.T) {
		usersRepository.On("AddUsers", mock.AnythingOfType("*request.InsertUsersRequest")).Return(nil).Once()

		err := usersService.AddUsers(&insertUsersRequest)

		assert.Nil(t, err)
	})

	t.Run("Expect add user not found", func(t *testing.T) {
		usersRepository.On("AddUsers", mock.AnythingOfType("*request.InsertUsersRequest")).Return(business.ErrInternalServerError).Once()

		err := usersService.AddUsers(&insertUsersRequest)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})
}

func TestService_UpdateUsers(t *testing.T) {
	t.Run("Expect update user success", func(t *testing.T) {
		usersRepository.On("GetUsers", mock.AnythingOfType("uint")).Return(&usersData, &rolesData, nil).Once()
		usersRepository.On("UpdateUsers", mock.AnythingOfType("uint"), mock.AnythingOfType("*request.UpdateUsersRequest")).Return(nil).Once()

		err := usersService.UpdateUsers(id, &updateUsersRequest)

		assert.Nil(t, err)

	})

	t.Run("Expect update user failed", func(t *testing.T) {
		usersRepository.On("GetUsers", mock.AnythingOfType("uint")).Return(&usersData, &rolesData, nil).Once()
		usersRepository.On("UpdateUsers", mock.AnythingOfType("uint"), mock.AnythingOfType("*request.UpdateUsersRequest")).Return(business.ErrInternalServerError).Once()

		err := usersService.UpdateUsers(id, &updateUsersRequest)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})
}

func TestService_DeleteUsers(t *testing.T) {
	t.Run("Expect delete user success", func(t *testing.T) {
		usersRepository.On("GetUsers", mock.AnythingOfType("uint")).Return(&usersData, &rolesData, nil).Once()
		usersRepository.On("DeleteUsers", mock.AnythingOfType("uint")).Return(nil).Once()

		err := usersService.DeleteUsers(id)

		assert.Nil(t, err)
	})

	t.Run("Expect delete user failed", func(t *testing.T) {
		usersRepository.On("GetUsers", mock.AnythingOfType("uint")).Return(&usersData, &rolesData, nil).Once()
		usersRepository.On("DeleteUsers", mock.AnythingOfType("uint")).Return(business.ErrInternalServerError).Once()

		err := usersService.DeleteUsers(id)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})
}
