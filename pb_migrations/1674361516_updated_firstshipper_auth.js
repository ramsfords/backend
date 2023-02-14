migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("rye5zosfxuhc1mj")

  collection.createRule = ""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("rye5zosfxuhc1mj")

  collection.createRule = null

  return dao.saveCollection(collection)
})
