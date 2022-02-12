package modules

import (
	"Deadcord/core"
	"Deadcord/requests"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func StartDisguiseThreads() {
	var wg sync.WaitGroup
	wg.Add(len(core.RawTokensLoaded))

	for _, token := range core.RawTokensLoaded {
		go func(token string) {
			disguiseWorker(token)
		}(token)
	}
}

func disguiseWorker(token string) {
	resp, err := http.Get("https://picsum.photos/512/512")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	image := "data:image/png;base64," + b64.StdEncoding.EncodeToString(body)

	status, status_code, _ := requests.SendDiscordRequest("users/@me", "PATCH", token, map[string]interface{}{
		"avatar": image,
	})

	if status {
		fmt.Println(status_code)
	}
}
