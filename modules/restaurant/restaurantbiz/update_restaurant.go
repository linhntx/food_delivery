package restaurantbiz

import (
	"context"
	"errors"

	"github.com/linhntx/food_delivery/modules/restaurant/restaurantmodel"
)

type UpdataRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	UpdateData(
		ctx 	context.Context,
		id		int,
		data	*restaurantmodel.RestaurantUpdate, 
	) error
}

type updateRestaurantBiz struct {
	store UpdataRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdataRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {

	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}