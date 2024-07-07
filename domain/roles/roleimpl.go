package role_domain

import (
	"context"
	role "fdms/domain/entities/roles"
	role_entity "fdms/domain/entities/roles"
	"time"

	"github.com/jackc/pgx/v5"
)

func (u *RoleImpl) GetRole(id int64) (*role.Role, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx,
		"SELECT id, role_name, st_role, access_schema, created_at, updated_at FROM users.roles WHERE id = $1", id,
	)
	if err != nil {
		return nil, err
	}

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[role.Role])
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, role.ErrorRoleNotFound
		}
		return nil, err
	}

	return &r, nil
}

func (u *RoleImpl) GetRoleSchema(id int64) (*string, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}
	defer conn.Release()

	row := conn.QueryRow(ctx, "select  access_schema from users.roles where id = $1", id)

	var schema string

	err = row.Scan(&schema)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, role.ErrorRoleNotFound
		}

		return nil, err
	}

	defer conn.Release()

	return &schema, nil
}

func (u *RoleImpl) GetAll() ([]role.Role, error) {
	ctx := context.Background() // Or use a specific context

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := "SELECT * FROM users.roles"

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []role.Role
	for rows.Next() {
		var role role.Role
		err := rows.Scan(&role.ID, &role.RoleName, &role.StRole, &role.AccessSchema, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (u *RoleImpl) Create(role *role.Role) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx,
		"INSERT INTO users.roles (role_name, access_schema, st_role) VALUES ($1, $2, $3)",
		role.RoleName, role.AccessSchema, role.StRole,
	)
	if err != nil {
		return role_entity.ErrorRoleNotCreated
	}

	return nil
}

func (u *RoleImpl) Update(role *role.Role) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx,
		"UPDATE users.roles SET role_name = $1, access_schema = $2, st_role = $3, updated_at = $4 WHERE id = $5",
		role.RoleName, role.AccessSchema, role.StRole, time.Now().UTC(), role.ID,
	)
	if err != nil {
		return err
	}
	if rows.RowsAffected() == 0 {
		return role_entity.ErrorRoleNotUpdated
	}

	return nil
}

func (u *RoleImpl) Delete(id int64) error {
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
