package spec

type (
	InputOrder struct {
		No                int     `json:"no"`
		PlatformProductId string  `json:"platformProductId"`
		Qty               int     `json:"qty"`
		UnitPrice         float64 `json:"unitPrice"`
		TotalPrice        float64 `json:"totalPrice"`
	}
)
