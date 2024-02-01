package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("e2fk3ns7qs81nda")
		if err != nil {
			return err
		}

		collection.Name = "build_cities"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("e2fk3ns7qs81nda")
		if err != nil {
			return err
		}

		collection.Name = "city_buildings"

		return dao.SaveCollection(collection)
	})
}
