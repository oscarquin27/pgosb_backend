package auth_routes

import "encoding/json"

const Users string = "users"
const Units string = "units"
const Stations string = "stations"

const Read string = "read"
const Write string = "write"
const Update string = "update"
const Delete string = "delete"
const Export string = "export"
const Print string = "print"

type Permissions struct {
	Write  bool `json:"add"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
	Export bool `json:"export"`
	Print  bool `json:"print"`
}

type UserPermissions map[string]Permissions

func UserPermissionFromJSONString(schema string) (UserPermissions, error) {
	var result UserPermissions

	err := json.Unmarshal([]byte(schema), &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

type LoginData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	SessionState string `json:"session_state"`
	ExpiresIn    int    `json:"expires_in"`
}

type LoginErr struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
