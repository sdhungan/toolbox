/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package web

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

type Film struct {
	Title    string
	Director string
}

var startWebCmd = &cobra.Command{
	Use:   "startWeb",
	Short: "Starts a web server on port 8080 in a new terminal",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Launching the web server in a new terminal window...")

		// Define the command to launch a new terminal and run the web server
		var terminalCommand *exec.Cmd

		if runtime.GOOS == "windows" {
			/* On Windows: Open a new `cmd` window on the current directory
			"cmd /C start" 							will start a new terminal
			"start cmd /K"							will keep the terminal open after the command finishes (use /C to auto close after command finishes)
			"toolbox.exe web startweb-server"		Is the actual command run in the new terminal
			*/
			terminalCommand = exec.Command("cmd", "/C", "start", "cmd", "/K", "toolbox.exe web startweb-server")
		} else {
			fmt.Println("This command is currently only supported on Windows.")
			return
		}

		// Start the new terminal process
		err := terminalCommand.Start()
		if err != nil {
			fmt.Printf("Failed to open new terminal: %s\n", err)
			return
		}

		fmt.Println("Web server is now running in a separate terminal.")
	},
}

// startWebCmd represents the startWeb command
var startWebServerCmd = &cobra.Command{
	Use:   "startweb-server",
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
			// Handles GET requests to the server on root URL (defined with http.HandleFunc("/", h1))

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

			fmt.Printf("[INFO] %s request received at %s - URL: %s\n",
				r.Method, time.Now().Format(time.RFC3339), r.URL.Path,
			)

			tmpl.Execute(w, films)

		}

		h2 := func(w http.ResponseWriter, r *http.Request) {
			// Handles POST requests to the server on /add-film/ URL (defined with http.HandleFunc("/add-film/", h2))
			title := r.PostFormValue("title")
			director := r.PostFormValue("director")

			fmt.Printf("[INFO] POST request received at %s -  URL: %s - New Film: Title: %s, Director: %s \n",
				time.Now().Format(time.RFC3339), r.URL.Path, title, director,
			)

			tmpl := template.Must(template.ParseFiles(indexHtmlPath))
			tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
		}

		// define handler functions based on the URL path
		http.HandleFunc("/", h1)
		http.HandleFunc("/add-film/", h2)

		err = http.ListenAndServe("localhost:8080", nil)
		if err != nil {
			fmt.Println("Failed to start server caused by ", err)
		}

	},
}

func init() {
	WebCmd.AddCommand(startWebCmd)
	WebCmd.AddCommand(startWebServerCmd)
}
