package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionInfrastructureRepository struct {
	db *pgxpool.Pool
}

func NewMissionInfrastructureService(db *pgxpool.Pool) services.MissionInfrastructureService {
	return &MissionInfrastructureRepository{
		db: db,
	}
}

func (u *MissionInfrastructureRepository) Get(id int) ([]models.MissionInfrastructure, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `
	SELECT id, 
	service_id, 
	build_type, 
	build_occupation, 
	build_area, 
	build_access, 
	levels, 
	people, 
	goods_type, 
	build_roof, 
	build_floor, 
	build_room_type, 
	observations
FROM missions.infrastructure
 	where service_id = $1;`, id)

	if err != nil {
		return nil, err
	}

	infra, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionInfrastructure])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionInfrastructureNotFound
		}

		return nil, err
	}

	return infra, nil
}

func (u *MissionInfrastructureRepository) Create(infra *models.MissionInfrastructure) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `
	insert into missions.infrastructure
	(service_id, build_type, build_occupation, build_area, build_access, levels, people, goods_type, build_roof, build_floor, build_room_type, observations)
	values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);
;
`, infra.ServiceId,
   infra.BuildType,
   infra.BuildOccupation,
   infra.BuildArea,
   infra.BuildAccess,
   infra.Levels,
   infra.People,
   infra.GoodsType,
   infra.BuildRoof,
   infra.BuildFloor,
   infra.BuildRoomType,
   infra.Observations)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorMissionInfrastructureNotCreated
}

func (u *MissionInfrastructureRepository) Update(infra *models.MissionInfrastructure) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `
		UPDATE missions.infrastructure
		SET service_id=$1, 
		build_type=$2, 
		build_occupation=$3, 
		build_area=$4, 
		build_access=$5, 
		levels=$6, 
		people=$7, 
		goods_type=$8, 
		build_roof=$9, 
		build_floor=$10, 
		build_room_type=$11, 
		observations=$12
		where id = $13;
		`,
   		infra.ServiceId,
   		infra.BuildType,
   		infra.BuildOccupation,
   		infra.BuildArea,
   		infra.BuildAccess,
   		infra.Levels,
   		infra.People,
   		infra.GoodsType,
   		infra.BuildRoof,
   		infra.BuildFloor,
   		infra.BuildRoomType,
   		infra.Observations,
		infra.Id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorMissionInfrastructureNotUpdated
}

func (u *MissionInfrastructureRepository) Delete(id int) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from missions.infrastructure where id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorVehicleNotDeleted
}
