package repository

import (
	"context"
	"fdms/src/infrastructure/keycloak"
	"fdms/src/models"
	"fdms/src/services"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

func (u *UserRepository) Get(id int64) (*models.User, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT id, id_role, 
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
	zip_code,
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
	institution,
	user_system, 
	skills,
	state,
	municipality,
	parish,
	sector,
	community,
	street,
	beach,
	address,
	legal_id
FROM users.user where id = $1;`, id)

	if err != nil {
		return nil, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetAll() ([]models.User, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT id, id_role, 
	user_name, 
	first_name, 
	last_name, 
	email, 
	photo, 
	gender, 
	phone, 
	secondary_phone, 
	birth_date::text, 
	age,
	zip_code,
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
	institution,
	user_system, 
	skills,
	state,
	municipality,
	parish,
	sector,
	community,
	street,
	beach,
	address,
	legal_id
FROM users.user`)

	if err != nil {
		return nil, err
	}

	user, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (u *UserRepository) Create(user *models.User) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var userId int

	err = tx.QueryRow(ctx, `insert into users.user (id_role, user_name, first_name, last_name, email, photo, gender, phone, secondary_phone, birth_date, age, residence, coordinates, marital_status, height, weight, shirt_size, pant_size, shoe_size, blood_type, allergies, code, personal_code, rank, promotion, condition, division, profession, institution, user_system, zip_code, skills, state, municipality, parish, sector, community, street, beach, address, legal_id)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10::date, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41) returning id;
`,
		user.UserIdentification.Id_role,
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
		user.UserProfile.Institution,
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
		return models.ErrorUserNotCreated
	}

	if user.User_system.Bool {
		keycloakId, err := u.auth.CreateUser(ctx, user.UserProfile.User_name.String, user.UserProfile.Email.String, strconv.Itoa(userId), "12345")

		if err != nil {
			return models.ErrorUserNotCreated
		}

		_, err = tx.Exec(ctx, `update users.user set id_keycloak = $1 where id = $2;`, keycloakId, userId)

		if err != nil {
			return models.ErrorUserNotCreated
		}
	}

	return nil
}

func (u *UserRepository) Update(user *models.User) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var keycloakId pgtype.Text

	previous, err := u.Get(user.Id)

	err = tx.QueryRow(ctx, `
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
			institution = $28,
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
		user.UserIdentification.Id_role,
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
		user.UserProfile.Institution,
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
		user.UserProfile.Legal_id).Scan(&keycloakId)

	if err != nil {
		return err
	}
	if previous.UserProfile.User_system.Bool && !user.UserProfile.User_system.Bool {
		err = u.auth.DeleteUser(ctx, keycloakId.String)

	} else if !previous.UserProfile.User_system.Bool && user.UserProfile.User_system.Bool || keycloakId.String == "" {
		keycloakId.String, err = u.auth.CreateUser(ctx, user.UserProfile.User_name.String, user.UserProfile.Email.String, strconv.Itoa(int(user.UserIdentification.Id)), "12345")

		_, err = tx.Exec(ctx, `update users.user set id_keycloak = $1 where id = $2;`, keycloakId, user.UserIdentification.Id)
	}

	if err != nil {
		return models.ErrorUserNotUpdated
	}

	return nil
}

func (u *UserRepository) Delete(id int64) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var keycloakId pgtype.Text

	err = tx.QueryRow(ctx, "select id_keycloak from users.user where id = $1", id).Scan(&keycloakId)

	if err != nil {
		return err
	}

	if keycloakId.String != "" {
		err = u.auth.DeleteUser(ctx, keycloakId.String)

		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(ctx, "delete from users.user where id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
