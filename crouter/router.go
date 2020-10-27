package crouter

import (
	"github.com/Starshine113/covebotnt/cbctx"
	"github.com/Starshine113/covebotnt/structs"
)

// NewRouter creates a Router object
func NewRouter() *Router {
	return &Router{}
}

// AddCommand adds a command to the router
func (r *Router) AddCommand(cmd *Command) {
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
			err = cmd.Command(ctx)
			return err
		}
	}
	return
}