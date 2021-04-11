package amazonmws

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ecommelite/go-mws-api/amazon"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"testing"
)

func TestRequestReport(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	api := AmazonMWSAPI{
		AccessKey:     os.Getenv("ACCESS_KEY"),
		SecretKey:     os.Getenv("SECRET_KEY"),
		Host:          amazon.Marketplace(amazon.UnitedStates).MWSEndpoint(),
		AuthToken:     "",
		MarketplaceId: "ATVPDKIKX0DER",
		SellerId:      os.Getenv("SELLER_ID"),
	}

	scenarios := []struct {
		Name string
	}{
		{
			Name: "it works",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.Name, func(t *testing.T) {
			// 1. Given

			// 2. Do this
			res, q, err := api.RequestReport(RequestReportRequest{
				ReportType:    "_GET_XML_BROWSE_TREE_DATA_",
				ReportOptions: String("MarketplaceId=ATVPDKIKX0DER"),
			})

			spew.Dump(res)
			spew.Dump(q)
			spew.Dump(err)

			// 3. Expect
		})
	}
}

func TestGetReportRequestList(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	api := AmazonMWSAPI{
		AccessKey:     os.Getenv("ACCESS_KEY"),
		SecretKey:     os.Getenv("SECRET_KEY"),
		Host:          amazon.Marketplace(amazon.UnitedStates).MWSEndpoint(),
		AuthToken:     "",
		MarketplaceId: "ATVPDKIKX0DER",
		SellerId:      os.Getenv("SELLER_ID"),
	}

	scenarios := []struct {
		Name string
	}{
		{
			Name: "it works",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.Name, func(t *testing.T) {
			// 1. Given

			// 2. Do this
			res, q, err := api.GetReportRequestList(GetReportRequestListRequest{
				ReportRequestIdList: []string{"1215824018728"},
				//ReportType: "_GET_XML_BROWSE_TREE_DATA_",
				//ReportOptions: String("MarketplaceId=ATVPDKIKX0DER"),
			})

			spew.Dump(res)
			spew.Dump(q)
			spew.Dump(err)

			// 3. Expect
		})
	}
}

func TestGetReport(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	api := AmazonMWSAPI{
		AccessKey:     os.Getenv("ACCESS_KEY"),
		SecretKey:     os.Getenv("SECRET_KEY"),
		Host:          amazon.Marketplace(amazon.UnitedStates).MWSEndpoint(),
		AuthToken:     "",
		MarketplaceId: "ATVPDKIKX0DER",
		SellerId:      os.Getenv("SELLER_ID"),
	}

	scenarios := []struct {
		Name string
	}{
		{
			Name: "it works",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.Name, func(t *testing.T) {
			// 1. Given

			// 2. Do this
			res, q, err := api.GetReport("28918622025018728")

			spew.Dump(res)
			spew.Dump(q)
			spew.Dump(err)

			ioutil.WriteFile("us-browse-tree.xml", []byte(res), 0644)
			// 3. Expect
		})
	}
}

func String(s string) *string {
	return &s
}
