package modules

import (
	"Deadcord/core"
	"Deadcord/requests"
	"Deadcord/util"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
)

var (
	JoinResults    int  = 0
	AttemptedJoins int  = 0
	HaltJoin       bool = false
)

func StartJoinGuildThreads(invite string, delay int) int {
	if JoinResults != 0 {
		JoinResults = 0
	}

	var wg sync.WaitGroup
	wg.Add(len(core.RawTokensLoaded))

	join_channel := make(chan int)

	for _, token := range core.RawTokensLoaded {
		util.Sleep(delay)
		go func(token string, invite string, join_channel chan int) {
			joinWorker(token, invite, join_channel)
		}(token, invite, join_channel)
	}

	join_channel_results := <-join_channel

	close(join_channel)
	wg.Done()

	return join_channel_results
}

func joinWorker(token string, invite string, join_results chan int) {
	invite_parts := strings.Split(invite, "/")
	invite_code := ""

	if invite_parts[3] == "invite" {
		invite_code = invite_parts[4]
	} else {
		invite_code = invite_parts[3]
	}

	status, status_code, join_json := requests.SendDiscordRequest("invites/"+invite_code, "POST", token, map[string]interface{}{})

	if status {
		if !HaltJoin {
			switch status_code {
			case 200:
				var guild_data core.GuildJoin
				if err := json.Unmarshal(join_json, &guild_data); err != nil {
					log.Fatal(err)
				}

				JoinResults++
				util.WriteToConsole("Bot successfully joined guild.", 2)
			case 400:
				util.WriteToConsole("Could not join guild, captcha detected! Attempting bypass.", 1)
				HaltJoin = true
			case 404:
				util.WriteToConsole("Guild not found, or invite invalid.", 1)
			case 429:
				util.WriteToConsole("IP ratelimited or Cloudflare banned.", 1)
			default:
				util.WriteToConsole("Token could not join guild, request failed. Code: "+strconv.Itoa(status_code), 3)
			}
		}
	}

	AttemptedJoins++

	if JoinResults == len(core.RawTokensLoaded) || AttemptedJoins == len(core.RawTokensLoaded) {
		join_results <- JoinResults
		return
	}

}
