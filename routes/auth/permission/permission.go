package permission

import "encoding/json"

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
