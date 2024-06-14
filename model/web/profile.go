package web

type Profile struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty" validate:"required"`
	Description string      `json:"description,omitempty" validate:"required"`
	Email       string      `json:"email,omitempty" validate:"required,email"`
	MediaSocial MediaSocial `json:"media_social,omitempty" validate:"required"`
	About       string      `json:"about,omitempty" validate:"required"`
}

type MediaSocial struct {
	LinkedIn  string `json:"linked_in,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	GitHub    string `json:"git_hub,omitempty"`
}
