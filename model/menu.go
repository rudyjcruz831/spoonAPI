package model

// TODO: make sure that all this structs are correct I believe the error is here
type MenuItemInfo struct {
	ID               int64     `json:"id"`
	Title            string    `json:"title"`
	RestaurantChain  string    `json:"restaurantChain"`
	Nutrition        Nutrition `josn:"nutrition"`
	Badges           []string  `json:"badges"` // maybe this is causing the error could be not string type
	BreadCrumbs      []string  `json:"breadcrumbs"`
	GeneratedText    string    `json:"generatedText"`
	InmageType       string    `json:"imageType"`
	Likes            float64   `json:"likes"`
	Servings         Servings  `json:"serving"`
	Price            float32   `json:"price"`
	SpoonacularScort float64   `json:"spoonacularScore"`
}

type Nutrition struct {
	Nutrients        []Nutrient       `json:"nutrients"`
	CaloricBreakdown CaloricBreakdown `json:"caloricBreakdown"`
}

type Nutrient struct {
	Name                string  `json:"name"`
	Amount              float64 `json:"amount"`
	Unit                string  `json:"unit"`
	PercentOfDailyNeeds float64 `json:"percentOfDailyNeeds"`
}

type CaloricBreakdown struct {
	PercentProtein float64 `json:"percentProtein"`
	PercnetFat     float64 `json:"percentFat"`
	PercentCarbs   float64 `json:"percentCarbs"`
}

type MenuItem struct {
	MenuItem       []MenuItems `json:"menuItems"`
	TotalMenuItems int64       `json:"totalMenuItems"`
	Type           string      `json:"type"`
	Offset         int64       `json:"offset"`
	Number         int         `json:"number"`
}

type MenuItems struct {
	ID              int64    `json:"id"`
	Title           string   `json:"title"`
	RestaurantChain string   `json:"restaurantChain"`
	Image           string   `json:"image"`
	ImageType       string   `json:"imageType"`
	Servings        Servings `json:"servings"`
}

type Servings struct {
	Number int64  `json:"totalMenuItems"`
	Size   int64  `json:"size"`
	Unit   string `json:"unit"`
}
