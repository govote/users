package controllers

import (
	"net/http"
	"testing"

	"github.com/deputadosemfoco/go-libs/messages"
	"github.com/deputadosemfoco/users/interactors"
	"github.com/deputadosemfoco/users/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWhenCreatingNewUserStatusCodeMustBeCreated201(t *testing.T) {
	// setup
	param := new(interactors.RegistrationRequest)
	param.Email = "test@test.com"
	param.Name = "vitor"
	param.FacebookID = "32132"
	param.PhotoURL = "cadsa"

	userCtrl := new(UserCtrl)
	fakeInteractor := new(FakeRegistrationInteractor)

	fakeInteractor.On("Register", param).Return(interactors.RegistrationResult{Response: messages.Response{Success: true}, Created: true})
	userCtrl.Interactor = fakeInteractor

	c, res, err := test.CreateJsonContext(param)

	test.PanicErr(err)

	// execution
	err = userCtrl.Post(c)

	// assertions
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestUserWithInvalidInformationShouldReturnBadRequest400(t *testing.T) {
	// setup
	param := new(interactors.RegistrationRequest)
	param.Name = "vitor"
	param.FacebookID = "32132"
	param.PhotoURL = "cadsa"

	userCtrl := new(UserCtrl)
	fakeInteractor := new(FakeRegistrationInteractor)

	fakeInteractor.On("Register", param).Return(interactors.RegistrationResult{Response: messages.Response{Success: false, Message: "error"}})
	userCtrl.Interactor = fakeInteractor

	c, res, err := test.CreateJsonContext(param)

	test.PanicErr(err)

	// execution
	err = userCtrl.Post(c)

	// assertions
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestAlreadyRegisteredUserShouldReturnOk200(t *testing.T) {
	// setup
	param := new(interactors.RegistrationRequest)
	param.Email = "test@test.com"
	param.Name = "vitor"
	param.FacebookID = "32132"
	param.PhotoURL = "cadsa"

	userCtrl := new(UserCtrl)
	fakeInteractor := new(FakeRegistrationInteractor)

	fakeInteractor.On("Register", param).Return(interactors.RegistrationResult{Response: messages.Response{Success: true}, Created: false})
	userCtrl.Interactor = fakeInteractor

	c, res, err := test.CreateJsonContext(param)

	test.PanicErr(err)

	// execution
	err = userCtrl.Post(c)

	// assertions
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
}

type FakeRegistrationInteractor struct{ mock.Mock }

func (fake *FakeRegistrationInteractor) Register(request *interactors.RegistrationRequest) interactors.RegistrationResult {
	args := fake.Called(request)

	return args.Get(0).(interactors.RegistrationResult)
}
