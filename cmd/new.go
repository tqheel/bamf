/*
Copyright © 2020 Todd E. Qualls <tqualls@gmail.com>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"

	// "github.com/gomarkdown/markdown"
	"gopkg.in/yaml.v2"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new blog post template in markdown format",
	Long:  `Creates a new blog post template`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello new post")
		var y postYaml
		y.getYamlTemplate()

		// fmt.Println("Title:", y.Title)
		// fmt.Println("Post Date:", y.PostDate)
		// fmt.Println("Post ID:", y.Id)

		postData := fmt.Sprintf("id: %s\ntitle: %s\npostDate:%s\n\n", y.Id, y.Title, y.PostDate)
		fmt.Printf(postData)

		fmt.Println("raw version:")
		fmt.Println(getRawYamlHeader())

		// ioutil.WriteFile("new-post.md")
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

type postYaml struct {
	Id       string `yaml:"id"`
	Title    string `yaml:"title"`
	PostDate string `yaml:"postDate"`
}

func (y *postYaml) getYamlTemplate() *postYaml {
	yamlFile, err := ioutil.ReadFile("./post-templates/post.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err    #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, y)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return y
}

func getRawYamlHeader() string {
	t, err := ioutil.ReadFile("./post-templates/post.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err    #%v ", err)
	}

	return string(t)
}
