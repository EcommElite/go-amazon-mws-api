package amazonmws

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestSimpleSignAmazonUrl(t *testing.T) {
	urlString := "https://mws-eu.amazonservices.com/Products/2011-10-01?ASINList.ASIN.1=1561706337&ASINList.ASIN.2=1561712027&ASINList.ASIN.10=1561711969"
	signedUrl := "https://mws-eu.amazonservices.com/Products/2011-10-01?ASINList.ASIN.1=1561706337&ASINList.ASIN.10=1561711969&ASINList.ASIN.2=1561712027&Signature=5aNYdV%2Fuc%2B98P2KwX8NhR6tzvKaHBHrLq5F4D7A9iH0%3D"

	var api AmazonMWSAPI

	api.SecretKey = "1234567890"

	url, err := url.Parse(urlString)
	if err != nil {
		t.Error("Could not parse urlstring")
	}

	resultUrl, err := SignAmazonUrl(url, api)
	if err != nil {
		t.Error("Signing failure:", err)
	}

	if signedUrl != resultUrl {
		t.Log(resultUrl, "\n")
		t.Error("Signed url does not match")
	}
}

func Test_SignAmazonUrl_1(t *testing.T) {

	urlString := `https://mws-eu.amazonservices.com/Products/2011-10-01?ASINList.ASIN.1=1561706337&ASINList.ASIN.10=1561711969&ASINList.ASIN.11=1561712027&ASINList.ASIN.12=2841051498&ASINList.ASIN.13=1561712043&ASINList.ASIN.14=1562413473&ASINList.ASIN.15=2729702806&ASINList.ASIN.16=2729702776&ASINList.ASIN.17=1561718939&ASINList.ASIN.18=2841090930&ASINList.ASIN.19=156171951X&ASINList.ASIN.2=1561712930&ASINList.ASIN.20=2729702032&ASINList.ASIN.3=1561713066&ASINList.ASIN.4=2729701737&ASINList.ASIN.5=1561711837&ASINList.ASIN.6=1561711845&ASINList.ASIN.7=1561711896&ASINList.ASIN.8=1561711810&ASINList.ASIN.9=1561712019&AWSAccessKeyId=AKIAJLUHOXLR5S2L6A6A&Action=GetLowestOfferListingsForASIN&MarketplaceId=APJ6JRA9NG5V4&SellerId=A2APQUVDBVWV7E&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=2013-03-29T02%3A16%3A18Z`

	signedUrl := `https://mws-eu.amazonservices.com/Products/2011-10-01?ASINList.ASIN.1=1561706337&ASINList.ASIN.10=1561711969&ASINList.ASIN.11=1561712027&ASINList.ASIN.12=2841051498&ASINList.ASIN.13=1561712043&ASINList.ASIN.14=1562413473&ASINList.ASIN.15=2729702806&ASINList.ASIN.16=2729702776&ASINList.ASIN.17=1561718939&ASINList.ASIN.18=2841090930&ASINList.ASIN.19=156171951X&ASINList.ASIN.2=1561712930&ASINList.ASIN.20=2729702032&ASINList.ASIN.3=1561713066&ASINList.ASIN.4=2729701737&ASINList.ASIN.5=1561711837&ASINList.ASIN.6=1561711845&ASINList.ASIN.7=1561711896&ASINList.ASIN.8=1561711810&ASINList.ASIN.9=1561712019&AWSAccessKeyId=AKIAJLUHOXLR5S2L6A6A&Action=GetLowestOfferListingsForASIN&MarketplaceId=APJ6JRA9NG5V4&SellerId=A2APQUVDBVWV7E&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=2013-03-29T02%3A16%3A18Z&Signature=rpDlBzpJ2t5nO3SLy66Y1oTAS9ZUhbH9kd639ed8g0w%3D`
	//signedUrl := `https://mws-eu.amazonservices.com/Products/2011-10-01?ASINList.ASIN.1=1561706337&ASINList.ASIN.2=1561712930&ASINList.ASIN.3=1561713066&ASINList.ASIN.4=2729701737&ASINList.ASIN.5=1561711837&ASINList.ASIN.6=1561711845&ASINList.ASIN.7=1561711896&ASINList.ASIN.8=1561711810&ASINList.ASIN.9=1561712019&ASINList.ASIN.10=1561711969&ASINList.ASIN.11=1561712027&ASINList.ASIN.12=2841051498&ASINList.ASIN.13=1561712043&ASINList.ASIN.14=1562413473&ASINList.ASIN.15=2729702806&ASINList.ASIN.16=2729702776&ASINList.ASIN.17=1561718939&ASINList.ASIN.18=2841090930&ASINList.ASIN.19=156171951X&ASINList.ASIN.20=2729702032&AWSAccessKeyId=AKIAJLUHOXLR5S2L6A6A&Action=GetLowestOfferListingsForASIN&MarketplaceId=APJ6JRA9NG5V4&SellerId=A2APQUVDBVWV7E&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=2013-03-29T02%3A16%3A18Z&Signature=2O9DpwF6%2F0x6dX6QQLMCETP42NRkqqAaOzFDsZdIGs8%3D`

	var api AmazonMWSAPI

	api.SecretKey = "1234567890"

	url, err := url.Parse(urlString)
	if err != nil {
		t.Error("Could not parse urlstring")
	}

	resultUrl, err := SignAmazonUrl(url, api)
	if err != nil {
		t.Error("Signing failure:", err)
	}

	if signedUrl != resultUrl {
		t.Log(resultUrl, "\n")
		t.Error("Signed url does not match")
	}
}

