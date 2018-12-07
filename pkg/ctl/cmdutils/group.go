package cmdutils

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Grouping struct {
	groups map[*cobra.Command]*NamedFlagSetGroup
}

type namedFlagSet struct {
	name string
	fs   *pflag.FlagSet
}

type NamedFlagSetGroup struct {
	list []namedFlagSet
}

func NewGrouping() *Grouping {
	return &Grouping{
		make(map[*cobra.Command]*NamedFlagSetGroup),
	}
}

func (g *Grouping) New(cmd *cobra.Command) *NamedFlagSetGroup {
	n := &NamedFlagSetGroup{}
	g.groups[cmd] = n
	return n
}

func (n *NamedFlagSetGroup) InFlagSet(name string, cb func(*pflag.FlagSet)) {
	for _, nfs := range n.list {
		if nfs.name == name {
			cb(nfs.fs)
			return
		}
	}

	nfs := namedFlagSet{
		name: name,
		fs:   &pflag.FlagSet{},
	}
	cb(nfs.fs)
	n.list = append(n.list, nfs)
}

func (n *NamedFlagSetGroup) AddTo(cmd *cobra.Command) {
	for _, nfs := range n.list {
		cmd.Flags().AddFlagSet(nfs.fs)
	}
}

func (g *Grouping) Usage(cmd *cobra.Command) error {
	group := g.groups[cmd]

	usage := []string{fmt.Sprintf("Usage: %s", cmd.UseLine())}

	if cmd.HasAvailableSubCommands() {
		usage = append(usage, "\nCommands:")
		for _, subCommand := range cmd.Commands() {
			usage = append(usage, fmt.Sprintf("  %s %-10s  %s", cmd.CommandPath(), subCommand.Name(), subCommand.Short))
		}
	}

	if len(cmd.Aliases) > 0 {
		usage = append(usage, "\nAliases: "+cmd.NameAndAliases())
	}

	if group != nil {
		for _, nfs := range group.list {
			usage = append(usage, fmt.Sprintf("\n%s flags:", nfs.name))
			usage = append(usage, strings.TrimRightFunc(nfs.fs.FlagUsages(), unicode.IsSpace))
		}
	}

	usage = append(usage, "\nCommon flags:")
	if len(cmd.PersistentFlags().FlagUsages()) != 0 {
		usage = append(usage, strings.TrimRightFunc(cmd.PersistentFlags().FlagUsages(), unicode.IsSpace))
	}
	if len(cmd.InheritedFlags().FlagUsages()) != 0 {
		usage = append(usage, strings.TrimRightFunc(cmd.InheritedFlags().FlagUsages(), unicode.IsSpace))
	}

	usage = append(usage, fmt.Sprintf("\nUse '%s [command] --help' for more information about a command.\n", cmd.CommandPath()))

	fmt.Printf(strings.Join(usage, "\n"))

	return nil
}
