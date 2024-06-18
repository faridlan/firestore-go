package web

type ExperienceCreate struct {
	ID          string   `json:"id,omitempty"`
	CompanyName string   `json:"company_name,omitempty"`
	Address     string   `json:"address,omitempty"`
	Title       string   `json:"title,omitempty"`
	EntryYear   string   `json:"entry_year,omitempty"`
	OutYear     string   `json:"out_year,omitempty"`
	JobDesk     []string `json:"job_desk,omitempty"`
}

type ExperienceResponse struct {
	ID          string   `json:"id,omitempty"`
	CompanyName string   `json:"company_name,omitempty"`
	Address     string   `json:"address,omitempty"`
	Title       string   `json:"title,omitempty"`
	EntryYear   string   `json:"entry_year,omitempty"`
	OutYear     string   `json:"out_year,omitempty"`
	JobDesk     []string `json:"job_desk,omitempty"`
}
