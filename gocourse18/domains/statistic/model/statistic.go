package model

type Statistics struct {
	ID   uint   `gorm:"primarykey;auto_increment" json:"id"`
	City string `gorm:"not null"`
	//AgeRange
	AverageTrips int64 `gorm:"default:0"`
}
