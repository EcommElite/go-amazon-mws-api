package amazonmws

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
	//"bitbucket.org/zombiezen/cardcpx/natsort"
)

var versions map[string]string

func init() {
	versions = make(map[string]string)
	versions["/Feeds/2009-01-01"] = "2009-01-01"
	versions["/Products/2011-10-01"] = "2011-10-01"
	versions["/Reports/2009-01-01"] = "2009-01-01"
}

type AmazonMWSAPI struct {
	AccessKey     string
	SecretKey     string
	Host          string
	AuthToken     string
	MarketplaceId string
	SellerId      string
}

var strPost = []byte("POST")
func (api AmazonMWSAPI) fastSignAndFetchViaPost(Action string, ActionPath string, Parameters map[string]string, body []byte) (string, error) {
	genUrl, err := GenerateAmazonUrlPost(api, ActionPath)
	if err != nil {
		return "", err
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req) // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	if api.AuthToken != "" {
		Parameters["MWSAuthToken"] = api.AuthToken
	}

	Parameters["Action"] = Action
	Parameters["AWSAccessKeyId"] = api.AccessKey
	Parameters["SellerId"] = api.SellerId
	Parameters["SignatureVersion"] = "2"
	Parameters["SignatureMethod"] = "HmacSHA256"
	Parameters["Version"] = versions[ActionPath]
	fmt.Println("Version for", ActionPath, "is", versions[ActionPath], "from", versions)
	Parameters["Timestamp"] = time.Now().UTC().Format(time.RFC3339)

	signature, err := sign("POST", genUrl, Parameters, api)
	if err != nil {
		return "", err
	}
	Parameters["Signature"] = signature
	req.Header.SetMethodBytes(strPost)
	req.SetRequestURI(genUrl.String())

	v := url.Values{}
	for index, value := range Parameters {
		v.Set(index, value)
	}
	s := v.Encode()

	if body != nil {
		req.SetRequestURI(string(req.RequestURI()) + "?" + s)

		hash := md5.Sum(body)
		MD5 := base64.StdEncoding.EncodeToString([]byte(hash[:]))

		req.Header.DisableNormalizing()
		req.Header.Set("Content-MD5", MD5)
		req.Header.Set("Content-Type", "text/xml; charset=iso-8859-1")
		req.SetBody(body)
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("ContentLength", strconv.Itoa(len([]byte(s))))
		req.SetBodyString(s)
	}

	err = fasthttp.Do(req, resp)
	if err != nil {
		return "", err
	}

	bodyBytes := resp.Body()
	return string(bodyBytes), nil
}

func GenerateAmazonUrlPost(api AmazonMWSAPI, ActionPath string) (finalUrl *url.URL, err error) {
	result, err := url.Parse(api.Host)
	if err != nil {
		return nil, err
	}

	result.Host = api.Host
	result.Scheme = "https"
	result.Path = ActionPath

	return result, nil
}

func SetTimestamp(origUrl *url.URL) (err error) {
	values, err := url.ParseQuery(origUrl.RawQuery)
	if err != nil {
		return err
	}
	values.Set("Timestamp", time.Now().UTC().Format(time.RFC3339))
	origUrl.RawQuery = values.Encode()

	return nil
}

func sign(method string, origUrl *url.URL, params map[string]string, api AmazonMWSAPI) (string, error) {
	paramMap := make(map[string]string)
	for key, value := range params {
		paramMap[key] = value
	}
	paramMap["Timestamp"] = url.QueryEscape(paramMap["Timestamp"])

	keys := make([]string, len(paramMap))
	count := 0
	for k, _ := range paramMap {
		keys[count] = k
		count++
	}
	sort.Strings(keys)

	sortedParams := make([]string, len(paramMap))
	count = 0
	for _, k := range keys {
		var buffer bytes.Buffer
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(paramMap[k])
		sortedParams[count] = buffer.String()
		count++
	}

	stringParams := strings.Join(sortedParams, "&")

	toSign := fmt.Sprintf("%s\n%s\n%s\n%s", method, origUrl.Host, origUrl.Path, stringParams)

	hasher := hmac.New(sha256.New, []byte(api.SecretKey))
	_, err := hasher.Write([]byte(toSign))
	if err != nil {
		return "", err
	}

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	return hash, nil
}

func SignAmazonUrl(origUrl *url.URL, api AmazonMWSAPI) (signedUrl string, err error) {
	escapeUrl := strings.Replace(origUrl.RawQuery, ",", "%2C", -1)
	escapeUrl = strings.Replace(escapeUrl, ":", "%3A", -1)

	params := strings.Split(escapeUrl, "&")
	paramMap := make(map[string]string)
	keys := make([]string, len(params))

	for k, v := range params {
		parts := strings.Split(v, "=")
		paramMap[parts[0]] = parts[1]
		keys[k] = parts[0]
	}
	sort.Strings(keys)

	sortedParams := make([]string, len(params))
	for k, _ := range params {
		var buffer bytes.Buffer
		buffer.WriteString(keys[k])
		buffer.WriteString("=")
		buffer.WriteString(paramMap[keys[k]])
		sortedParams[k] = buffer.String()
	}

	stringParams := strings.Join(sortedParams, "&")

	toSign := fmt.Sprintf("GET\n%s\n%s\n%s", origUrl.Host, origUrl.Path, stringParams)

	hasher := hmac.New(sha256.New, []byte(api.SecretKey))
	_, err = hasher.Write([]byte(toSign))
	if err != nil {
		return "", err
	}

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	newParams := fmt.Sprintf("%s&Signature=%s", stringParams, hash)

	origUrl.RawQuery = newParams

	return origUrl.String(), nil
}
