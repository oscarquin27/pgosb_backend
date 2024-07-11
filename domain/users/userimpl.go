package user_domain

import (
	"context"
	entities "fdms/domain/entities/users"
	"fdms/infra/config"
	authentication "fdms/infra/keycloak"
	"fdms/src/utils"
	"strconv"

	"github.com/Nerzal/gocloak/v13"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var keycloakAuthService authentication.KeycloakAuthenticationService
var keycloakClient gocloak.GoCloak

func init() {
	keycloakClient = *gocloak.NewClient(config.Get().Keycloak.Address)
	keycloakAuthService = *authentication.NewService(&keycloakClient)
}

func (u *UserImpl) GetUser(id int64) (*entities.User, error) {
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
		return entities.ErrorUserNotCreated
	}

	if user.User_system.Bool {
		keycloakId, err := keycloakAuthService.CreateUser(ctx, user.UserProfile.User_name.String, user.UserProfile.Email.String, strconv.Itoa(userId), "12345")

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

	var keycloakId pgtype.Text

	previous, err := u.GetUser(user.Id)

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
		err = keycloakAuthService.DeleteUser(ctx, keycloakId.String)

	} else if !previous.UserProfile.User_system.Bool && user.UserProfile.User_system.Bool || keycloakId.String == "" {
		keycloakId.String, err = keycloakAuthService.CreateUser(ctx, user.UserProfile.User_name.String, user.UserProfile.Email.String, strconv.Itoa(int(user.UserIdentification.Id)), "12345")

		_, err = tx.Exec(ctx, `update users.user set id_keycloak = $1 where id = $2;`, keycloakId, user.UserIdentification.Id)
	}

	if err != nil {
		return entities.ErrorUserNotUpdated
	}

	return nil
}

