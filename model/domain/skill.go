package domain

type Skill struct {
	ID   string `firestore:"id,omitempty"`
	Name string `firestore:"name"`
}
