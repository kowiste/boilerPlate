print('===============================')
print('===       INIT SCRIPT       ===')
print('===============================')

conn = new Mongo('localhost:27017')

// Asset Database
db = conn.getDB('user')
db.asset.createIndex({ id: 1 }, { unique: true })

// Inserting Assets with fixed UUID-like string IDs
db.asset.insertMany([
  {
    id: '550e8400-e29b-41d4-a716-446655440000',
    parentID: '650e8400-e29b-41d4-a716-446655440000',
    description: 'First Asset',
  },
  {
    id: '660e8400-e29b-41d4-a716-446655440000',
    parentID: '550e8400-e29b-41d4-a716-446655440000',
    description: 'Second Asset',
  },
  {
    id: '770e8400-e29b-41d4-a716-446655440000',
    parentID: '550e8400-e29b-41d4-a716-446655440000',
    description: 'Third Asset',
  },
])

// User Database
db.user.createIndex(
  { id: 1 },
  { unique: true }
)

// Inserting Users with fixed UUID-like string IDs
db.user.insertMany([
  {
    id: '880e8400-e29b-41d4-a716-446655440000',
    name: 'John',
    lastName: 'Doe',
  },
  {
    id: '990e8400-e29b-41d4-a716-446655440000',
    name: 'Jane',
    lastName: 'Smith',
  },
  {
    id: 'aaa0e8400-e29b-41d4-a716-446655440000',
    name: 'Alice',
    lastName: 'Johnson',
  },
])

print('Initialization completed.')
