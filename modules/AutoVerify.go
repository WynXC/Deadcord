package modules

import (
	"Deadcord/core"
	"Deadcord/requests"
	"Deadcord/util"
	"encoding/json"
	"log"
	"strings"
	"sync"
)

func StartAutoVerifyThreads(server_id string) int {
	var channel_data core.GuildChannels

	var wg sync.WaitGroup
	wg.Add(len(core.RawTokensLoaded))

	for _, token := range core.RawTokensLoaded {
		go func(server_id string, token string) {
			membershipScreenBypassWorker(server_id, token)
		}(server_id, token)
	}

	wg.Done()

	status, status_code, channel_json := requests.SendDiscordRequest("guilds/"+server_id+"/channels", "GET", core.RawTokensLoaded[0], map[string]interface{}{})

	if status && status_code == 200 {
		if err := json.Unmarshal(channel_json, &channel_data); err != nil {
			log.Fatal(err)
		}

		for _, channel_object := range channel_data {
			if strings.Contains(channel_object.Name, "verify") || strings.Contains(channel_object.Name, "verification") || strings.Contains(channel_object.Name, "prove-human") {
				util.WriteToConsole("Found verification channel. Attempting to verify.", 2)

				scraped_messages, err := GetMessages(channel_object.ID, 50, core.RawTokensLoaded[0])

				if err != nil {
					return 1
				}

				if len(scraped_messages) > 0 {
					var messages core.Message
					if err := json.Unmarshal(scraped_messages, &messages); err != nil {
						log.Fatal(err)
					}

					for _, message := range messages {
						if strings.Contains(message.Content, "verify") || strings.Contains(message.Content, "verification") {

							for _, reaction := range message.Reactions {
								StartReactThreads(message.ChannelID, message.ID, reaction.Emoji.Name, false)
							}

						}
					}
				} else {
					return 2
				}
			}
		}
	} else {
		return 3
	}

	return 0
}

func membershipScreenBypassWorker(server_id string, token string) {
	status, status_code, _ := requests.SendDiscordRequest("guilds/"+server_id+"/requests/@me", "PUT", token, map[string]interface{}{})

	if status {
		if status_code == 201 {
			util.WriteToConsole("Token bypassed member screening.", 2)
		}
	}

}
