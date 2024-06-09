package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	projectType string
	projectPath string
)

var rootCmd = &cobra.Command{
	Use:   "gen-dir",
	Short: "Generate a project directory structure",
	Long:  "Generate a predefined project directory structure based on the specified project type.",
	Run: func(cmd *cobra.Command, args []string) {
		generateStructure(projectType, projectPath)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&projectType, "type", "t", "", "Type of project (go, backend, cli)")
	rootCmd.Flags().StringVarP(&projectPath, "path", "p", "", "Path to generate")
	rootCmd.MarkFlagRequired("type")
}

func generateStructure(projectType, projectPath string) {
	structure, exists := projectStructures[projectType]
	if !exists {
		fmt.Printf("Unknown project type: %s\n", projectType)
		return
	}

	for _, dir := range structure {
		fullPath := filepath.Join(projectPath, dir)
		err := os.Mkdir(fullPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			return
		}
		fmt.Printf("Created directory: %s\n", dir)
	}

	fmt.Printf("Project structure for '%s' generated successfully.\n", projectType)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error to execute the rootcmd")
		os.Exit(1)
	}
}
