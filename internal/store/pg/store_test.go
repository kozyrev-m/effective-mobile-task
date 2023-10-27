package pg

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

// test db address.
var testDatabaseDSN = "postgres://effmobile:effmobile@localhost:15439/effmobile_test"

type storeTestSuite struct {
	suite.Suite
	store *Store
}

// SetupSuite executes before the test suite begins execution.
func (suite *storeTestSuite) SetupSuite() {
	// test will be ignored if db is unavailable
	if err := suite.isDBAvailable(testDatabaseDSN); err != nil {
		suite.T().Skipf("skip db tests: database is not available: %v", err)
		return
	}

}

// TearDownSuite executes after all tests executed.
func (suite *storeTestSuite) TearDownSuite() {
}

// SetupTest executes before each test cases.
func (suite *storeTestSuite) SetupTest() {
	// clear test bd
	err := suite.clearDB(testDatabaseDSN)
	suite.NoError(err)

	suite.store, err = NewStore(testDatabaseDSN)

	suite.NoError(err)
}

// TearDownTest executes after each test case.
func (suite *storeTestSuite) TearDownTest() {
	suite.store.Close()
}

// isDBAvailable check db connection.
func (suite *storeTestSuite) isDBAvailable(dsn string) error {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	defer db.Close()

	return db.Ping()
}

// clearDB rolls back db.
func (suite *storeTestSuite) clearDB(dsn string) error {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	defer db.Close()

	return MigrationsDown(db)
}

// TestStoreTestSuite is entry point for testing the store.
func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(storeTestSuite))
}
