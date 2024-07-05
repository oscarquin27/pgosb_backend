package postgres_respository

import (
	"context"
	user_domain "fdms/domain/entities/users"
	"fdms/repository/user_repository/postgres_respository/user_dto"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserImpl struct {
	db *pgxpool.Pool
}

func NewUserImplDb(db *pgxpool.Pool) *UserImpl {
	return &UserImpl{
		db: db,
	}
}

func (u *UserImpl) GetUser(id int64) (*user_domain.User, error) {

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

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[user_dto.UserDto])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, user_domain.ErrorUserNotFound
		}

		return nil, err
	}

	userEntity := user_dto.MapToEntity(&user)

	return &userEntity, nil
}

func (u *UserImpl) GetAll() ([]user_domain.User, error) {
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

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[user_dto.UserDto])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, user_domain.ErrorUserNotFound
		}

		return nil, err
	}

	usersEntityList := []user_domain.User{}

	for _, u := range users {
		usersEntityList = append(usersEntityList, user_dto.MapToEntity(&u))
	}

	return usersEntityList, nil
}

func (u *UserImpl) Create(userEntity *user_domain.User) error {

	user := user_dto.MapFromEntity(userEntity)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	if user.UserProfileDto.Promotion_date == "" {
		user.UserProfileDto.Promotion_date = time.Now().Format("2000-01-01")
	}

	rows, err := conn.Exec(ctx, `insert into users.user (id_role, user_name, first_name, last_name, email, photo, gender, phone, secondary_phone, birth_date, age, residence, coordinates, marital_status, height, weight, shirt_size, pant_size, shoe_size, blood_type, allergies, code, personal_code, rank, promotion_date, promotion, condition, division, profession, institution, user_system, zip_code)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10::date, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25::date, $26, $27, $28, $29, $30, $31, $32);
`,
		user.UserIdentificationDto.Id_role,
		user.UserProfileDto.User_name,
		user.UserProfileDto.First_name,
		user.UserProfileDto.Last_name,
		user.UserProfileDto.Email,
		user.UserProfileDto.Photo,
		user.UserProfileDto.Gender,
		user.UserProfileDto.Phone,
		user.UserProfileDto.Secondary_Phone,
		user.UserProfileDto.Birth_date,
		user.UserProfileDto.Age,
		user.UserProfileDto.Residence,
		user.UserProfileDto.Coordinates,
		user.UserProfileDto.Marital_status,
		user.UserProfileDto.Height,
		user.UserProfileDto.Weight,
		user.UserProfileDto.Shirt_size,
		user.UserProfileDto.Pant_size,
		user.UserProfileDto.Shoe_size,
		user.UserProfileDto.Blood_type,
		user.UserProfileDto.Allergies,
		user.UserProfileDto.Code,
		user.UserProfileDto.Personal_code,
		user.UserProfileDto.Rank,
		user.UserProfileDto.Promotion_date,
		user.UserProfileDto.Promotion,
		user.UserProfileDto.Condition,
		user.UserProfileDto.Division,
		user.UserProfileDto.Profession,
		user.UserProfileDto.Institution,
		user.UserProfileDto.User_system,
		user.UserProfileDto.Zip_code)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return user_domain.ErrorUserNotCreated
}

func (u *UserImpl) Update(userEntity *user_domain.User) error {

	user := user_dto.MapFromEntity(userEntity)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `
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
		user.UserIdentificationDto.Id_role,
		user.UserProfileDto.User_name,
		user.UserProfileDto.First_name,
		user.UserProfileDto.Last_name,
		user.UserProfileDto.Email,
		user.UserProfileDto.Photo,
		user.UserProfileDto.Gender,
		user.UserProfileDto.Phone,
		user.UserProfileDto.Secondary_Phone,
		user.UserProfileDto.Birth_date,
		user.UserProfileDto.Age,
		user.UserProfileDto.Residence,
		user.UserProfileDto.Coordinates,
		user.UserProfileDto.Marital_status,
		user.UserProfileDto.Height,
		user.UserProfileDto.Weight,
		user.UserProfileDto.Shirt_size,
		user.UserProfileDto.Pant_size,
		user.UserProfileDto.Shoe_size,
		user.UserProfileDto.Blood_type,
		user.UserProfileDto.Allergies,
		user.UserProfileDto.Code,
		user.UserProfileDto.Personal_code,
		user.UserProfileDto.Rank,
		user.Promotion_date,
		user.UserProfileDto.Promotion,
		user.UserProfileDto.Condition,
		user.UserProfileDto.Division,
		user.UserProfileDto.Profession,
		user.UserProfileDto.Institution,
		user.UserProfileDto.User_system,
		user.UserProfileDto.Zip_code,
		user.UserIdentificationDto.Id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return user_domain.ErrorUserNotUpdated
}

func (u *UserImpl) Delete(id int64) error {
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

	return user_domain.ErrorUserNotDeleted
}
