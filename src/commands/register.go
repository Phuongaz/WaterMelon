package commands

import (
	"github.com/phuongaz/minecraft-bedrock-server/src/commands/economy"
	"github.com/phuongaz/minecraft-bedrock-server/src/commands/npc"
	"github.com/phuongaz/minecraft-bedrock-server/src/commands/world"
	"github.com/phuongaz/minecraft-bedrock-server/src/skyblock/command"
	"github.com/df-mc/dragonfly/server/cmd"
)

func Setup() {
	cmd.Register(cmd.New("help", "Provides help/list of commands.", []string{"?"}, Help{}))

	cmd.Register(cmd.New("version", "Gets the version of this server in use.", []string{"ver", "about"}, Version{}))
	cmd.Register(cmd.New("status", "Reads back the server's performance.", []string{"stat"}, Status{}))
	cmd.Register(cmd.New("teleport", "Teleport to x y z", []string{"tp"}, Teleport{}))
	cmd.Register(cmd.New("list", "Lists all online players", nil, List{}))
	cmd.Register(cmd.New("gc", "Fires garbage collection tasks.", nil, GC{}))
	cmd.Register(cmd.New("stop", "Stops the server.", nil, Stop{}))
	cmd.Register(cmd.New("op", "Grants operator status to a player.", nil, Op{}))
	cmd.Register(cmd.New("deop", "Revokes operator status from a player.", nil, DeOp{}))
	cmd.Register(cmd.New("banlist", "View all players banned from this server", nil, BanList{}))
	cmd.Register(cmd.New("ban", "Adds player to banlist.", nil, Ban{}))
	cmd.Register(cmd.New("unban", "Removes player from banlist.", nil, Unban{}))
	cmd.Register(cmd.New("kick", "Kicks a player from the server.", nil, Kick{}))
	cmd.Register(cmd.New("difficulty", "Sets the game difficulty", nil, Difficulty{}))
	cmd.Register(cmd.New("defaultgamemode", "Sets the default game mode.", nil, DefaultGameMode{}))
	cmd.Register(cmd.New("gamemode", "Sets your game mode.", []string{"gm"}, GameMode{}))
	cmd.Register(cmd.New("setworldspawn", "Sets the world spawn.", nil, SetWorldSpawn{}))

	cmd.Register(cmd.New("mymoney", "Your balance.", nil, economy.MyBalance{}))
	cmd.Register(cmd.New("setmoney", "Set balance.", nil, economy.SetBalance{}))
	cmd.Register(cmd.New("givemoney", "Give balance.", nil, economy.GiveBalance{}))
	cmd.Register(cmd.New("paymoney", "Pay balance to player", []string{"pay"}, economy.PayBalance{}))

	cmd.Register(cmd.New("npc", "NPC commands", nil, npc.Create{}, npc.Delete{}))

	cmd.Register(cmd.New("skyblock", "Manages skyblock and their settings.", []string{"sb"},
		command.Claim{},
		command.List{},
		command.Teleport{},
		command.Clear{},
		command.Auto{},
	))

	cmd.Register(cmd.New("world", "World commands", nil, world.List{}, world.Teleport{}))
}
