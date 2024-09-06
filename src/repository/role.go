package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RoleRepositoy struct {
	connPool *pgxpool.Pool
}

func NewRoleService(connPool *pgxpool.Pool) services.RoleService {
	return &RoleRepositoy{
		connPool: connPool,
	}
}

func (u *RoleRepositoy) Get(id int64) *results.ResultWithValue[*models.Role] {

	r := results.NewResultWithValue[*models.Role]("Get-Role", false, nil, nil).Failure()

	ctx := context.Background()

	conn, err := u.connPool.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx,
		"SELECT id, role_name, st_role, access_schema, created_at, updated_at FROM users.roles WHERE id = $1", id,
	)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	role, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Role])

	if err != nil {
		if err == pgx.ErrNoRows {
			return r.WithError(results.NewNotFoundError("no se encontro el role espeificado", err))
		}

		return r.WithError(
			results.NewUnknowError("error obteniendo unidad", err))
	}

	return r.Success().WithValue(&role)
}

func (u *RoleRepositoy) GetAll(params ...string) ([]models.Role, *results.GeneralError) {
	var roles []models.Role = make([]models.Role, 0)

	ctx := context.Background() // Or use a specific context

	conn, err := u.connPool.Acquire(ctx)

	if err != nil {
		return roles, results.
			NewUnknowError("no se pudo adquirir conexion", err)
	}

	defer conn.Release()

	query := "SELECT * FROM users.roles"

	rows, err := conn.Query(ctx, query)

	if err != nil {
		return roles, results.
			NewUnknowError("no se pudo ejecutar el query", err)
	}

	defer rows.Close()

	rolesValus, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Role])

	if err != nil {
		if err == pgx.ErrNoRows {
			return roles, results.NewNotFoundError("no encontraron registros", err)
		}

		return roles, results.
			NewUnknowError("no se pudo ejecutar el query", err)
	}

	return rolesValus, nil
}

func (u *RoleRepositoy) Create(role *models.Role) *results.ResultWithValue[*models.Role] {

	r := results.NewResultWithValue[*models.Role]("Create-Role", false, nil, nil).Failure()

	ctx := context.Background()

	conn, err := u.connPool.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	row := conn.QueryRow(ctx,
		"INSERT INTO users.roles (role_name, access_schema, st_role) VALUES ($1, $2, $3) RETURNING id",
		role.RoleName, role.AccessSchema, role.StRole,
	)

	err = row.Scan(&role.ID)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	return r.Success().WithValue(role)
}

func (u *RoleRepositoy) Update(role *models.Role) *results.ResultWithValue[*models.Role] {

	r := results.NewResultWithValue[*models.Role]("Update-Unit", false, nil, nil).Failure()

	ctx := context.Background()

	conn, err := u.connPool.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Exec(ctx,
		"UPDATE users.roles SET role_name = $1, access_schema = $2, st_role = $3, updated_at = $4 WHERE id = $5",
		role.RoleName, role.AccessSchema, role.StRole, time.Now().UTC(), role.ID,
	)

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

	return r.Success().WithValue(role)
}

func (u *RoleRepositoy) Delete(id int64) *results.Result {

	r := results.NewResult("Delete-Role", false, nil).Failure()

	ctx := context.Background()

	conn, err := u.connPool.Acquire(ctx)

	if err != nil {
		return r.FailureWithError(err)
	}

	defer conn.Release()

	rows, err := conn.Exec(ctx, "delete from users.roles where id = $1", id)

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

func (u *RoleRepositoy) GetSchema(id int64) *results.ResultWithValue[string] {

	r := results.NewResultWithValue[string]("GetSchema-Role", false, "", nil).Failure()

	ctx := context.Background()

	conn, err := u.connPool.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	row := conn.QueryRow(ctx, "select  access_schema from users.roles where id = $1", id)

	var schema string

	err = row.Scan(&schema)

	if err != nil {
		if err == pgx.ErrNoRows {
			if err == pgx.ErrNoRows {
				return r.WithError(results.NewNotFoundError("no encontro el schema", err))
			}
		}

		return r.WithError(
			results.NewUnknowError("error obteniendo registro", err))
	}

	return r.Success().WithValue(schema)
}
