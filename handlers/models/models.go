package models

type TrackInfo struct {
	ID         int
	TrackID    string
	TrackName  string
	ArtistName string
	ImageLink  string
}

type HandlerData struct {
	TokenCookie bool
	RedirectURL string
	IsPremium   bool
	DeviceName  string
	Tracks      []TrackInfo
}
