package role_domain

import (
	"context"
	role "fdms/domain/entities/roles"
	"time"

	"github.com/jackc/pgx/v5"
)


func (u *RoleImpl) GetRole(id int64) (*role.Role, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil,err
	}

	rows, err := conn.Query(ctx, "select id, role_name, st_role, access_schema, created_at, updated_at from users.roles where id = $1", id)

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[role.Role])
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, role.ErrorRoleNotFound
		}

		return nil, err
	}

	defer conn.Release()

	return &r,nil
}

func (u *RoleImpl) GetAll() ([]role.Role, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil,err
	}

	rows, err := conn.Query(ctx, "select id, role_name, st_role, access_schema, created_at, updated_at from users.roles")

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[role.Role])
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, role.ErrorRoleNotFound
		}

		return nil, err
	}

	defer conn.Release()

	return r,nil
}
func (u *RoleImpl) Create(r *role.Role) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "insert into users.roles (role_name, access_schema) values ($1, $2)", r.Role_name, r.Access_schema)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return role.ErrorRoleNotCreated
}

func (u *RoleImpl) Update(r *role.Role) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "update users.roles set role_name = $1, access_schema = $2, st_role = $3,  updated_at = $4 where id = $5", r.Role_name, r.Access_schema, r.St_role, time.Now().UTC(), r.Id)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return role.ErrorRoleNotUpdated
}

func (u *RoleImpl) Delete(id int64) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from users.roles where id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return role.ErrorRoleNotDeleted
}
