package mysql_test

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"go_docker/user/repository/mysql"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "name1").
		AddRow(2, "name1")

	mock.ExpectQuery("SELECT \\* FROM user WHERE id = \\?").
		WithArgs(1).
		WillReturnRows(rows)

	repo := mysql.NewMysqlUserRepository(db)
	user, err := repo.GetById(context.Background(), 1)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if user.ID != 1 {
		t.Fatalf("ID is not 1")
	}
	if user.Name != "name1" {
		t.Fatalf("Name is not name1")
	}
}