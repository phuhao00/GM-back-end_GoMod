package models
type Comment struct {
	ID     int    `gorm:"column:id;primary_key" json:"id"`
	GameID  int    `gorm:"column:gameId" json:"gameId"`
	Comment string `gorm:"column:comment" json:"comment"`
}
