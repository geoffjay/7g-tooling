package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

$ source <(7g completion bash)

# To load completions for each session, execute once:
Linux:
  $ 7g completion bash > /etc/bash_completion.d/7g
MacOS:
  $ 7g completion bash > /usr/local/etc/bash_completion.d/7g

Zsh:

$ source <(7g completion zsh)

# To load completions for each session, execute once:
$ 7g completion zsh > "${fpath[1]}/_7g"

Fish:

$ 7g completion fish | source

# To load completions for each session, execute once:
$ 7g completion fish > ~/.config/fish/completions/7g.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			_ = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			_ = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			_ = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			_ = cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}
