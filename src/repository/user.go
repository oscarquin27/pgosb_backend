package repository

import (
	"context"
	"database/sql"
	"fdms/src/infrastructure/keycloak"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils"
	"fdms/src/utils/results"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const RolNotFound = "el rol no fue encontrado"

type UserRepository struct {
	db   *pgxpool.Pool
	auth *keycloak.KeycloakAuthenticationService
}

func NewUserService(db *pgxpool.Pool, auth *keycloak.KeycloakAuthenticationService) services.UserService {
	return &UserRepository{
		db:   db,
		auth: auth,
	}
}

func (u *UserRepository) Get(id int64) *results.ResultWithValue[*models.User] {

	r := results.NewResultWithValue[*models.User]("Get-User", false, nil, nil).Failure()

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT u.id, u.id_role,u.user_name,u.first_name,u.last_name,u.email,u.photo,u.gender,u.phone,
	u.secondary_phone,
	u.birth_date,
	u.age,
	u.zip_code,
	u.residence,
	u.coordinates,
	u.marital_status,
	u.height,
	u.weight,
	u.shirt_size,
	u.pant_size,
	u.shoe_size,
	u.blood_type,
	u.allergies,
	u.code,
	u.personal_code,
	u.rank,
	u.promotion,
	u.condition,
	u.division,
	u.profession,
	u.station,
	u.user_system,
	u.skills,
	u.state,
	u.municipality,
	u.parish,
	u.sector,
	u.community,
	u.street,
	u.beach,
	u.address,
	u.legal_id,
	u.status_user,
	ra.role_name as role
FROM users.user u
left join users.roles ra on ra.id = u.id_role
where u.id = $1`, id)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])

	if err != nil {
		if err == pgx.ErrNoRows {
			return r.WithError(results.
				NewNotFoundError("no se encontro el usuario espeificado", err))
		}

		return r.WithError(
			results.NewUnknowError("error obteniendo usuario", err))
	}

	return r.Success().WithValue(&user)
}

func (u *UserRepository) GetAll() ([]models.User, *results.GeneralError) {

	var usersDefault []models.User = make([]models.User, 0)

	ctx := context.Background() // Or use a specific context

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return usersDefault, results.
			NewUnknowError("no se pudo adquirir conexion", err)
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT u.id, u.id_role,
	u.user_name,
	u.first_name,
	u.last_name,
	u.email,
	u.photo,
	u.gender,
	u.phone,
	u.secondary_phone,
	u.birth_date,
	u.age,
	u.zip_code,
	u.residence,
	u.coordinates,
	u.marital_status,
	u.height,
	u.weight,
	u.shirt_size,
	u.pant_size,
	u.shoe_size,
	u.blood_type,
	u.allergies,
	u.code,
	u.personal_code,
	u.rank,
	u.promotion,
	u.condition,
	u.division,
	u.profession,
	u.station,
	u.user_system,
	u.skills,
	u.state,
	u.municipality,
	u.parish,
	u.sector,
	u.community,
	u.street,
	u.beach,
	u.address,
	u.legal_id,
	u.status_user,
	
	ra.role_name as role
FROM users.user u
left join users.roles ra on ra.id = u.id_role`)

	if err != nil {
		return usersDefault, results.
			NewUnknowError("no se pudo ejecutar el query", err)
	}

	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])

	if err != nil {
		if err == pgx.ErrNoRows {
			return usersDefault, results.NewNotFoundError("no encontraron registros", err)
		}

		return usersDefault, results.
			NewUnknowError("no se pudo ejecutar el query", err)
	}

	return users, nil
}

