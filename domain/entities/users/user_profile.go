package user_entity

import (
	"fdms/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserProfile struct {
	User_name       pgtype.Text    `json:"user_name" binding:"required"`
	First_name      pgtype.Text    `json:"first_name"`
	Last_name       pgtype.Text    `json:"last_name"`
	Email           pgtype.Text    `json:"email"`
	Photo           pgtype.Text    `json:"photo"`
	Gender          pgtype.Text    `json:"gender"`
	Phone           pgtype.Text    `json:"phone"`
	Secondary_Phone pgtype.Text    `json:"secondary_phone"`
	Birth_date      pgtype.Text     `json:"birth_date"`
	Age             pgtype.Int2    `json:"age"`
	Zip_code        pgtype.Text    `json:"zip_code"`
	Residence       pgtype.Text    `json:"residence"`
	Coordinates     pgtype.Text    `json:"coordinates"`
	Marital_status  pgtype.Text    `json:"marital_status"`
	Height          pgtype.Numeric `json:"height"`
	Weight          pgtype.Numeric `json:"weight"`
	Shirt_size      pgtype.Text `json:"shirt_size"`
	Pant_size       pgtype.Text `json:"pant_size"`
	Shoe_size       pgtype.Numeric `json:"shoe_size"`
	Blood_type      pgtype.Text    `json:"blood_type"`
	Allergies       []string      `json:"allergies"`

	Code           pgtype.Text `json:"code"`
	Personal_code  pgtype.Text `json:"personal_code"`
	Rank           pgtype.Text `json:"rank"`
	Promotion      pgtype.Text `json:"promotion"`
	Condition      pgtype.Text `json:"condition"`
	Division       pgtype.Text `json:"division"`
	Profession     pgtype.Text `json:"profession"`
	Institution    pgtype.Text `json:"institution"`
	User_system    pgtype.Bool `json:"user_system"`
	Skills		   []string	   `json:"skills"`
	State 		   pgtype.Text `json:"state"`
	Municipality   pgtype.Text `json:"municipality"`
	Parish         pgtype.Text `json:"parish"`
	Sector         pgtype.Text `json:"sector"`
	Community      pgtype.Text `json:"community"`
	Street         pgtype.Text `json:"street"`
	Beach          pgtype.Text `json:"beach"`
	Address        pgtype.Text `json:"address"`
	Legal_id       pgtype.Text `json:"legal_id"`
	// Skills          []string       `json:"skills"`

}

type UserProfileDto struct {
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
	Institution     string   `json:"institution"`
	User_system     bool     `json:"user_system"`
	Skills		   []string	 `json:"skills"`
	State 		    string   `json:"state"`
	Municipality    string   `json:"municipality"`
	Parish          string   `json:"parish"`
	Sector          string   `json:"sector"`
	Community       string   `json:"community"`
	Street          string   `json:"street"`
	Beach           string   `json:"beach"`
	Address         string   `json:"address"`
	Legal_id       	string	 `json:"legal_id"`

}

type Allergy struct {
	name pgtype.Text
}

func MapAFromArray(allergies []string) []Allergy {
	all := []Allergy{}

	if len(allergies) == 0 {
		return all
	}

	for _, value := range allergies {
		allergie := Allergy{
			name: utils.ConvertToPgTypeText(value),
		}
		all = append(all, allergie)
	}

	return all

}

func MapToArray(data []Allergy) []string {
	all := []string{}

	if len(data) == 0 {
		return all
	}

	for _, value := range data {
		allergie := utils.ConvertFromText(value.name)
		all = append(all, allergie)
	}

	return all
}
