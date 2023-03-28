package gpt

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func printSuggestionsTable(contents string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"original name", "improved name", "reason"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	)

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.FgGreenColor},
		tablewriter.Colors{},
	)

	suggestions := strings.Split(contents, "\n")
	for _, suggestion := range suggestions {
		if strings.HasPrefix(suggestion, "{") && strings.HasSuffix(suggestion, "}") {
			suggestion = suggestion[1 : len(suggestion)-1]
		}

		fields := strings.Split(suggestion, " | ")
		if len(fields) == 3 {
			table.Append(fields)
		}
	}
	table.SetRowLine(true) // Enable row line
	table.Render()
}
