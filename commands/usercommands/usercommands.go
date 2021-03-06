package usercommands

import (
	"time"

	"github.com/starshine-sys/covebotnt/crouter"
)

// Init adds all the commands from this package to the router
func Init(router *crouter.Router) {
	router.AddCommand(&crouter.Command{
		Name:        "Ping",
		Description: "Ping pong!",
		Permissions: crouter.PermLevelNone,
		Command:     Ping,
	})

	router.AddCommand(&crouter.Command{
		Name:        "Invite",
		Description: "Send an invite link for the bot",
		Permissions: crouter.PermLevelNone,
		Command:     Invite,
	})

	router.AddCommand(&crouter.Command{
		Name:        "Enlarge",
		Aliases:     []string{"E", "Emote", "Emoji", "Enlorge"},
		Description: "Enlarges up to 10 custom emoji",
		Usage:       "<emoji...>",
		Permissions: crouter.PermLevelNone,
		Command:     Enlarge,
		Cooldown:    5 * time.Second,
	})

	router.AddCommand(&crouter.Command{
		Name:        "EmojiInfo",
		Aliases:     []string{"EI", "EmoteInfo"},
		Description: "Get detailed info about a custom emoji",
		Usage:       "<emoji>",
		Permissions: crouter.PermLevelNone,
		Command:     EmojiInfo,
	})

	router.AddCommand(&crouter.Command{
		Name:        "Color",
		Aliases:     []string{"Colour"},
		Description: "Previews a color",
		Usage:       "<color>",
		Permissions: crouter.PermLevelNone,
		Command:     Color,
	})

	router.AddCommand(&crouter.Command{
		Name:        "Avatar",
		Aliases:     []string{"Pfp", "A"},
		Description: "Show a user's avatar",
		Usage:       "[user]",
		Permissions: crouter.PermLevelNone,
		Command:     Avatar,
		GuildOnly:   true,
	})

	router.AddCommand(&crouter.Command{
		Name:        "Meow",
		Description: "Show a random meowmoji",
		Permissions: crouter.PermLevelNone,
		Command:     meow,
	})

	router.AddCommand(&crouter.Command{
		Name:    "Bubble",
		Aliases: []string{"Pop"},

		Description: "Generate spoilered bubble wrap",
		Usage:       "[-prepop] [size: 10]",

		Permissions: crouter.PermLevelNone,
		Command:     bubble,
	})

	s := router.AddGroup(&crouter.Group{
		Name:        "Snowflake",
		Aliases:     []string{"IDTime"},
		Description: "Get a snowflake/ID's timestamp",
		Command: &crouter.Command{
			Name: "Discord",

			Description: "Get the timestamps for Discord snowflakes",
			Usage:       "[...IDs]",

			Permissions: crouter.PermLevelNone,
			Command:     Snowflake,
		},
	})

	s.AddCommand(&crouter.Command{
		Name:    "CoveBot",
		Aliases: []string{"CB"},

		Description: "Get the timestamps for CoveBot snowflakes",
		Usage:       "[...IDs]",

		Permissions: crouter.PermLevelNone,
		Command:     covebotSnowflake,
	})

	router.AddCommand(&crouter.Command{
		Name:        "Snowflake",
		Aliases:     []string{"IDTime"},
		Description: "Get timestamps from the given ID(s)",
		Usage:       "[...IDs]",
		Permissions: crouter.PermLevelNone,
		Command:     Snowflake,
	})

	router.AddCommand(&crouter.Command{
		Name:        "UserInfo",
		Aliases:     []string{"I", "Info", "Whois", "Profile"},
		Description: "Show information about a user (or yourself)",
		Usage:       "[user]",
		Permissions: crouter.PermLevelNone,
		Command:     UserInfo,
		GuildOnly:   true,
	})

	router.AddCommand(&crouter.Command{
		Name:        "PKInfo",
		Aliases:     []string{"PKI", "PKUserInfo", "PKWhois", "PKProfile"},
		Description: "Show information about the user who sent a PluralKit-proxied message",
		Usage:       "<message ID>",
		Permissions: crouter.PermLevelNone,
		Command:     PKUserInfo,
		GuildOnly:   true,
	})

	router.AddCommand(&crouter.Command{
		Name:        "RoleInfo",
		Aliases:     []string{"Ri"},
		Description: "Show information about a role",
		Usage:       "<role>",
		Permissions: crouter.PermLevelNone,
		Command:     RoleInfo,
		GuildOnly:   true,
	})

	router.AddCommand(&crouter.Command{
		Name:        "ServerInfo",
		Aliases:     []string{"Si", "GuildInfo"},
		Description: "Show information about the current server",
		Usage:       "",
		Permissions: crouter.PermLevelNone,
		Command:     GuildInfo,
		GuildOnly:   true,
	})

	router.AddCommand(&crouter.Command{
		Name:        "Hello",
		Aliases:     []string{"Hi", "Henlo", "Heya", "Heyo", "Hewwo"},
		Description: "Say hi to the bot",
		Permissions: crouter.PermLevelNone,
		Command:     Hello,
		GuildOnly:   true,
	})

	router.AddCommand(&crouter.Command{
		Name:        "About",
		Description: "Show some info about the bot",
		Permissions: crouter.PermLevelNone,
		Command:     about,
	})

	router.AddCommand(&crouter.Command{
		Name:        "Quickpoll",
		Aliases:     []string{"qp"},
		Description: "Start a poll on the triggering message",
		Permissions: crouter.PermLevelNone,
		Command:     quickpoll,
	})

	router.AddCommand(&crouter.Command{
		Name:        "Poll",
		Description: "Create a poll, either yes/no or with a number of options",
		Usage:       "[--options <num>]",
		Permissions: crouter.PermLevelNone,
		Command:     poll,
	})
}
