package database_test

import (
	"database/sql"
	"testing"

	"github.com/cesar-marino/fc_ms_wallet/internal/database"
	"github.com/cesar-marino/fc_ms_wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDbTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *database.ClientDB
}

func (s *ClientDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients(id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATE)")
	s.clientDB = database.NewClientDB(db)
}

func (s *ClientDbTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDbTestSuite))
}

func (s *ClientDbTestSuite) TestSave() {
	client := &entity.Client{
		ID:    "1",
		Name:  "César Marino",
		Email: "cesar@teste.com",
	}

	err := s.clientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDbTestSuite) TestGet() {
	client, _ := entity.NewClient("César Marino", "cesar@test.com")
	s.clientDB.Save(client)

	clientDB, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}
