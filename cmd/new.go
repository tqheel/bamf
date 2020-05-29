/*
Copyright © 2020 Todd E. Qualls <tqualls@gmail.com>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new blog post template in markdown format",
	Long:  `Creates a new blog post template`,
	Run: func(cmd *cobra.Command, args []string) {
		postTemplate := getRawYamlHeader()
		// Note have to grant read/write/execute permissions on destination folder on *nix file systems
		err := ioutil.WriteFile("posts/new-post.md", postTemplate, os.ModeExclusive)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getRawYamlHeader() []byte {
	t, err := ioutil.ReadFile("./post-templates/post.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err    #%v ", err)
	}

	return t
}
