package user_entity

import "github.com/jackc/pgx/v5/pgtype"

type UserProfile struct {
	User_name  pgtype.Text `json:"user_name" binding:"required"`
	First_name pgtype.Text `json:"first_name" binding:"required"`
	Last_name  pgtype.Text `json:"last_name" binding:"required"`
	Email      pgtype.Text `json:"email" binding:"required"`
	Photo 	   pgtype.Text `json:"photo"`
	Gender     pgtype.Text `json:"gender"`
	Phone	   pgtype.Text `json:"phone"`
	Secondary_Phone pgtype.Text `json:"secondary_phone"`
	Birth_date pgtype.Text `json:"birth_date"`
	Age 	   pgtype.Int2 `json:"age"`
	Zip_code   pgtype.Text `json:"zip_code"`
	Residence  pgtype.Text `json:"residence"`
	Coordinates pgtype.Text `json:"coordinates"`
	Marital_status pgtype.Text `json:"marital_status"`
	Height 		pgtype.Numeric `json:"height"`
	Weight 		pgtype.Numeric `json:"weight"`
	Shirt_size  pgtype.Numeric `json:"shirt_size"`
	Pant_size 	pgtype.Numeric	`json:"pant_size"`
	Shoe_size	pgtype.Numeric `json:"shoe_size"`
	Blood_type 	pgtype.Text `json:"blood_type"`
	Allergies 	[]Allergy    `json:"allergies"`
	Code 		pgtype.Text `json:"code"`
	Personal_code pgtype.Text `json:"personal_code"`
	Rank 		pgtype.Text `json:"rank"`
	Promotion_date pgtype.Date `json:"promotion_date"`
	Promotion   pgtype.Text `json:"promotion"`
	Condition pgtype.Text `json:"condition"`
	Division pgtype.Text `json:"division"`
	Profession pgtype.Text `json:"profession"`
	Institution pgtype.Text `json:"institution"`
	User_system pgtype.Bool `json:"user_system"`
 }


type Allergy struct {
	name pgtype.Text
}