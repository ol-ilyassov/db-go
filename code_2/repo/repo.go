package repo

import (
	"context"
	"database/sql"
	"dbgo/code_2/model"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	Conn *pgxpool.Pool
	QB   squirrel.StatementBuilderType
}

func New(conn *pgxpool.Pool) *Repo {
	return &Repo{
		Conn: conn,
		QB:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *Repo) Get(ctx context.Context, id int) (*model.Book, error) {
	query := `
		SELECT id, title, price
		FROM books
		WHERE id = $1
	`

	var b model.Book

	err := r.Conn.QueryRow(ctx, query, id).
		Scan(&b.ID, &b.Title, &b.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrRecordNotFound
		}
		return nil, err
	}

	return &b, nil
}

func (r *Repo) Update(ctx context.Context, obj model.Book) error {
	query := `
		UPDATE books
		SET title = $2, price = $3
		WHERE id = $1
	`
	args := []any{obj.Title, obj.Price, obj.ID}

	_, err := r.Conn.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) UpdateV2(ctx context.Context, id int, obj *model.Edit) error {
	qBuilder := r.QB.Update("books")

	if obj.Price != nil {
		qBuilder = qBuilder.Set("price", *obj.Price)
	}
	if obj.Title != nil {
		qBuilder = qBuilder.Set("title", *obj.Title)
	}

	query, args, err := qBuilder.Where("id = $1", id).ToSql()
	if err != nil {
		return err
	}

	_, err = r.Conn.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
