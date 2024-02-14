package food

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rudyjcruz831/spoonAPI/model"
	"github.com/rudyjcruz831/spoonAPI/utils/errors"
)

func newGetRequest(endpoint string) *http.Request {
	// env :

	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s", endpoint), nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return req
}

type menuAPI struct {
	ApiKey string
}

func NewMenuAPI(apikey string) model.MenuAPI {
	return &menuAPI{ApiKey: apikey}
}

/*
GET https://api.spoonacular.com/food/menuItems/{id}

Response Headers:
	Content-Type: application/json

Parameters
	Name	Type	Example	Description
	id		number	424571	The menu item id.

{
    "id": 424571,
    "title": "Bacon King Burger",
    "restaurantChain": "Burger King",
    "nutrition": {
        "nutrients": [
            {
                "name": "Fat",
                "amount": 69,
                "unit": "g",
                "percentOfDailyNeeds": 30
            },
            {
                "name": "Protein",
                "amount": 57,
                "unit": "g",
                "percentOfDailyNeeds": 35
            },
            {
                "name": "Calories",
                "amount": 1040,
                "unit": "cal",
                "percentOfDailyNeeds": 50
            },
            {
                "name": "Carbohydrates",
                "amount": 48,
                "unit": "g",
                "percentOfDailyNeeds": 35
            }
        ],
        "caloricBreakdown": {
            "percentProtein": 35,
            "percentFat": 30,
            "percentCarbs": 35
        }
    },
    "badges": [],
    "breadcrumbs": [
        "burger",
        "main course",
        "food product category"
    ],
    "generatedText": null,
    "imageType": "png",
    "likes": 0,
    "servings": {
        "number": 1,
        "size": 2,
        "unit": "oz"
    },
    "price": null,
    "spoonacularScore": null
}


*/

func (a *menuAPI) GetMenuItemInfo(id int64) (*model.MenuItemInfo, *errors.FoodError) {

	s := fmt.Sprintf("api.spoonacular.com/food/menuItems/%d?apiKey=%s", id, a.ApiKey)
	req := newGetRequest(s)

	res, err := http.DefaultClient.Do(req)
	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		log.Println(res)
		foodErr := errors.NewBadRequestError("bad request from food api")
		return nil, foodErr
	}
	if err != nil {
		foodErr := errors.NewInternalServerError("Error" + err.Error())
		return nil, foodErr
	}
	data, errBody := ioutil.ReadAll(res.Body)
	if errBody != nil {
		log.Println("Error reading data from body")
		foodErr := errors.NewInternalServerError(errBody.Error())
		return nil, foodErr
	}
	res.Body.Close()

	// fmt.Println(data)

	var result model.MenuItemInfo
	if err := json.Unmarshal(data, &result); err != nil {
		log.Println("can not unmarshal JSON MenuItemInfo")
		foodErr := errors.NewInternalServerError("unmarshal error")
		return nil, foodErr
	}
	// fmt.Println(&result)

	return &result, nil
}

/*

   GET https://api.spoonacular.com/food/menuItems/search
   Header : Content-Type: application/json
   Parameters:
   Name					Type			Example	   		Description
   query					string			snickers	The search query.
   minCalories				number			50			The minimum amount of calories the menu item must have per serving.
   maxCalories				number			800			The maximum amount of calories the menu item can have per serving.
   minCarbs				number			10			The minimum amount of carbohydrates in grams the menu item must have per serving.
   maxCarbs				number			100			The maximum amount of carbohydrates in grams the menu item can have per serving.
   minProtein				number			10			The minimum amount of protein in grams the menu item must have per serving.
   maxProtein				number			100			The maximum amount of protein in grams the menu item can have per serving.
   minFat					number			1			The minimum amount of fat in grams the menu item must have per serving.
   maxFat					number			100			The maximum amount of fat in grams the menu item can have per serving.
   addMenuItemInformation	boolean			false		If set to true, you get more information about the menu items returned.
   offset					number			0			The offset number for paging (between 0 and 990).
   number					number			100			The number of expected results (between 1 and 10).


   https://api.spoonacular.com/recipes/716429/information?apiKey=YOUR-API-KEY&includeNutrition=true.

   {
       "menuItems": [
           {
               "id": 419357,
               "title": "Burger Sliders",
               "restaurantChain": "Hooters",
               "image": "https://images.spoonacular.com/file/wximages/419357-312x231.png",
               "imageType": "png",
               "servings": {
                   "number": 1,
                   "size": 2,
                   "unit": "oz"
               }
           },
           {
               "id": 424571,
               "title": "Bacon King Burger",
               "restaurantChain": "Burger King",
               "image": "https://images.spoonacular.com/file/wximages/424571-312x231.png",
               "imageType": "png",
               "servings": {
                   "number": 1,
                   "size": 2,
                   "unit": "oz"
               }
           }
       ],
       "totalMenuItems": 6749,
       "type": "menuItem",
       "offset": 0,
       "number": 2
   }
*/

func (a *menuAPI) GetMenuByName(itemName string) (*model.MenuItem, *errors.FoodError) {
	item := itemName
	log.Println("item: ", item)
	num := 3
	s := fmt.Sprintf("api.spoonacular.com/food/menuItems/search?apiKey=%s&query=%s&number=%v", a.ApiKey, item, num)
	// s := "api.spoonacular.com/food/menuItems/search?apiKey=a2854e9d3a2a49cdb5436e06a36752c8&query=burger&number=2"

	req := newGetRequest(s)

	res, err := http.DefaultClient.Do(req)
	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		fmt.Println(res)
		foodErr := errors.NewBadRequestError("bad request from food api")
		return nil, foodErr
	}
	if err != nil {
		foodErr := errors.NewInternalServerError("Error" + err.Error())
		return nil, foodErr
	}

	data, errBody := ioutil.ReadAll(res.Body)
	if errBody != nil {
		log.Println("Error reading data from body")
		foodErr := errors.NewInternalServerError(errBody.Error())
		return nil, foodErr
	}
	res.Body.Close()

	var result model.MenuItem
	if err := json.Unmarshal(data, &result); err != nil {
		log.Println("Can not unmarshal JSON MenuItemResponce")
		foodErr := errors.NewInternalServerError("unmarshal error")
		return nil, foodErr
	}

	log.Println("result: ", &result)
	return &result, nil
}
