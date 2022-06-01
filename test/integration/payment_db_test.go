package integration

import (
	mongoPayment "shopping-cart/pkg/payment/mongo"
	"testing"

	mongoUtils "shopping-cart/pkg/utils/mongo"

	"github.com/stretchr/testify/assert"
)

var (
	paymentConfig = mongoUtils.Config{
		Host:     "localhost",
		Port:     0,
		Username: "root",
		Password: "LoFiBeats",
		Database: "payment",
	}
)

func TestCreateUser(t *testing.T) {
	port, teardownSuite := setupSuite(t)
	paymentConfig.Port = port
	mongoClient := mongoPayment.Connect(&paymentConfig)
	defer teardownSuite(t, mongoClient)

	paymentConn := mongoPayment.Init(&mongoClient)

	userId, err := paymentConn.CreateUser()
	assert.True(t, err == nil)
	assert.True(t, len(userId) > 0)
}

func TestFindUser(t *testing.T) {
	port, teardownSuite := setupSuite(t)
	paymentConfig.Port = port
	mongoClient := mongoPayment.Connect(&paymentConfig)
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
	port, teardownSuite := setupSuite(t)
	paymentConfig.Port = port
	mongoClient := mongoPayment.Connect(&paymentConfig)
	defer teardownSuite(t, mongoClient)

	paymentConn := mongoPayment.Init(&mongoClient)
	userId, _ := paymentConn.CreateUser()

	amount := int64(69)

	errFunds := paymentConn.AddFunds(userId, amount)
	assert.True(t, errFunds == nil)

	userObject, _ := paymentConn.FindUser(userId)
	assert.Equal(t, amount, userObject.Credit)
}
