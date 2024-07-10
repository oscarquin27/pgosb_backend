package layout_entity

type Layout struct {
	Id           int    `json:"id" db:"id"`
	Column_name  string `json:"column_name"`
	Display_name string `json:"display_name"`
	Group_name   string `json:"group_name"`
	Entity_name  string `json:"entity_name"`
	Visibility   bool   `json:"visibility"`
	Type         string `json:"type"`
}