func Test_GetLowestOfferListingForAsin(t *testing.T) {
}

func TestGetMyFeesEstimateQuery(t *testing.T) {
	//item := FeeEstimateRequest{ IdValue: "BOOKBOOK12", PriceToEstimateFees: 10.11 }

	//request := item.requestString(1, "US");

	//result := item.toQuery(1, "US");

	//fmt.Println(request);
	//fmt.Println(result);
	if false {
		fmt.Println("")
	}
}

func TestSign(t *testing.T) {
	api := AmazonMWSAPI{}
	api.AccessKey = "TEST"
	api.SecretKey = "TEST"
	api.SellerId = "TEST"
	api.MarketplaceId = "ATVPDKIKX0DER"
	api.Host = "mws.amazonservices.com"

	link, err := url.Parse(api.Host)
	link.Host = api.Host
	link.Scheme = "https"
	link.Path = "/Products/2011-10-01"
	assert.Nil(t, err)

	params := make(map[string]string)
	params["Action"] = "GetMyFeesEstimate"
	params["AWSAccessKeyId"] = api.AccessKey
	params["SellerId"] = api.SellerId
	params["SignatureVersion"] = "2"
	params["SignatureMethod"] = "HmacSHA256"
	params["Version"] = "2011-10-01"
	params["FeesEstimateRequestList.FeesEstimateRequest.1.MarketplaceId"] = "ATVPDKIKX0DER"
	params["FeesEstimateRequestList.FeesEstimateRequest.1.IdType"] = "ASIN"
	params["FeesEstimateRequestList.FeesEstimateRequest.1.IdValue"] = "B06XPRCY44"
	params["FeesEstimateRequestList.FeesEstimateRequest.1.IsAmazonFulfilled"] = "true"
	params["FeesEstimateRequestList.FeesEstimateRequest.1.Identifier"] = "B06XPRCY44"
	params["FeesEstimateRequestList.FeesEstimateRequest.1.PriceToEstimateFees.ListingPrice.Amount"] = "8.86"
	params["FeesEstimateRequestList.FeesEstimateRequest.1.PriceToEstimateFees.ListingPrice.CurrencyCode"] = "USD"
	params["FeesEstimateRequestList.FeesEstimateRequest.1.PriceToEstimateFees.Shipping.Amount"] = "0"
	params["FeesEstimateRequestList.FeesEstimateRequest.1.PriceToEstimateFees.Shipping.CurrencyCode"] = "USD"
	params["FeesEstimateRequestList.FeesEstimateRequest.1.PriceToEstimateFees.Points.PointsNumber"] = "0"
	params["FeesEstimateRequestList.FeesEstimateRequest.1.PriceToEstimateFees.Points.PointsMonetaryValue.Amount"] = "0"
	params["FeesEstimateRequestList.FeesEstimateRequest.1.PriceToEstimateFees.Points.PointsMonetaryValue.CurrencyCode"] = "USD"
	params["Timestamp"] = "2018-07-31T19:52:07Z"

	// 2. Do this
	signature, err := sign("POST", link, params, api)

	// 3.1 Expect no error
	assert.Nil(t, err)

	// 3.2 Expect signature to be correct
	assert.Equal(t, "9D6Lv2KDXcsJ5aWvityyJhEav1EV0tgjpPHI7w7hTIc=", signature)
}

func TestSignFromRequest(t *testing.T) {
	api := AmazonMWSAPI{}
	api.AccessKey = "TEST"
	api.SecretKey = "TEST"
	api.SellerId = "TEST"
	api.MarketplaceId = "ATVPDKIKX0DER"
	api.Host = "mws.amazonservices.com"

	link, err := url.Parse(api.Host)
	link.Host = api.Host
	link.Scheme = "https"
	link.Path = "/Products/2011-10-01"
	assert.Nil(t, err)

	request := FeeEstimateRequest{
		IdValue:             "B06XPRCY44",
		PriceToEstimateFees: 8.86,
		Currency:            "USD",
		MarketplaceId:       "ATVPDKIKX0DER",
		IdType:              "ASIN",
		Identifier:          "B06XPRCY44",
		IsAmazonFulfilled:   true,
	}

	params := make(map[string]string)
	queryItems := request.toQuery(0, api.MarketplaceId)

	for key, value := range queryItems {
		params[key] = value
	}

	params["Action"] = "GetMyFeesEstimate"
	params["AWSAccessKeyId"] = api.AccessKey
	params["SellerId"] = api.SellerId
	params["SignatureVersion"] = "2"
	params["SignatureMethod"] = "HmacSHA256"
	params["Version"] = "2011-10-01"
	params["Timestamp"] = "2018-07-31T19:52:07Z"

	// 2. Do this
	signature, err := sign("POST", link, params, api)

	// 3.1 Expect no error
	assert.Nil(t, err)

	// 3.2 Expect signature to be correct
	assert.Equal(t, "9D6Lv2KDXcsJ5aWvityyJhEav1EV0tgjpPHI7w7hTIc=", signature)
}
