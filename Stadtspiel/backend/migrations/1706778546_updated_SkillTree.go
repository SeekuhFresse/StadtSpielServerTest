package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("btmvdnx0migokuu")
		if err != nil {
			return err
		}

		collection.Name = "skill_trees"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("btmvdnx0migokuu")
		if err != nil {
			return err
		}

		collection.Name = "SkillTree"

		return dao.SaveCollection(collection)
	})
}
