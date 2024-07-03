package user_domain

import (
	"context"
	entities "fdms/domain/entities/users"
	"fdms/utils"
	"time"

	"github.com/jackc/pgx/v5"
)


func (u *UserImpl) GetUser(id int64) (*entities.User, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil,err
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

	return &user,nil
}

func (u *UserImpl) GetAll() ([]entities.User, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil,err
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

	return user,nil
}

func (u *UserImpl) Create(user *entities.User) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	if user.UserProfile.Promotion_date == "" {
		user.UserProfile.Promotion_date = time.Now().Format("2000-01-01")
	}

	rows, err := conn.Exec(ctx, `insert into users.user (id_role, user_name, first_name, last_name, email, photo, gender, phone, secondary_phone, birth_date, age, residence, coordinates, marital_status, height, weight, shirt_size, pant_size, shoe_size, blood_type, allergies, code, personal_code, rank, promotion_date, promotion, condition, division, profession, institution, user_system, zip_code)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10::date, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25::date, $26, $27, $28, $29, $30, $31, $32);
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
	user.UserProfile.Zip_code)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return entities.ErrorUserNotCreated
}

func (u *UserImpl) Update(user *entities.User) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx,`
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
		WHERE id = $33`,
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
	user.UserIdentification.Id)
	
	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return entities.ErrorUserNotUpdated
}

func (u *UserImpl) Delete(id int64) (error) {
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

func (u *UserImpl) MapFromDto(userDto *entities.UserDto) (entities.User) {
	user := entities.User{}

	id := utils.ParseInt(userDto.Id)

	id_role := utils.ParseInt(userDto.Id_role)

	age := utils.ParseInt(userDto.Age)

	height := utils.ParseInt(userDto.Height)

	weight := utils.ParseInt(userDto.Weight)

	shirt := utils.ParseInt(userDto.Shirt_size)

	pant := utils.ParseInt(userDto.Pant_size)

	shoe := utils.ParseInt(userDto.Shoe_size)

	if len(user.Promotion_date) <= 0 {
		user.Promotion_date = time.Now().String()
	}

	user.UserIdentification.Id = utils.ConvertToPgTypeInt4(id)
	user.UserIdentification.Id_role = utils.ConvertToPgTypeInt4(id_role)
	user.UserProfile.User_name = utils.ConvertToPgTypeText(userDto.User_name)
	user.UserProfile.First_name = utils.ConvertToPgTypeText(userDto.First_name)
	user.UserProfile.Last_name = utils.ConvertToPgTypeText(userDto.Last_name)
	user.UserProfile.Email = utils.ConvertToPgTypeText(userDto.Email)
	user.UserProfile.Photo = utils.ConvertToPgTypeText(userDto.Photo) 
	user.UserProfile.Gender = utils.ConvertToPgTypeText(userDto.Gender)
	user.UserProfile.Phone = utils.ConvertToPgTypeText(userDto.Phone)
	user.UserProfile.Secondary_Phone = utils.ConvertToPgTypeText(userDto.Secondary_Phone)
	user.UserProfile.Birth_date = userDto.Birth_date
	user.UserProfile.Age = utils.ConvertToPgTypeInt2(age)
	user.UserProfile.Residence = utils.ConvertToPgTypeText(userDto.Residence)
	user.UserProfile.Coordinates = utils.ConvertToPgTypeText(userDto.Coordinates)
	user.UserProfile.Marital_status = utils.ConvertToPgTypeText(userDto.Marital_status)
	user.UserProfile.Height = utils.ConvertToPgTypeNumeric(height)
	user.UserProfile.Weight = utils.ConvertToPgTypeNumeric(weight)
	user.UserProfile.Shirt_size = utils.ConvertToPgTypeNumeric(shirt)
	user.UserProfile.Pant_size = utils.ConvertToPgTypeNumeric(pant)
	user.UserProfile.Shoe_size = utils.ConvertToPgTypeNumeric(shoe)
	user.UserProfile.Blood_type = utils.ConvertToPgTypeText(userDto.Blood_type)
	user.UserProfile.Allergies = userDto.Allergies
	user.UserProfile.Code = utils.ConvertToPgTypeText(userDto.Code)
	user.UserProfile.Personal_code = utils.ConvertToPgTypeText(userDto.Personal_code)
	user.UserProfile.Rank = utils.ConvertToPgTypeText(userDto.Rank)
	user.Promotion_date = userDto.Promotion_date
	user.UserProfile.Promotion = utils.ConvertToPgTypeText(userDto.Promotion)
	user.UserProfile.Condition = utils.ConvertToPgTypeText(userDto.Condition)
	user.UserProfile.Division = utils.ConvertToPgTypeText(userDto.Division)
	user.UserProfile.Profession = utils.ConvertToPgTypeText(userDto.Profession)
	user.UserProfile.Institution = utils.ConvertToPgTypeText(userDto.Profession)
	user.UserProfile.User_system = utils.ConvertToPgTypeBool(userDto.User_system)
	user.UserProfile.Zip_code = utils.ConvertToPgTypeText(userDto.Zip_code)

	return user
}

func (u *UserImpl) MapToDto(user *entities.User) (entities.UserDto) {
	userDto := entities.UserDto{}
	//var err error
	userDto.Id = utils.ConvertFromInt4(user.Id)
	userDto.UserIdentificationDto.Id_role = utils.ConvertFromInt4(user.Id_role)
	userDto.UserProfileDto.User_name = utils.ConvertFromText(user.User_name)
	userDto.UserProfileDto.First_name = utils.ConvertFromText(user.First_name)
	userDto.UserProfileDto.Last_name = utils.ConvertFromText(user.Last_name)
	userDto.UserProfileDto.Email = utils.ConvertFromText(user.Email)
	userDto.UserProfileDto.Photo = utils.ConvertFromText(user.Photo) 
	userDto.UserProfileDto.Gender = utils.ConvertFromText(user.Gender)
	userDto.UserProfileDto.Phone = utils.ConvertFromText(user.Phone)
	userDto.UserProfileDto.Secondary_Phone = utils.ConvertFromText(user.Secondary_Phone)
	userDto.UserProfileDto.Birth_date = user.Birth_date
	userDto.UserProfileDto.Age = utils.ConvertFromInt2(user.Age)
	userDto.UserProfileDto.Residence = utils.ConvertFromText(user.Residence)
	userDto.UserProfileDto.Coordinates = utils.ConvertFromText(user.Coordinates)
	userDto.UserProfileDto.Marital_status = utils.ConvertFromText(user.Marital_status)
	userDto.UserProfileDto.Height = utils.ConvertFromNumeric(user.Height)
	userDto.UserProfileDto.Weight = utils.ConvertFromNumeric(user.Weight)
	userDto.UserProfileDto.Shirt_size = utils.ConvertFromNumeric(user.Shirt_size)
	userDto.UserProfileDto.Pant_size = utils.ConvertFromNumeric(user.Pant_size)
	userDto.UserProfileDto.Shoe_size = utils.ConvertFromNumeric(user.Shoe_size)
	userDto.UserProfileDto.Blood_type = utils.ConvertFromText(user.Blood_type)
	userDto.UserProfileDto.Allergies = userDto.Allergies
	userDto.UserProfileDto.Code = utils.ConvertFromText(user.Code)
	userDto.UserProfileDto.Personal_code = utils.ConvertFromText(user.Personal_code)
	userDto.UserProfileDto.Rank = utils.ConvertFromText(user.Rank)
	userDto.Promotion_date = user.Promotion_date
	userDto.UserProfileDto.Promotion = utils.ConvertFromText(user.Promotion)
	userDto.UserProfileDto.Condition = utils.ConvertFromText(user.Condition)
	userDto.UserProfileDto.Division = utils.ConvertFromText(user.Division)
	userDto.UserProfileDto.Profession = utils.ConvertFromText(user.Profession)
	userDto.UserProfileDto.Institution = utils.ConvertFromText(user.Profession)
	userDto.UserProfileDto.User_system = utils.ConvertFromBool(user.User_system)
	userDto.UserProfileDto.Zip_code = utils.ConvertFromText(user.Zip_code)

	return userDto
}
