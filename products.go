package products

type Product struct {
	Id         int     `json:"id"`
	Name       string  `json:"name" binding:"required"`
	Quantity   float64 `json:"quantity" binding:"required"`
	Unit_coast float64 `json:"unit_coast" binding:"required"`
	Measure    int     `json:"measure" binding:"required"`
}

type Measure struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}
