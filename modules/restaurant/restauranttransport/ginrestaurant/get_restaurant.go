package ginrestaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/linhntx/food_delivery/common"
	"github.com/linhntx/food_delivery/component"
	"github.com/linhntx/food_delivery/modules/restaurant/restaurantbiz"
	"github.com/linhntx/food_delivery/modules/restaurant/restaurantstorage"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
			// c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			// return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			// c.JSON(400, err)
			// return
			panic(err)
		}
		
		c.JSON(http.StatusOK, common.SimpleSuccessRespone(data))
	}	
}