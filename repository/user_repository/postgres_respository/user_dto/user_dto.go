package user_dto

import (
	user_domain "fdms/domain/entities/users"
	"fdms/utils"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserStatusDto struct {
	Status_user     pgtype.Text      `json:"status_user"`
	Last_connection pgtype.Timestamp `json:"last_connection"`
	Created_at      pgtype.Timestamp `json:"created_at"`
	Updated_at      pgtype.Timestamp `json:"updated_at"`
	Ip              pgtype.Text      `json:"ip"`
}

type UserProfileDto struct {
	User_name       pgtype.Text    `json:"user_name" binding:"required"`
	First_name      pgtype.Text    `json:"first_name"`
	Last_name       pgtype.Text    `json:"last_name"`
	Email           pgtype.Text    `json:"email"`
	Photo           pgtype.Text    `json:"photo"`
	Gender          pgtype.Text    `json:"gender"`
	Phone           pgtype.Text    `json:"phone"`
	Secondary_Phone pgtype.Text    `json:"secondary_phone"`
	Birth_date      string         `json:"birth_date"`
	Age             pgtype.Int2    `json:"age"`
	Zip_code        pgtype.Text    `json:"zip_code"`
	Residence       pgtype.Text    `json:"residence"`
	Coordinates     pgtype.Text    `json:"coordinates"`
	Marital_status  pgtype.Text    `json:"marital_status"`
	Height          pgtype.Numeric `json:"height"`
	Weight          pgtype.Numeric `json:"weight"`
	Shirt_size      pgtype.Numeric `json:"shirt_size"`
	Pant_size       pgtype.Numeric `json:"pant_size"`
	Shoe_size       pgtype.Numeric `json:"shoe_size"`
	Blood_type      pgtype.Text    `json:"blood_type"`
	Allergies       []string       `json:"allergies"`
	Code            pgtype.Text    `json:"code"`
	Personal_code   pgtype.Text    `json:"personal_code"`
	Rank            pgtype.Text    `json:"rank"`
	Promotion_date  string         `json:"promotion_date"`
	Promotion       pgtype.Text    `json:"promotion"`
	Condition       pgtype.Text    `json:"condition"`
	Division        pgtype.Text    `json:"division"`
	Profession      pgtype.Text    `json:"profession"`
	Institution     pgtype.Text    `json:"institution"`
	User_system     pgtype.Bool    `json:"user_system"`
}

type UserIdentificationDto struct {
	Id      pgtype.Int4 `json:"id"`
	Id_role pgtype.Int4 `json:"id_role"`
}

type UserDto struct {
	UserStatusDto
	UserProfileDto
	UserIdentificationDto
}

func MapFromEntity(userDto *user_domain.User) UserDto {
	user := UserDto{}

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

	user.UserIdentificationDto.Id = utils.ConvertToPgTypeInt4(id)
	user.UserIdentificationDto.Id_role = utils.ConvertToPgTypeInt4(id_role)
	user.UserProfileDto.User_name = utils.ConvertToPgTypeText(userDto.User_name)
	user.UserProfileDto.First_name = utils.ConvertToPgTypeText(userDto.First_name)
	user.UserProfileDto.Last_name = utils.ConvertToPgTypeText(userDto.Last_name)
	user.UserProfileDto.Email = utils.ConvertToPgTypeText(userDto.Email)
	user.UserProfileDto.Photo = utils.ConvertToPgTypeText(userDto.Photo)
	user.UserProfileDto.Gender = utils.ConvertToPgTypeText(userDto.Gender)
	user.UserProfileDto.Phone = utils.ConvertToPgTypeText(userDto.Phone)
	user.UserProfileDto.Secondary_Phone = utils.ConvertToPgTypeText(userDto.Secondary_Phone)
	user.UserProfileDto.Birth_date = userDto.Birth_date
	user.UserProfileDto.Age = utils.ConvertToPgTypeInt2(age)
	user.UserProfileDto.Residence = utils.ConvertToPgTypeText(userDto.Residence)
	user.UserProfileDto.Coordinates = utils.ConvertToPgTypeText(userDto.Coordinates)
	user.UserProfileDto.Marital_status = utils.ConvertToPgTypeText(userDto.Marital_status)
	user.UserProfileDto.Height = utils.ConvertToPgTypeNumeric(height)
	user.UserProfileDto.Weight = utils.ConvertToPgTypeNumeric(weight)
	user.UserProfileDto.Shirt_size = utils.ConvertToPgTypeNumeric(shirt)
	user.UserProfileDto.Pant_size = utils.ConvertToPgTypeNumeric(pant)
	user.UserProfileDto.Shoe_size = utils.ConvertToPgTypeNumeric(shoe)
	user.UserProfileDto.Blood_type = utils.ConvertToPgTypeText(userDto.Blood_type)
	user.UserProfileDto.Allergies = userDto.Allergies
	user.UserProfileDto.Code = utils.ConvertToPgTypeText(userDto.Code)
	user.UserProfileDto.Personal_code = utils.ConvertToPgTypeText(userDto.Personal_code)
	user.UserProfileDto.Rank = utils.ConvertToPgTypeText(userDto.Rank)
	user.Promotion_date = userDto.Promotion_date
	user.UserProfileDto.Promotion = utils.ConvertToPgTypeText(userDto.Promotion)
	user.UserProfileDto.Condition = utils.ConvertToPgTypeText(userDto.Condition)
	user.UserProfileDto.Division = utils.ConvertToPgTypeText(userDto.Division)
	user.UserProfileDto.Profession = utils.ConvertToPgTypeText(userDto.Profession)
	user.UserProfileDto.Institution = utils.ConvertToPgTypeText(userDto.Profession)
	user.UserProfileDto.User_system = utils.ConvertToPgTypeBool(userDto.User_system)
	user.UserProfileDto.Zip_code = utils.ConvertToPgTypeText(userDto.Zip_code)

	return user
}

func MapToEntity(user *UserDto) user_domain.User {
	userEntity := user_domain.User{}
	//var err error
	userEntity.Id = utils.ConvertFromInt4(user.Id)
	userEntity.UserIdentificationEntity.Id_role = utils.ConvertFromInt4(user.Id_role)
	userEntity.UserProfileEntity.User_name = utils.ConvertFromText(user.User_name)
	userEntity.UserProfileEntity.First_name = utils.ConvertFromText(user.First_name)
	userEntity.UserProfileEntity.Last_name = utils.ConvertFromText(user.Last_name)
	userEntity.UserProfileEntity.Email = utils.ConvertFromText(user.Email)
	userEntity.UserProfileEntity.Photo = utils.ConvertFromText(user.Photo)
	userEntity.UserProfileEntity.Gender = utils.ConvertFromText(user.Gender)
	userEntity.UserProfileEntity.Phone = utils.ConvertFromText(user.Phone)
	userEntity.UserProfileEntity.Secondary_Phone = utils.ConvertFromText(user.Secondary_Phone)
	userEntity.UserProfileEntity.Birth_date = user.Birth_date
	userEntity.UserProfileEntity.Age = utils.ConvertFromInt2(user.Age)
	userEntity.UserProfileEntity.Residence = utils.ConvertFromText(user.Residence)
	userEntity.UserProfileEntity.Coordinates = utils.ConvertFromText(user.Coordinates)
	userEntity.UserProfileEntity.Marital_status = utils.ConvertFromText(user.Marital_status)
	userEntity.UserProfileEntity.Height = utils.ConvertFromNumeric(user.Height)
	userEntity.UserProfileEntity.Weight = utils.ConvertFromNumeric(user.Weight)
	userEntity.UserProfileEntity.Shirt_size = utils.ConvertFromNumeric(user.Shirt_size)
	userEntity.UserProfileEntity.Pant_size = utils.ConvertFromNumeric(user.Pant_size)
	userEntity.UserProfileEntity.Shoe_size = utils.ConvertFromNumeric(user.Shoe_size)
	userEntity.UserProfileEntity.Blood_type = utils.ConvertFromText(user.Blood_type)
	userEntity.UserProfileEntity.Allergies = user.Allergies
	userEntity.UserProfileEntity.Code = utils.ConvertFromText(user.Code)
	userEntity.UserProfileEntity.Personal_code = utils.ConvertFromText(user.Personal_code)
	userEntity.UserProfileEntity.Rank = utils.ConvertFromText(user.Rank)
	userEntity.Promotion_date = user.Promotion_date
	userEntity.UserProfileEntity.Promotion = utils.ConvertFromText(user.Promotion)
	userEntity.UserProfileEntity.Condition = utils.ConvertFromText(user.Condition)
	userEntity.UserProfileEntity.Division = utils.ConvertFromText(user.Division)
	userEntity.UserProfileEntity.Profession = utils.ConvertFromText(user.Profession)
	userEntity.UserProfileEntity.Institution = utils.ConvertFromText(user.Profession)
	userEntity.UserProfileEntity.User_system = utils.ConvertFromBool(user.User_system)
	userEntity.UserProfileEntity.Zip_code = utils.ConvertFromText(user.Zip_code)

	return userEntity
}
