print('===============================')
print('===       INIT SCRIPT       ===')
print('===============================')

conn = new Mongo('localhost:27017')

//ACTIONS
db = conn.getDB('test')
db.Other.createIndex({ name: 1 }, { unique: true })
db.Other.insert([
  {
    field1: 2,
    name: 'Pablo',
  },
  {
    field1: 6,
    name: 'Nam',
  },
  {
    field1: 34,
    name: 'Pete',
  },
])
