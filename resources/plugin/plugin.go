package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "prismacloud"
	Kind    = "source"
	Team    = "nronix"
	Version = "1.0.0"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(Name,
		Version, Configure,
		plugin.WithKind(Kind),

		plugin.WithTeam(Team))
}
