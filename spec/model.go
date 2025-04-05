package spec

type (
	CleanedOrder struct {
		No         int     `json:"no"`
		ProductId  string  `json:"productId"`
		MaterialId string  `json:"materialId"`
		ModelId    string  `json:"modelId"`
		Qty        int     `json:"qty"`
		UnitPrice  float64 `json:"unitPrice"`
		TotalPrice float64 `json:"totalPrice"`
	}

	ProductsModel struct {
		Data       []ProductModel `json:"data"`
		Total      int            `json:"total"`
		TotalPrice float64        `json:"totalPrice"`
	}

	ProductModel struct {
		Code       string `json:"code"`
		Qty        int    `json:"qty"`
		MaterialId string `json:"materialId"`
		ModelId    string `json:"modelId"`
	}
)
