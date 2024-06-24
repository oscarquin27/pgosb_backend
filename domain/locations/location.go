package location_domain

import (
	city "fdms/domain/entities/cities"
	municipality "fdms/domain/entities/municipalities"
	parish "fdms/domain/entities/parish"
	state "fdms/domain/entities/states"
	station "fdms/domain/entities/stations"

	"github.com/jackc/pgx/v5/pgxpool"
)
type LocationsRepository interface {
	GetState(id int64) (*state.State, error)
	GetAllStates()([]state.State, error)
	CreateState(state *state.State) (error)
	UpdateState(state *state.State) (error)
	DeleteState(id int64) (error)
	GetCity(id int64) (*city.City, error)
	GetAllCity()([]city.City, error)
	CreateCity(city *city.City) (error)
	UpdateCity(city *city.City) (error)
	DeleteCity(id int64) (error)
	GetMunicipality(id int64) (*municipality.Municipality, error)
	GetAllMunicipality()([]municipality.Municipality, error)
	CreateMunicipality(municipality *municipality.Municipality) (error)
	UpdateMunicipality(municipality *municipality.Municipality) (error)
	DeleteMunicipality(id int64) (error)
	GetParish(id int64) (*parish.Parish, error)
	GetAllParish()([]parish.Parish, error)
	CreateParish(municipality *parish.Parish) (error)
	UpdateParish(municipality *parish.Parish) (error)
	DeleteParish(id int64) (error)
	GetStation(id int64) (*station.Station, error)
	GetAllStations()([]station.Station, error)
	CreateStation(state *station.Station) (error)
	UpdateStation(state *station.Station) (error)
	DeleteStation(id int64) (error)
}

type LocationsImpl struct {
	db *pgxpool.Pool
}

func NewLocationService(db *pgxpool.Pool) LocationsRepository {
	return &LocationsImpl{
		db : db,
	}
}