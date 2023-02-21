package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	medapp "github.com/mnogohoddovochka/med-app"
	"github.com/mnogohoddovochka/med-app/pkg/service"
	mock_service "github.com/mnogohoddovochka/med-app/pkg/service/mocks"
)

// go install github.com/golang/mock/mockgen@v1.6.0

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorisation, doctor medapp.Doctor)
	testTable := []struct {
		name                string
		inputBody           string
		inputDoctor         medapp.Doctor
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test", "surname":"test", "login":"testLogin", "password":"qwerty"}`,
			inputDoctor: medapp.Doctor{
				Name:     "Test",
				Surname:  "test",
				Login:    "testLogin",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorisation, doctor medapp.Doctor) {
				s.EXPECT().CreateDoctor(doctor).Return(3, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":3}`,
		},
		{
			name:                "Empty Fields",
			inputBody:           `{"name":"Test", "surname":"test", "password":"qwerty"}`,
			mockBehavior:        func(s *mock_service.MockAuthorisation, doctor medapp.Doctor) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"name":"Test", "surname":"test", "login":"testLogin", "password":"qwerty"}`,
			inputDoctor: medapp.Doctor{
				Name:     "Test",
				Surname:  "test",
				Login:    "testLogin",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorisation, doctor medapp.Doctor) {
				s.EXPECT().CreateDoctor(doctor).Return(3, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorisation(c)
			testCase.mockBehavior(auth, testCase.inputDoctor)

			services := &service.Service{Authorisation: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
