package ds

type Spectrum struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `gorm:"type:text" json:"name"`
	Len         float64 `gorm:"not null"  json:"len"`
	Freq        float64 `gorm:"not null"  json:"freq"`
	Description string  `gorm:"type:text" json:"description"`
	Image       string  `json:"image_url" gorm:"type:varchar(1000);default:'http://172.21.0.3:9000/spectrumbucket/default.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=A4WCECNHB9ST6BB1A0AY%2F20231024%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20231024T120117Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiJBNFdDRUNOSEI5U1Q2QkIxQTBBWSIsImV4cCI6MTY5ODE1MjI2NiwicGFyZW50IjoibWluaW8ifQ.qGdue-B-8G6glv9OYVgSSBFbNILWcZ69u6zJ0u32YNxYZAGdBWd-N4E6xCiCmEBe9BlA5RfGtMEVeB9uOVT8NA&X-Amz-SignedHeaders=host&versionId=197e7072-8b67-49a5-ac09-94bd70eacd08&X-Amz-Signature=13a719da1d8fcad3654e0f15d38e5efb09e191af614325152d696cd831a75dbd'"`
	IsDelete    bool    `json:"is_delete"`
}
