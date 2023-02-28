package cmd

import (
	"api-server/api"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	port int
)
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		api.StartServer(port)
	},
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "Default Port For HTTP Server")
	if err := serverCmd.MarkFlagRequired("port"); err != nil {
		fmt.Println(err)
	}
	rootCmd.AddCommand(serverCmd)

}
