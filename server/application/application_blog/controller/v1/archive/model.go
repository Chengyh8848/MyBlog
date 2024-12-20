package archive

type ArchiveBlog struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Day      string `json:"day"`
	Password string `json:"password"`
	Private  bool   `json:"private"`
}
