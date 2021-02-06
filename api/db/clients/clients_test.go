package clients_test

import (
	"testing"

	"github.com/allez-chauffe/marcel/api/db"
	"github.com/allez-chauffe/marcel/api/db/clients"
	"github.com/allez-chauffe/marcel/api/db/internal/test_utils"
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

func testAllMethods(runner test_utils.Runner, t *testing.T) {
	t.Run("GetNotFound", testGetNotFound(runner))
	t.Run("InsertAndGet", testInsertAndGet(runner))
	t.Run("List", testList(runner))
	t.Run("Update", testUpdate(runner))
	t.Run("Delete", testDelete(runner))
	t.Run("DeleteAll", testDeleteAll(runner))
}

func testGetNotFound(runner test_utils.Runner) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				client, err := db.Clients().Get("not found")
				if err != nil {
					t.Fatalf("Get failed: %s", err)
				}

				if client != nil {
					t.Fatalf("Client should be nil if not exists")
				}
				return nil
			})
		})
	}
}

func testInsertAndGet(runner test_utils.Runner) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				client := clients.New()
				t.Log("before insert")
				if err := tx.Clients().Insert(client); err != nil {
					t.Fatalf("Insertion failed: %s", err)
				}
				t.Log("after insert")
				if inserted, _ := tx.Clients().Exists(client.ID); !inserted {
					t.Fatalf("Client should have been inserted")
				}
				t.Log("after get")
				return nil
			})
		})
	}
}

func testList(runner test_utils.Runner) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				tx.Clients().Insert(clients.New())
				tx.Clients().Insert(clients.New())
				tx.Clients().Insert(clients.New())

				list, err := tx.Clients().List()

				if err != nil {
					t.Fatalf("List failed: %s", err)
				}

				if len(list) != 3 {
					t.Fatalf("Expected a list of 3 clients, %d found", len(list))
				}

				if list[0] == list[1] || list[1] == list[2] || list[0] == list[2] {
					t.Fatalf("Each clients should be different")
				}
				return nil
			})
		})
	}
}

func testUpdate(runner test_utils.Runner) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				client := clients.New()
				tx.Clients().Insert(client)

				client.Name = "test"

				if err := tx.Clients().Update(client); err != nil {
					t.Fatalf("Update failed: %s", err)
				}

				if updated, _ := tx.Clients().Get(client.ID); updated.Name != "test" {
					t.Fatalf("Expected name to be update to 'test', '%s' found", updated.Name)
				}
				return nil
			})
		})
	}
}

func testDelete(runner test_utils.Runner) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				client := clients.New()
				tx.Clients().Insert(client)

				if err := tx.Clients().Delete(client.ID); err != nil {
					t.Fatalf("Delete failed: %s", err)
				}

				if exists, _ := tx.Clients().Exists(client.ID); exists {
					t.Error("Client should have been deleted")
				}
				return nil
			})
		})
	}
}

func testDeleteAll(runner test_utils.Runner) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				tx.Clients().Insert(clients.New())
				tx.Clients().Insert(clients.New())
				tx.Clients().Insert(clients.New())

				if err := tx.Clients().DeleteAll(); err != nil {
					t.Fatalf("DeleteAll failed: %s", err)
				}

				if list, _ := tx.Clients().List(); len(list) != 0 {
					t.Error("All client should have been deleted")
				}
				return nil
			})
		})
	}
}
