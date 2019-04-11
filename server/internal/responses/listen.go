package responses

type Listen struct {
	AlbumFound  bool   `json:"album_found"`
	AlbumTitle  string `json:"album_title"`
	Image       string `json:"image"`
	SongTitle   string `json:"song_title"`
	ArtistTitle string `json:"artist_title"`
	DiscogsId   uint   `json:"discogs_id"`
}
