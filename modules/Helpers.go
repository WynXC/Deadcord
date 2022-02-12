package modules

import (
	"Deadcord/core"
	"Deadcord/requests"
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

func InServer(server_id string, token string) bool {
	status, status_code, _ := requests.SendDiscordRequest("guilds/"+server_id, "GET", token, map[string]interface{}{})

	if status {

		if status_code == 200 {
			return true
		} else {
			return false
		}

	} else {
		return false
	}
}

func GetMessages(channel_id string, amount int, token string) []byte {
	status, status_code, messages_json := requests.SendDiscordRequest("channels/"+channel_id+"/messages?limit="+strconv.Itoa(amount), "GET", token, map[string]interface{}{})

	if status {
		if status_code == 200 {
			return messages_json
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func GetGuildIdFromInvite(invite string) string {
	invite_parts := strings.Split(invite, "/")
	invite_code := ""

	if invite_parts[3] == "invite" {
		invite_code = invite_parts[4]
	} else {
		invite_code = invite_parts[3]
	}

	status, status_code, invite_json := requests.SendDiscordRequest("invites/"+invite_code, "GET", core.RawTokensLoaded[0], map[string]interface{}{})

	if status && status_code == 200 {
		var invite core.Invite
		if err := json.Unmarshal(invite_json, &invite); err != nil {
			log.Fatal(err)
		}

		return invite.Guild.ID
	} else {
		return ""
	}
}

func GetChannels(server_id string) []string {
	var channels []string

	status, status_code, channel_json := requests.SendDiscordRequest("guilds/"+server_id+"/channels", "GET", core.RawTokensLoaded[0], map[string]interface{}{})

	var result core.GuildChannels
	if err := json.Unmarshal(channel_json, &result); err != nil {
		return nil
	}

	if status {
		switch status_code {
		case 200:
			for _, channel := range result {
				if channel.Type == 0 {
					channels = append(channels, channel.ID)
				}
			}
		default:
			return channels
		}
	}

	return channels
}
