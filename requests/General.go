package requests

import (
	"Deadcord/util"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func ReadyRequestCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func GetDiscordCookies() string {
	resp, err := http.Get("https://discord.com")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return "__dcfduid=" + resp.Cookies()[0].Value + "; __sdcfduid=" + resp.Cookies()[1].Value + "; locale=en-GB;"

}

func JsonResponse(code int, message string, data map[string]interface{}) []byte {
	response := make(map[string]interface{})
	response["code"] = code
	response["message"] = message
	response["data"] = data

	switch code {
	case 200:
		util.WriteToConsole(response["message"].(string), 2)
	case 400:
		util.WriteToConsole(response["message"].(string), 1)
	case 500:
		util.WriteToConsole(response["message"].(string), 3)
	default:
		util.WriteToConsole(response["message"].(string), 0)
	}

	raw_json_response, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	return raw_json_response
}

func GetNonce() int64 {
	nonce_raw := strconv.FormatInt((time.Now().UTC().UnixNano()/1000000)-1420070400000, 2) + "0000000000000000000000"
	nonce, _ := strconv.ParseInt(nonce_raw, 2, 64)
	return nonce
}

func ErrorResponse(message string) []byte {
	return []byte(JsonResponse(500, message, map[string]interface{}{}))
}

func AllParametersError() []byte {
	return []byte(JsonResponse(400, "All parameters must be provided.", map[string]interface{}{}))
}
