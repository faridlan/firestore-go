package domain

type Profile struct {
	Name        string
	Description string
	Email       string
	MediaSocial MediaSocial
	About       string
}

type MediaSocial struct {
	LinkedIn  string
	Instagram string
	GitHub    string
}
