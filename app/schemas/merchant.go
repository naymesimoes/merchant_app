package schemas

type Merchant struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	LastName     string `json:"last_name"`
	SignupedAt   string `json:"signuped_at"`
	Cpf          string `json:"cpf"`
	PhoneNumber1 string `json:"phone_number1"`
	PhoneNumber2 string `json:"phone_number2"`
	Email        string `json:"email"`
}
