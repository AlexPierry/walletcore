package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "x@x.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("John Doe 2", "y@y.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, float64(900), account1.Balance)
	assert.Equal(t, float64(1100), account2.Balance)
}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	client1, _ := NewClient("John Doe", "x@x.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("John Doe 2", "y@y.com")
	account2 := NewAccount(client2)

	account1.Credit(50)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100)
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Error(t, err, "Insuficient funds")
	assert.Equal(t, float64(50), account1.Balance)
	assert.Equal(t, float64(1000), account2.Balance)

}
