migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("e5q7sfgmigaonrb")

  collection.createRule = null

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("e5q7sfgmigaonrb")

  collection.createRule = ""

  return dao.saveCollection(collection)
})
