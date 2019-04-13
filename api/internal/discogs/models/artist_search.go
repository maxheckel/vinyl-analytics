package models

type ArtistSearch struct {
	Thumb          string   `json:"thumb"`
	Title          string   `json:"title"`
	MasterURL      *string  `json:"master_url"`
	URI            string   `json:"uri"`
	CoverImage     string   `json:"cover_image"`
	ResourceURL    string   `json:"resource_url"`
	MasterId       *int     `json:"master_id"`
	UserData       UserData `json:"user_data"`
	ID             uint     `json:"id"`
}
