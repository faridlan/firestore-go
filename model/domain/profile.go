package domain

type Profile struct {
	ID          string      `firestore:"id,omitempty"`
	Name        string      `firestore:"name"`
	Description string      `firestore:"description"`
	Email       string      `firestore:"email"`
	MediaSocial MediaSocial `firestore:"media_social"`
	About       string      `firestore:"about"`
}

type MediaSocial struct {
	LinkedIn  string `firestore:"linkedIn"`
	Instagram string `firestore:"instagram"`
	GitHub    string `firestore:"github"`
}
