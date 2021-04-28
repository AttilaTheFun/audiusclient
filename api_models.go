package audiusclient

type APIUser struct {
	ID string `json:"id"`

	Name       string `json:"name"`
	Handle     string `json:"handle"`
	Bio        string `json:"bio"`
	Location   string `json:"location"`
	IsVerified bool   `json:"is_verified"`

	ProfilePicture map[string]string `json:"profile_picture"`
	CoverPhoto     map[string]string `json:"cover_photo"`

	FollowerCount int `json:"follower_count"`
	FolloweeCount int `json:"followee_count"`

	AlbumCount    int `json:"album_count"`
	TrackCount    int `json:"track_count"`
	PlaylistCount int `json:"playlist_count"`
	RepostCount   int `json:"repost_count"`
}

type APITrack struct {
	ID string `json:"id"`

	Title          string `json:"title"`
	Description    string `json:"description"`
	Duration       int    `json:"duration"`
	Mood           string `json:"mood"`
	Genre          string `json:"genre"`
	Tags           string `json:"tags"`
	IsDownloadable bool   `json:"downloadable"`
	ReleaseDate    string `json:"release_date"`

	Artwork map[string]string `json:"artwork"`

	PlayCount     int `json:"play_count"`
	RepostCount   int `json:"repost_count"`
	FavoriteCount int `json:"favorite_count"`

	User APIUser `json:"user"`

	StreamURL string `json:"-"`
}

type APIPlaylist struct {
	ID string `json:"id"`

	PlaylistName string `json:"playlist_name"`
	Description  string `json:"description"`
	IsAlbum      bool   `json:"is_album"`

	Artwork map[string]string `json:"artwork"`

	TotalPlayCount int `json:"total_play_count"`
	RepostCount    int `json:"repost_count"`
	FavoriteCount  int `json:"favorite_count"`

	User APIUser `json:"user"`
}
