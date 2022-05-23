package ginrestaurant

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linhntx/food_delivery/modules/restaurant/restaurantbiz"
	"github.com/linhntx/food_delivery/modules/restaurant/restaurantmodel"
	"github.com/linhntx/food_delivery/modules/restaurant/restaurantstorage"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		fmt.Println("DB main", &data)
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStore(db)
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
