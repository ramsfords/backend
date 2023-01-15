migrate((db) => {
  const collection = new Collection({
    "id": "rye5zosfxuhc1mj",
    "created": "2023-01-15 02:50:23.427Z",
    "updated": "2023-01-15 02:50:23.427Z",
    "name": "firstshipper_auth",
    "type": "auth",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "66os2vfg",
        "name": "type",
        "type": "text",
        "required": false,
        "unique": false,
        "options": {
          "min": null,
          "max": 100,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "xnvr6jya",
        "name": "origin",
        "type": "url",
        "required": false,
        "unique": false,
        "options": {
          "exceptDomains": null,
          "onlyDomains": null
        }
      },
      {
        "system": false,
        "id": "w7fam5ig",
        "name": "avatar",
        "type": "text",
        "required": false,
        "unique": false,
        "options": {
          "min": null,
          "max": 300,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "xtzf5nog",
        "name": "name",
        "type": "text",
        "required": false,
        "unique": false,
        "options": {
          "min": 2,
          "max": 200,
          "pattern": ""
        }
      }
    ],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {
      "allowEmailAuth": true,
      "allowOAuth2Auth": true,
      "allowUsernameAuth": true,
      "exceptEmailDomains": null,
      "manageRule": null,
      "minPasswordLength": 8,
      "onlyEmailDomains": null,
      "requireEmail": false
    }
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("rye5zosfxuhc1mj");

  return dao.deleteCollection(collection);
})
