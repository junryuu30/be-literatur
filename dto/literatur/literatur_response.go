package literaturdto

type LiteraturResponse struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	UserID          int    `json:"-"`
	PublicationDate string `json:"publicationdate"`
	Pages           string `json:"pages"`
	ISBN            string `json:"isbn"`
	Author          string `json:"author"`
	Attache         string `json:"attache"`
}
