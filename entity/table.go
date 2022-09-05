package entity

type AvList struct {
	Linkid      string `gorm:"column:linkid;primary_key" json:"linkid"`
	Title       string `gorm:"column:title" json:"title"`
	AvId        string `gorm:"column:av_id" json:"av_id"`
	ReleaseDate string `gorm:"column:release_date" json:"release_date"`
	Len         int    `gorm:"column:len" json:"len"`
	Director    string `gorm:"column:director" json:"director"`
	Studio      string `gorm:"column:studio" json:"studio"`
	Label       string `gorm:"column:label" json:"label"`
	Series      string `gorm:"column:series" json:"series"`
	Genre       string `gorm:"column:genre" json:"genre"`
	Stars       string `gorm:"column:stars" json:"stars"`
	DirectorUrl string `gorm:"column:director_url" json:"director_url"`
	StudioUrl   string `gorm:"column:studio_url" json:"studio_url"`
	LabelUrl    string `gorm:"column:label_url" json:"label_url"`
	SeriesUrl   string `gorm:"column:series_url" json:"series_url"`
	StarsUrl    string `gorm:"column:stars_url" json:"stars_url"`
	Bigimage    string `gorm:"column:bigimage" json:"bigimage"`
	ImageLen    int    `gorm:"column:image_len" json:"image_len"`
}
