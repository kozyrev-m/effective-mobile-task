package httpserver

import (
	"testing"

	mock_integration "github.com/kozyrev-m/effective-mobile-task/internal/integration/mocks"
	mock_store "github.com/kozyrev-m/effective-mobile-task/internal/store/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type handlerTestSuite struct {
	suite.Suite
	store *mock_store.MockStore
	agent *mock_integration.MockIntegration
	srv   *HTTPServer
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
	suite.store = mock_store.NewMockStore(ctrl)
	suite.agent = mock_integration.NewMockIntegration(ctrl)
	suite.srv = New(suite.store, suite.agent)
}

// TearDownTest executes after each test case.
func (suite *handlerTestSuite) TearDownTest() {
}

// TestHandlerTestSuite is entry point for testing the handlers.
func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(handlerTestSuite))
}
