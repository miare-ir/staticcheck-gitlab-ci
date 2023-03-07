package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type StaticCheckEntry struct {
	Code     string `json:"code"`
	Severity string `json:"severity"`
	Location struct {
		File   string `json:"file"`
		Line   int    `json:"line"`
		Column int    `json:"column"`
	} `json:"location"`
	End     interface{} `json:"end"`
	Message string      `json:"message"`
}

type GitlabCIEntry struct {
	Description string `json:"description"`
	Fingerprint string `json:"fingerprint"`
	Severity    string `json:"severity"`
	Location    struct {
		Path  string `json:"path"`
		Lines struct {
			Begin int `json:"begin"`
		} `json:"lines"`
	} `json:"location"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var gitlabEntries = make([]GitlabCIEntry, 0)
	for scanner.Scan() {
		var entry StaticCheckEntry
		err := json.Unmarshal([]byte(scanner.Text()), &entry)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			continue
		}

		// Create a new GitlabCIEntry object from the StaticCheckEntry
		var gitlabEntry GitlabCIEntry
		gitlabEntry.Description = entry.Message
		gitlabEntry.Fingerprint = entry.Code
		gitlabEntry.Severity = entry.Severity

		gitlabEntry.Location.Path = entry.Location.File
		gitlabEntry.Location.Lines.Begin = entry.Location.Line
		gitlabEntries = append(gitlabEntries, gitlabEntry)
	}
	gitlabJson, err := json.Marshal(gitlabEntries)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
	fmt.Println(string(gitlabJson))
}
