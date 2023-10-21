package httpserver

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type handlerTestSuite struct {
	suite.Suite
	srv *HTTPServer
}

// SetupSuite executes before the test suite begins execution.
func (suite *handlerTestSuite) SetupSuite() {
}

// TearDownSuite executes after all tests executed.
func (suite *handlerTestSuite) TearDownSuite() {
}

// SetupTest executes before each test cases.
func (suite *handlerTestSuite) SetupTest() {
	suite.srv = New()
}

// TearDownTest executes after each test case.
func (suite *handlerTestSuite) TearDownTest() {
}

// TestHandlerTestSuite is entry point for testing the handlers
func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(handlerTestSuite))
}
