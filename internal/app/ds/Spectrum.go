package ds

type Spectrum struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `gorm:"type:text" json:"name"`
	Len         float64 `gorm:"not null"  json:"len"`
	Freq        float64 `gorm:"not null"  json:"freq"`
	Description string  `gorm:"type:text" json:"description"`
	Image       string  `json:"image_url" gorm:"type:varchar(1000);default:'http://localhost:9000/spectrumbucket/default.jpeg'"`
	IsDelete    bool    `json:"is_delete"`
}
