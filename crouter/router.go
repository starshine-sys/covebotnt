package crouter

import (
	"github.com/Starshine113/covebotnt/cbctx"
	"github.com/Starshine113/covebotnt/structs"
)

// NewRouter creates a Router object
func NewRouter() *Router {
	router := &Router{}

	router.AddCommand(&Command{
		Name:        "help",
		Aliases:     []string{"usage", "hlep"},
		Description: "Show info about how to use the bot",
		Usage:       "help [command]",
		Permissions: PermLevelNone,
		Command:     router.dummy,
	})

	return router
}

// dummy is used when a command isn't handled with the normal process
func (r *Router) dummy(ctx *cbctx.Ctx) error {
	return nil
}

// AddCommand adds a command to the router
func (r *Router) AddCommand(cmd *Command) {
	r.Commands = append(r.Commands, cmd)
}

// CreateCommand creates a command and adds it to the router
func (r *Router) CreateCommand(names []string, description, usage string, perms PermLevel, command func(ctx *cbctx.Ctx) error) {
	name := names[0]
	var aliases []string
	if len(names) >= 1 {
		aliases = names[1:]
	}
	cmd := &Command{
		Name:        name,
		Aliases:     aliases,
		Description: description,
		Usage:       usage,
		Permissions: perms,
		Command:     command,
	}
	r.Commands = append(r.Commands, cmd)
}

// Execute actually executes the router
func (r *Router) Execute(ctx *cbctx.Ctx, guildSettings *structs.GuildSettings, ownerIDs []string) (err error) {
	if ctx.Match("help") {
		ctx.TriggerTyping()
		err = r.Help(ctx, guildSettings)
		return
	}
	for _, cmd := range r.Commands {
		if ctx.Match(append([]string{cmd.Name}, cmd.Aliases...)...) {
			ctx.TriggerTyping()
			if cmd.Permissions == PermLevelHelper {
				perms := checkHelperPerm(ctx, guildSettings)
				if perms != nil {
					ctx.CommandError(perms)
					return nil
				}
			} else if cmd.Permissions == PermLevelMod {
				perms := checkModPerm(ctx, guildSettings)
				if perms != nil {
					ctx.CommandError(perms)
					return nil
				}
			} else if cmd.Permissions == PermLevelAdmin {
				perms := checkAdmin(ctx)
				if perms != nil {
					ctx.CommandError(perms)
					return nil
				}
			} else if cmd.Permissions == PermLevelOwner {
				perms := checkOwner(ctx.Author.ID, ownerIDs)
				if perms != nil {
					ctx.CommandError(perms)
					return nil
				}
			}
			// add the guild settings to the additional parameters
			ctx.AdditionalParams["guildSettings"] = guildSettings

			err = cmd.Command(ctx)
			return err
		}
	}
	return
}