func (u *UserRepository) Create(user *models.User) *results.ResultWithValue[*models.User] {

	r := results.NewResultWithValue[*models.User]("Create-User", false, nil, nil).Failure()

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer tx.Rollback(ctx)

	var userId int64

	var idRol int

	err = tx.QueryRow(ctx, `select id from users.roles where role_name = $1`, user.Role).Scan(&idRol)

	if err != nil {

		if err == pgx.ErrNoRows {
			return r.Failure().
				WithError(
					results.NewErrorWithCode(RolNotFound, err.Error(), err))
		}

		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	err = tx.QueryRow(ctx, `insert into users.user
	(id_role,
	user_name,
	first_name,
	last_name,
	email,
	photo,
	gender,
	phone,
	secondary_phone,
	birth_date,
	age,
	residence,
	coordinates,
	marital_status,
	height,
	weight,
	shirt_size,
	pant_size,
	shoe_size,
	blood_type,
	allergies,
	code,
	personal_code,
	rank,
	promotion,
	condition,
	division,
	profession,
	station,
	user_system,
	zip_code,
	skills,
	state,
	municipality,
	parish,
	sector,
	community,
	street,
	beach,
	address,
	legal_id)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41) returning id;
`,
		idRol,
		user.UserProfile.User_name,
		user.UserProfile.First_name,
		user.UserProfile.Last_name,
		user.UserProfile.Email,
		user.UserProfile.Photo,
		user.UserProfile.Gender,
		user.UserProfile.Phone,
		user.UserProfile.Secondary_Phone,
		user.UserProfile.Birth_date,
		user.UserProfile.Age,
		user.UserProfile.Residence,
		user.UserProfile.Coordinates,
		user.UserProfile.Marital_status,
		user.UserProfile.Height,
		user.UserProfile.Weight,
		user.UserProfile.Shirt_size,
		user.UserProfile.Pant_size,
		user.UserProfile.Shoe_size,
		user.UserProfile.Blood_type,
		user.UserProfile.Allergies,
		user.UserProfile.Code,
		user.UserProfile.Personal_code,
		user.UserProfile.Rank,
		user.UserProfile.Promotion,
		user.UserProfile.Condition,
		user.UserProfile.Division,
		user.UserProfile.Profession,
		user.UserProfile.Station,
		user.UserProfile.User_system,
		user.UserProfile.Zip_code,
		user.UserProfile.Skills,
		user.UserProfile.State,
		user.UserProfile.Municipality,
		user.UserProfile.Parish,
		user.UserProfile.Sector,
		user.UserProfile.Community,
		user.UserProfile.Street,
		user.UserProfile.Beach,
		user.UserProfile.Address,
		user.UserProfile.Legal_id).Scan(&userId)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	user.Id = userId

	if user.User_system.Bool {

		keycloakId, err := u.auth.CreateUser(ctx, user.UserProfile.User_name.String,
			user.UserProfile.Email.String, utils.ParseInt64Sring(userId), "12345")

		if err != nil {
			return r.WithError(
				results.NewUnknowError("no se pudo adquirir conexion", err))
		}

		_, err = tx.Exec(ctx, `update users.user set id_keycloak = $1 where id = $2;`, keycloakId, userId)

		if err != nil {
			return r.WithError(
				results.NewUnknowError("erro actualizando", err))
		}

	}

	err = tx.Commit(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	return r.Success().WithValue(user)
}

func (u *UserRepository) Update(user *models.User) *results.ResultWithValue[*models.User] {

	r := results.NewResultWithValue[*models.User]("Update-User", false, nil, nil).Failure()

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer tx.Rollback(ctx)

	var keycloakId sql.NullString

	userResult := u.Get(user.Id)

	if !userResult.IsSuccessful {
		return userResult
	}

	var idRol int

	err = tx.QueryRow(ctx, `select id from users.roles where role_name = $1`, user.Role).Scan(&idRol)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	row := tx.QueryRow(ctx, `
		UPDATE users.user
		SET id_role = $1,
			first_name = $2,
			last_name = $3,
			email = $4,
			photo = $5,
			gender = $6,
			phone = $7,
			secondary_phone = $8,
			birth_date = $9,
			age = $10,
			residence = $11,
			coordinates = $12,
			marital_status = $13,
			height = $14,
			weight = $15,
			shirt_size = $16,
			pant_size = $17,
			shoe_size = $18,
			blood_type = $19,
			allergies = $20,
			code = $21,
			personal_code = $22,
			rank = $23,
			promotion = $24,
			condition = $25,
			division = $26,
			profession = $27,
			station = $28,
			user_system = $29,
			zip_code = $30,
			skills = $32,
			state = $33,
			municipality = $34,
			parish = $35,
			sector = $36,
			community = $37,
			street = $38,
			beach = $39,
			address = $40,
			legal_id = $41
		WHERE id = $31 returning id_keycloak`,
		idRol,
		user.UserProfile.First_name,
		user.UserProfile.Last_name,
		user.UserProfile.Email,
		user.UserProfile.Photo,
		user.UserProfile.Gender,
		user.UserProfile.Phone,
		user.UserProfile.Secondary_Phone,
		user.UserProfile.Birth_date,
		user.UserProfile.Age,
		user.UserProfile.Residence,
		user.UserProfile.Coordinates,
		user.UserProfile.Marital_status,
		user.UserProfile.Height,
		user.UserProfile.Weight,
		user.UserProfile.Shirt_size,
		user.UserProfile.Pant_size,
		user.UserProfile.Shoe_size,
		user.UserProfile.Blood_type,
		user.UserProfile.Allergies,
		user.UserProfile.Code,
		user.UserProfile.Personal_code,
		user.UserProfile.Rank,
		user.UserProfile.Promotion,
		user.UserProfile.Condition,
		user.UserProfile.Division,
		user.UserProfile.Profession,
		user.UserProfile.Station,
		user.UserProfile.User_system,
		user.UserProfile.Zip_code,
		user.UserIdentification.Id,
		user.UserProfile.Skills,
		user.UserProfile.State,
		user.UserProfile.Municipality,
		user.UserProfile.Parish,
		user.UserProfile.Sector,
		user.UserProfile.Community,
		user.UserProfile.Street,
		user.UserProfile.Beach,
		user.UserProfile.Address,
		user.UserProfile.Legal_id)

	err = row.Scan(&keycloakId)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	previous := userResult.Value

	if previous.UserProfile.User_system.Bool && !user.UserProfile.User_system.Bool {
		err = u.auth.DeleteUser(ctx, keycloakId.String)

		if err != nil {
			return r.WithError(
				results.NewUnknowError("no se pudo ejecutar query", err))
		}

	} else if !previous.UserProfile.User_system.Bool && user.UserProfile.User_system.Bool {
		keycloakId.String, err = u.auth.CreateUser(ctx, user.UserProfile.User_name.String,
			user.UserProfile.Email.String, strconv.Itoa(int(user.UserIdentification.Id)), "12345")

		if err != nil {
			return r.WithError(
				results.NewUnknowError("no se pudo ejecutar query", err))
		}

		row, err := tx.Exec(ctx, `update users.user set id_keycloak = $1 where id = $2;`, keycloakId, user.UserIdentification.Id)

		if err != nil {
			return r.WithError(
				results.NewUnknowError("no se pudo ejecutar query", err))
		}

		if row.RowsAffected() != 1 {
			return r.WithError(
				results.NewUnknowError("no se pudo ejecutar query", err))
		}

	}

	err = tx.Commit(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	return r.Success().WithValue(user)
}

func (u *UserRepository) Delete(id int64) *results.Result {

	r := results.NewResult("Delete-User", false, nil).Failure()

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return r.FailureWithError(err)
	}

	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return r.FailureWithError(err)
	}

	defer tx.Rollback(ctx)

	keycloakId := ""

	_ = tx.QueryRow(ctx, "select id_keycloak from users.user where id = $1", id).Scan(&keycloakId)

	if keycloakId != "" {
		err = u.auth.DeleteUser(ctx, keycloakId)

		if err != nil {
			return r.FailureWithError(err)
		}
	}

	rows, err := tx.Exec(ctx, "delete from users.user where id = $1", id)

	if err != nil {
		return r.FailureWithError(err)
	}

	if rows.RowsAffected() == 0 {
		return r.WithError(results.NewNotFoundError("no se consiguio unidad", err))
	}

	if rows.RowsAffected() > 1 {
		return r.WithError(results.NewUnknowError("se borraron multiples registros", err))
	}

	err = tx.Commit(ctx)

	if err != nil {
		return r.FailureWithError(err)
	}

	return r.Success()
}

func (u *UserRepository) GetAllSimple() *results.ResultWithValue[[]models.UserSimple] {

	r := results.NewResultWithValue[[]models.UserSimple]("Get-All-Simple", false, make([]models.UserSimple, 0), nil).
		Failure()

	allUsers, err := u.GetAll()

	if err != nil {
		return r.WithError(err)
	}

	var usersSimples []models.UserSimple = make([]models.UserSimple, 0)

	for _, user := range allUsers {

		if user.StatusUser == nil {
			continue
		}

		status := *user.StatusUser

		if status != "ACTIVO" {
			continue
		}

		userSimple := models.UserSimple{}

		userSimple = *userSimple.UserSimpleFromUser(&user)

		usersSimples = append(usersSimples, userSimple)
	}

	return r.Success().WithValue(usersSimples)
}
