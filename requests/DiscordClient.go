package requests

import (
	"Deadcord/core"
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var (
	CookieString = ""
	BaseURL      = "https://discord.com/api/v9/"
)

func SendDiscordRequest(endpoint string, method string, token string, data map[string]interface{}) (bool, int, []byte) {

	if CookieString == "" {
		CookieString = GetDiscordCookies()
	}

	_, token_data := core.GetTokenInfo(token)

	x_super_props := map[string]string{
		"os":                       token_data["os"],
		"browser":                  token_data["browser"],
		"device":                   "",
		"system_locale":            "en-US",
		"browser_user_agent":       token_data["agent"],
		"browser_version":          token_data["browser_version"],
		"os_version":               token_data["os_version"],
		"referrer":                 "",
		"referring_domain":         "",
		"referrer_current":         "",
		"referring_domain_current": "",
		"release_channel":          "stable",
		"client_build_number":      "107767",
		"client_event_source":      "None",
	}

	x_super_props_json, err := json.Marshal(x_super_props)

	if err != nil {
		log.Fatal(err)
	}

	x_super_props_refined := b64.StdEncoding.EncodeToString(x_super_props_json)

	discord_headers := http.Header{
		"Accept":             []string{"*/*"},
		"Accept-language":    []string{"en-GB"},
		"Authorization":      []string{token},
		"Alt-Used":           []string{"discord.com"},
		"Content-type":       []string{"application/json"},
		"Cookie":             []string{CookieString},
		"DNT":                []string{"1"},
		"Origin":             []string{"https://discord.com"},
		"Referer":            []string{"https://discord.com/channels/@me"},
		"Sec-fetch-dest":     []string{"empty"},
		"Sec-fetch-mode":     []string{"cors"},
		"Sec-fetch-site":     []string{"same-origin"},
		"Sec-ch-ua":          []string{"Not A;Brand';v='99', 'Chromium';v='96', 'Google Chrome';v='96'"},
		"Sec-ch-ua-mobile":   []string{"0"},
		"Sec-ch-ua-platform": []string{"Windows"},
		"TE":                 []string{"Trailers"},
		"User-Agent":         []string{token_data["agent"]},
		"X-debug-options":    []string{"bugReporterEnabled"},
		"X-discord-locale":   []string{"en-US"},
		"X-super-properties": []string{x_super_props_refined},
	}

	switch method {
	case "GET":
		status, status_code, body := GetRequestTemplate(BaseURL+endpoint, discord_headers)
		return status, status_code, body
	case "POST":
		status, status_code, body := RequestTemplate("POST", BaseURL+endpoint, discord_headers, data)
		return status, status_code, body
	case "PUT":
		status, status_code, body := RequestTemplate("PUT", BaseURL+endpoint, discord_headers, data)
		return status, status_code, body
	case "PATCH":
		status, status_code, body := RequestTemplate("PATCH", BaseURL+endpoint, discord_headers, data)
		return status, status_code, body
	case "DELETE":
		status, status_code, body := RequestTemplate("DELETE", BaseURL+endpoint, discord_headers, data)
		return status, status_code, body
	}

	return false, 0, nil
}

func RequestTemplate(request_type string, url string, headers http.Header, json_payload map[string]interface{}) (bool, int, []byte) {

	discord_client := http.Client{}

	patch_json, err := json.Marshal(json_payload)

	req, err := http.NewRequest(request_type, url, bytes.NewBuffer(patch_json))

	if err != nil {
		log.Fatal(err)
	}

	req.Header = headers

	res, err := discord_client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return true, res.StatusCode, []byte(body)
}

func GetRequestTemplate(url string, headers http.Header) (bool, int, []byte) {
	discord_client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header = headers

	res, err := discord_client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return true, res.StatusCode, []byte(body)
}
