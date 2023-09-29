package cli

import (
	"errors"
	"github.com/FirinKinuo/rename4plex/cli/operation"
	"github.com/FirinKinuo/rename4plex/config"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"io/fs"
)

type CommandProvider interface {
	Command() *cobra.Command
}

type Rename4Plex struct {
	command *cobra.Command
	config  *config.Config
}

func NewRename4Plex(version string, cfg *config.Config) *Rename4Plex {
	r4p := &Rename4Plex{
		config: cfg,
	}

	r4p.command = &cobra.Command{
		Use: "rename4plex [FILE/FOLDER]...",
		Long: `Rename video files for Plex.

For example, rename Ugly Name Season 2 [10] [BDRip 1080p HEVC].mkv to Ugly Name s02e10.mkv

OR for all files in folder:
  UglyNameFolder
    |- Ugly Name Season 2 [5] [BDRip 1080p HEVC].mkv
    |- Ugly Name Season 2 [10] [BDRip 1080p HEVC].mkv
to:
  UglyName
    |- Ugly Name s02e5.mkv
    |- Ugly Name s02e10.mkv
`,
		Version: version,
		Args:    cobra.MinimumNArgs(1),
		Run:     r4p.run,
	}

	availableCommands := r4p.availableCommands()
	r4p.AddCommand(availableCommands...)

	return r4p
}

func (r *Rename4Plex) availableCommands() []CommandProvider {
	init := operation.NewInit(r.config)

	return []CommandProvider{
		init,
	}
}

func (r *Rename4Plex) AddCommand(provider ...CommandProvider) {
	for _, commandProvider := range provider {
		r.command.AddCommand(commandProvider.Command())
	}

}

func (r *Rename4Plex) run(_ *cobra.Command, _ []string) {
	err := r.config.Read()
	if errors.Is(err, fs.ErrNotExist) {
		log.Fatal("no configuration, need init, see --help")
	}
	if err != nil {
		log.Fatal("config", "err", err)
	}

	for s, match := range r.config.MatchGroups {
		libPath := r.config.DefaultLibraryPath
		if match.HasLibraryPath() {
			libPath = match.LibraryPath
		}

		log.Infof("group: %s, library: %s, patters: %v", s, libPath, match.Patterns)
	}
}

func (r *Rename4Plex) Execute() error {
	return r.command.Execute()
}
