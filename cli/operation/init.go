package operation

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

type ConfigCreator interface {
	Source() string
	Create() error
}

type Init struct {
	command *cobra.Command
	config  ConfigCreator
}

func NewInit(config ConfigCreator) *Init {
	init := &Init{
		config: config,
	}

	init.command = &cobra.Command{
		Use:   "init",
		Short: "Initializes configuration",
		Run:   init.run,
	}

	return init
}

func (i *Init) run(_ *cobra.Command, _ []string) {
	err := i.config.Create()
	if err != nil {
		log.Fatal("create config", "err", err)
	}

	log.Info("success creating config", "path", i.config.Source())
}

func (i *Init) Command() *cobra.Command {
	return i.command
}
