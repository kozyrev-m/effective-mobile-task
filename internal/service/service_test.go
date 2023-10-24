package service

import (
	"testing"

	mocks "github.com/kozyrev-m/effective-mobile-task/internal/store/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type serviceTestSuite struct {
	suite.Suite
	svc  *Service
	mock *mocks.MockStore
}

// SetupSuite executes before the test suite begins execution.
func (suite *serviceTestSuite) SetupSuite() {
}

// TearDownSuite executes after all tests executed.
func (suite *serviceTestSuite) TearDownSuite() {
}

// SetupTest executes before each test cases.
func (suite *serviceTestSuite) SetupTest() {
	suite.mock = mocks.NewMockStore(gomock.NewController(suite.T()))
	suite.svc = NewService(suite.mock)
}

// TearDownTest executes after each test case.
func (suite *serviceTestSuite) TearDownTest() {
}

// TestServiceTestSuite is entry point for testing layer service.
func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(serviceTestSuite))
}
