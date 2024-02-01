package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("g19jljehartx1lw")
		if err != nil {
			return err
		}

		// add
		new_skill_trees_id := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "z7guwzht",
			"name": "skill_trees_id",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "btmvdnx0migokuu",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_skill_trees_id)
		collection.Schema.AddField(new_skill_trees_id)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("g19jljehartx1lw")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("z7guwzht")

		return dao.SaveCollection(collection)
	})
}
