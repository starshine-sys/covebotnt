package crouter

import (
	"errors"
	"strings"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/Starshine113/covebotnt/bot"
	"github.com/Starshine113/covebotnt/cbdb"
	"github.com/Starshine113/covebotnt/etc"
	"github.com/Starshine113/covebotnt/structs"
	"github.com/bwmarrin/discordgo"
)

const (
	// SuccessEmoji is the emoji used to designate a successful action
	SuccessEmoji = "✅"
	// ErrorEmoji is the emoji used for errors
	ErrorEmoji = "❌"
	// WarnEmoji is the emoji used to warn that a command failed
	WarnEmoji = "⚠️"
)

// Ctx is the context for a command
type Ctx struct {
	GuildPrefix      string
	Command          string
	Args             []string
	Session          *discordgo.Session
	Bot              *bot.Bot
	Message          *discordgo.MessageCreate
	Channel          *discordgo.Channel
	Author           *discordgo.User
	BotUser          *discordgo.User
	Database         *cbdb.Db
	BoltDb           *cbdb.BoltDb
	Handlers         *ttlcache.Cache
	AdditionalParams map[string]interface{}
	GuildSettings    *structs.GuildSettings
	Cmd              *Command
}

// Context creates a new Ctx
func Context(prefix, messageContent string, m *discordgo.MessageCreate, b *bot.Bot) (ctx *Ctx, err error) {
	botUser := b.Session.State.User
	if botUser == nil {
		return
	}

	messageContent = etc.TrimPrefixesSpace(messageContent, prefix, "<@"+botUser.ID+">", "<@!"+botUser.ID+">")
	message := strings.Split(messageContent, " ")
	command := message[0]
	args := []string{}
	if len(message) > 1 {
		args = message[1:]
	}

	ctx = &Ctx{GuildPrefix: prefix, Command: command, Args: args, Session: b.Session, Message: m, Author: m.Author, Database: b.Pool, BoltDb: b.Bolt, Handlers: b.Handlers, Bot: b}

	channel, err := b.Session.Channel(m.ChannelID)
	if err != nil {
		return ctx, err
	}
	ctx.Channel = channel
	ctx.BotUser = botUser

	return ctx, nil
}

func combineQuotedItems(in []string) (out []string, err error) {
	var matchedQuote bool
	var beginQuote int
	for i, item := range in {
		if matchedQuote {
			if strings.HasSuffix(item, "\"") {
				item = strings.Join(in[beginQuote:i+1], " ")
				item = strings.Trim(item, "\"")
				matchedQuote = false
				out = append(out, item)
			}
			if matchedQuote && i == len(in)-1 {
				err = errors.New("unexpected end of input")
			}
			continue
		}
		if strings.HasPrefix(item, "\"") {
			matchedQuote = true
			beginQuote = i
			continue
		}
		out = append(out, item)
	}
	return out, err
}
