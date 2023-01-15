migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("e5q7sfgmigaonrb")

  collection.listRule = "id = @request.auth.id"
  collection.viewRule = "id = @request.auth.id"
  collection.updateRule = "id = @request.auth.id"
  collection.deleteRule = "id = @request.auth.id"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("e5q7sfgmigaonrb")

  collection.listRule = null
  collection.viewRule = null
  collection.updateRule = null
  collection.deleteRule = null

  return dao.saveCollection(collection)
})
