package restaurantbiz

import (
	"context"
	"errors"
	"fmt"

	"github.com/linhntx/food_delivery/modules/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	fmt.Println("data", data)
	if data.Name == "" {
		return errors.New("restaurant name can not be blank")
	}

	err := biz.store.Create(ctx, data)

	return err
}
