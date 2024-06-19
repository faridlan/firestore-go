package domain

type Education struct {
	ID        string   `firestore:"id,omitempty"`
	EduName   string   `firestore:"edu_name"`
	Address   string   `firestore:"address"`
	EntryYear string   `firestore:"entry_year"`
	OutYear   string   `firestore:"out_year"`
	Title     string   `firestore:"title"`
	Achiev    []string `firestore:"achiev"`
}
