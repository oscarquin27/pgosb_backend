package location_domain

import (
	"context"
	city "fdms/domain/entities/cities"
	municipality "fdms/domain/entities/municipalities"
	parish "fdms/domain/entities/parish"
	state "fdms/domain/entities/states"
	station "fdms/domain/entities/stations"
	"fdms/utils"

	"github.com/jackc/pgx/v5"
)

func (u *LocationsImpl) GetState(id int64) (*state.State, error) {
	
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select state_id, name, coordinates from locations.states where state_id = $1", id)

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[state.State])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, state.ErrorStateFound
		}

		return nil, err
	}

	return &r, nil
}

func (u *LocationsImpl) GetAllStates()([]state.State, error){
	
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select state_id, name, coordinates from locations.states")

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[state.State])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, state.ErrorStateFound
		}

		return nil, err
	}

	return r, nil
}

func (u *LocationsImpl) CreateState(r *state.State) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "insert into locations.state (name, coordinates) values ($1, $2)", r.Coordinates, r.Name)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}


	return state.ErrorStateNotCreated
}

func (u *LocationsImpl) UpdateState(r *state.State) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "update locations.state set name = $1, coordinates = $2 where state_id = $3", r.Name, r.Coordinates, r.Id)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return state.ErrorStateNotUpdated
}

func (u *LocationsImpl) DeleteState(id int64) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from locations.state where state_id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return state.ErrorStateNotDeleted
}

func (u *LocationsImpl) GetCity(id int64) (*city.City, error){
	
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select city_id, state_id, name, area_code, zip_code, coordinates from locations.cities where city_id = $1", id)

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[city.City])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, city.ErrorCityFound
		}

		return nil, err
	}

	return &r, nil	
}

func (u *LocationsImpl) GetAllCity() ([]city.City, error){
	
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select city_id, state_id, name, area_code, zip_code, coordinates from locations.cities")

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[city.City])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, city.ErrorCityFound
		}

		return nil, err
	}

	return r, nil	
}

func (u *LocationsImpl) CreateCity(r *city.City) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "insert into locations.state (state_id, name, area_code, zip_code, coordinates) values ($1, $2, $3, $4)", r.State_Id, r.Name, r.Area_Code, r.Zip_Code, r.Coordinates)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return city.ErrorCityNotCreated
}

func (u *LocationsImpl) UpdateCity(r *city.City) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "update locations.cities set name = $1, area_code = $2, zip_code = $3, coordinates = $4 where city_id = $5", r.Name, r.Area_Code, r.Zip_Code, r.Coordinates, r.Id)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return city.ErrorCityNotUpdated
}

func (u *LocationsImpl) DeleteCity(id int64) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from locations.cities where city_id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return city.ErrorCityNotDeleted
}

func (u *LocationsImpl) GetMunicipality(id int64) (*municipality.Municipality, error){
	
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select municipality_id, state_id, name, coordinates from locations.municipalities where municipality_id = $1", id)

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[municipality.Municipality])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, municipality.ErrorMunicipalityFound
		}

		return nil, err
	}

	return &r, nil	
}

func (u *LocationsImpl) GetAllMunicipality() ([]municipality.Municipality, error){
	
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select municipality_id, state_id, name, coordinates from locations.municipalities")

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[municipality.Municipality])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, municipality.ErrorMunicipalityFound
		}

		return nil, err
	}

	return r, nil	
}

func (u *LocationsImpl) CreateMunicipality(r *municipality.Municipality) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "insert into locations.municipalities (state_id, name, coordinates) values ($1, $2, $3, $4)", r.State_Id, r.Name, r.Coordinates)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return city.ErrorCityNotCreated
}

func (u *LocationsImpl) UpdateMunicipality(r *municipality.Municipality) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "update locations.mnunicipalities set state_id = $1, name = $2, coordinates = $3 where municipality_id = $4", r.Name, r.State_Id, r.Coordinates, r.Id)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return municipality.ErrorMunicipalityNotCreated
}

func (u *LocationsImpl) DeleteMunicipality(id int64) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from locations.municipalities where municipality_id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return municipality.ErrorMunicipalityNotDeleted
}

func (u *LocationsImpl) GetParish(id int64) (*parish.Parish, error){
	
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select parish_id, state_id, municipality_id, name, coordinates from locations.parish where parish_id = $1", id)

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[parish.Parish])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, parish.ErrorParishFound
		}

		return nil, err
	}

	return &r, nil	
}

func (u *LocationsImpl) GetAllParish() ([]parish.Parish, error){
	
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select parish_id, state_id, municipality_id, name, coordinates from locations.parish")

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[parish.Parish])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, parish.ErrorParishFound
		}

		return nil, err
	}

	return r, nil	
}

