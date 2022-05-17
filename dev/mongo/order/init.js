orderDb = db.getSiblingDB('order')

if (orderDb.system.users.find({user: 'order'}).count() === 0) {
    orderDb.createUser({
        user: 'order',
        pwd: 'password',
        roles: [{role: 'dbOwner', db: 'order'}]
    });
}
