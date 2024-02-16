package mysql

import (
	"context"
	"database/sql"
	"go_docker/domain"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) domain.UserRepository {
	return &mysqlUserRepository{
		db: db,
	}
}

func (m * mysqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.User, error) {
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		res = append(res, user)
	}
	return res, nil
}

func (m *mysqlUserRepository) GetById(ctx context.Context, id int64) (res domain.User, err error) {
	rows, err := m.fetch(ctx, "SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		return domain.User{}, err
	}

	if len(rows) > 0 {
		res = rows[0]
	}else{
		return res, domain.ErrUserNotFound
	}

	return
}