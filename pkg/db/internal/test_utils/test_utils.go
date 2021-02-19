package test_utils

import (
	"database/sql"
	"fmt"
	"path"
	"testing"

	"github.com/allez-chauffe/marcel/pkg/config"
	"github.com/allez-chauffe/marcel/pkg/db"
	"github.com/sirupsen/logrus"
)

var logLevel = logrus.WarnLevel

func DatabaseTest(t *testing.T, tests map[string]DatabaseTestFunc) {
	t.Helper()
	config.SetDefault(config.New())
	if testing.Verbose() {
		config.Default().SetLogLevel(logrus.DebugLevel)
	} else {
		config.Default().SetLogLevel(logrus.WarnLevel)
	}

	testEveryDatabases(t, tests)
}

func testEveryDatabases(t *testing.T, tests map[string]DatabaseTestFunc) {
	t.Helper()
	t.Run("Bolt", testDatabaseTx(boltTester(t.Name()), tests))
	t.Run("Postgres", testDatabaseTx(postgresTester(t.Name()), tests))
}

func testDatabaseTx(tester databaseTester, tests map[string]DatabaseTestFunc) TestFunc {
	return func(t *testing.T) {
		t.Helper()
		t.Run("WithoutTransactions", runTests(tests, func(t *testing.T, test DatabaseTestFunc) {
			t.Helper()
			tester(t, func() {
				t.Helper()
				test(&db.Tx{}, t)
			})
		}))

		t.Run("WithTransaction", runTests(tests, func(t *testing.T, test DatabaseTestFunc) {
			t.Helper()
			tester(t, func() {
				t.Helper()
				db.Transactional(func(tx *db.Tx) error {
					t.Helper()
					test(tx, t)
					return nil
				})
			})
		}))
	}
}

func runTests(tests map[string]DatabaseTestFunc, runner func(*testing.T, DatabaseTestFunc)) TestFunc {
	return func(t *testing.T) {
		t.Helper()
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				t.Helper()
				runner(t, test)
			})
		}
	}
}

type Runner func(task func(tx *db.Tx) error) error
type TestFunc func(*testing.T)
type DatabaseTestFunc func(tx *db.Tx, t *testing.T)
type databaseTester func(*testing.T, func())

func boltTester(name string) func(*testing.T, func()) {
	return func(t *testing.T, test func()) {
		t.Helper()
		config.Default().API().DB().Bolt().SetFile(path.Join(t.TempDir(), "marcel.db"))
		config.Default().API().DB().SetDriver("bolt")
		if err := db.Open(); err != nil {
			panic(err)
		}
		defer db.Close()
		test()
	}
}

func postgresTester(name string) func(*testing.T, func()) {
	return func(t *testing.T, test func()) {
		// WORKAROUND: Generate uniq db name for each test package to allow concurent runs
		dbName := fmt.Sprintf("marce_test_%s", name)
		t.Helper()
		config.Default().API().DB().SetDriver("postgres")
		pgConfig := config.Default().API().DB().Postgres()
		pgConfig.SetDBName(dbName)
		pgConfig.SetUsername("postgres")
		pgConfig.SetPassword("password")
		pgConfig.SetHost("localhost")
		pgConfig.SetPort("5432")

		if err := db.Open(); err != nil {
			panic(err)
		}
		defer postgresCleanup(t, dbName)
		defer db.Close()
		test()
	}
}

func postgresCleanup(t *testing.T, dbName string) {
	t.Helper()
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err = db.Exec(fmt.Sprintf(`DROP DATABASE "%s"`, dbName)); err != nil {
		panic(err)
	}
}
