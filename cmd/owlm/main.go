package main

import (
	"OWLm/configs"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

const help = `
OpenWRT 
`

var rootCmd = &cobra.Command{
	Use:   "owlm",
	Short: "Wrapper to Openwrt build-system",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to OWLm! Use --help for usage.")
	},
}
var graphCmd = &cobra.Command{
	Use:   "graph [directory]",
	Short: "Create dependency graph",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		outputFormat, _ := cmd.Flags().GetString("format")
		outputFile, _ := cmd.Flags().GetString("output")
		layout, _ := cmd.Flags().GetString("layout")

		validFormats := map[string]bool{
			"dot": true, "png": true, "svg": true,
			"pdf": true, "json": true,
		}
		if !validFormats[outputFormat] {
			fmt.Printf("Invalid format: %s. Supported: dot, html, svg, pdf, json \n", outputFormat)
			return
		}

		projectDir := getProjectDir(args)
		if projectDir == "" {
			return
		}

		python, err := findPython()
		if err != nil {
			fmt.Println(err)
			return
		}

		if outputFile == "" && outputFormat != "dot" && outputFormat != "json" && outputFormat != "txt" {
			outputFile = fmt.Sprintf("dependency_graph.%s", outputFormat)
			fmt.Printf("Output will be saved to: %s\n", outputFile)
		}

		scriptPath := filepath.Join("scripts", "internal", "graph.py")
		args = []string{scriptPath, projectDir,
			"--format", outputFormat,
			"--layout", layout}

		if outputFile != "" {
			args = append(args, "--output", outputFile)
		}

		exe := exec.Command(python, args...)
		exe.Stdout = os.Stdout
		exe.Stderr = os.Stderr

		if err := exe.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func getProjectDir(args []string) string {
	if len(args) > 0 && args[0] != "" {
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			fmt.Printf("Directory does not exist: %s\n", args[0])
			return ""
		}
		return args[0]
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return ""
	}
	return dir
}

func findPython() (string, error) {
	if python, err := exec.LookPath("python"); err == nil {
		return python, nil
	}
	if python, err := exec.LookPath("python3"); err == nil {
		return python, nil
	}
	return "", fmt.Errorf("python not found in PATH")
}

func init() {
	//--Commands-

	rootCmd.AddCommand(graphCmd)

	//--Flags--

	rootCmd.Flags().String("help", "h", help)
	rootCmd.Flags().String("version", "v", fmt.Sprintf("versino: %s", configs.CFG.Get("version")))
	graphCmd.Flags().StringP("format", "f", "dot", "Output format (dot, png, svg, pdf, json)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
