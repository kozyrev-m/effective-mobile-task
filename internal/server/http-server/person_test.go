package httpserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"

	"time"

	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
	"go.uber.org/mock/gomock"
)

var (
	errTestCreatePerson = errors.New("Can't create person")
	errTestFindPerson   = errors.New("Can't find person")
	errTestDeletePerson = errors.New("Can't delete person")
	errTestUpdatePerson = errors.New("Can't update person")
)

func (suite *handlerTestSuite) TestHandlerFindPerson() {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	testCases := []struct {
		name         string
		expectedCode int
		initMock     func() string
	}{
		{
			name:         "success",
			expectedCode: http.StatusOK,
			initMock: func() string {
				// init person id for test
				id := rnd.Intn(32)

				// init person for test
				person := entities.TestPerson()
				person.ID = uint64(id)

				// init mock for current test case
				suite.store.EXPECT().GetPersonByID(gomock.Any(), uint64(id)).Return(&person, nil)

				return strconv.Itoa(id)
			},
		},
		{
			name:         "fail",
			expectedCode: http.StatusUnprocessableEntity,
			initMock: func() string {
				// init person id for test
				id := rnd.Intn(32)

				// init mock for current test case
				suite.store.EXPECT().GetPersonByID(gomock.Any(), uint64(id)).Return(nil, errTestFindPerson)

				return strconv.Itoa(id)
			},
		},
		{
			name:         "invalid person id",
			expectedCode: http.StatusBadRequest,
			initMock: func() string {
				return "something_bad"
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			expectedID := tc.initMock()

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/find/%s", expectedID), nil)

			suite.srv.ServeHTTP(w, req)
			suite.Equal(tc.expectedCode, w.Code)
		})
	}
}

func (suite *handlerTestSuite) TestHandlerDeletePerson() {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	testCases := []struct {
		name         string
		expectedCode int
		initMock     func() string
	}{
		{
			name:         "success",
			expectedCode: http.StatusOK,
			initMock: func() string {
				// init person id for test
				id := rnd.Intn(32)

				// init mock for current test case
				suite.store.EXPECT().DeletePerson(gomock.Any(), uint64(id)).Return(uint64(id), nil)

				return strconv.Itoa(id)
			},
		},
		{
			name:         "fail",
			expectedCode: http.StatusUnprocessableEntity,
			initMock: func() string {
				// init person id for test
				id := rnd.Intn(32)

				// init mock for current test case
				suite.store.EXPECT().DeletePerson(gomock.Any(), uint64(id)).Return(uint64(0), errTestDeletePerson)

				return strconv.Itoa(id)
			},
		},
		{
			name:         "invalid person id",
			expectedCode: http.StatusBadRequest,
			initMock: func() string {
				return "something_bad"
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			expectedID := tc.initMock()

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/delete/%s", expectedID), nil)

			suite.srv.ServeHTTP(w, req)
			suite.Equal(tc.expectedCode, w.Code)
		})
	}
}

func (suite *handlerTestSuite) TestHandlerUpdatePerson() {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	testCases := []struct {
		name         string
		initMock     func() (string, []byte)
		expectedCode int
	}{
		{
			name:         "success",
			expectedCode: http.StatusOK,
			initMock: func() (string, []byte) {
				// init person id for test
				id := rnd.Intn(32)

				// init person for test
				person := entities.TestPerson()
				person.ID = uint64(id)

				// init params for updating
				newName := "Sergey"
				newAge := 45
				newNationality := "USA"
				params := entities.Person{Name: &newName, Age: &newAge, Nationality: &newNationality}

				// init updated person data for mock
				updatedPerson := entities.TestPerson()
				updatedPerson.ID = uint64(id)
				updatedPerson.Name = &newName
				updatedPerson.Age = &newAge
				updatedPerson.Nationality = &newNationality

				// init mock for current test case
				suite.store.EXPECT().GetPersonByID(gomock.Any(), uint64(id)).Return(&person, nil)
				suite.store.EXPECT().UpdatePerson(gomock.Any(), uint64(id), params).Return(&updatedPerson, nil)

				// make body
				body, _ := json.Marshal(params)

				return strconv.Itoa(id), body
			},
		},
		{
			name:         "fail",
			expectedCode: http.StatusUnprocessableEntity,
			initMock: func() (string, []byte) {
				// init person id for test
				id := rnd.Intn(32)

				// init person for test
				person := entities.TestPerson()
				person.ID = uint64(id)

				// init params for updating
				newName := "Sergey"
				newAge := 45
				newNationality := "USA"
				params := entities.Person{Name: &newName, Age: &newAge, Nationality: &newNationality}

				// init mock for current test case
				suite.store.EXPECT().GetPersonByID(gomock.Any(), uint64(id)).Return(&person, nil)
				suite.store.EXPECT().UpdatePerson(gomock.Any(), uint64(id), params).Return(nil, errTestUpdatePerson)

				// make body
				body, _ := json.Marshal(params)

				return strconv.Itoa(id), body
			},
		},
		{
			name:         "invalid person id",
			expectedCode: http.StatusBadRequest,
			initMock: func() (string, []byte) {
				return "something_bad", []byte("")
			},
		},
		{
			name:         "bad request",
			expectedCode: http.StatusBadRequest,
			initMock: func() (string, []byte) {
				id := rnd.Intn(32)
				body := "{bad_request}"

				return strconv.Itoa(id), []byte(body)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			id, body := tc.initMock()

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("/update/%s", id), bytes.NewBuffer(body))

			suite.srv.ServeHTTP(w, req)
			suite.Equal(tc.expectedCode, w.Code)
		})
	}
}

func (suite *handlerTestSuite) TestHandlerAddPerson() {
	testCases := []struct {
		name         string
		expectedCode int
		// body         string
		initTestCase func() []byte
	}{
		{
			name:         "success",
			expectedCode: http.StatusCreated,
			initTestCase: func() []byte {
				// init person for test
				person := entities.TestPerson()

				// init mock store for current test case
				suite.store.EXPECT().CreatePerson(gomock.Any(), person).Return(uint64(1), nil)

				suite.agent.EXPECT().ReceiveAndSet(gomock.Any(), person).Return(&person, nil)

				// make body
				body, _ := json.Marshal(person)

				return body
			},
		},
		{
			name:         "fail",
			expectedCode: http.StatusUnprocessableEntity,
			initTestCase: func() []byte {
				// init person for test
				person := entities.Person{}

				// init mock store for current test case
				suite.store.EXPECT().CreatePerson(gomock.Any(), person).Return(uint64(0), errTestCreatePerson)

				// init mock agent for current test case
				suite.agent.EXPECT().ReceiveAndSet(gomock.Any(), person).Return(&person, nil)

				// make body
				body, _ := json.Marshal(person)

				return body
			},
		},
		{
			name:         "bad request",
			expectedCode: http.StatusBadRequest,
			initTestCase: func() []byte {
				// make body
				body := []byte(`{bad_request}`)

				return body
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			body := tc.initTestCase()

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(body))

			suite.srv.ServeHTTP(w, req)
			suite.Equal(tc.expectedCode, w.Code)
		})
	}
}
