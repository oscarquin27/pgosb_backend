package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func notFoundUnitError(r *results.ResultWithValue[*models.Unit], err error) *results.ResultWithValue[*models.Unit] {
	return r.WithError(
		results.NewNotFoundError("no se encontro la unidad", err))
}

const selectQuery = ` SELECT 
           id, 
           plate, 
           station, 
           unit_type, 
           make, 
           unit_condition, 
           vehicle_serial, 
		   motor_serial, 
		   capacity, 
		   details, 
		   fuel_type, 
		   water_capacity, 
		   observations,
		   hurt_capacity, 
		   doors, 
		   performance, 
		   load_capacity,
		   model, 
		   alias, 
		   color,
		   year, 
		   purpose, 
		   init_kilometer
	FROM vehicles.unit
	WHERE id = $1 `

const updateQuery = `
	UPDATE vehicles.unit
	 SET 
	
	  plate = $1, 
	  station = $2, 
	  unit_type = $3, 
	  make = $4, 
	  unit_condition = $5,
	  vehicle_serial = $6, 
	  motor_serial = $7, 
	  capacity = $8, 
	  details = $9, 
	  fuel_type = $10, 
	  water_capacity = $11, 
	  observations = $12, 
	  hurt_capacity = $13,
	  doors = $14, 
	  performance = $15, 
	  load_capacity = $16, 
	  model = $17, 
	  alias = $18,
	  color = $19, 
	  year = $20, 
	  purpose = $21, 
	  init_kilometer = $22
	
	  WHERE id = $23`

const selectAllQuery = ` SELECT 
               id, 
			   plate, 
			   station, 
			   unit_type, 
			   make, 
			   unit_condition, 
			   vehicle_serial, 
               motor_serial, 
			   capacity, 
			   details, 
			   fuel_type, 
			   water_capacity, 
			   observations,
               hurt_capacity, 
			   doors, 
			   performance, 
			   load_capacity, 
			   model, 
			   alias, 
			   color,
               year, 
			   purpose, 
			   init_kilometer
        FROM vehicles.unit ORDER BY id ASC`

const insertQuery = `
	INSERT INTO vehicles.unit (
		plate, 
		station, 
		unit_type, 
		make, 
		unit_condition, 
		vehicle_serial, 
		motor_serial, 
		capacity, 
		details, 
		fuel_type, 
		water_capacity, 
		observations, 
		hurt_capacity, 
		doors, 
		performance, 
		load_capacity, 
		model, 
		alias, 
		color, 
		year, 
		purpose, 
		init_kilometer
	) VALUES (
	 $1, 
	 $2, 
	 $3, 
	 $4, 
	 $5, 
	 $6, 
	 $7, 
	 $8, 
	 $9, 
	 $10, 
	 $11, 
	 $12, 
	 $13, 
	 $14, 
	 $15, 
	 $16, 
	 $17, 
	 $18, 
	 $19, 
	 $20, 
	 $21, 
	 $22
	 )
	 RETURNING id
	 `

type UnitRepository struct {
	db *pgxpool.Pool
}

func NewUnityService(db *pgxpool.Pool) services.UnitService {
	return &UnitRepository{
		db: db,
	}
}

func (u *UnitRepository) Get(id int64) *results.ResultWithValue[*models.Unit] {

	r := results.NewResultWithValue[*models.Unit]("Get-Unit", false, nil, nil).Failure()

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

	unity, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Unit])

	if err != nil {

		if err == pgx.ErrNoRows {
			return notFoundUnitError(r, err)
		}

		return r.WithError(
			results.NewUnknowError("error obteniendo unidad", err))
	}

	return r.Success().WithValue(&unity)
}

func (u *UnitRepository) GetAll(parmas ...string) ([]models.Unit, *results.GeneralError) {
	var units []models.Unit = make([]models.Unit, 0)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return units, results.
			NewUnknowError("no se pudo adquirir conexion", err)
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, selectAllQuery)

	if err != nil {
		return units, results.
			NewUnknowError("no se pudo ejecutar el query", err)
	}

	defer rows.Close()

	unity, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Unit])

	if err != nil {
		if err == pgx.ErrNoRows {
			return units, results.NewNotFoundError("no encontraron registros", err)
		}

		return units, results.
			NewUnknowError("no se pudo ejecutar el query", err)
	}

	return unity, nil
}

func (u *UnitRepository) Create(unit *models.Unit) *results.ResultWithValue[*models.Unit] {

	r := results.NewResultWithValue[*models.Unit]("Create-Unit", false, nil, nil).Failure()

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	row := conn.QueryRow(ctx, insertQuery,
		unit.Plate,
		unit.Station,
		unit.Unit_type,
		unit.Make,
		unit.Unit_condition,
		unit.Vehicle_serial,
		unit.Motor_serial,
		unit.Capacity,
		unit.Details,
		unit.Fuel_type,
		unit.Water_capacity,
		unit.Observations,
		unit.Hurt_capacity,
		unit.Doors,
		unit.Performance,
		unit.Load_capacity,
		unit.Model,
		unit.Alias,
		unit.Color,
		unit.Year,
		unit.Purpose,
		unit.Init_kilometer,
	)

	err = row.Scan(&unit.Id)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	return r.Success().WithValue(unit)
}

func (u *UnitRepository) Update(unit *models.Unit) *results.ResultWithValue[*models.Unit] {
	r := results.NewResultWithValue[*models.Unit]("Update-Unit", false, nil, nil).Failure()

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Exec(ctx, updateQuery,
		unit.Plate,
		unit.Station,
		unit.Unit_type,
		unit.Make,
		unit.Unit_condition,
		unit.Vehicle_serial,

		unit.Motor_serial,
		unit.Capacity,
		unit.Details,
		unit.Fuel_type,
		unit.Water_capacity,

		unit.Observations,
		unit.Hurt_capacity,
		unit.Doors,
		unit.Performance,
		unit.Load_capacity,

		unit.Model,
		unit.Alias,
		unit.Color,
		unit.Year,
		unit.Purpose,
		unit.Init_kilometer,
		unit.Id,
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

	return r.Success().WithValue(unit)
}

func (u *UnitRepository) Delete(id int64) *results.Result {

	r := results.NewResult("Delete-Unit", false, nil).Failure()

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return r.FailureWithError(err)
	}

	defer conn.Release()

	rows, err := conn.Exec(ctx, "delete from vehicles.unit where id = $1", id)

	if err != nil {
		return r.FailureWithError(err)
	}

	if rows.RowsAffected() == 0 {
		return r.WithError(results.NewNotFoundError("no se consiguio unidad", err))
	}

	if rows.RowsAffected() > 1 {
		return r.WithError(results.NewUnknowError("se borraron multiples registros", err))
	}

	return r.Success()
}

func (u *UnitRepository) GetAllSimple() *results.ResultWithValue[[]models.UnitSimple] {

	r := results.NewResultWithValue[[]models.UnitSimple]("Get-All-Simple", false, make([]models.UnitSimple, 0), nil).
		Failure()

	allUnist, err := u.GetAll()

	if err != nil {
		return r.WithError(err)
	}

	var unitSimples []models.UnitSimple = make([]models.UnitSimple, 0)

	for _, unit := range allUnist {

		unitSimple := models.UnitSimple{}
		unitSimple = *unitSimple.UnitSimpleFromUnit(&unit)
		unitSimples = append(unitSimples, unitSimple)
	}

	return r.Success().WithValue(unitSimples)
}
