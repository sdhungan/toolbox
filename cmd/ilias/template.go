/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package ilias

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var (
	hclDir    string
	excelDir  string
	outputDir string
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Create a template file based on the excel file and hcl file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Flags are required, input files should be checked if they exist and output file should be checked if directory exits
		err := validateInput(hclDir, excelDir, outputDir)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	templateCmd.Flags().StringVarP(&hclDir, "hclDir", "H", "", "Directory of HCL file")
	templateCmd.Flags().StringVarP(&excelDir, "excelDir", "e", "", "Directory of excel file")
	templateCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "", "Directory of output file")

	templateCmd.MarkFlagRequired("hclDir")
	templateCmd.MarkFlagRequired("excelDir")
	templateCmd.MarkFlagRequired("outputDir")

	HclTempCmd.AddCommand(templateCmd)
}

func validateInput(hcldir string, exceldir string, outputdir string) error {
	var validationErrors []string

	if err := validatePath(hcldir, ".hcl"); err != nil {
		validationErrors = append(validationErrors, err.Error())
	}
	if err := validatePath(exceldir, ".xlsx"); err != nil {
		validationErrors = append(validationErrors, err.Error())
	}
	if err := validatePath(outputdir, ""); err != nil {
		validationErrors = append(validationErrors, err.Error())
	}

	if len(validationErrors) > 0 {
		return errors.New("validation errors:\n" + strings.Join(validationErrors, "\n"))
	}

	return nil
}

func validatePath(path string, extension string) error {
	// Checks if path exists as a file with the given extension or as a folder if extension == ""

	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if extension == "" {
		// Check if path is a directory
		if !info.IsDir() {
			return errors.New(path + " should be the output directory for the template not a file")
		}
		return nil
	} else {
		// Search for a specific file extension
		if info.IsDir() {
			return errors.New(path + " is a directory but should be a file")
		}
		if filepath.Ext(path) != extension {
			return errors.New(path + " should end with " + extension)
		}
		return nil
	}
}
