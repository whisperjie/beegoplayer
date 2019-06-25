package structs

type Album struct {
	Id int64
	Artist string
	ArticleId int64
}
type Artist struct {
	Id int64
	Name string
}
type QuerySong struct {
	Artist string `json:"ARTIST"`
	SongName string `json:"SONGNAME"`
	MusicRID string `json:"MUSICRID"`
}
type QueryResult struct {
	Code int `json:"code"`
	Data []QuerySong `json:"data"`
}
type Song struct {
	Name string
	Artist string
	AlbumId string
	SongId string
	Pic string
	SongTimeMinutes string
}


