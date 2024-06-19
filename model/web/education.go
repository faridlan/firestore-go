package web

type EducationResponse struct {
	ID        string   `json:"id,omitempty"`
	EduName   string   `json:"edu_name,omitempty"`
	Address   string   `json:"address,omitempty"`
	EntryYear string   `json:"entry_year,omitempty"`
	OutYear   string   `json:"out_year,omitempty"`
	Title     string   `json:"title,omitempty"`
	Achiev    []string `json:"achiev,omitempty"`
}

type EducationCreate struct {
	EduName   string   `json:"edu_name,omitempty"`
	Address   string   `json:"address,omitempty"`
	EntryYear string   `json:"entry_year,omitempty"`
	OutYear   string   `json:"out_year,omitempty"`
	Title     string   `json:"title,omitempty"`
	Achiev    []string `json:"achiev,omitempty"`
}
