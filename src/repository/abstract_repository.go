package repository

import (
	"context"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// type AbstractRepoType interface {
// 	SetId(id int64)
// }

type AbstractRepository[T any] struct {
	db *pgxpool.Pool
}

func NewAbstractRepository[T any](connPool *pgxpool.Pool) AbstractRepository[T] {
	return AbstractRepository[T]{
		db: connPool,
	}
}

func (u *AbstractRepository[T]) Get(id int64, selectQuery string) *results.ResultWithValue[T] {

	r := results.NewResultWithZeroValue[T]("Get", false, nil).Failure()

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, selectQuery, id)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	register, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[T])

	if err != nil {

		if err == pgx.ErrNoRows {
			return r.WithError(results.NewNotFoundError("no se encontro registro", err))
		}

		return r.WithError(
			results.NewUnknowError("error obteniendo unidad", err))
	}

	return r.Success().WithValue(register)
}

func (u *AbstractRepository[T]) GetAll(query string) ([]T, *results.GeneralError) {
	var registersDefault []T = make([]T, 0)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return registersDefault, results.
			NewUnknowError("no se pudo adquirir conexion", err)
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, query)

	if err != nil {
		return registersDefault, results.
			NewUnknowError("no se pudo ejecutar el query", err)
	}

	defer rows.Close()

	registers, err := pgx.CollectRows(rows, pgx.RowToStructByName[T])

	if err != nil {
		if err == pgx.ErrNoRows {
			return registersDefault, results.NewNotFoundError("no encontraron registros", err)
		}

		return registersDefault, results.
			NewUnknowError("no se pudo ejecutar el query", err)
	}

	return registers, nil
}

func (u *AbstractRepository[T]) Create(register T, insertQuery string, args pgx.NamedArgs, SetId func(int64)) *results.ResultWithValue[T] {

	r := results.NewResultWithZeroValue[T]("Create-Unit", false, nil).Failure()

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	row := conn.QueryRow(ctx, insertQuery, args)

	var id int64 = 0

	err = row.Scan(&id)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	SetId(id)

	return r.Success().WithValue(register)
}

func (u *AbstractRepository[T]) Update(register T, updateQuery string, args pgx.NamedArgs) *results.ResultWithValue[T] {
	r := results.NewResultWithZeroValue[T]("Update-Register", false, nil).Failure()

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Exec(ctx, updateQuery, args)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	if rows.RowsAffected() == 0 {
		return r.WithError(
			results.NewNotFoundError("no se consiguio registro", err))
	}

	if rows.RowsAffected() > 1 {
		return r.WithError(
			results.NewUnknowError("se afecto mas de un registro", err))
	}

	return r.Success().WithValue(register)
}

func (u *AbstractRepository[T]) Delete(id int64, deleteInsert string) *results.Result {

	r := results.NewResult("Delete-Register", false, nil).Failure()

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return r.FailureWithError(err)
	}

	defer conn.Release()

	rows, err := conn.Exec(ctx, deleteInsert, id)

	if err != nil {
		return r.FailureWithError(err)
	}

	if rows.RowsAffected() == 0 {
		return r.WithError(results.NewNotFoundError("no se consiguio registro", err))
	}

	if rows.RowsAffected() > 1 {
		return r.WithError(results.NewUnknowError("se borraron multiples registros", err))
	}

	return r.Success()
}
