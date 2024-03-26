package cmd

import (
	"context"
	"fmt"
	"starfish/starfish/globalflag"
	"starfish/starfish/loginit"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starfish app",
	Long:  `The starfish program .`,
	Run: func(cmd *cobra.Command, args []string) {
		isDaemon, _ := cmd.Flags().GetBool("daemon")
		globalflag.IsDaemon = isDaemon
		fmt.Println("starfish app with isDaemon", isDaemon)
		implFunc()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func implFunc() {

	loginit.InitLog()
	logrus.Infoln("starfish is launch")
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	logrus.Infoln("starfish graceful closed.")
}
