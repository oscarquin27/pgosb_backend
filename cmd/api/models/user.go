package api_models

import (
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserIdentificationJson struct {
	Id      string `json:"id"`
	Id_role string `json:"id_role"`
}

type UserProfileJson struct {
	User_name       string   `json:"user_name" binding:"required"`
	First_name      string   `json:"first_name"`
	Last_name       string   `json:"last_name"`
	Email           string   `json:"email"`
	Photo           string   `json:"photo"`
	Gender          string   `json:"gender"`
	Phone           string   `json:"phone"`
	Secondary_Phone string   `json:"secondary_phone"`
	Birth_date      string   `json:"birth_date"`
	Age             string   `json:"age"`
	Zip_code        string   `json:"zip_code"`
	Residence       string   `json:"residence"`
	Coordinates     string   `json:"coordinates"`
	Marital_status  string   `json:"marital_status"`
	Height          string   `json:"height"`
	Weight          string   `json:"weight"`
	Shirt_size      string   `json:"shirt_size"`
	Pant_size       string   `json:"pant_size"`
	Shoe_size       string   `json:"shoe_size"`
	Blood_type      string   `json:"blood_type"`
	Allergies       []string `json:"allergies"`
	Code            string   `json:"code"`
	Personal_code   string   `json:"personal_code"`
	Rank            string   `json:"rank"`
	Promotion       string   `json:"promotion"`
	Condition       string   `json:"condition"`
	Division        string   `json:"division"`
	Profession      string   `json:"profession"`
	Station         string   `json:"station"`
	User_system     bool     `json:"user_system"`
	Skills          []string `json:"skills"`
	State           string   `json:"state"`
	Municipality    string   `json:"municipality"`
	Parish          string   `json:"parish"`
	Sector          string   `json:"sector"`
	Community       string   `json:"community"`
	Street          string   `json:"street"`
	Beach           string   `json:"beach"`
	Address         string   `json:"address"`
	Legal_id        string   `json:"legal_id"`
	Role            string   `json:"role"`
}

type UserStatusJson struct {
	Status_user     string `json:"status_user"`
	Last_connection string `json:"last_connection"`
	Created_at      string `json:"created_at"`
	Updated_at      string `json:"updated_at"`
	Ip              string `json:"ip"`
}

type UserJson struct {
	UserIdentificationJson
	UserProfileJson
}

func (userDto *UserJson) ToModel() models.User {
	user := models.User{}

	layout := "02-01-2006" // Layout for both dates

	// Parse birth date
	birthDate, _ := time.Parse(layout, userDto.Birth_date)

	// Get current date
	currentDate := time.Now()

	// Calculate years (simplified)
	years := currentDate.Year() - birthDate.Year()

	id_role := utils.ParseInt(userDto.Id_role)

	//height := utils.ParseFloat(userDto.Height)

	weight := utils.ParseFloat(userDto.Weight)

	shoe := utils.ParseInt(userDto.Shoe_size)

	var altura pgtype.Numeric

	err := altura.Scan(userDto.Height)

	if err != nil {
		logger.Error().Err(err).Msg("Error escanenado altura")
	}

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
	user.UserProfile.Age = utils.ConvertToPgTypeInt2(years)
	user.UserProfile.Residence = utils.ConvertToPgTypeText(userDto.Residence)
	user.UserProfile.Coordinates = utils.ConvertToPgTypeText(userDto.Coordinates)
	user.UserProfile.Marital_status = utils.ConvertToPgTypeText(userDto.Marital_status)
	user.UserProfile.Height = altura
	user.UserProfile.Weight = &weight
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
	user.UserProfile.Station = utils.ConvertToPgTypeText(userDto.Station)
	user.UserProfile.User_system = utils.ConvertToPgTypeBool(userDto.User_system)
	user.UserProfile.Zip_code = utils.ConvertToPgTypeText(userDto.Zip_code)
	user.Skills = userDto.Skills
	user.UserProfile.State = utils.ConvertToPgTypeText(userDto.State)
	user.UserProfile.Municipality = utils.ConvertToPgTypeText(userDto.Municipality)
	user.UserProfile.Parish = utils.ConvertToPgTypeText(userDto.Parish)
	user.UserProfile.Sector = utils.ConvertToPgTypeText(userDto.Sector)
	user.UserProfile.Community = utils.ConvertToPgTypeText(userDto.Community)
	user.UserProfile.Street = utils.ConvertToPgTypeText(userDto.Street)
	user.UserProfile.Beach = utils.ConvertToPgTypeText(userDto.Beach)
	user.UserProfile.Address = utils.ConvertToPgTypeText(userDto.Address)
	user.UserProfile.Legal_id = utils.ConvertToPgTypeText(userDto.Legal_id)
	user.UserProfile.Role = utils.ConvertToPgTypeText(userDto.Role)
	return user
}

func (userDto *UserJson) ToModelSimple() models.UserSimple {
	user := models.UserSimple{}

	user.Id = userDto.Id
	user.Name = fmt.Sprintf("%s %s", userDto.First_name, userDto.Last_name)
	user.Legal_id = userDto.Legal_id
	user.PersonalCode = userDto.Code
	user.Rank = userDto.Rank
	user.User_name = userDto.User_name

	return user
}

func (userDto *UserJson) FromModel(user *models.User) *UserJson {
	f, err := user.Height.Float64Value()

	if err != nil {
		logger.Error().Err(err).Msg("Error conviertiendo")
	}

	value := f.Float64
	strValue := strconv.FormatFloat(value, 'f', 2, 64)

	peso := 0.00

	if user.Weight != nil {
		peso = *user.Weight
	}

	//var err error
	userDto.Id = strconv.FormatInt(user.Id, 10)
	userDto.UserIdentificationJson.Id_role = utils.ConvertFromInt4(user.Id_role)
	userDto.UserProfileJson.User_name = utils.ConvertFromText(user.User_name)
	userDto.UserProfileJson.First_name = utils.ConvertFromText(user.First_name)
	userDto.UserProfileJson.Last_name = utils.ConvertFromText(user.Last_name)
	userDto.UserProfileJson.Email = utils.ConvertFromText(user.Email)
	userDto.UserProfileJson.Photo = utils.ConvertFromText(user.Photo)
	userDto.UserProfileJson.Gender = utils.ConvertFromText(user.Gender)
	userDto.UserProfileJson.Phone = utils.ConvertFromText(user.Phone)
	userDto.UserProfileJson.Secondary_Phone = utils.ConvertFromText(user.Secondary_Phone)
	userDto.UserProfileJson.Birth_date = utils.ConvertFromText(user.Birth_date)
	userDto.UserProfileJson.Age = utils.ConvertFromInt2(user.Age)
	userDto.UserProfileJson.Residence = utils.ConvertFromText(user.Residence)
	userDto.UserProfileJson.Coordinates = utils.ConvertFromText(user.Coordinates)
	userDto.UserProfileJson.Marital_status = utils.ConvertFromText(user.Marital_status)
	userDto.UserProfileJson.Height = strValue
	userDto.UserProfileJson.Weight = strconv.FormatFloat(peso, 'f', 2, 32)
	userDto.UserProfileJson.Shirt_size = utils.ConvertFromText(user.Shirt_size)
	userDto.UserProfileJson.Pant_size = utils.ConvertFromText(user.Pant_size)
	userDto.UserProfileJson.Shoe_size = utils.ConvertFromNumeric(user.Shoe_size)
	userDto.UserProfileJson.Blood_type = utils.ConvertFromText(user.Blood_type)
	userDto.UserProfileJson.Allergies = user.Allergies
	userDto.UserProfileJson.Code = utils.ConvertFromText(user.Code)
	userDto.UserProfileJson.Personal_code = utils.ConvertFromText(user.Personal_code)
	userDto.UserProfileJson.Rank = utils.ConvertFromText(user.Rank)
	userDto.UserProfileJson.Promotion = utils.ConvertFromText(user.Promotion)
	userDto.UserProfileJson.Condition = utils.ConvertFromText(user.Condition)
	userDto.UserProfileJson.Division = utils.ConvertFromText(user.Division)
	userDto.UserProfileJson.Profession = utils.ConvertFromText(user.Profession)
	userDto.UserProfileJson.Station = utils.ConvertFromText(user.Station)
	userDto.UserProfileJson.User_system = utils.ConvertFromBool(user.User_system)
	userDto.UserProfileJson.Zip_code = utils.ConvertFromText(user.Zip_code)
	userDto.UserProfileJson.Skills = user.Skills
	userDto.UserProfileJson.State = utils.ConvertFromText(user.State)
	userDto.UserProfileJson.Municipality = utils.ConvertFromText(user.Municipality)
	userDto.UserProfileJson.Parish = utils.ConvertFromText(user.Parish)
	userDto.UserProfileJson.Sector = utils.ConvertFromText(user.Sector)
	userDto.UserProfileJson.Community = utils.ConvertFromText(user.Community)
	userDto.UserProfileJson.Street = utils.ConvertFromText(user.Street)
	userDto.UserProfileJson.Beach = utils.ConvertFromText(user.Beach)
	userDto.UserProfileJson.Address = utils.ConvertFromText(user.Address)
	userDto.UserProfileJson.Legal_id = utils.ConvertFromText(user.Legal_id)
	userDto.UserProfileJson.Role = utils.ConvertFromText(user.Role)

	return userDto
}

func ModelToUserJson(user *models.User) *UserJson {
	f, err := user.Height.Float64Value()

	if err != nil {
		logger.Error().Err(err).Msg("Error conviertiendo")
	}

	value := f.Float64
	strValue := strconv.FormatFloat(value, 'f', 2, 64)

	peso := 0.00

	if user.Weight != nil {
		peso = *user.Weight
	}

	userDto := UserJson{}
	//var err error
	userDto.Id = strconv.FormatInt(user.Id, 10)
	userDto.UserIdentificationJson.Id_role = utils.ConvertFromInt4(user.Id_role)
	userDto.UserProfileJson.User_name = utils.ConvertFromText(user.User_name)
	userDto.UserProfileJson.First_name = utils.ConvertFromText(user.First_name)
	userDto.UserProfileJson.Last_name = utils.ConvertFromText(user.Last_name)
	userDto.UserProfileJson.Email = utils.ConvertFromText(user.Email)
	userDto.UserProfileJson.Photo = utils.ConvertFromText(user.Photo)
	userDto.UserProfileJson.Gender = utils.ConvertFromText(user.Gender)
	userDto.UserProfileJson.Phone = utils.ConvertFromText(user.Phone)
	userDto.UserProfileJson.Secondary_Phone = utils.ConvertFromText(user.Secondary_Phone)
	userDto.UserProfileJson.Birth_date = utils.ConvertFromText(user.Birth_date)
	userDto.UserProfileJson.Age = utils.ConvertFromInt2(user.Age)
	userDto.UserProfileJson.Residence = utils.ConvertFromText(user.Residence)
	userDto.UserProfileJson.Coordinates = utils.ConvertFromText(user.Coordinates)
	userDto.UserProfileJson.Marital_status = utils.ConvertFromText(user.Marital_status)
	userDto.UserProfileJson.Height = strValue
	userDto.UserProfileJson.Weight = strconv.FormatFloat(peso, 'f', 2, 32)
	userDto.UserProfileJson.Shirt_size = utils.ConvertFromText(user.Shirt_size)
	userDto.UserProfileJson.Pant_size = utils.ConvertFromText(user.Pant_size)
	userDto.UserProfileJson.Shoe_size = utils.ConvertFromNumeric(user.Shoe_size)
	userDto.UserProfileJson.Blood_type = utils.ConvertFromText(user.Blood_type)
	userDto.UserProfileJson.Allergies = user.Allergies
	userDto.UserProfileJson.Code = utils.ConvertFromText(user.Code)
	userDto.UserProfileJson.Personal_code = utils.ConvertFromText(user.Personal_code)
	userDto.UserProfileJson.Rank = utils.ConvertFromText(user.Rank)
	userDto.UserProfileJson.Promotion = utils.ConvertFromText(user.Promotion)
	userDto.UserProfileJson.Condition = utils.ConvertFromText(user.Condition)
	userDto.UserProfileJson.Division = utils.ConvertFromText(user.Division)
	userDto.UserProfileJson.Profession = utils.ConvertFromText(user.Profession)
	userDto.UserProfileJson.Station = utils.ConvertFromText(user.Station)
	userDto.UserProfileJson.User_system = utils.ConvertFromBool(user.User_system)
	userDto.UserProfileJson.Zip_code = utils.ConvertFromText(user.Zip_code)
	userDto.UserProfileJson.Skills = user.Skills
	userDto.UserProfileJson.State = utils.ConvertFromText(user.State)
	userDto.UserProfileJson.Municipality = utils.ConvertFromText(user.Municipality)
	userDto.UserProfileJson.Parish = utils.ConvertFromText(user.Parish)
	userDto.UserProfileJson.Sector = utils.ConvertFromText(user.Sector)
	userDto.UserProfileJson.Community = utils.ConvertFromText(user.Community)
	userDto.UserProfileJson.Street = utils.ConvertFromText(user.Street)
	userDto.UserProfileJson.Beach = utils.ConvertFromText(user.Beach)
	userDto.UserProfileJson.Address = utils.ConvertFromText(user.Address)
	userDto.UserProfileJson.Legal_id = utils.ConvertFromText(user.Legal_id)
	userDto.UserProfileJson.Role = utils.ConvertFromText(user.Role)

	return &userDto
}
