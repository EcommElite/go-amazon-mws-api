// amazonmws provides methods for interacting with the Amazon Marketplace Services API.
package amazonmws

import (
	"bytes"
	"fmt"
	"strconv"
)

type FeeEstimateRequest struct {
	IdValue             string
	PriceToEstimateFees float64
	Currency            string
	MarketplaceId       string
	IdType              string
	Identifier          string
	IsAmazonFulfilled   bool
}

func (f *FeeEstimateRequest) requestString(index int, key string) string {
	var buffer bytes.Buffer
	buffer.WriteString("FeesEstimateRequestList.FeesEstimateRequest.")
	buffer.WriteString(strconv.Itoa(index))
	buffer.WriteString(".")
	buffer.WriteString(key)
	return buffer.String()
}

func (f *FeeEstimateRequest) setDefaults(mid string) {
	if f.Currency == "" {
		f.Currency = "USD"
	}

	if f.MarketplaceId == "" {
		f.MarketplaceId = mid
	}

	if f.IdType == "" {
		f.IdType = "ASIN"
	}

	if f.Identifier == "" {
		f.Identifier = f.IdValue
	}

	f.IsAmazonFulfilled = true
}

func (f *FeeEstimateRequest) toQuery(index int, marketplaceId string) map[string]string {
	output := make(map[string]string)

	f.setDefaults(marketplaceId)
	output[f.requestString(index+1, "IdValue")] = f.IdValue
	output[f.requestString(index+1, "PriceToEstimateFees.ListingPrice.Amount")] = strconv.FormatFloat(f.PriceToEstimateFees, 'f', 2, 32)
	output[f.requestString(index+1, "PriceToEstimateFees.ListingPrice.CurrencyCode")] = f.Currency
	output[f.requestString(index+1, "PriceToEstimateFees.Shipping.Amount")] = "0"
	output[f.requestString(index+1, "PriceToEstimateFees.Shipping.CurrencyCode")] = f.Currency
	output[f.requestString(index+1, "PriceToEstimateFees.Points.PointsNumber")] = "0"
	output[f.requestString(index+1, "PriceToEstimateFees.Points.PointsMonetaryValue.Amount")] = "0"
	output[f.requestString(index+1, "PriceToEstimateFees.Points.PointsMonetaryValue.CurrencyCode")] = f.Currency
	output[f.requestString(index+1, "MarketplaceId")] = f.MarketplaceId
	output[f.requestString(index+1, "IdType")] = f.IdType
	output[f.requestString(index+1, "Identifier")] = f.Identifier

	var isFba string
	if f.IsAmazonFulfilled {
		isFba = "true"
	} else {
		isFba = "false"
	}

	output[f.requestString(index+1, "IsAmazonFulfilled")] = isFba

	return output
}

// ListMatchingProducts - returns a list of products and their attributes, based on a search query.
func (api AmazonMWSAPI) ListMatchingProducts(query, queryContextID string) (string, Quota, error) {
	params := make(map[string]string)

	params["MarketplaceId"] = string(api.MarketplaceId)
	params["Query"] = query

	if queryContextID != "" {
		params["QueryContextId"] = queryContextID
	}

	return api.fastSignAndFetchViaPost("ListMatchingProducts", "/Products/2011-10-01", params, nil)
}

/*
GetLowestOfferListingsForASIN takes a list of ASINs and returns the result.
*/
func (api AmazonMWSAPI) GetLowestOfferListingsForASIN(items []string) (string, Quota, error) {
	params := make(map[string]string)

	for k, v := range items {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = string(api.MarketplaceId)

	return api.fastSignAndFetchViaPost("GetLowestOfferListingsForASIN", "/Products/2011-10-01", params, nil)
}

/*
GetCompetitivePricingForAsin takes a list of ASINs and returns the result.
*/
func (api AmazonMWSAPI) GetCompetitivePricingForASIN(items []string) (string, Quota, error) {
	params := make(map[string]string)

	for k, v := range items {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = string(api.MarketplaceId)

	return api.fastSignAndFetchViaPost("GetCompetitivePricingForASIN", "/Products/2011-10-01", params, nil)
}

func (api AmazonMWSAPI) GetMatchingProductForId(idType string, idList []string) (string, Quota, error) {
	params := make(map[string]string)

	for k, v := range idList {
		key := fmt.Sprintf("IdList.Id.%d", (k + 1))
		params[key] = string(v)
	}

	params["IdType"] = idType
	params["MarketplaceId"] = string(api.MarketplaceId)

	return api.fastSignAndFetchViaPost("GetMatchingProductForId", "/Products/2011-10-01", params, nil)
}

func (api AmazonMWSAPI) GetMyFeesEstimate(items []FeeEstimateRequest) (string, Quota, error) {
	params := make(map[string]string)

	for index, item := range items {
		queryItems := item.toQuery(index, api.MarketplaceId)

		for key, value := range queryItems {
			params[key] = value
		}
	}

	return api.fastSignAndFetchViaPost("GetMyFeesEstimate", "/Products/2011-10-01", params, nil)
}

func (api AmazonMWSAPI) GetReportRequestStatus(reportID string) (string, Quota, error) {
	params := make(map[string]string)

	params["ReportRequestIdList.Id.1"] = reportID

	return api.fastSignAndFetchViaPost("GetReportRequestList", "/Reports/2009-01-01", params, nil)
}

func (api AmazonMWSAPI) SubmitFeed(content []byte, feedType string) (string, Quota, error) {
	params := make(map[string]string)

	params["FeedType"] = feedType

	return api.fastSignAndFetchViaPost("SubmitFeed", "/Feeds/2009-01-01", params, content)
}

func (api AmazonMWSAPI) ListMarketplaceParticipations(content []byte, feedType string) (string, Quota, error) {
	return api.fastSignAndFetchViaPost("ListMarketplaceParticipations", "/Sellers/2011-07-01", nil, nil)
}