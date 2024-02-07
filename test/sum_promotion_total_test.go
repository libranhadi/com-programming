package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CartItem struct {
	ItemName string  `json:"item_name"`
	Price    float32 `json`
	Qty      int     `json`
}

type ShoppingCart struct {
	Items     []CartItem `json:"items"`
	PromoCode string     `json:"promo_code"`
}

func TestSumPromotionTotalWithoutPromo(t *testing.T) {
	cartWithoutPromo := `{
		"items": [
		  {"item_name" : "lays", "price" : 14000, "qty" : 3},
		  {"item_name" : "banana", "price" : 9000, "qty" : 2},
		  {"item_name" : "chair", "price" : 100000, "qty" : 2},
		  {"item_name" : "minerale", "price" : 4000, "qty" : 3}
		]
	  }`
	totalCartWithoutPromo := SumPromotionTotal(cartWithoutPromo)
	assert.Equal(t, float32(0), totalCartWithoutPromo)

}

func TestSumPromotionTotalWithPromo(t *testing.T) {
	cartWithPromo := `{
		"items": [
		  {"item_name" : "lays", "price" : 14000, "qty" : 3},
		  {"item_name" : "banana", "price" : 9000, "qty" : 2},
		  {"item_name" : "chair", "price" : 100000, "qty" : 2},
		  {"item_name" : "minerale", "price" : 4000, "qty" : 3}
		],
		"promo_code" : "BUY NOW DISC 10%"
	  }`
	result := SumPromotionTotal(cartWithPromo)
	assert.Equal(t, float32(11400), result)

}

/*
*
  - @Author:
  - @Date: 2024-02-07 19:21:19
  - @Desc: Sum Promotion total
  - @param: cart jsonstring and promotion code
  - @Case:
    1. promotion 10%
    2. promotion active when any promotion code
    3. promotion active if price equal or above 10000
    4. if promotion total less than 100000
*/
func SumPromotionTotal(jsonString string) float32 {
	shoppingCart := convertStringJson(jsonString)
	var result float32
	limit := 10000
	for _, v := range shoppingCart.Items {
		if shoppingCart.PromoCode == "" {
			continue
		}
		if result >= float32(limit) {
			continue
		}
		if v.Price < float32(10000) {
			continue
		}

		result += float32(v.Price * (10.0 / 100))
	}
	return result
}

/**
 * @Author:
 * @Date: 2024-02-07 19:21:30
 * @Desc: format string json from frontend
 */
func convertStringJson(items string) ShoppingCart {
	var shoppingCart ShoppingCart
	if err := json.Unmarshal([]byte(items), &shoppingCart); err != nil {
		fmt.Println(err.Error())
	}

	return shoppingCart
}
