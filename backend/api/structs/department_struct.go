package structs

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Manager int    `json:"manager"`
}