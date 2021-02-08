package medias_test

import (
	"fmt"
	"testing"

	"github.com/allez-chauffe/marcel/api/db"
	"github.com/allez-chauffe/marcel/api/db/internal/test_utils"
	"github.com/allez-chauffe/marcel/api/db/medias"
)

func TestMedias(t *testing.T) {
	test_utils.DatabaseTest(t, map[string]test_utils.DatabaseTestFunc{
		"GetNotFound":  testGetNotFound,
		"InsertAndGet": testInsertAndGet,
		"List":         testList,
		"Update":       testUpdate,
		"Delete":       testDelete,
		"Exists":       testExists,
		"UpsertAll":    testUpsertAll,
	})
}

func testGetNotFound(tx *db.Tx, t *testing.T) {
	media, err := tx.Medias().Get(0)
	if err != nil {
		t.Fatalf("Get failed: %s", err)
	}

	if media != nil {
		t.Fatal("Should return nil for a media when it is not found")
	}
}

func testInsertAndGet(tx *db.Tx, t *testing.T) {
	media := medias.New("test")
	if err := tx.Medias().Insert(media); err != nil {
		t.Fatalf("Insert failed: %s", err)
	}

	saved, err := tx.Medias().Get(media.ID)
	if err != nil {
		t.Fatalf("Get failed: %s", err)
	}

	if saved == nil {
		t.Fatal("Should find newly inserted media")
	}
}

func testList(tx *db.Tx, t *testing.T) {
	tx.Medias().Insert(medias.New("test"))
	tx.Medias().Insert(medias.New("test"))
	tx.Medias().Insert(medias.New("test"))

	list, err := tx.Medias().List()

	if err != nil {
		t.Fatalf("List failed: %s", err)
	}

	if len(list) != 3 {
		t.Fatalf("Expected a list of 3 medias, %d found", len(list))
	}

	if list[0].ID == list[1].ID ||
		list[1].ID == list[2].ID ||
		list[0].ID == list[2].ID {
		t.Fatal("Each medias should be different")
	}
}

func testUpdate(tx *db.Tx, t *testing.T) {
	media := medias.New("test")
	tx.Medias().Insert(media)

	media.Name = "test"

	if err := tx.Medias().Update(media); err != nil {
		t.Fatalf("Update failed: %s", err)
	}

	if media, _ = tx.Medias().Get(media.ID); media.Name != "test" {
		t.Fatal("Expect media name to be 'test', '%' found", media.Name)
	}
}

func testDelete(tx *db.Tx, t *testing.T) {
	media := medias.New("test")
	tx.Medias().Insert(media)

	if err := tx.Medias().Delete(media.ID); err != nil {
		t.Fatalf("Delete failed: %s", err)
	}

	if exists, _ := tx.Medias().Exists(media.ID); exists {
		t.Fatal("Media should have been deleted")
	}
}

func testExists(tx *db.Tx, t *testing.T) {
	exists, err := tx.Medias().Exists(1)

	if err != nil {
		t.Fatalf("Exists failed: %s", err)
	}

	if exists {
		t.Fatal("Should not return true with an empty database")
	}

	if err := tx.Medias().Insert(medias.New("test")); err != nil {
		t.Fatalf("Insertion failed: %s", err)
	}

	if exists, err = tx.Medias().Exists(1); err != nil {
		t.Fatalf("Exists failed: %s", err)
	}

	fmt.Println(exists)
	if !exists {
		t.Fatal("Media should have been created")
	}
}

func testUpsertAll(tx *db.Tx, t *testing.T) {
	media := medias.New("test")
	tx.Medias().Insert(media)
	media.Name = "test"

	err := tx.Medias().UpsertAll([]medias.Media{
		*media,
		*medias.New("test"),
		*medias.New("test"),
	})

	if err != nil {
		t.Fatalf("UpsertAll failed: %s", err)
	}

	list, err := tx.Medias().List()
	if err != nil {
		t.Fatalf("List failed: %s", err)
	}

	if len(list) != 3 {
		t.Fatalf("Expected 3 medias, %d found", len(list))
	}

	if list[0].Name != "test" {
		t.Fatalf("Expected media name to be 'test', '%s' found", list[0].Name)
	}

	if list[0].ID == list[1].ID ||
		list[1].ID == list[2].ID ||
		list[0].ID == list[2].ID {
		t.Fatalf("All ids should be diferent")
	}
}
