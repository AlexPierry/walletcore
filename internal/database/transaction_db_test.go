package database

import (
	"database/sql"
	"testing"

	"github.com.br/alexpierry/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client1       *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	TransactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATE)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance INT, created_at DATE)")
	db.Exec("CREATE TABLE transactions (id VARCHAR(255), account_id_from VARCHAR(255), account_id_to VARCHAR(255), amount INT, created_at DATE)")

	client1, err := entity.NewClient("John Doe", "x@x.com")
	s.Nil(err)
	s.client1 = client1

	client2, err := entity.NewClient("John Doe 2", "y@y.com")
	s.Nil(err)
	s.client2 = client2

	accountFrom := entity.NewAccount(s.client1)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(s.client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.TransactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.TransactionDB.Create(transaction)
	s.Nil(err)
}
