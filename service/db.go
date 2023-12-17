package service

import (
	"fmt"
	"log"

	"github.com/hashicorp/go-memdb"
)

type DBService struct {
	db *memdb.MemDB
}

type Pack struct {
	Size int
}

func NewDBService() *DBService {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"pack": {
				Name: "pack",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Size"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Fatalf("could not create in-mem db")
	}

	service := &DBService{db: db}

	if err := service.initializePackSizes(); err != nil {
		log.Panicln("failed to insert pack size", err)
	}

	return service
}

func (db *DBService) AddPackSize(size int) error {
	txn := db.db.Txn(true)
	defer txn.Abort()

	found, err := txn.First("pack", "id", size)

	if err != nil {
		return fmt.Errorf("error with query: %v", err)
	}
	if found != nil {
		return fmt.Errorf("pack size %d already exists", size)
	}

	if err := txn.Insert("pack", &Pack{Size: size}); err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func (db *DBService) RemovePackSize(size int) error {
	txn := db.db.Txn(true)
	defer txn.Abort()

	packSizeCount := 0
	iterator, err := txn.Get("pack", "id")
	if err != nil {
		return err
	}

	for pack := iterator.Next(); pack != nil; pack = iterator.Next() {
		packSizeCount++
	}

	if packSizeCount <= 1 {
		return fmt.Errorf("cannot remove pack size need at least one size to pack items")
	}

	if _, err := txn.First("pack", "id", size); err != nil {
		return fmt.Errorf("pack size %d does not exist", size)
	}

	if err := txn.Delete("pack", Pack{Size: size}); err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func (db *DBService) ListAllPackSizes() ([]int, error) {
	txn := db.db.Txn(false)
	iterator, err := txn.Get("pack", "id")
	if err != nil {
		return nil, err
	}

	var sizes []int
	for size := iterator.Next(); size != nil; size = iterator.Next() {
		s := size.(*Pack)
		sizes = append(sizes, s.Size)
	}

	return sizes, nil
}

func (db *DBService) ResetPackSize() error {

	if err := db.clearPackSizes(); err != nil {
		return err
	}

	if err := db.initializePackSizes(); err != nil {
		return err
	}

	return nil
}

func (db *DBService) NewPackSizes(packSizes []int) error {

	if err := db.clearPackSizes(); err != nil {
		return err
	}

	for _, size := range packSizes {
		if err := db.AddPackSize(size); err != nil {
			return err
		}
	}
	return nil
}

func (db *DBService) clearPackSizes() error {
	txn := db.db.Txn(true)
	defer txn.Abort()

	iterator, err := txn.Get("pack", "id")
	if err != nil {
		return err
	}

	for size := iterator.Next(); size != nil; size = iterator.Next() {
		if err := txn.Delete("pack", size); err != nil {
			return err
		}
	}
	txn.Commit()
	return nil
}

func (db *DBService) initializePackSizes() error {
	initialPackSizes := []int{5000, 2000, 1000, 500, 250}
	for _, size := range initialPackSizes {
		if err := db.AddPackSize(size); err != nil {
			return err
		}
	}
	return nil
}
