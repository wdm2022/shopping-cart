stockDb = db.getSiblingDB('stock')

if (stockDb.system.users.find({user: 'stock'}).count() === 0) {
    stockDb.createUser({
        user: 'stock',
        pwd: 'password',
        roles: [{role: 'dbOwner', db: 'stock'}]
    });
}
