package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "rmoiu3xq54ynksf",
			"created": "2024-02-01 09:08:10.942Z",
			"updated": "2024-02-01 09:08:10.942Z",
			"name": "users",
			"type": "auth",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "jjkauh0w",
					"name": "field",
					"type": "file",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"mimeTypes": [
							"image/jpeg",
							"image/png",
							"image/svg+xml",
							"image/gif",
							"image/webp"
						],
						"thumbs": [],
						"maxSelect": 1,
						"maxSize": 5242880,
						"protected": false
					}
				}
			],
			"indexes": [],
			"listRule": "id = @request.auth.id",
			"viewRule": "id = @request.auth.id",
			"createRule": null,
			"updateRule": "id = @request.auth.id",
			"deleteRule": null,
			"options": {
				"allowEmailAuth": true,
				"allowOAuth2Auth": true,
				"allowUsernameAuth": false,
				"exceptEmailDomains": null,
				"manageRule": null,
				"minPasswordLength": 8,
				"onlyEmailDomains": null,
				"onlyVerified": false,
				"requireEmail": true
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("rmoiu3xq54ynksf")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
