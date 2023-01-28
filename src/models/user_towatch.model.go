package models

type UserTowatch struct {
	ID                int             `gorm:"primaryKey" json:"id"`
	UserID            int             `json:"user_id"`
	Towatches         []Towatch       `gorm:"many2many:user_towatch_cards; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"towatches"`
	TowatchCategoryID int             `json:"-"`
	TowatchCategory   TowatchCategory `gorm:"foreignKey:TowatchCategoryID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"towatch_category"`
}
