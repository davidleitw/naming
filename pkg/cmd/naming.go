package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/davidleitw/naming/pkg/gpt"

	"github.com/spf13/cobra"
)

func Execute() error {
	return namingCmdRoot.Execute()
}

var namingCmdRoot = &cobra.Command{
	Use:   "naming [-f filepath]",
	Short: "🎨 naming is a command line tool that can provide program naming suggestions through the ChatGPT API.",
	Long: `🎨 naming is a command line tool based on ChatGPT API, 
designed to improve the readability of your code by suggesting 
intuitive and descriptive names for your functions and variables.`,
	Example: `
🎨 Provide naming suggestions for main.go
naming -f main.go

🎨 Provide naming suggestions for pkg/gpt/namegpt.go
naming -f pkg/gpt/namegpt.go
	`,
	RunE: runCmd,
}

var filePathFlag string

func init() {
	namingCmdRoot.Flags().StringVarP(&filePathFlag, "file", "f", "", "input file path")
}

func runCmd(cmd *cobra.Command, args []string) error {
	apiKey := os.Getenv("CHATGPT_API_KEY")
	if apiKey == "" {
		return fmt.Errorf(" CHATGPT_API_KEY is not missing, please set it in your environment variables")
	}

	if filePathFlag == "" {
		return fmt.Errorf("please provide a file path using the -f flag")
	}

	filePath, err := filepath.Abs(filePathFlag)
	if err != nil {
		return err
	}

	if fileInfo, err := os.Stat(filePath); err != nil || fileInfo.IsDir() {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// call GetNamingSuggestions to get naming suggestions for the specified file path
	gpt.NewConsultant(apiKey).GetNamingSuggestions(filePath)
	return nil
}
