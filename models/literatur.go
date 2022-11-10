package models

type Literatur struct {
	ID              int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title           string               `json:"title" form:"title" gorm:"type: varchar(255)"`
	UserID          int                  `json:"user_id" form:"user_id"`
	User            UsersProfileResponse `json:"user"`
	PublicationDate string               `json:"publicationdate" form:"publicationdate" gorm:"type: varchar(255)"`
	Pages           string               `json:"pages" form:"pages" gorm:"type: varchar(255)"`
	ISBN            string               `json:"isbn" form:"isbn" gorm:"type: varchar(255)"`
	Author          string               `json:"author" form:"author" gorm:"type: varchar(255)"`
	Attache         string               `json:"attache" form:"attache" gorm:"type: varchar(255)"`
}

type LiteraturResponse struct {
	ID              int                  `json:"id"`
	Title           string               `json:"title"`
	UserID          int                  `json:"-"`
	User            UsersProfileResponse `json:"user"`
	PublicationDate string               `json:"publicationdate"`
	Pages           string               `json:"pages"`
	ISBN            string               `json:"isbn"`
	Author          string               `json:"author"`
	Attache         string               `json:"attache"`
}

type LiteraturUserResponse struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	PublicationDate string `json:"publicationdate"`
	Pages           string `json:"pages"`
	ISBN            string `json:"isbn"`
	Author          string `json:"author"`
	Attache         string `json:"attache"`
	UserID          int    `json:"-"`
}

func (LiteraturResponse) TableName() string {
	return "literaturs"
}

func (LiteraturUserResponse) TableName() string {
	return "literaturs"
}
