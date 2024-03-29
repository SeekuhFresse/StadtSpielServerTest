package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("g19jljehartx1lw")
		if err != nil {
			return err
		}

		collection.Name = "skills"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("g19jljehartx1lw")
		if err != nil {
			return err
		}

		collection.Name = "Skill"

		return dao.SaveCollection(collection)
	})
}
