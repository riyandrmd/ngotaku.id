package models

type Anime struct {
	Anime_Id   int     `gorm:"primaryKey;autoIncreament" json:"anime_id"`
	Title      string  `json:"title"`
	Status     string  `json:"status"`
	Rating     float32 `json:"rating"`
	Poster_url string  `json:"poster_url"`
	Trailer    string  `json:"trailer"`
	Genre_Id   int     `json:"genre_id"`
	Studio_Id  int     `json:"studio_id"`
}

type Genre struct {
	Genre_Id int     `gorm:"primaryKey;autoIncreament" json:"genre_id"`
	Genre    string  `json:"genre"`
	Anime    []Anime `gorm:"foreignKey:Genre_Id;refrences:Genre_Id; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"anime"`
}

type Studio struct {
	Studio_Id   int     `gorm:"primaryKey;autoIncreament" json:"studio_id"`
	Studio_Name string  `json:"Studio_Name"`
	Studio_loc  string  `json:"studio_loc"`
	Anime       []Anime `gorm:"foreignKey:Studio_Id;refrences:Studio_Id; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"anime"`
}

type AnimeFull struct {
	Anime_Id    int     `json:"anime_id"`
	Title       string  `json:"title"`
	Genre       string  `json:"genre"`
	Status      string  `json:"status"`
	Rating      float32 `json:"rating"`
	Studio_Name string  `json:"Studio_Name"`
	Poster_url  string  `json:"poster_url"`
	Trailer     string  `json:"trailer"`
}
