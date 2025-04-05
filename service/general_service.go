package service

import (
	"bw-test/spec"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type GeneralService struct {
	outputs []spec.CleanedOrder
	total   int
}

func NewGeneralService() *GeneralService {

	return &GeneralService{
		outputs: []spec.CleanedOrder{},
		total:   0,
	}
}

func (g *GeneralService) Output(inputs []spec.InputOrder) []spec.CleanedOrder {
	g.outputs = []spec.CleanedOrder{}
	g.total = 0

	for _, input := range inputs {
		g.Generate(input)
	}

	g.AddOnMaterial()
	return g.outputs
}

func (g *GeneralService) Generate(input spec.InputOrder) {
	products := g.TransformData(strings.Split(input.PlatformProductId, "/"), input)

	for _, product := range products.Data {
		unit, price := getPrice(products.Total, product.Qty, products.TotalPrice)
		g.outputs = append(g.outputs, g.AddData(product, "main", len(g.outputs), unit, price))
	}

	g.total = g.total + products.Total
}

func (g *GeneralService) TransformData(tmpProducts []string, input spec.InputOrder) spec.ProductsModel {
	models := []string{}
	products := []spec.ProductModel{}
	total := 0
	for _, product := range tmpProducts {
		re := regexp.MustCompile(`([A-Z0-9]+(?:-[A-Z0-9]+){2,})`)
		arrInput := re.FindAllString(product, -1)

		mainStr := product
		if len(arrInput) == 0 {
			product = fmt.Sprintf("%s-%s", findMaterial(product), models[0])
		} else {
			product = arrInput[0]
		}

		arr := strings.Split(product, "-")
		materialID := strings.Join(arr[:2], "-")
		model := strings.Join(arr[2:], "-")
		qty := getQty(mainStr, input.Qty)

		models = append(models, model)
		total = total + qty

		products = append(products, spec.ProductModel{
			Code:       product,
			Qty:        qty,
			MaterialId: materialID,
			ModelId:    model,
		})
	}

	return spec.ProductsModel{
		Data:       products,
		Total:      total,
		TotalPrice: input.TotalPrice,
	}
}
func (g *GeneralService) AddOnMaterial() {
	keys := []string{"WIPING-CLOTH", "FG0A-CLEAR", "FG0A-MATTE", "FG0A-PRIVACY"}

	for _, key := range keys {
		qty := 0
		value := "add1"
		if key != "WIPING-CLOTH" {
			value = "add2"
		}

		if value == "add2" {
			if count, ok := g.GetByMaterial(key); ok {
				qty = count
			} else {
				continue
			}
		} else {
			qty = g.total
		}

		g.outputs = append(g.outputs, g.AddData(spec.ProductModel{
			Qty:        qty,
			MaterialId: key,
		}, value, len(g.outputs), 0, 0))

	}
}

func (g *GeneralService) GetByMaterial(key string) (int, bool) {
	count := 0
	check := false
	for _, output := range g.outputs {
		if output.MaterialId == key {
			count = count + (output.Qty * 1)
			check = true
		}
	}

	return count, check
}

func (g *GeneralService) AddData(product spec.ProductModel, key string, i int, unitPrice, total float64) spec.CleanedOrder {
	output := spec.CleanedOrder{
		No:         i + 1,
		ProductId:  "",
		MaterialId: "",
		ModelId:    "",
		Qty:        product.Qty,
		UnitPrice:  0.00,
		TotalPrice: 0.00,
	}
	//
	switch key {
	case "main":
		output.ProductId = product.Code
		output.MaterialId = product.MaterialId
		output.ModelId = product.ModelId
		output.UnitPrice = unitPrice
		output.TotalPrice = total
		break
	case "add1":
		output.ProductId = "WIPING-CLOTH"
	case "add2":
		output.ProductId = getAddOn(product.MaterialId)
	default:

	}

	return output
}

func getAddOn(materialID string) string {
	switch materialID {
	case "FG0A-MATTE":
		return "MATTE-CLEANNER"
	case "FG0A-CLEAR":
		return "CLEAR-CLEANNER"
	case "FG0A-PRIVACY":
		return "PRIVACY-CLEANNER"
	}
	return ""
}

func findMaterial(materialID string) string {
	materials := []string{"FG0A-MATTE", "FG0A-CLEAR", "FG0A-PRIVACY"}
	for _, material := range materials {
		if strings.Contains(material, materialID) {
			return material
		}
	}

	return materialID
}

func getQty(key string, qty int) int {
	arr := strings.SplitAfter(key, "*")
	if len(arr) == 2 {
		newQty, err := strconv.Atoi(arr[1])
		if err != nil {
			fmt.Println(err.Error())
			return qty
		}
		return newQty
	}

	return qty
}

func getPrice(total, qty int, allPrice float64) (float64, float64) {
	unitPrice := allPrice / float64(total)

	return unitPrice, unitPrice * float64(qty)
}
