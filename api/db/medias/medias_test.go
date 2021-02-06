package medias_test

import (
	"testing"

	"github.com/allez-chauffe/marcel/api/db"
	"github.com/allez-chauffe/marcel/api/db/internal/test_utils"
	"github.com/allez-chauffe/marcel/api/db/medias"
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
	t.Run("List", testList(runner))
	t.Run("Update", testUpdate(runner))
	t.Run("Delete", testDelete(runner))
	t.Run("Exists", testExists(runner))
	t.Run("UpsertAll", testUpsertAll(runner))
}

func testGetNotFound(runner func(func(tx *db.Tx) error) error) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				media, err := tx.Medias().Get(0)
				if err != nil {
					t.Fatalf("Get failed: %s", err)
				}

				if media != nil {
					t.Fatal("Should return nil for a media when it is not found")
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
				return nil
			})
		})
	}
}

func testList(runner func(func(tx *db.Tx) error) error) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
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
				return nil
			})
		})
	}
}

func testUpdate(runner func(func(tx *db.Tx) error) error) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				media := medias.New("test")
				tx.Medias().Insert(media)

				media.Name = "test"

				if err := tx.Medias().Update(media); err != nil {
					t.Fatalf("Update failed: %s", err)
				}

				if media, _ = tx.Medias().Get(media.ID); media.Name != "test" {
					t.Fatal("Expect media name to be 'test', '%' found", media.Name)
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
				media := medias.New("test")
				tx.Medias().Insert(media)

				if err := tx.Medias().Delete(media.ID); err != nil {
					t.Fatalf("Delete failed: %s", err)
				}

				if exists, _ := tx.Medias().Exists(media.ID); exists {
					t.Fatal("Media should have been deleted")
				}
				return nil
			})
		})
	}
}

func testExists(runner func(func(tx *db.Tx) error) error) func(t *testing.T) {
	return func(t *testing.T) {
		test_utils.DatabaseTest(func() {
			runner(func(tx *db.Tx) error {
				exists, err := tx.Medias().Exists(0)

				if err != nil {
					t.Fatalf("Exists failed: %s", err)
				}

				if exists {
					t.Fatal("Should not return true with an empty database")
				}

				tx.Medias().Insert(medias.New("test"))

				if exists, err = tx.Medias().Exists(0); err != nil {
					t.Fatalf("Exists failed: %s", err)
				}

				if !exists {
					t.Fatal("Media should have been created")
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
				return nil
			})
		})
	}
}
