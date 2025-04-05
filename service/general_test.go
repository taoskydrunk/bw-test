package service

import (
	"bw-test/spec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	t.Run("Case 1 : Only one product", func(t *testing.T) {
		generalService := NewGeneralService()

		input := []spec.InputOrder{{
			No:                1,
			PlatformProductId: "FG0A-CLEAR-IPHONE16PROMAX",
			Qty:               2,
			UnitPrice:         50,
			TotalPrice:        100,
		}}

		outputs := generalService.Output(input)
		assert.Len(t, outputs, 3)
		assert.Equal(t, outputs[0], spec.CleanedOrder{
			No:         1,
			ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
			MaterialId: "FG0A-CLEAR",
			ModelId:    "IPHONE16PROMAX",
			Qty:        2,
			UnitPrice:  50,
			TotalPrice: 100,
		})
		assert.Equal(t, outputs[1], spec.CleanedOrder{
			No:         2,
			ProductId:  "WIPING-CLOTH",
			MaterialId: "",
			ModelId:    "",
			Qty:        2,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		assert.Equal(t, outputs[2], spec.CleanedOrder{
			No:         3,
			ProductId:  "CLEAR-CLEANNER",
			MaterialId: "",
			ModelId:    "",
			Qty:        2,
			UnitPrice:  0,
			TotalPrice: 0,
		})
	})
}

