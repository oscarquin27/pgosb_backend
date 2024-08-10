package repository

// type MunicipalityRepository struct {
// 	*AbstractRepository[models.Municipality]
// }

// func NewMunicipalityService(db *pgxpool.Pool) services.MunicipalityService {

// 	abstractImplent := NewAbstractRepository[models.Municipality](db)

// 	return &MunicipalityRepository{
// 		&abstractImplent,
// 	}
// }

// const selectMunicipalityQuery = "SELECT id, name, coordinates , capital FROM locations.states WHERE id = $1"

// const selectAllMunicipalityQuery = "SELECT id, name, coordinates , capital FROM locations.states"

// const insertMunicipalityQuery = `INSERT INTO locations.states(
// 	name, coordinates, capital)
// 	VALUES (@name, @coordinates, @capital)`

// const updateMunicipalityQuery = `UPDATE locations.states
// 	SET  name=@name, coordinates=@coordinates, capital=@capital
// 	WHERE id = @id `

// const deleteMunicipalityQuery = `DELETE FROM locations.states WHERE id = $1`

// func (u *MunicipalityRepository) Get(id int64) *results.ResultWithValue[*models.Municipality] {

// 	r := u.AbstractRepository.Get(id, selectStateQuery)

// 	results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, r.Value, r.Err)
// }
// func (u *MunicipalityRepository) GetAll() ([]models.Municipality, *results.GeneralError) {

// 	var municipalities []models.Municipality = make([]models.Municipality, 0)

// 	values, err := u.AbstractRepository.GetAll(selectAllStateQuery)

// 	if err != nil {
// 		return municipalities, err
// 	}

// 	return values, nil
// }

// func (u *MunicipalityRepository) Create(state *models.Municipality) *results.ResultWithValue[models.Municipality] {

// 	r := u.AbstractRepository.Create(*state, insertStateQuery, state.GetNameArgs(), state.SetId)

// 	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
// }

// func (u *MunicipalityRepository) Update(state models.Municipality) *results.ResultWithValue[models.Municipality] {
// 	r := u.AbstractRepository.Update(*state, updateStateQuery, state.GetNameArgs())

// 	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
// }

// func (u *MunicipalityRepository) Delete(id int64) *results.Result {

// 	return u.AbstractRepository.Delete(id, deleteStateQuery)
// }
