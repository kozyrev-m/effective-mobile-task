package httpserver

import (
	"testing"

	mocks "github.com/kozyrev-m/effective-mobile-task/internal/store/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type handlerTestSuite struct {
	suite.Suite
	mock *mocks.MockStore
	srv  *HTTPServer
}

// SetupSuite executes before the test suite begins execution.
func (suite *handlerTestSuite) SetupSuite() {
}

// TearDownSuite executes after all tests executed.
func (suite *handlerTestSuite) TearDownSuite() {
}

// SetupTest executes before each test cases.
func (suite *handlerTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.mock = mocks.NewMockStore(ctrl)
	suite.srv = New(suite.mock)
}

// TearDownTest executes after each test case.
func (suite *handlerTestSuite) TearDownTest() {
}

// TestHandlerTestSuite is entry point for testing the handlers.
func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(handlerTestSuite))
}
