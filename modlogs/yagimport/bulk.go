package yagimport

import (
	"fmt"

	"github.com/starshine-sys/covebotnt/crouter"
	"github.com/bwmarrin/discordgo"
)

func (y *yag) bulk(ctx *crouter.Ctx) (err error) {
	gs, err := ctx.Database.GetGuildSettings(ctx.Message.GuildID)
	if err != nil {
		_, err = ctx.CommandError(err)
		return err
	}

	if gs.YAG.Channel == "" {
		_, err = ctx.SendfNoAddXHandler("No YAGPDB.xyz mod log channel set. Set one with `%vyagimport channel`.", ctx.GuildPrefix)
		return err
	}

	messages := make([]*discordgo.Message, 0)
	var before string
	for {
		msgs, err := ctx.Session.ChannelMessages(gs.YAG.Channel, 100, before, "", "")
		fmt.Printf("Messages before ID %v, got %v messages\n", before, len(msgs))
		if len(msgs) == 0 {
			break
		}
		if err != nil {
			_, err = ctx.CommandError(err)
			return err
		}
		before = msgs[len(msgs)-1].ID
		messages = append(messages, msgs...)
	}

	messages = reverseMsgs(messages)
	total := len(messages)
	var processed int

	for _, m := range messages {
		if m.Author.ID != yagID {
			continue
		}
		y.process(ctx.Message.GuildID, &discordgo.MessageCreate{Message: m}, gs)
		processed++
	}

	_, err = ctx.SendfNoAddXHandler("Got %v messages; processed %v messages.", total, processed)
	return
}

func reverseMsgs(s []*discordgo.Message) []*discordgo.Message {
	a := make([]*discordgo.Message, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
