package core

import (
	"time"
)

type GuildChannels []struct {
	ID                   string        `json:"id"`
	Type                 int           `json:"type"`
	Name                 string        `json:"name"`
	Position             int           `json:"position"`
	ParentID             interface{}   `json:"parent_id"`
	GuildID              string        `json:"guild_id"`
	PermissionOverwrites []interface{} `json:"permission_overwrites"`
	Nsfw                 bool          `json:"nsfw"`
	LastMessageID        string        `json:"last_message_id,omitempty"`
	Topic                interface{}   `json:"topic,omitempty"`
	RateLimitPerUser     int           `json:"rate_limit_per_user,omitempty"`
	Banner               interface{}   `json:"banner,omitempty"`
	Bitrate              int           `json:"bitrate,omitempty"`
	UserLimit            int           `json:"user_limit,omitempty"`
	RtcRegion            interface{}   `json:"rtc_region,omitempty"`
}

type RateLimit struct {
	Code       int     `json:"code"`
	Global     bool    `json:"global"`
	Message    string  `json:"message"`
	RetryAfter float64 `json:"retry_after"`
}

type GuildJoin struct {
	Code      string    `json:"code"`
	Type      int       `json:"type"`
	ExpiresAt time.Time `json:"expires_at"`
	Guild     Guild     `json:"guild"`
	Channel   Channel   `json:"channel"`
	Inviter   Inviter   `json:"inviter"`
}

type GuildJoinFail struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Guild struct {
	ID                string        `json:"id"`
	Name              string        `json:"name"`
	Splash            interface{}   `json:"splash"`
	Banner            interface{}   `json:"banner"`
	Description       interface{}   `json:"description"`
	Icon              interface{}   `json:"icon"`
	Features          []interface{} `json:"features"`
	VerificationLevel int           `json:"verification_level"`
	VanityURLCode     interface{}   `json:"vanity_url_code"`
	Nsfw              bool          `json:"nsfw"`
	NsfwLevel         int           `json:"nsfw_level"`
}

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

type Inviter struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   int    `json:"public_flags"`
}

type Invite struct {
	Code      string    `json:"code"`
	Type      int       `json:"type"`
	ExpiresAt time.Time `json:"expires_at"`
	Guild     Guild     `json:"guild"`
	Channel   Channel   `json:"channel"`
	Inviter   Inviter   `json:"inviter"`
}

type Message []struct {
	ID              string        `json:"id"`
	Type            int           `json:"type"`
	Content         string        `json:"content"`
	ChannelID       string        `json:"channel_id"`
	Author          Author        `json:"author"`
	Attachments     []interface{} `json:"attachments"`
	Embeds          []interface{} `json:"embeds"`
	Mentions        []interface{} `json:"mentions"`
	MentionRoles    []interface{} `json:"mention_roles"`
	Pinned          bool          `json:"pinned"`
	MentionEveryone bool          `json:"mention_everyone"`
	Tts             bool          `json:"tts"`
	Timestamp       time.Time     `json:"timestamp"`
	EditedTimestamp interface{}   `json:"edited_timestamp"`
	Flags           int           `json:"flags"`
	Components      []interface{} `json:"components"`
	Reactions       []Reactions   `json:"reactions"`
}

type GuildMessages []struct {
	ID              string        `json:"id"`
	Type            int           `json:"type"`
	Content         string        `json:"content"`
	ChannelID       string        `json:"channel_id"`
	Author          Author        `json:"author"`
	Attachments     []interface{} `json:"attachments"`
	Embeds          []interface{} `json:"embeds"`
	Mentions        []interface{} `json:"mentions"`
	MentionRoles    []interface{} `json:"mention_roles"`
	Pinned          bool          `json:"pinned"`
	MentionEveryone bool          `json:"mention_everyone"`
	Tts             bool          `json:"tts"`
	Timestamp       time.Time     `json:"timestamp"`
	EditedTimestamp interface{}   `json:"edited_timestamp"`
	Flags           int           `json:"flags"`
	Components      []interface{} `json:"components"`
}

type Author struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   int    `json:"public_flags"`
}

type Emoji struct {
	ID   interface{} `json:"id"`
	Name string      `json:"name"`
}

type Reactions struct {
	Emoji Emoji `json:"emoji"`
	Count int   `json:"count"`
	Me    bool  `json:"me"`
}
