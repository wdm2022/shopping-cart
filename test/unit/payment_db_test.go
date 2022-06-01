package unit

import (
	mongoPayment "shopping-cart/pkg/payment/mongo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	teardownSuite := setupSuite(t)
	mongoClient := mongoPayment.Connect(&mongoTestConfigPayment)
	defer teardownSuite(t, mongoClient)

	paymentConn := mongoPayment.Init(&mongoClient)

	userId, err := paymentConn.CreateUser()
	assert.True(t, err == nil)
	assert.True(t, len(userId) > 0)
}

func TestFindUser(t *testing.T) {
	teardownSuite := setupSuite(t)
	mongoClient := mongoPayment.Connect(&mongoTestConfigPayment)
	defer teardownSuite(t, mongoClient)

	paymentConn := mongoPayment.Init(&mongoClient)
	userId, err := paymentConn.CreateUser()
	assert.True(t, err == nil)
	assert.True(t, len(userId) > 0)

	userObject, err := paymentConn.FindUser(userId)
	assert.True(t, err == nil)
	assert.True(t, len(userObject.UserId) > 0)
	assert.Equal(t, int64(0), userObject.Credit)
}

func TestAddCredit(t *testing.T) {
	teardownSuite := setupSuite(t)
	mongoClient := mongoPayment.Connect(&mongoTestConfigPayment)
	defer teardownSuite(t, mongoClient)

	paymentConn := mongoPayment.Init(&mongoClient)
	userId, _ := paymentConn.CreateUser()

	amount := int64(69)

	errFunds := paymentConn.AddFunds(userId, amount)
	assert.True(t, errFunds == nil)

	userObject, _ := paymentConn.FindUser(userId)
	assert.Equal(t, amount, userObject.Credit)
}
