package models

type ListenResponse struct {
	Metadata   Metadata `json:"metadata"`
	CostTime   float64  `json:"cost_time"`
	Status     Status   `json:"status"`
	ResultType int      `json:"result_type"`
}
type Genres struct {
	Name string `json:"name"`
}
type ExternalIds struct {
}
type Artists struct {
	Name string `json:"name"`
}
type AlbumName struct {
	Name string `json:"name"`
}
type ExternalMetadata struct {
}
type Music struct {
	Label            string           `json:"label"`
	PlayOffsetMs     int              `json:"play_offset_ms"`
	Genres           []Genres         `json:"genres"`
	ExternalIds      ExternalIds      `json:"external_ids"`
	ResultFrom       int              `json:"result_from"`
	Artists          []Artists        `json:"artists"`
	Acrid            string           `json:"acrid"`
	Title            string           `json:"title"`
	DurationMs       int              `json:"duration_ms"`
	AlbumName        AlbumName        `json:"album"`
	Score            int              `json:"score"`
	ExternalMetadata ExternalMetadata `json:"external_metadata"`
	ReleaseDate      string           `json:"release_date"`
}
type Metadata struct {
	TimestampUtc string  `json:"timestamp_utc"`
	Music        []Music `json:"music"`
}
type Status struct {
	Msg     string `json:"msg"`
	Version string `json:"version"`
	Code    int    `json:"code"`
}
