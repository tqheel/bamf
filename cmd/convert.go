/*
Copyright © 2020 Todd E. Qualls <tqualls@gmail.com>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		mdf := getMarkdown("posts/new-post.md")

		headerLines := getHeaderLines(mdf)
		fmt.Printf("%s", headerLines)
		md := getBodyLines(mdf)
		html := markdown.ToHTML(md, nil, nil)
		fmt.Printf("%s", html)

		// TODO: create func to read html template
		getHtmlTemplate()
		// iterate through each line
		// read header rows into yaml struct
		// find hidden data elements and assign from properties of struct
		// append body of converted markdown inside of main div tag

	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getHtmlTemplate() {
	f, err := os.Open("post-templates/post.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(doc.Html())
	fmt.Println()

}

func getMarkdown(fileName string) []byte {
	m, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	return m
}

func getHeaderLines(rawBytes []byte) []byte {
	lines := strings.Split(string(rawBytes), "\n")
	var buf []byte
	for i, line := range lines {
		if i < 3 {
			// println(line)
			buf = append(buf, line+"\n"...)
		}
	}
	buf = append(buf, "\n"...)
	return buf
}

func getBodyLines(rawBytes []byte) []byte {
	lines := strings.Split(string(rawBytes), "\n")
	var buf []byte
	for i, line := range lines {
		if i >= 3 {
			// println(line)
			buf = append(buf, line+"\n"...)
		}
	}
	buf = append(buf, "\n"...)
	return buf
}

func readHtmlTemplate() {
	// TODO: read html template file and return []byte
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
