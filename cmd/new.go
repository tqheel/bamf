/*
Copyright © 2020 Todd E. Qualls <tqualls@gmail.com>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/hectane/go-acl"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new blog post template in markdown format",
	Long:  `Creates a new blog post template`,
	Run: func(cmd *cobra.Command, args []string) {
		newMdFile := "posts/new-post.md"
		postTemplate := getRawYamlHeader()
		// Note have to grant read/write/execute permissions on destination folder on *nix file systems
		err := ioutil.WriteFile(newMdFile, postTemplate, os.ModeExclusive)
		if err != nil {
			fmt.Println(err)
		}
		if runtime.GOOS == "windows" {
			err := acl.Chmod(newMdFile, 0777)
			if err != nil {
				panic(err)
			}
		} else {
			err := os.Chmod(newMdFile, 0777)
			if err != nil {
				panic(err)
			}
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
