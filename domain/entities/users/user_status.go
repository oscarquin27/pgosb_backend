package user_domain

type UserStatusEntity struct {
	Status_user     string `json:"status_user"`
	Last_connection string `json:"last_connection"`
	Created_at      string `json:"created_at"`
	Updated_at      string `json:"updated_at"`
	Ip              string `json:"ip"`
}
