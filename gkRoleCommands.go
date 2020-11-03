package main

import "github.com/Starshine113/covebotnt/cbctx"

func commandGkRoles(ctx *cbctx.Ctx) (err error) {
	if len(ctx.Args) < 1 {
		ctx.CommandError(&cbctx.ErrorNotEnoughArgs{
			NumRequiredArgs: 1,
			SuppliedArgs:    0,
		})
		return nil
	}

	var roles []string
	for _, r := range ctx.Args {
		role, err := ctx.ParseRole(r)
		if err != nil {
			ctx.CommandError(err)
			return nil
		}
		roles = append(roles, role.ID)
	}
	err = ctx.Database.SetGatekeeperRoles(ctx.Message.GuildID, roles)
	if err != nil {
		ctx.CommandError(err)
		return err
	}
	err = getSettingsForGuild(ctx.Message.GuildID)
	if err != nil {
		return err
	}

	_, err = ctx.Send(cbctx.SuccessEmoji + " Updated the list of gatekeeper roles.")
	return
}

func commandMemberRoles(ctx *cbctx.Ctx) (err error) {
	if len(ctx.Args) < 1 {
		ctx.CommandError(&cbctx.ErrorNotEnoughArgs{
			NumRequiredArgs: 1,
			SuppliedArgs:    0,
		})
		return nil
	}

	var roles []string
	for _, r := range ctx.Args {
		role, err := ctx.ParseRole(r)
		if err != nil {
			ctx.CommandError(err)
			return nil
		}
		roles = append(roles, role.ID)
	}
	err = ctx.Database.SetMemberRoles(ctx.Message.GuildID, roles)
	if err != nil {
		ctx.CommandError(err)
		return err
	}
	err = getSettingsForGuild(ctx.Message.GuildID)
	if err != nil {
		return err
	}

	_, err = ctx.Send(cbctx.SuccessEmoji + " Updated the list of member roles.")
	return
}
