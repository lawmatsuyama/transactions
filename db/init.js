db = db.getSiblingDB('account')
db.createCollection('transaction');
db.createCollection('user');
db.user.insertOne({
    "_id": "52814c2d-657b-4e7b-be5c-9f28e59253f8",
    "name": "Jose",
    "active": true
})

db.user.insertOne({
    "_id": "74211f47-9c6f-4648-88a3-cb7d0614b5fe",
    "name": "Lawrence",
    "active": true
})

db.user.insertOne({
    "_id": "355daea3-bfdc-41d5-8ecf-c9bcd21f4dbf",
    "name": "João",
    "active": false
})