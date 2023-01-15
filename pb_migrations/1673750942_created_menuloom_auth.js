migrate((db) => {
  const collection = new Collection({
    "id": "e5q7sfgmigaonrb",
    "created": "2023-01-15 02:49:02.837Z",
    "updated": "2023-01-15 02:49:02.837Z",
    "name": "menuloom_auth",
    "type": "auth",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "p5mwtdj0",
        "name": "avatar",
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
        "id": "dquy4fuq",
        "name": "name",
        "type": "text",
        "required": false,
        "unique": false,
        "options": {
          "min": 2,
          "max": 100,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "hsguvloe",
        "name": "origin",
        "type": "url",
        "required": false,
        "unique": false,
        "options": {
          "exceptDomains": [],
          "onlyDomains": []
        }
      },
      {
        "system": false,
        "id": "z7joypbg",
        "name": "type",
        "type": "text",
        "required": false,
        "unique": false,
        "options": {
          "min": null,
          "max": 100,
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
  const collection = dao.findCollectionByNameOrId("e5q7sfgmigaonrb");

  return dao.deleteCollection(collection);
})
