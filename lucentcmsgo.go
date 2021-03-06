package lucentcmsgo

import (
	"fmt"
	"net/url"
	"time"
)

const (
	lucentBaseUrl = "https://api.lucentcms.com/api/"
)

var (
	validEndpoints = map[string]bool{
		"documents":  true,
		"documents/": true,
		"channels":   true,
		"channels/":  true,
		"files":      true,
		"files/":     true,
	}

	protectedHeaders = map[string]bool{
		"Lucent-Channel": true,
		"Lucent-User":    true,
	}
)

type LucentClient struct {
	Channel, Token, LucentUser, BaseUrl string
	DefaultHeaders                      map[string]string
	RequestTimeout                      time.Duration
}

// Creates a new lucent struct
// Recommend to use NewLucentClient instead of populating the fields
func NewLucentClient(channel, token, lucentUser, locale string, duration time.Duration) *LucentClient {
	headers := make(map[string]string)

	headers["Accept"] = "application/json"

	headers["Lucent-Channel"] = channel
	headers["Authorization"] = "Bearer " + token
	headers["Accept-Language"] = locale

	if lucentUser != "" {
		headers["Lucent-User"] = lucentUser
	}

	lucentClient := &LucentClient{
		Channel:        channel,
		Token:          token,
		LucentUser:     lucentUser,
		DefaultHeaders: headers,
		BaseUrl:        lucentBaseUrl,
		RequestTimeout: duration,
	}

	return lucentClient
}

func (lc *LucentClient) NewRequest(endpoint string, data map[string]interface{}) (*Request, error) {

	if _, ok := validEndpoints[endpoint]; !ok {
		return nil, fmt.Errorf("unsupported out of scope. can not create request endpoint %v", endpoint)
	}

	endpoint = lc.BaseUrl + endpoint

	_, err := url.ParseRequestURI(endpoint)

	if err != nil {
		return nil, err
	}

	req := &Request{
		EndPoint: endpoint,
		Data:     data,
		Headers:  lc.DefaultHeaders,
		Timeout:  lc.RequestTimeout,
		Limit:    10,
		Skip:     0,
	}

	return req, nil
}
