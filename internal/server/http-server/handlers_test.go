package httpserver

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	// "strings"
	"time"
)

func (suite *handlerTestSuite) TestHandlerFindPerson() {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	testCases := []struct {
		name         string
		initPersonID func() string
		expectedCode int
	}{
		{
			name: "success",
			initPersonID: func() string {
				personID := rnd.Intn(32)
				return strconv.Itoa(personID)
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid person id",
			initPersonID: func() string {
				return "invalid person id"
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			id := tc.initPersonID()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/find/%s", id), nil)

			suite.srv.ServeHTTP(w, req)
			suite.Equal(tc.expectedCode, w.Code)
		})
	}
}

func (suite *handlerTestSuite) TestHandlerDeletePerson() {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	testCases := []struct {
		name         string
		initPersonID func() string
		expectedCode int
	}{
		{
			name: "success",
			initPersonID: func() string {
				personID := rnd.Intn(32)
				return strconv.Itoa(personID)
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid person id",
			initPersonID: func() string {
				return "invalid person id"
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			id := tc.initPersonID()

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/delete/%s", id), nil)

			suite.srv.ServeHTTP(w, req)
			suite.Equal(tc.expectedCode, w.Code)
		})
	}
}

func (suite *handlerTestSuite) TestHandlerUpdatePerson() {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	testCases := []struct {
		name         string
		initPersonID func() string
		expectedCode int
	}{
		{
			name: "success",
			initPersonID: func() string {
				personID := rnd.Intn(32)
				return strconv.Itoa(personID)
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid person id",
			initPersonID: func() string {
				return "invalid person id"
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			id := tc.initPersonID()

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("/update/%s", id), nil)

			suite.srv.ServeHTTP(w, req)
			suite.Equal(tc.expectedCode, w.Code)
		})
	}
}

func (suite *handlerTestSuite) TestHandlerAddPerson() {
	testCases := []struct {
		name         string
		body         string
		expectedCode int
	}{
		{
			name: "success",
			body: `{
				"name": "Dmitriy",
				"surname": "Ushakov",
				"patronymic": "Vasilevich"
			}`,
			expectedCode: http.StatusCreated,
		},
		{
			name:         "bad request",
			body:         `{bad_request}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			w := httptest.NewRecorder()

			b := strings.NewReader(tc.body)
			req, _ := http.NewRequest(http.MethodPost, "/add", b)

			suite.srv.ServeHTTP(w, req)
			suite.Equal(tc.expectedCode, w.Code)
		})
	}
}
