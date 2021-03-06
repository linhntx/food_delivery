package restaurantmodel

import "github.com/linhntx/food_delivery/common"

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Addr            string `json:"address" gorm:"column:addr;"`
	Logo 			*common.Image `json:"logo" gorm:"column:logo;"`
	Cover 			*common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name 	*string `json:"name" gorm:"column:name;"`
	Addr 	*string `json:"address" gorm:"column:addr;"`
	Logo 	*common.Image `json:"logo" gorm:"column:logo;"`
	Cover 	*common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Id   	int    `json:"id" gorm:"column:id;"`
	Name 	string `json:"name" gorm:"column:name;"`
	Addr 	string `json:"address" gorm:"column:addr;"`
	Logo 	*common.Image `json:"logo" gorm:"column:logo;"`
	Cover 	*common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}
