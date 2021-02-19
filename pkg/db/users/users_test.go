package users_test

import (
	"testing"

	"github.com/allez-chauffe/marcel/pkg/db"
	"github.com/allez-chauffe/marcel/pkg/db/internal/test_utils"
	"github.com/allez-chauffe/marcel/pkg/db/users"
)

func TestUsers(t *testing.T) {
	test_utils.DatabaseTest(t, map[string]test_utils.DatabaseTestFunc{
		"GetNotFound":   testGetNotFound,
		"InsertAndGet":  testInsertAndGet,
		"List":          testList,
		"Update":        testUpdate,
		"Delete":        testDelete,
		"UpsertAll":     testUpsertAll,
		"GetByLogin":    testGetByLogin,
		"Disconnect":    testDisconnect,
		"EnsureOneUser": testEnsureOneUser,
	})
}

func testGetNotFound(tx *db.Tx, t *testing.T) {
	user, err := tx.Users().Get("not found")
	if err != nil {
		t.Fatalf("Get failed: %s", err)
	}

	if user != nil {
		t.Fatal("Should return nil when it is not found")
	}
}

func testInsertAndGet(tx *db.Tx, t *testing.T) {
	user := users.New()
	if err := tx.Users().Insert(user); err != nil {
		t.Fatalf("Insert failed: %s", err)
	}

	saved, err := tx.Users().Get(user.ID)
	if err != nil {
		t.Fatalf("Get failed: %s", err)
	}

	if saved == nil {
		t.Fatal("Should find newly inserted user")
	}
}

func testList(tx *db.Tx, t *testing.T) {
	tx.Users().Insert(users.New())
	tx.Users().Insert(users.New())
	tx.Users().Insert(users.New())

	list, err := tx.Users().List()

	if err != nil {
		t.Fatalf("List failed: %s", err)
	}

	if len(list) != 3 {
		t.Fatalf("Expected a list of 3 users, %d found", len(list))
	}

	if list[0] == list[1] ||
		list[1] == list[2] ||
		list[0] == list[2] {
		t.Fatal("Each user should be different")
	}
}

func testUpdate(tx *db.Tx, t *testing.T) {
	user := users.New()
	tx.Users().Insert(user)

	user.DisplayName = "test"

	if err := tx.Users().Update(user); err != nil {
		t.Fatalf("Update failed: %s", err)
	}

	if user, _ = tx.Users().Get(user.ID); user.DisplayName != "test" {
		t.Fatal("Expect user display name to be 'test', '%' found", user.DisplayName)
	}
}

func testDelete(tx *db.Tx, t *testing.T) {
	user := users.New()
	tx.Users().Insert(user)

	if err := tx.Users().Delete(user.ID); err != nil {
		t.Fatalf("Delete failed: %s", err)
	}

	if exists, _ := tx.Users().Exists(user.ID); exists {
		t.Fatal("User should have been deleted")
	}
}

func testUpsertAll(tx *db.Tx, t *testing.T) {
	user := users.New()
	tx.Users().Insert(user)
	user.DisplayName = "test"

	err := tx.Users().UpsertAll([]users.User{
		*user,
		*users.New(),
		*users.New(),
	})

	if err != nil {
		t.Fatalf("UpsertAll failed: %s", err)
	}

	list, err := tx.Users().List()
	if err != nil {
		t.Fatalf("List failed: %s", err)
	}

	if len(list) != 3 {
		t.Fatalf("Expected 3 users, %d found", len(list))
	}

	for _, u := range list {
		if u.ID == user.ID && u.DisplayName != "test" {
			t.Fatalf("Expected user display name to be 'test', '%s' found", u.DisplayName)
		}
	}

	if list[0] == list[1] ||
		list[1] == list[2] ||
		list[0] == list[2] {
		t.Fatalf("All ids should be diferent")
	}
}

func testGetByLogin(tx *db.Tx, t *testing.T) {
	user := users.New()
	user.Login = "test@marcel.com"
	tx.Users().Insert(user)

	found, err := tx.Users().GetByLogin("test@marcel.com")
	if err != nil {
		t.Fatalf("GetByLogin failed: %s", err)
	}

	if found == nil {
		t.Fatal("Should find user by login")
	}
}

func testDisconnect(tx *db.Tx, t *testing.T) {
	user := users.New()
	originalTime := user.LastDisconnection
	tx.Users().Insert(user)

	err := tx.Users().Disconnect(user.ID)
	if err != nil {
		t.Fatalf("Disconnect failed: %s", err)
	}

	if saved, _ := tx.Users().Get(user.ID); saved.LastDisconnection == originalTime {
		t.Fatal("Should have updated the last disconnection timestamp")
	}
}

func testEnsureOneUser(tx *db.Tx, t *testing.T) {
	if err := tx.Users().EnsureOneUser(); err != nil {
		t.Fatalf("EnsureOneUser failed: %s", err)
	}

	if list, _ := tx.Users().List(); len(list) != 1 {
		t.Fatalf("Expected 1 user, %d found", len(list))
	}

	if err := tx.Users().EnsureOneUser(); err != nil {
		t.Fatalf("EnsureOneUser failed: %s", err)
	}

	list, _ := tx.Users().List()
	if len(list) != 1 {
		t.Fatalf("Expected 1 user, %d found", len(list))
	}

	if list[0].ID == "" {
		t.Fatalf("Expected a generated uuid, got empty string")
	}
}
