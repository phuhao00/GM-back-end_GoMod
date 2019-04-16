package models
type Game struct {
	ID               int    `gorm:"column:id;primary_key" json:"id"`
	Name             string `gorm:"column:name" json:"name"`
	Price            int    `gorm:"column:price" json:"price"`
	CommentID        int    `gorm:"column:commentId" json:"commentId"`
	DownloadQuantity int    `gorm:"column:downloadQuantity" json:"downloadQuantity"`
	Score            int    `gorm:"column:score" json:"score"`
	URL              string `gorm:"column:url" json:"url"`
	SupplierID       int    `gorm:"column:supplierId" json:"supplierId"`
}
