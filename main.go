package main

import (
	"bw-test/service"
	"bw-test/spec"
	"encoding/json"
	"flag"
	"fmt"
)

func main() {
	caseTest := flag.Int("case", 0, "case")
	flag.Parse()

	generalService := service.NewGeneralService()
	var output interface{}
	var input []spec.InputOrder

	switch *caseTest {
	case 1:
		input = []spec.InputOrder{{
			No:                1,
			PlatformProductId: "FG0A-CLEAR-IPHONE16PROMAX",
			Qty:               2,
			UnitPrice:         50,
			TotalPrice:        100,
		}}
		break
	case 2:
		input = []spec.InputOrder{{
			No:                1,
			PlatformProductId: "x2-3&FG0A-CLEAR-IPHONE16PROMAX",
			Qty:               2,
			UnitPrice:         50,
			TotalPrice:        100,
		}}
	case 3:
		input = []spec.InputOrder{{
			No:                1,
			PlatformProductId: "x2-3&FG0A-MATTE-IPHONE16PROMAX*3",
			Qty:               1,
			UnitPrice:         90,
			TotalPrice:        90,
		}}

	case 4:
		input = []spec.InputOrder{{
			No:                1,
			PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B",
			Qty:               1,
			UnitPrice:         80,
			TotalPrice:        80,
		}}

	case 5:
		input = []spec.InputOrder{{
			No:                1,
			PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B/FG0A-MAT",
			Qty:               1,
			UnitPrice:         120,
			TotalPrice:        120,
		}}

	case 6:
		input = []spec.InputOrder{{
			No:                1,
			PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3",
			Qty:               1,
			UnitPrice:         120,
			TotalPrice:        120,
		}}

	case 7:
		input = []spec.InputOrder{{
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

	default:
		panic("Not found Case!!!")
	}

	output = generalService.Output(input)
	byteData, _ := json.MarshalIndent(output, "", " ")
	fmt.Println(string(byteData))

}