func (u *UserImpl) Delete(id int64) error {
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
		err = keycloakAuthService.DeleteUser(ctx, keycloakId.String)

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

func (u *UserImpl) MapFromDto(userDto *entities.UserDto) entities.User {
	user := entities.User{}

	id_role := utils.ParseInt(userDto.Id_role)

	age := utils.ParseInt(userDto.Age)

	height := utils.ParseInt(userDto.Height)

	weight := utils.ParseInt(userDto.Weight)

	shoe := utils.ParseInt(userDto.Shoe_size)

	user.UserIdentification.Id = int64(utils.ParseInt(userDto.Id))
	user.UserIdentification.Id_role = utils.ConvertToPgTypeInt4(id_role)
	user.UserProfile.User_name = utils.ConvertToPgTypeText(userDto.User_name)
	user.UserProfile.First_name = utils.ConvertToPgTypeText(userDto.First_name)
	user.UserProfile.Last_name = utils.ConvertToPgTypeText(userDto.Last_name)
	user.UserProfile.Email = utils.ConvertToPgTypeText(userDto.Email)
	user.UserProfile.Photo = utils.ConvertToPgTypeText(userDto.Photo)
	user.UserProfile.Gender = utils.ConvertToPgTypeText(userDto.Gender)
	user.UserProfile.Phone = utils.ConvertToPgTypeText(userDto.Phone)
	user.UserProfile.Secondary_Phone = utils.ConvertToPgTypeText(userDto.Secondary_Phone)
	user.UserProfile.Birth_date = utils.ConvertToPgTypeText(userDto.Birth_date)
	user.UserProfile.Age = utils.ConvertToPgTypeInt2(age)
	user.UserProfile.Residence = utils.ConvertToPgTypeText(userDto.Residence)
	user.UserProfile.Coordinates = utils.ConvertToPgTypeText(userDto.Coordinates)
	user.UserProfile.Marital_status = utils.ConvertToPgTypeText(userDto.Marital_status)
	user.UserProfile.Height = utils.ConvertToPgTypeNumeric(height)
	user.UserProfile.Weight = utils.ConvertToPgTypeNumeric(weight)
	user.UserProfile.Shirt_size = utils.ConvertToPgTypeText(userDto.Shirt_size)
	user.UserProfile.Pant_size = utils.ConvertToPgTypeText(userDto.Pant_size)
	user.UserProfile.Shoe_size = utils.ConvertToPgTypeNumeric(shoe)
	user.UserProfile.Blood_type = utils.ConvertToPgTypeText(userDto.Blood_type)
	user.UserProfile.Allergies = userDto.Allergies
	user.UserProfile.Code = utils.ConvertToPgTypeText(userDto.Code)
	user.UserProfile.Personal_code = utils.ConvertToPgTypeText(userDto.Personal_code)
	user.UserProfile.Rank = utils.ConvertToPgTypeText(userDto.Rank)
	user.UserProfile.Promotion = utils.ConvertToPgTypeText(userDto.Promotion)
	user.UserProfile.Condition = utils.ConvertToPgTypeText(userDto.Condition)
	user.UserProfile.Division = utils.ConvertToPgTypeText(userDto.Division)
	user.UserProfile.Profession = utils.ConvertToPgTypeText(userDto.Profession)
	user.UserProfile.Institution = utils.ConvertToPgTypeText(userDto.Profession)
	user.UserProfile.User_system = utils.ConvertToPgTypeBool(userDto.User_system)
	user.UserProfile.Zip_code = utils.ConvertToPgTypeText(userDto.Zip_code)
	userDto.UserProfileDto.Skills = user.Skills
	user.UserProfile.State = utils.ConvertToPgTypeText(userDto.State)
	user.UserProfile.Municipality = utils.ConvertToPgTypeText(userDto.Municipality)
	user.UserProfile.Parish = utils.ConvertToPgTypeText(userDto.Parish)
	user.UserProfile.Sector = utils.ConvertToPgTypeText(userDto.Sector)
	user.UserProfile.Community = utils.ConvertToPgTypeText(userDto.Community)
	user.UserProfile.Street = utils.ConvertToPgTypeText(userDto.Street)
	user.UserProfile.Beach = utils.ConvertToPgTypeText(userDto.Beach)
	user.UserProfile.Address = utils.ConvertToPgTypeText(userDto.Address)
	user.UserProfile.Legal_id = utils.ConvertToPgTypeText(userDto.Legal_id)

	return user
}

func (u *UserImpl) MapToDto(user *entities.User) entities.UserDto {
	userDto := entities.UserDto{}
	//var err error
	userDto.Id = strconv.FormatInt(user.Id, 10)
	userDto.UserIdentificationDto.Id_role = utils.ConvertFromInt4(user.Id_role)
	userDto.UserProfileDto.User_name = utils.ConvertFromText(user.User_name)
	userDto.UserProfileDto.First_name = utils.ConvertFromText(user.First_name)
	userDto.UserProfileDto.Last_name = utils.ConvertFromText(user.Last_name)
	userDto.UserProfileDto.Email = utils.ConvertFromText(user.Email)
	userDto.UserProfileDto.Photo = utils.ConvertFromText(user.Photo)
	userDto.UserProfileDto.Gender = utils.ConvertFromText(user.Gender)
	userDto.UserProfileDto.Phone = utils.ConvertFromText(user.Phone)
	userDto.UserProfileDto.Secondary_Phone = utils.ConvertFromText(user.Secondary_Phone)
	userDto.UserProfileDto.Birth_date = utils.ConvertFromText(user.Birth_date)
	userDto.UserProfileDto.Age = utils.ConvertFromInt2(user.Age)
	userDto.UserProfileDto.Residence = utils.ConvertFromText(user.Residence)
	userDto.UserProfileDto.Coordinates = utils.ConvertFromText(user.Coordinates)
	userDto.UserProfileDto.Marital_status = utils.ConvertFromText(user.Marital_status)
	userDto.UserProfileDto.Height = utils.ConvertFromNumeric(user.Height)
	userDto.UserProfileDto.Weight = utils.ConvertFromNumeric(user.Weight)
	userDto.UserProfileDto.Shirt_size = utils.ConvertFromText(user.Shirt_size)
	userDto.UserProfileDto.Pant_size = utils.ConvertFromText(user.Pant_size)
	userDto.UserProfileDto.Shoe_size = utils.ConvertFromNumeric(user.Shoe_size)
	userDto.UserProfileDto.Blood_type = utils.ConvertFromText(user.Blood_type)
	userDto.UserProfileDto.Allergies = user.Allergies
	userDto.UserProfileDto.Code = utils.ConvertFromText(user.Code)
	userDto.UserProfileDto.Personal_code = utils.ConvertFromText(user.Personal_code)
	userDto.UserProfileDto.Rank = utils.ConvertFromText(user.Rank)
	userDto.UserProfileDto.Promotion = utils.ConvertFromText(user.Promotion)
	userDto.UserProfileDto.Condition = utils.ConvertFromText(user.Condition)
	userDto.UserProfileDto.Division = utils.ConvertFromText(user.Division)
	userDto.UserProfileDto.Profession = utils.ConvertFromText(user.Profession)
	userDto.UserProfileDto.Institution = utils.ConvertFromText(user.Institution)
	userDto.UserProfileDto.User_system = utils.ConvertFromBool(user.User_system)
	userDto.UserProfileDto.Zip_code = utils.ConvertFromText(user.Zip_code)
	userDto.UserProfileDto.Skills = user.Skills
	userDto.UserProfileDto.State = utils.ConvertFromText(user.State)
	userDto.UserProfileDto.Municipality = utils.ConvertFromText(user.Municipality)
	userDto.UserProfileDto.Parish = utils.ConvertFromText(user.Parish)
	userDto.UserProfileDto.Sector = utils.ConvertFromText(user.Sector)
	userDto.UserProfileDto.Community = utils.ConvertFromText(user.Community)
	userDto.UserProfileDto.Street = utils.ConvertFromText(user.Street)
	userDto.UserProfileDto.Beach = utils.ConvertFromText(user.Beach)
	userDto.UserProfileDto.Address = utils.ConvertFromText(user.Address)
	userDto.UserProfileDto.Legal_id = utils.ConvertFromText(user.Legal_id)

	return userDto
}
