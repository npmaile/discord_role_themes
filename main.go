package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/bwmarrin/discordgo"
)

func main() {
	guildID := flag.String("guildID", "", "guild you want the color scheme set for")
	flagToken := flag.String("token", "", "discord api token (you can get one of these if you register a new bot at https://discord.com/developers/applications )")
	color := flag.String("colorScheme", "TokyoNight", "colorscheme you wish to set your roles to")
	printColors := flag.Bool("getColors", false, "ignore everything else and just print a list of available colorschemes")
	flag.Parse()
	if *printColors {
		printAllColors()
		os.Exit(0)
	}
	token := *flagToken

	if *guildID == "" {
		panic("guildID not passed in")
	}
	if token == "" {
		token = os.Getenv("DISCORD_API_TOKEN")
		if token == "" {
			panic("no token found. Please specify one on the command line or as the DISCORD_API_TOKEN environment variable")
		}
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(fmt.Sprintf("dg failed to initialize: %s", err.Error()))
	}
	dg.State.TrackRoles = true
	existingRoles, err := dg.GuildRoles(*guildID)
	if err != nil {
		panic(fmt.Sprintf("unable to get roles in server: %s", err.Error()))
	}
	sort.Slice(existingRoles, func(i int, j int) bool {
		return existingRoles[i].Position > existingRoles[j].Position
	})
	colorscheme, ok := schemes[*color]
	if !ok {
		panic(fmt.Sprintf("colorscheme %s not found", *color))
	}
	for i, role := range existingRoles {
		var colorToPick int
		if i < len(colorscheme.Restricted) {
			colorToPick = colorscheme.Restricted[i]
		} else {
			colorToPick = colorscheme.Unrestricted[(i-len(colorscheme.Restricted))%len(colorscheme.Unrestricted)]
		}
		dg.GuildRoleEdit(*guildID, role.ID, &discordgo.RoleParams{
			Color: &colorToPick,
		})
	}

}

func printAllColors() {
	for name := range schemes {
		fmt.Println(name)
	}
}
