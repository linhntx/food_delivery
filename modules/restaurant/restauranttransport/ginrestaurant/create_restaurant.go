package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linhntx/food_delivery/component"
	"github.com/linhntx/food_delivery/modules/restaurant/restaurantbiz"
	"github.com/linhntx/food_delivery/modules/restaurant/restaurantmodel"
	"github.com/linhntx/food_delivery/modules/restaurant/restaurantstorage"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, data)
	}
}