func (u *LocationsImpl) CreateParish(r *parish.Parish) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "insert into locations.parish (state_id, municipality_id, name, coordinates) values ($1, $2, $3, $4)", r.State_Id, r.Municipality_Id, r.Name)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return city.ErrorCityNotCreated
}

func (u *LocationsImpl) UpdateParish(r *parish.Parish) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "update locations.parish set state_id = $1, municipality_id = $2, name = $3, coordinates = $4 where parish_id = $4", r.State_Id, r.Municipality_Id, r.Name, r.Coordinates, r.Id)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return parish.ErrorParishNotUpdated
}

func (u *LocationsImpl) DeleteParish(id int64) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from locations.parish where parish_id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return parish.ErrorParishNotDeleted
}

func (u *LocationsImpl) GetStation(id int64) (*station.Station, error){
	
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select station_id, municipality_id, name, coordinates, description, code, abbreviation, phones, state_id, parish_id, sector, community, street, address from locations.fire_stations where station_id = $1", id)

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[station.Station])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, station.ErrorStationFound
		}

		return nil, err
	}

	return &r, nil	
}

func (u *LocationsImpl) GetAllStations() ([]station.Station, error){
	
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select station_id, municipality_id, name, coordinates, description, code, abbreviation, phones, state_id, parish_id, sector, community, street, address from locations.fire_stations")

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[station.Station])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, station.ErrorStationFound
		}

		return nil, err
	}

	return r, nil	
}

func (u *LocationsImpl) CreateStation(r *station.Station) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `insert into locations.fire_station 
	(municipality_id, name, coordinates, description, code, abbreviation, phones, state_id, parish_id, sector, community, street, address) 
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`, r.Municipality_id, r.Name, r.Coordinates, r.Description, r.Code, r.Abbreviation, r.Phones, r.State_id, r.Parish_id, r.Sector, r.Community, r.Street, r.Address)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return station.ErrorStationFound
}

func (u *LocationsImpl) UpdateStation(r *station.Station) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "update locations.fire_station set municipality_id = $1, name = $2, coordinates = $3, description = $4, code = $5, abbreviation = $6, phones = $7, state_id = $8, parish_id = $9, sector = $10, community = $11, street = $12, address = $13 where station_id = $14", r.Municipality_id, r.Name, r.Coordinates, r.Description, r.Code, r.Abbreviation, r.Phones, r.State_id, r.Parish_id, r.Sector, r.Community, r.Street, r.Address ,r.Id)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return station.ErrorStationNotUpdated
}

func (u *LocationsImpl) DeleteStation(id int64) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from locations.fire_station where station_id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return station.ErrorStationNotDeleted
}

func (u *LocationsImpl) MapToDto(s station.Station) (station.StationDto) {
	station := station.StationDto{}

	station.Municipality_id = utils.ConvertFromInt4(s.Municipality_id)
	station.Name = utils.ConvertFromText(s.Name)
	station.Coordinates = utils.ConvertFromText(s.Coordinates)
	station.Description = utils.ConvertFromText(s.Description)
	station.Code = utils.ConvertFromText(s.Code)
	station.Abbreviation = utils.ConvertFromText(s.Abbreviation)
	station.Phones = s.Phones
	station.State_id = utils.ConvertFromInt4(s.State_id)
	station.Parish_id = utils.ConvertFromInt4(s.Parish_id)
	station.Sector = utils.ConvertFromText(s.Sector)
	station.Community = utils.ConvertFromText(s.Community)
	station.Street = utils.ConvertFromText(s.Street)
	station.Address = utils.ConvertFromText(s.Address)

	return station
}

func (u *LocationsImpl) MapFromDto(s station.StationDto) (station.Station) {
	station := station.Station{}

	station.Municipality_id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Municipality_id))
	station.Name = utils.ConvertToPgTypeText(s.Name)
	station.Coordinates = utils.ConvertToPgTypeText(s.Coordinates)
	station.Description = utils.ConvertToPgTypeText(s.Description)
	station.Code = utils.ConvertToPgTypeText(s.Code)
	station.Abbreviation = utils.ConvertToPgTypeText(s.Abbreviation)
	station.Phones = s.Phones
	station.State_id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.State_id))
	station.Parish_id = utils.ConvertToPgTypeInt4(utils.ParseInt(s.Parish_id))
	station.Sector = utils.ConvertToPgTypeText(s.Sector)
	station.Community = utils.ConvertToPgTypeText(s.Community)
	station.Street = utils.ConvertToPgTypeText(s.Street)
	station.Address = utils.ConvertToPgTypeText(s.Address)

	return station
}

