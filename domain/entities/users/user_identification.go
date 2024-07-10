package user_entity

type UserIdentification struct {
	Id      int `json:"id"`
	Id_role int `json:"id_role"`
}

type UserIdentificationDto struct {
	Id      string `json:"id"`
	Id_role string `json:"id_role"`
}
