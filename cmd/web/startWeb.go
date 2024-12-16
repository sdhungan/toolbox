/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package web

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

type Film struct {
	Title    string
	Director string
}

// startWebCmd represents the startWeb command
var startWebCmd = &cobra.Command{
	Use:   "startWeb",
	Short: "Stars a web server on port 8080 in a new ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Server is running on port 8080")

		exePath, err := os.Executable()
		if err != nil {
			fmt.Println("Failed to get executable path caused by ", err)
			return
		}

		exeDir := filepath.Dir(exePath)

		indexHtmlPath := filepath.Join(exeDir, "cmd", "web", "index.html")

		h1 := func(w http.ResponseWriter, r *http.Request) {

			if r.Method == "GET" {
				tmpl := template.Must(template.ParseFiles(indexHtmlPath))

				films := map[string][]Film{
					"Films": {
						{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
						{Title: "The Godfather", Director: "Francis Ford Coppola"},
						{Title: "The Dark Knight", Director: "Christopher Nolan"},
						{Title: "The Lord of the Rings: The Return of the King", Director: "Peter Jackson"},
						{Title: "Pulp Fiction", Director: "Quentin Tarantino"},
					},
				}

				tmpl.Execute(w, films)

				fmt.Printf("[INFO] %s request received at %s - URL: %s\n",
					r.Method, time.Now().Format(time.RFC3339), r.URL.Path,
				)
			} else if r.Method == "POST" {
				title := r.PostFormValue("title")
				director := r.PostFormValue("director")

				fmt.Printf("[INFO] POST request received at %s - New Film: Title: %s, Director: %s  ***\n",
					time.Now().Format(time.RFC3339), title, director,
				)

				tmpl := template.Must(template.ParseFiles(indexHtmlPath))
				tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
			}
		}

		// define handlers
		http.HandleFunc("/", h1)
		//http.HandleFunc("/add-film/", h2)

		http.ListenAndServe("localhost:8080", nil)
		fmt.Println("startWeb called")
	},
}

func init() {
	WebCmd.AddCommand(startWebCmd)
}
