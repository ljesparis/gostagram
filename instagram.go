// gostagram is an unofficial go client for
// the instagram api.
//
// note: there's two way to get the
// access token, using oauth2 golang library
// or using curl (both ways need a local server
// to get the token).
package gostagram

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/parnurzeal/gorequest"
)

const (
	apiUrl    = "https://api.instagram.com/v1/"
	version   = "1.0.0"
	userAgent = "gostagram/v1"
)

var (
	ClientSecretNotSetError     = errors.New("Client secret not set for signed request.")
	WrongInstagramEndpointError = errors.New("Wrong instagram endpoint.")
)

type Parameters map[string]string
type bodyData map[string]interface{}
type httpResponse map[string]interface{}

// Instagram client.
type Client struct {
	clientId      string // necesary to create subcriptions.
	clientSecret  string // necesary to send signed request.
	access_token  string
	signedRequest bool
}

func NewClient(access_token string) *Client {
	return &Client{
		clientId:      "",
		clientSecret:  "",
		access_token:  access_token,
		signedRequest: false,
	}
}

func (c *Client) SetClientId(id string) {
	c.clientId = id
}

func (c *Client) SetClientSecret(cs string) {
	c.clientSecret = cs
}

func (c *Client) SetSignedRequest(sr bool) {
	c.signedRequest = sr
}

func (c Client) newRequest(method, uri string, res *httpResponse, dataToSend ...bodyData) error {
	if c.signedRequest && len(c.clientSecret) == 0 {
		return ClientSecretNotSetError
	} else if c.signedRequest {
		// if you plan to use signed request,
		// you should check 'Enforce signed requests'
		// option at instagram client settings.
		//
		// Right here, endpoint and query parameters are extracted
		// from the url, to generate a hash and add it to the current url.
		//
		// Check this https://www.instagram.com/developer/secure-api-requests/
		// for more information, about it.
		tmpUrl, err := url.Parse(uri)

		if err != nil {
			return err
		}

		// deleting /v1 from endpoint.
		endpoint := strings.Replace(tmpUrl.EscapedPath(), "/v1", "", 1)

		// if endpoint ends in '/' character
		// should be deleted to correctly generate signed request.
		if endpoint[len(endpoint)-1] == '/' {
			endpoint = endpoint[0 : len(endpoint)-1]
		}

		params := make(Parameters)
		for paramName, paramValue := range tmpUrl.Query() {
			params[paramName] = paramValue[0]
		}

		sig, err := c.generateSignature(endpoint, params)

		if err != nil {
			return err
		}

		uri = uri + "&sig=" + sig
	}

	// set gostagram user agent.
	request := gorequest.New().Set("User-Agent", userAgent)

	switch method {
	case gorequest.POST:
		request = request.Post(uri).Type("form")

		// if there's some data to send,
		// add it to the body.
		if dataToSend != nil && len(dataToSend) > 0 {
			request = request.SendMap(dataToSend[0])
		}

		break
	case gorequest.GET:
		request = request.Get(uri).Type("json")
		break
	case gorequest.DELETE:
		request = request.Delete(uri).Type("json")
		break
	}

	_, body, errs := request.EndBytes()
	json.Unmarshal(body, &res)

	if len(errs) > 0 {
		return errs[0]
	}

	// checking if exists response content,
	// if not, endpoint may be wrong.
	if len(*res) > 0 {
		var code int
		var errorType string = ""
		var errorMessage string = ""

		// some responses may come with some
		// meta information and some don't.
		if (*res)["meta"] != nil {
			meta := (*res)["meta"].(map[string]interface{})
			code = int(meta["code"].(float64))

			if code != 200 {
				errorType = meta["error_type"].(string)
				errorMessage = meta["error_message"].(string)
			}
		} else {
			code = int((*res)["code"].(float64))

			if code != 200 {
				errorType = (*res)["error_type"].(string)
				errorMessage = (*res)["error_message"].(string)
			}
		}

		// if response code, ins't 200,
		// an error need be returned with it's
		// respective error message.
		if code != 200 {
			return fmt.Errorf("[%s]: %s", errorType, errorMessage)
		}
	} else {
		return WrongInstagramEndpointError
	}

	return nil
}

func (c Client) get(url string) (*interface{}, *interface{}, error) {
	var res httpResponse
	err := c.newRequest(gorequest.GET, url, &res)

	if err != nil {
		return nil, nil, err
	}

	tmp := res["data"]

	// check if response has pagination,
	// if it has, return it along response data.
	tmpPagination := res["pagination"]
	if tmpPagination != nil {
		return &tmp, &tmpPagination, nil
	}

	return &tmp, nil, nil
}

func (c Client) post(url string, dataToSend ...bodyData) (*interface{}, error) {
	var res httpResponse
	err := c.newRequest(gorequest.POST, url, &res, dataToSend...)

	if err != nil {
		return nil, err
	}

	tmp := res["data"]
	return &tmp, nil
}

func (c Client) delete(url string) (*interface{}, error) {
	var res httpResponse
	err := c.newRequest(gorequest.DELETE, url, &res)

	if err != nil {
		return nil, err
	}

	tmp := res["data"]
	return &tmp, nil
}
