paymentDb = db.getSiblingDB('payment')

if (paymentDb.system.users.find({user: 'payment'}).count() === 0) {
    paymentDb.createUser({
        user: 'payment',
        pwd: 'password',
        roles: [{role: 'dbOwner', db: 'payment'}]
    });
}
