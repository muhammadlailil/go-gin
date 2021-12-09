package entity

//https://forum.golangbridge.org/t/format-gorm-struct-to-proper-json/4907
type Todos struct {
	ID        uint   `gorm:"primaryKey"json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
