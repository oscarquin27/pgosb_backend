package role_domain

import (
	role "fdms/domain/entities/roles"

	"github.com/jackc/pgx/v5/pgxpool"
)
type RoleRepository interface {
	GetRole(id int64) (*role.Role, error)
	GetAll() ([]role.Role, error)
	Create(user *role.Role) (error)
	Update(user *role.Role) (error)
	Delete(id int64) (error)
}

type RoleImpl struct {
	db *pgxpool.Pool
}

func NewRoleService(db *pgxpool.Pool) RoleRepository {
	return &RoleImpl{
		db : db,
	}
}