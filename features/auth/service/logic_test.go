package service

import (
	"14-api-clean-arch/features/auth"
	"14-api-clean-arch/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	repo := new(mocks.AuthRepository)

	t.Run("Success Login", func(t *testing.T) {
		inputData := auth.Core{Email: "Budi@gmail.com", Password: "13123123"}
		returnData := auth.Core{ID: 1, FullName: "Budi", Email: "Budi@gmail.com", Password: "13123123", Role: "user"}

		repo.On("FindUser").Return(returnData, nil).Once()
		srv := New(repo)
		token, errLogin := srv.Login(inputData)

		assert.NoError(t, errLogin, "Error is null")
		assert.NotEmpty(t, token, "Token generated.")
		repo.AssertExpectations(t)
	})

	t.Run("Failed Login - Empty Data", func(t *testing.T) {
		inputData1 := auth.Core{Email: "Budi@gmail.com", Password: ""}
		inputData2 := auth.Core{Email: "Budi@gmail.com", Password: ""}
		inputData3 := auth.Core{Email: "", Password: ""}
		returnData := auth.Core{}

		repo.On("FindUser").Return(returnData, errors.New("Failed process query")).Once()
		srv := New(repo)
		_, errLogin1 := srv.Login(inputData1)
		_, errLogin2 := srv.Login(inputData2)
		_, errLogin3 := srv.Login(inputData3)

		assert.Contains(t, errLogin1, "validate input", "Failed login. Empty Email.")
		assert.Contains(t, errLogin2, "validate input", "Failed login. Empty Password.")
		assert.Contains(t, errLogin3, "validate input", "Failed login. Empty Email and Password.")
		repo.AssertExpectations(t)
	})
}