func TestCase2(t *testing.T) {
	t.Run("Case 2: One product with wrong prefix", func(t *testing.T) {
		generalService := NewGeneralService()

		input := []spec.InputOrder{{
			No:                1,
			PlatformProductId: "x2-3&FG0A-CLEAR-IPHONE16PROMAX",
			Qty:               2,
			UnitPrice:         50,
			TotalPrice:        100,
		}}

		outputs := generalService.Output(input)
		assert.Len(t, outputs, 3)
		assert.Equal(t, outputs[0], spec.CleanedOrder{
			No:         1,
			ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
			MaterialId: "FG0A-CLEAR",
			ModelId:    "IPHONE16PROMAX",
			Qty:        2,
			UnitPrice:  50,
			TotalPrice: 100,
		})
		assert.Equal(t, outputs[1], spec.CleanedOrder{
			No:         2,
			ProductId:  "WIPING-CLOTH",
			MaterialId: "",
			ModelId:    "",
			Qty:        2,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		assert.Equal(t, outputs[2], spec.CleanedOrder{
			No:         3,
			ProductId:  "CLEAR-CLEANNER",
			MaterialId: "",
			ModelId:    "",
			Qty:        2,
			UnitPrice:  0,
			TotalPrice: 0,
		})
	})
}

func TestCase3(t *testing.T) {
	t.Run("Case 3: One product with wrong prefix and has * symbol that indicates the quantity", func(t *testing.T) {
		generalService := NewGeneralService()

		input := []spec.InputOrder{{
			No:                1,
			PlatformProductId: "x2-3&FG0A-MATTE-IPHONE16PROMAX*3",
			Qty:               1,
			UnitPrice:         90,
			TotalPrice:        90,
		}}

		outputs := generalService.Output(input)
		assert.Len(t, outputs, 3)
		assert.Equal(t, outputs[0], spec.CleanedOrder{
			No:         1,
			ProductId:  "FG0A-MATTE-IPHONE16PROMAX",
			MaterialId: "FG0A-MATTE",
			ModelId:    "IPHONE16PROMAX",
			Qty:        3,
			UnitPrice:  30,
			TotalPrice: 90,
		})
		assert.Equal(t, outputs[1], spec.CleanedOrder{
			No:         2,
			ProductId:  "WIPING-CLOTH",
			MaterialId: "",
			ModelId:    "",
			Qty:        3,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		assert.Equal(t, outputs[2], spec.CleanedOrder{
			No:         3,
			ProductId:  "MATTE-CLEANNER",
			MaterialId: "",
			ModelId:    "",
			Qty:        3,
			UnitPrice:  0,
			TotalPrice: 0,
		})
	})
}

func TestCase4(t *testing.T) {
	t.Run("Case 4: One bundle product with wrong prefix and split by / symbol into two product", func(t *testing.T) {
		generalService := NewGeneralService()

		input := []spec.InputOrder{{
			No:                1,
			PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B",
			Qty:               1,
			UnitPrice:         80,
			TotalPrice:        80,
		}}

		outputs := generalService.Output(input)
		assert.Len(t, outputs, 4)
		assert.Equal(t, outputs[0], spec.CleanedOrder{
			No:         1,
			ProductId:  "FG0A-CLEAR-OPPOA3",
			MaterialId: "FG0A-CLEAR",
			ModelId:    "OPPOA3",
			Qty:        1,
			UnitPrice:  40,
			TotalPrice: 40,
		})
		assert.Equal(t, outputs[1], spec.CleanedOrder{
			No:         2,
			ProductId:  "FG0A-CLEAR-OPPOA3-B",
			MaterialId: "FG0A-CLEAR",
			ModelId:    "OPPOA3-B",
			Qty:        1,
			UnitPrice:  40,
			TotalPrice: 40,
		})
		assert.Equal(t, outputs[2], spec.CleanedOrder{
			No:         3,
			ProductId:  "WIPING-CLOTH",
			MaterialId: "",
			ModelId:    "",
			Qty:        2,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		assert.Equal(t, outputs[3], spec.CleanedOrder{
			No:         4,
			ProductId:  "CLEAR-CLEANNER",
			MaterialId: "",
			ModelId:    "",
			Qty:        2,
			UnitPrice:  0,
			TotalPrice: 0,
		})
	})
}

func TestCase5(t *testing.T) {
	t.Run("Case 5: One bundle product with wrong prefix and split by / symbol into three product", func(t *testing.T) {
		generalService := NewGeneralService()

		input := []spec.InputOrder{{
			No:                1,
			PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B/FG0A-MAT",
			Qty:               1,
			UnitPrice:         120,
			TotalPrice:        120,
		}}

		outputs := generalService.Output(input)
		assert.Len(t, outputs, 6)
		assert.Equal(t, outputs[0], spec.CleanedOrder{
			No:         1,
			ProductId:  "FG0A-CLEAR-OPPOA3",
			MaterialId: "FG0A-CLEAR",
			ModelId:    "OPPOA3",
			Qty:        1,
			UnitPrice:  40,
			TotalPrice: 40,
		})
		assert.Equal(t, outputs[1], spec.CleanedOrder{
			No:         2,
			ProductId:  "FG0A-CLEAR-OPPOA3-B",
			MaterialId: "FG0A-CLEAR",
			ModelId:    "OPPOA3-B",
			Qty:        1,
			UnitPrice:  40,
			TotalPrice: 40,
		})
		assert.Equal(t, outputs[2], spec.CleanedOrder{
			No:         3,
			ProductId:  "FG0A-MATTE-OPPOA3",
			MaterialId: "FG0A-MATTE",
			ModelId:    "OPPOA3",
			Qty:        1,
			UnitPrice:  40,
			TotalPrice: 40,
		})
		assert.Equal(t, outputs[3], spec.CleanedOrder{
			No:         4,
			ProductId:  "WIPING-CLOTH",
			MaterialId: "",
			ModelId:    "",
			Qty:        3,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		assert.Equal(t, outputs[4], spec.CleanedOrder{
			No:         5,
			ProductId:  "CLEAR-CLEANNER",
			MaterialId: "",
			ModelId:    "",
			Qty:        2,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		assert.Equal(t, outputs[5], spec.CleanedOrder{
			No:         6,
			ProductId:  "MATTE-CLEANNER",
			MaterialId: "",
			ModelId:    "",
			Qty:        1,
			UnitPrice:  0,
			TotalPrice: 0,
		})
	})
}

func TestCase6(t *testing.T) {
	t.Run("Case 6: One bundle product with wrong prefix and have / symbol and * symbol", func(t *testing.T) {
		generalService := NewGeneralService()

		input := []spec.InputOrder{{
			No:                1,
			PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3",
			Qty:               1,
			UnitPrice:         120,
			TotalPrice:        120,
		}}

		outputs := generalService.Output(input)
		assert.Len(t, outputs, 5)
		assert.Equal(t, outputs[0], spec.CleanedOrder{
			No:         1,
			ProductId:  "FG0A-CLEAR-OPPOA3",
			MaterialId: "FG0A-CLEAR",
			ModelId:    "OPPOA3",
			Qty:        2,
			UnitPrice:  40,
			TotalPrice: 80,
		})
		assert.Equal(t, outputs[1], spec.CleanedOrder{
			No:         2,
			ProductId:  "FG0A-MATTE-OPPOA3",
			MaterialId: "FG0A-MATTE",
			ModelId:    "OPPOA3",
			Qty:        1,
			UnitPrice:  40,
			TotalPrice: 40,
		})
		assert.Equal(t, outputs[2], spec.CleanedOrder{
			No:         3,
			ProductId:  "WIPING-CLOTH",
			MaterialId: "",
			ModelId:    "",
			Qty:        3,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		assert.Equal(t, outputs[3], spec.CleanedOrder{
			No:         4,
			ProductId:  "CLEAR-CLEANNER",
			MaterialId: "",
			ModelId:    "",
			Qty:        2,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		assert.Equal(t, outputs[4], spec.CleanedOrder{
			No:         5,
			ProductId:  "MATTE-CLEANNER",
			MaterialId: "",
			ModelId:    "",
			Qty:        1,
			UnitPrice:  0,
			TotalPrice: 0,
		})
	})
}

func TestCase7(t *testing.T) {
	t.Run("Case 7: one product and one bundle product with wrong prefix and have / symbol and *\nsymbol", func(t *testing.T) {
		generalService := NewGeneralService()

		input := []spec.InputOrder{{
			No:                1,
			PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2",
			Qty:               1,
			UnitPrice:         160,
			TotalPrice:        160,
		},
			{
				No:                2,
				PlatformProductId: "FG0A-PRIVACY-IPHONE16PROMAX",
				Qty:               1,
				UnitPrice:         50,
				TotalPrice:        50,
			},
		}

		outputs := generalService.Output(input)
		assert.Len(t, outputs, 7)
		assert.Equal(t, outputs[0], spec.CleanedOrder{
			No:         1,
			ProductId:  "FG0A-CLEAR-OPPOA3",
			MaterialId: "FG0A-CLEAR",
			ModelId:    "OPPOA3",
			Qty:        2,
			UnitPrice:  40,
			TotalPrice: 80,
		})
		assert.Equal(t, outputs[1], spec.CleanedOrder{
			No:         2,
			ProductId:  "FG0A-MATTE-OPPOA3",
			MaterialId: "FG0A-MATTE",
			ModelId:    "OPPOA3",
			Qty:        2,
			UnitPrice:  40,
			TotalPrice: 80,
		})
		assert.Equal(t, outputs[2], spec.CleanedOrder{
			No:         3,
			ProductId:  "FG0A-PRIVACY-IPHONE16PROMAX",
			MaterialId: "FG0A-PRIVACY",
			ModelId:    "IPHONE16PROMAX",
			Qty:        1,
			UnitPrice:  50,
			TotalPrice: 50,
		})
		assert.Equal(t, outputs[3], spec.CleanedOrder{
			No:         4,
			ProductId:  "WIPING-CLOTH",
			MaterialId: "",
			ModelId:    "",
			Qty:        5,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		assert.Equal(t, outputs[4], spec.CleanedOrder{
			No:         5,
			ProductId:  "CLEAR-CLEANNER",
			MaterialId: "",
			ModelId:    "",
			Qty:        2,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		assert.Equal(t, outputs[5], spec.CleanedOrder{
			No:         6,
			ProductId:  "MATTE-CLEANNER",
			MaterialId: "",
			ModelId:    "",
			Qty:        2,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		assert.Equal(t, outputs[6], spec.CleanedOrder{
			No:         7,
			ProductId:  "PRIVACY-CLEANNER",
			MaterialId: "",
			ModelId:    "",
			Qty:        1,
			UnitPrice:  0,
			TotalPrice: 0,
		})
	})
}
