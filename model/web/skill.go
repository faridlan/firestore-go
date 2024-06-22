package web

type SkillCreate struct {
	ID   string `firestore:"id,omitempty"`
	Name string `firestore:"name" validate:"required"`
}

type SkillResponse struct {
	ID   string `firestore:"id,omitempty"`
	Name string `firestore:"name"`
}
