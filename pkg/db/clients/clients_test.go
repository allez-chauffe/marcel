package clients_test

import (
	"testing"

	"github.com/allez-chauffe/marcel/pkg/db"
	"github.com/allez-chauffe/marcel/pkg/db/clients"
	"github.com/allez-chauffe/marcel/pkg/db/internal/test_utils"
	uuid "github.com/satori/go.uuid"
)

var drivers [2]string = [2]string{"bolt", "postgres"}

func TestClients(t *testing.T) {
	test_utils.DatabaseTest(t, map[string]test_utils.DatabaseTestFunc{
		"GetNotFound":  testGetNotFound,
		"InsertAndGet": testInsertAndGet,
		"List":         testList,
		"Update":       testUpdate,
		"Delete":       testDelete,
		"DeleteAll":    testDeleteAll,
	})
}

func testGetNotFound(tx *db.Tx, t *testing.T) {
	client, err := db.Clients().Get(uuid.NewV4().String())
	if err != nil {
		t.Fatalf("Get failed: %s", err)
	}

	if client != nil {
		t.Fatalf("Client should be nil if not exists")
	}
}

func testInsertAndGet(tx *db.Tx, t *testing.T) {
	client := clients.New()
	if err := tx.Clients().Insert(client); err != nil {
		t.Fatalf("Insertion failed: %s", err)
	}
	if inserted, _ := tx.Clients().Exists(client.ID); !inserted {
		t.Fatalf("Client should have been inserted")
	}
}

func testList(tx *db.Tx, t *testing.T) {
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
}

func testUpdate(tx *db.Tx, t *testing.T) {
	client := clients.New()
	tx.Clients().Insert(client)

	client.Name = "test"

	if err := tx.Clients().Update(client); err != nil {
		t.Fatalf("Update failed: %s", err)
	}

	if updated, _ := tx.Clients().Get(client.ID); updated.Name != "test" {
		t.Fatalf("Expected name to be update to 'test', '%s' found", updated.Name)
	}
}

func testDelete(tx *db.Tx, t *testing.T) {
	client := clients.New()
	tx.Clients().Insert(client)

	if err := tx.Clients().Delete(client.ID); err != nil {
		t.Fatalf("Delete failed: %s", err)
	}

	if exists, _ := tx.Clients().Exists(client.ID); exists {
		t.Error("Client should have been deleted")
	}
}

func testDeleteAll(tx *db.Tx, t *testing.T) {
	tx.Clients().Insert(clients.New())
	tx.Clients().Insert(clients.New())
	tx.Clients().Insert(clients.New())

	if err := tx.Clients().DeleteAll(); err != nil {
		t.Fatalf("DeleteAll failed: %s", err)
	}

	if list, _ := tx.Clients().List(); len(list) != 0 {
		t.Error("All client should have been deleted")
	}
}
