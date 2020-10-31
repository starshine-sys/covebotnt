package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/Starshine113/covebotnt/cbctx"
	"github.com/Starshine113/covebotnt/cbdb"
	"github.com/bwmarrin/discordgo"
)

// command handler
func messageCreateCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// if message was sent by a bot return; not only to ignore bots, but also to make sure PluralKit users don't trigger commands twice.
	if m.Author.Bot {
		allowed := false
		for _, bot := range config.Bot.AllowedBots {
			if bot == m.Author.ID {
				allowed = true
			}
		}
		if allowed != true {
			return
		}
	}

	err := router.Respond(s, m)
	if err != nil {
		sugar.Errorf("Error sending autoresponse: %v", err)
	}

	// get prefix for the guild
	prefix := getPrefix(m.GuildID)

	botUser, err := s.User("@me")
	if err != nil {
		sugar.Errorf("Error fetching bot user: %v", err)
	}

	ctx, err := cbctx.Context(prefix, m.Content, s, m, &cbdb.Db{Pool: db}, boltDb)
	if err != nil {
		sugar.Errorf("Error getting context: %v", err)
	}
	// check if the message might be a command
	if ctx.MatchPrefix() {
		commandTree(ctx)
		return
	}

	// if not, check if the message contains a bot mention, and ends with "hello"
	content := strings.ToLower(m.Content)
	match, _ := regexp.MatchString(fmt.Sprintf("<@!?%v>.*hello$", botUser.ID), content)
	match2, _ := regexp.MatchString(fmt.Sprintf("%vhello$", regexp.QuoteMeta(prefix)), content)
	if match || match2 {
		ctx, err = cbctx.Context("--", "--hello", s, m, &cbdb.Db{Pool: db}, boltDb)
		if err != nil {
			sugar.Errorf("Error getting context: %v", err)
			return
		}
		commandTree(ctx)
		return
	}

	match, _ = regexp.MatchString(fmt.Sprintf("^<@!?%v>", botUser.ID), content)
	if match {
		_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The current prefix is `%v`", prefix))
		if err != nil {
			sugar.Errorf("Error sending message: %v", err)
			return
		}
		return
	}
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
