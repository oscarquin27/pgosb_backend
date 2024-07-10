package user_domain

import (
	"context"
	entities "fdms/domain/entities/users"
	"fdms/infra/config"
	authentication "fdms/infra/keycloak"
	"fdms/utils"
	"strconv"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"github.com/jackc/pgx/v5"
)

var keycloakAuthService authentication.KeycloakAuthenticationService
var keycloakClient gocloak.GoCloak

func init() {
	keycloakClient = *gocloak.NewClient(config.Get().Keycloak.Address)
	keycloakAuthService = *authentication.NewService(&keycloakClient)
}

func (u *UserImpl) GetUser(id int) (*entities.User, error) {
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
	promotion_date::text, 
	promotion, 
	condition, 
	division, 
	profession, 
	institution,
	user_system, 
	status_user, 
	created_at, 
	updated_at, 
	last_connection, 
	ip
FROM users.user where id = $1;`, id)

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entities.User])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entities.ErrorUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserImpl) GetAll() ([]entities.User, error) {
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
	promotion_date::text, 
	promotion, 
	condition, 
	division, 
	profession, 
	institution,
	user_system, 
	status_user, 
	created_at, 
	updated_at, 
	last_connection, 
	ip
FROM users.user`)

	user, err := pgx.CollectRows(rows, pgx.RowToStructByName[entities.User])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entities.ErrorUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (u *UserImpl) Create(user *entities.User) error {
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

	if user.UserProfile.Promotion_date == "" {
		user.UserProfile.Promotion_date = time.Now().Format("2000-01-01")
	}
	
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	
	var userId int

	err = tx.QueryRow(ctx, `insert into users.user (id_role, user_name, first_name, last_name, email, photo, gender, phone, secondary_phone, birth_date, age, residence, coordinates, marital_status, height, weight, shirt_size, pant_size, shoe_size, blood_type, allergies, code, personal_code, rank, promotion_date, promotion, condition, division, profession, institution, user_system, zip_code)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10::date, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25::date, $26, $27, $28, $29, $30, $31, $32) returning id;
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
		user.UserProfile.Promotion_date,
		user.UserProfile.Promotion,
		user.UserProfile.Condition,
		user.UserProfile.Division,
		user.UserProfile.Profession,
		user.UserProfile.Institution,
		user.UserProfile.User_system,
		user.UserProfile.Zip_code).Scan(&userId)

	if user.User_system {
		keycloakId, err := keycloakAuthService.CreateUser(ctx, user.UserProfile.User_name, user.UserProfile.Email, strconv.Itoa(userId), "12345")
		
		if err != nil {
			return entities.ErrorUserNotCreated
		}

		_, err = tx.Exec(ctx, `update users.user set id_keycloak = $1 where id = $2;`, keycloakId, userId)

		if err != nil {
			return entities.ErrorUserNotCreated
		}
	}
	
	return nil
}

func (u *UserImpl) Update(user *entities.User) error {
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
	
	var keycloakId string

	previous, err := u.GetUser(user.Id)

	err = tx.QueryRow(ctx, `
		UPDATE users.user
		SET id_role = $1,
			user_name = $2,
			first_name = $3,
			last_name = $4,
			email = $5,
			photo = $6,
			gender = $7,
			phone = $8,
			secondary_phone = $9,
			birth_date = $10,
			age = $11,
			residence = $12,
			coordinates = $13,
			marital_status = $14,
			height = $15,
			weight = $16,
			shirt_size = $17,
			pant_size = $18,
			shoe_size = $19,
			blood_type = $20,
			allergies = $21,
			code = $22,
			personal_code = $23,
			rank = $24,
			promotion_date = $25,
			promotion = $26,
			condition = $27,
			division = $28,
			profession = $29,
			institution = $30,
			user_system = $31,
			zip_code = $32
		WHERE id = $33 returning id_keycloak`,
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
		user.Promotion_date,
		user.UserProfile.Promotion,
		user.UserProfile.Condition,
		user.UserProfile.Division,
		user.UserProfile.Profession,
		user.UserProfile.Institution,
		user.UserProfile.User_system,
		user.UserProfile.Zip_code,
		user.UserIdentification.Id).Scan(&keycloakId)

	if previous.UserProfile.User_system && !user.UserProfile.User_system {
		err = keycloakAuthService.DeleteUser(ctx, keycloakId)

	} else if !previous.UserProfile.User_system && user.UserProfile.User_system || keycloakId == "" {
		keycloakId, err = keycloakAuthService.CreateUser(ctx, user.UserProfile.User_name, user.UserProfile.Email, strconv.Itoa(int(user.UserIdentification.Id)), "12345")		

		_, err = tx.Exec(ctx, `update users.user set id_keycloak = $1 where id = $2;`, keycloakId, user.UserIdentification.Id)
	}

	if err != nil {
		return entities.ErrorUserNotUpdated
	}

	return nil
}

func (u *UserImpl) Delete(id int) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from users.user where id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return entities.ErrorUserNotDeleted
}

func (u *UserImpl) MapFromDto(userDto *entities.UserDto) entities.User {
	user := entities.User{}

	if len(user.Promotion_date) <= 0 {
		user.Promotion_date = time.Now().String()
	}

	user.UserIdentification.Id = utils.ParseInt(userDto.Id)
	user.UserIdentification.Id_role = utils.ParseInt(userDto.Id_role)
	user.UserProfile.User_name = userDto.User_name
	user.UserProfile.First_name = userDto.First_name
	user.UserProfile.Last_name = userDto.Last_name
	user.UserProfile.Email = userDto.Email
	user.UserProfile.Photo = userDto.Photo
	user.UserProfile.Gender = userDto.Gender
	user.UserProfile.Phone = userDto.Phone
	user.UserProfile.Secondary_Phone = userDto.Secondary_Phone
	user.UserProfile.Birth_date = userDto.Birth_date
	user.UserProfile.Age = utils.ParseInt(userDto.Age)
	user.UserProfile.Residence = userDto.Residence
	user.UserProfile.Coordinates = userDto.Coordinates
	user.UserProfile.Marital_status = userDto.Marital_status
	user.UserProfile.Height = utils.ParseFloat(userDto.Height)
	user.UserProfile.Weight = utils.ParseFloat(userDto.Weight)
	user.UserProfile.Shirt_size = userDto.Shirt_size
	user.UserProfile.Pant_size = userDto.Pant_size
	user.UserProfile.Shoe_size = utils.ParseInt(userDto.Shoe_size)
	user.UserProfile.Blood_type = userDto.Blood_type
	user.UserProfile.Allergies = entities.MapAFromArray(userDto.Allergies)
	user.UserProfile.Code = userDto.Code
	user.UserProfile.Personal_code = userDto.Personal_code
	user.UserProfile.Rank = userDto.Rank
	user.Promotion_date = userDto.Promotion_date
	user.UserProfile.Promotion = userDto.Promotion
	user.UserProfile.Condition = userDto.Condition
	user.UserProfile.Division = userDto.Division
	user.UserProfile.Profession = userDto.Profession
	user.UserProfile.Institution = userDto.Profession
	user.UserProfile.User_system = userDto.User_system
	user.UserProfile.Zip_code = userDto.Zip_code

	// user.UserProfile.Skills = userDto.Skills

	return user
}

func (u *UserImpl) MapToDto(user *entities.User) entities.UserDto {
	userDto := entities.UserDto{}
	//var err error
	userDto.Id = utils.ConvertFromInt(user.Id)
	userDto.UserIdentificationDto.Id_role = utils.ConvertFromInt(user.Id_role)
	userDto.UserProfileDto.User_name = user.User_name
	userDto.UserProfileDto.First_name = user.First_name
	userDto.UserProfileDto.Last_name = user.Last_name
	userDto.UserProfileDto.Email = user.Email
	userDto.UserProfileDto.Photo = user.Photo
	userDto.UserProfileDto.Gender = user.Gender
	userDto.UserProfileDto.Phone = user.Phone
	userDto.UserProfileDto.Secondary_Phone = user.Secondary_Phone
	userDto.UserProfileDto.Birth_date = user.Birth_date
	userDto.UserProfileDto.Age = utils.ConvertFromInt(user.Age)
	userDto.UserProfileDto.Residence = user.Residence
	userDto.UserProfileDto.Coordinates = user.Coordinates
	userDto.UserProfileDto.Marital_status = user.Marital_status
	userDto.UserProfileDto.Height = utils.ConvertFromDecimal(user.Height)
	userDto.UserProfileDto.Weight = utils.ConvertFromDecimal(user.Weight)
	userDto.UserProfileDto.Shirt_size = user.Shirt_size
	userDto.UserProfileDto.Pant_size = user.Pant_size
	userDto.UserProfileDto.Shoe_size = utils.ConvertFromInt(user.Shoe_size)
	userDto.UserProfileDto.Blood_type = user.Blood_type
	userDto.UserProfileDto.Allergies = entities.MapToArray(user.Allergies)
	userDto.UserProfileDto.Code = user.Code
	userDto.UserProfileDto.Personal_code = user.Personal_code
	userDto.UserProfileDto.Rank = user.Rank
	userDto.Promotion_date = user.Promotion_date
	userDto.UserProfileDto.Promotion = user.Promotion
	userDto.UserProfileDto.Condition = user.Condition
	userDto.UserProfileDto.Division = user.Division
	userDto.UserProfileDto.Profession = user.Profession
	userDto.UserProfileDto.Institution = user.Profession
	userDto.UserProfileDto.User_system = user.User_system
	userDto.UserProfileDto.Zip_code = user.Zip_code

	// userDto.UserProfileDto.Skills = user.Skills

	return userDto
}
