package user_messages

type User struct {
	UserName       string   `json:"user_name"`
	Email          string   `json:"email"`
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	Phone          string   `json:"phone"`
	SecondaryPhone string   `json:"secondary_phone"`
	ZipCode        string   `json:"zip_code"`
	MaritalStatus  string   `json:"marital_status"`
	BirthDate      string   `json:"birth_date"` // Consider using time.Time for date handling
	Gender         string   `json:"gender"`
	Sector         string   `json:"sector"`
	Community      string   `json:"community"`
	Street         string   `json:"street"`
	Beach          string   `json:"beach"`
	Address        string   `json:"address"`
	State          string   `json:"state"`
	Municipality   string   `json:"municipality"`
	Parish         string   `json:"parish"`
	Height         string   `json:"height"` // Consider using float64 for numerical values
	Weight         string   `json:"weight"` // Consider using float64 for numerical values
	BloodType      string   `json:"blood_type"`
	ShirtSize      string   `json:"shirt_size"`
	PantSize       string   `json:"pant_size"`
	ShoeSize       string   `json:"shoe_size"`
	Skills         []string `json:"skills"`
	Allergies      []string `json:"allergies"`
	Code           string   `json:"code"`
	InDate         string   `json:"in_date"`  // Consider using time.Time for date handling
	OutDate        string   `json:"out_date"` // Consider using time.Time for date handling
	Rank           string   `json:"rank"`
	Institution    string   `json:"institution"`
	UserSystem     bool     `json:"user_system"`
	Profesion      string   `json:"profesion"`
	IDRol          string   `json:"id_rol"`
	Division       string   `json:"division"`
}
