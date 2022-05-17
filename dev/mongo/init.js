orderDb = db.getSiblingDB('order')
paymentDb = db.getSiblingDB('payment')
stockDb = db.getSiblingDB('stock')

if (orderDb.system.users.find({user: 'order'}).count() === 0) {
    orderDb.createUser({
        user: 'order',
        pwd: 'password',
        roles: [{role: 'dbOwner', db: 'order'}]
    });
}

if (paymentDb.system.users.find({user: 'payment'}).count() === 0) {
    paymentDb.createUser({
        user: 'payment',
        pwd: 'password',
        roles: [{role: 'dbOwner', db: 'payment'}]
    });
}

if (stockDb.system.users.find({user: 'stock'}).count() === 0) {
    stockDb.createUser({
        user: 'stock',
        pwd: 'password',
        roles: [{role: 'dbOwner', db: 'stock'}]
    });
}
