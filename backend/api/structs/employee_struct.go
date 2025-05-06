package structs

type Employee struct {
	ID       int    `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	BirthDate string `json:"birth_date"`
	HireDate  string `json:"hire_date"`
	Salary    float64 `json:"salary"`
	Department string `json:"department"`
	Position   string `json:"position"`
	Project    int `json:"project"`
}