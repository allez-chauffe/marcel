package plugins_test

import (
	"testing"

	"github.com/allez-chauffe/marcel/api/db"
	"github.com/allez-chauffe/marcel/api/db/internal/test_utils"
	"github.com/allez-chauffe/marcel/api/db/plugins"
)

func TestPlugins(t *testing.T) {
	test_utils.DatabaseTest(t, map[string]test_utils.DatabaseTestFunc{
		"GetNotFound":  testGetNotFound,
		"InsertAndGet": testInsertAndGet,
		"Exists":       testExists,
		"List":         testList,
		"Update":       testUpdate,
		"Delete":       testDelete,
		"UpsertAll":    testUpsertAll,
	})
}

func testExists(tx *db.Tx, t *testing.T) {
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
}

func testGetNotFound(tx *db.Tx, t *testing.T) {
	plugin, err := tx.Plugins().Get("not found")
	if err != nil {
		t.Fatalf("Get failed: %s", err)
	}

	if plugin != nil {
		t.Fatal("Should return nil when it is not found")
	}
}

func testInsertAndGet(tx *db.Tx, t *testing.T) {
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
}

func testList(tx *db.Tx, t *testing.T) {
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
}

func testUpdate(tx *db.Tx, t *testing.T) {
	plugin := plugins.New("test")
	tx.Plugins().Insert(plugin)

	plugin.Description = "test"

	if err := tx.Plugins().Update(plugin); err != nil {
		t.Fatalf("Update failed: %s", err)
	}

	if plugin, _ = tx.Plugins().Get(plugin.EltName); plugin.Description != "test" {
		t.Fatal("Expect plugin description to be 'test', '%' found", plugin.Description)
	}
}

func testDelete(tx *db.Tx, t *testing.T) {
	plugin := plugins.New("test")
	tx.Plugins().Insert(plugin)

	if err := tx.Plugins().Delete(plugin.EltName); err != nil {
		t.Fatalf("Delete failed: %s", err)
	}

	if exists, _ := tx.Plugins().Exists(plugin.EltName); exists {
		t.Fatal("Plugin should have been deleted")
	}
}

func testUpsertAll(tx *db.Tx, t *testing.T) {
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
}
