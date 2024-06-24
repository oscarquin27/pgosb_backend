package user_domain

import (
	"context"
	entities "fdms/domain/entities/users"

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
	promotion_date, 
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
	promotion_date, 
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

func (u *UserImpl) Create(user *entities.UserCreateDto) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `insert into users.user (id_role, user_name, first_name, last_name, email, photo, gender, phone, secondary_phone, birth_date, age, residence, coordinates, marital_status, height, weight, shirt_size, pant_size, shoe_size, blood_type, allergies, code, personal_code, rank, promotion_date, promotion, condition, division, profession, institution, user_system, zip_code)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32);
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
	user.Promotion_date,
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

func (u *UserImpl) Update(user *entities.UserUpdateDto) (error) {
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

	return entities.ErrorUserNotCreated
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
