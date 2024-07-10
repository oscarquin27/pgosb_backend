package user_entity

type UserProfile struct {
	User_name       string    `json:"user_name" binding:"required"`
	First_name      string    `json:"first_name"`
	Last_name       string    `json:"last_name"`
	Email           string    `json:"email"`
	Photo           string    `json:"photo"`
	Gender          string    `json:"gender"`
	Phone           string    `json:"phone"`
	Secondary_Phone string    `json:"secondary_phone"`
	Birth_date      string    `json:"birth_date"`
	Age             int       `json:"age" binding:"exists"`
	Zip_code        string    `json:"zip_code"`
	Residence       string    `json:"residence"`
	Coordinates     string    `json:"coordinates"`
	Marital_status  string    `json:"marital_status"`
	Height          float32   `json:"height" binding:"exists"`
	Weight          float32   `json:"weight" binding:"exists"`
	Shirt_size      string    `json:"shirt_size"`
	Pant_size       string    `json:"pant_size"`
	Shoe_size       int       `json:"shoe_size" binding:"exists"`
	Blood_type      string    `json:"blood_type"`
	Allergies       []Allergy `json:"allergies"`

	Code           string `json:"code"`
	Personal_code  string `json:"personal_code"`
	Rank           string `json:"rank"`
	Promotion_date string `json:"promotion_date"`
	Promotion      string `json:"promotion"`
	Condition      string `json:"condition"`
	Division       string `json:"division"`
	Profession     string `json:"profession"`
	Institution    string `json:"institution"`
	User_system    bool   `json:"user_system"`

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
	Promotion_date  string   `json:"promotion_date"`
	Promotion       string   `json:"promotion"`
	Condition       string   `json:"condition"`
	Division        string   `json:"division"`
	Profession      string   `json:"profession"`
	Institution     string   `json:"institution"`
	User_system     bool     `json:"user_system"`

	// Skills          []string `json:"skills"`

}

type Allergy struct {
	name string
}

func MapAFromArray(allergies []string) []Allergy {
	all := []Allergy{}

	if len(allergies) == 0 {
		return all
	}

	for _, value := range allergies {
		allergie := Allergy{
			name: value,
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
		allergie := value.name
		all = append(all, allergie)
	}

	return all
}
