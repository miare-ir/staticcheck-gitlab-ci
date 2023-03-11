package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
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
			log.Fatal(err)
		}

		var gitlabEntry GitlabCIEntry
		gitlabEntry.Description = entry.Message
		gitlabEntry.Fingerprint = fmt.Sprintf("%s%s%d%d", entry.Code, entry.Location.File, entry.Location.Line, entry.Location.Column)
		gitlabEntry.Severity = entry.Severity

		gitlabEntry.Location.Path = getRelativePath(entry.Location.File)
		gitlabEntry.Location.Lines.Begin = entry.Location.Line
		gitlabEntries = append(gitlabEntries, gitlabEntry)
	}
	gitlabJson, err := json.Marshal(gitlabEntries)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
	fmt.Println(string(gitlabJson))
}

func getRelativePath(absolutePath string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.ReplaceAll(absolutePath, path+"/", "")
}
