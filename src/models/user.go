package models

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorUserNotFound   = errors.New("usuario no encontrado")
	ErrorUserNotCreated = errors.New("usuario no creado")
	ErrorUserNotUpdated = errors.New("el usuario no pudo ser actualizado")
	ErrorUserNotDeleted = errors.New("el usuario no pudo ser eliminado")
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
	Birth_date      pgtype.Text    `json:"birth_date"`
	Age             pgtype.Int2    `json:"age"`
	Zip_code        pgtype.Text    `json:"zip_code"`
	Residence       pgtype.Text    `json:"residence"`
	Coordinates     pgtype.Text    `json:"coordinates"`
	Marital_status  pgtype.Text    `json:"marital_status"`
	Height          pgtype.Numeric `json:"height"`
	Weight          pgtype.Numeric `json:"weight"`
	Shirt_size      pgtype.Text    `json:"shirt_size"`
	Pant_size       pgtype.Text    `json:"pant_size"`
	Shoe_size       pgtype.Numeric `json:"shoe_size"`
	Blood_type      pgtype.Text    `json:"blood_type"`
	Allergies       []string       `json:"allergies"`

	Code          pgtype.Text `json:"code"`
	Personal_code pgtype.Text `json:"personal_code"`
	Rank          pgtype.Text `json:"rank"`
	Promotion     pgtype.Text `json:"promotion"`
	Condition     pgtype.Text `json:"condition"`
	Division      pgtype.Text `json:"division"`
	Profession    pgtype.Text `json:"profession"`
	Institution   pgtype.Text `json:"institution"`
	User_system   pgtype.Bool `json:"user_system"`
	Skills        []string    `json:"skills"`
	State         pgtype.Text `json:"state"`
	Municipality  pgtype.Text `json:"municipality"`
	Parish        pgtype.Text `json:"parish"`
	Sector        pgtype.Text `json:"sector"`
	Community     pgtype.Text `json:"community"`
	Street        pgtype.Text `json:"street"`
	Beach         pgtype.Text `json:"beach"`
	Address       pgtype.Text `json:"address"`
	Legal_id      pgtype.Text `json:"legal_id"`
	// Skills          []string       `json:"skills"`
}
type UserIdentification struct {
	Id      int64       `json:"id"`
	Id_role pgtype.Int4 `json:"id_role"`
}

type UserStatus struct {
	Status_user     pgtype.Text      `json:"status_user"`
	Last_connection pgtype.Timestamp `json:"last_connection"`
	Created_at      pgtype.Timestamp `json:"created_at"`
	Updated_at      pgtype.Timestamp `json:"updated_at"`
	Ip              pgtype.Text      `json:"ip"`
}

type User struct {
	UserIdentification
	UserProfile
}