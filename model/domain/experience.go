package domain

type Experience struct {
	ID          string   `firestore:"id,omitempty"`
	CompanyName string   `firestore:"company_name"`
	Address     string   `firestore:"address"`
	Title       string   `firestore:"title"`
	EntryYear   string   `firestore:"entry_year"`
	OutYear     string   `firestore:"out_year"`
	JobDesk     []string `firestore:"job_desk"`
}
