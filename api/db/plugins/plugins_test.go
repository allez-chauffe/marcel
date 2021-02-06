package plugins_test

import (
	"testing"

	"github.com/allez-chauffe/marcel/api/db"
	"github.com/allez-chauffe/marcel/api/db/internal/test_utils"
	"github.com/allez-chauffe/marcel/api/db/plugins"
)

func TestWithoutTransaction(t *testing.T) {
	tx := &db.Tx{}
	testAllMethods(func(test func(*db.Tx) error) error {
		test(tx)
		return nil
	}, t)
}

func TestWithTransaction(t *testing.T) {
	testAllMethods(db.Transactional, t)
}

func testAllMethods(runner func(func(tx *db.Tx) error) error, t *testing.T) {
	t.Run("GetNotFound", testGetNotFound(runner))
	t.Run("InsertAndGet", testInsertAndGet(runner))
	t.Run("Exists", testExists(runner))
	t.Run("List", testList(runner))
	t.Run("Update", testUpdate(runner))
	t.Run("Delete", testDelete(runner))
	t.Run("UpsertAll", testUpsertAll(runner))
}

func testExists(runner func(func(*db.Tx) error) error) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				exists, err := tx.Plugins().Exists("test")
				if err != nil {
					t.Errorf("Existence check failed: %s", err)
				}

				if exists {
					t.Error("Should not exist in an empty database")
				}

				tx.Plugins().Insert(plugins.New("test"))

				exists, err = tx.Plugins().Exists("test")

				if err != nil {
					t.Errorf("Existence check failed: %s", err)
				}

				if !exists {
					t.Errorf("Should exists")
				}
				return nil
			})
		})
	}
}

func testGetNotFound(runner func(func(tx *db.Tx) error) error) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				plugin, err := tx.Plugins().Get("not found")
				if err != nil {
					t.Fatalf("Get failed: %s", err)
				}

				if plugin != nil {
					t.Fatal("Should return nil when it is not found")
				}
				return nil
			})
		})
	}
}

func testInsertAndGet(runner func(func(tx *db.Tx) error) error) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				plugin := plugins.New("test")
				if err := tx.Plugins().Insert(plugin); err != nil {
					t.Fatalf("Insert failed: %s", err)
				}

				saved, err := tx.Plugins().Get(plugin.EltName)
				if err != nil {
					t.Fatalf("Get failed: %s", err)
				}

				if saved == nil {
					t.Fatal("Should find newly inserted plugin")
				}
				return nil
			})
		})
	}
}

func testList(runner func(func(tx *db.Tx) error) error) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				tx.Plugins().Insert(plugins.New("test1"))
				tx.Plugins().Insert(plugins.New("test2"))
				tx.Plugins().Insert(plugins.New("test3"))

				list, err := tx.Plugins().List()

				if err != nil {
					t.Fatalf("List failed: %s", err)
				}

				if len(list) != 3 {
					t.Fatalf("Expected a list of 3 plugins, %d found", len(list))
				}

				if list[0].EltName == list[1].EltName ||
					list[1].EltName == list[2].EltName ||
					list[0].EltName == list[2].EltName {
					t.Fatal("Each plugin should be different")
				}
				return nil
			})
		})
	}
}

func testUpdate(runner func(func(tx *db.Tx) error) error) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				plugin := plugins.New("test")
				tx.Plugins().Insert(plugin)

				plugin.Description = "test"

				if err := tx.Plugins().Update(plugin); err != nil {
					t.Fatalf("Update failed: %s", err)
				}

				if plugin, _ = tx.Plugins().Get(plugin.EltName); plugin.Description != "test" {
					t.Fatal("Expect plugin description to be 'test', '%' found", plugin.Description)
				}
				return nil
			})
		})
	}
}

func testDelete(runner func(func(tx *db.Tx) error) error) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				plugin := plugins.New("test")
				tx.Plugins().Insert(plugin)

				if err := tx.Plugins().Delete(plugin.EltName); err != nil {
					t.Fatalf("Delete failed: %s", err)
				}

				if exists, _ := tx.Plugins().Exists(plugin.EltName); exists {
					t.Fatal("Plugin should have been deleted")
				}
				return nil
			})
		})
	}
}

func testUpsertAll(runner func(func(tx *db.Tx) error) error) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				plugin := plugins.New("test1")
				tx.Plugins().Insert(plugin)
				plugin.Description = "test"

				err := tx.Plugins().UpsertAll([]plugins.Plugin{
					*plugin,
					*plugins.New("test2"),
					*plugins.New("test3"),
				})

				if err != nil {
					t.Fatalf("UpsertAll failed: %s", err)
				}

				list, err := tx.Plugins().List()
				if err != nil {
					t.Fatalf("List failed: %s", err)
				}

				if len(list) != 3 {
					t.Fatalf("Expected 3 medias, %d found", len(list))
				}

				if list[0].Description != "test" {
					t.Fatalf("Expected plugin description to be 'test', '%s' found", list[0].Description)
				}

				if list[0].EltName == list[1].EltName ||
					list[1].EltName == list[2].EltName ||
					list[0].EltName == list[2].EltName {
					t.Fatalf("All ids should be diferent")
				}
				return nil
			})
		})
	}
}
