package message

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// Generic message type as specified in assignment
type Message struct {
	To      string
	From    string
	Title   string
	Content string
	Date    time.Time
}

// UnmarshalJSON is a custom Unmarshaler for the Date field in Message
func (m *Message) UnmarshalJSON(j []byte) error {
	var rawStrings map[string]string

	err := json.Unmarshal(j, &rawStrings)
	if err != nil {
		return err
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "to" {
			m.To = v
		}
		if strings.ToLower(k) == "from" {
			m.From = v
		}
		if strings.ToLower(k) == "title" {
			m.Title = v
		}
		if strings.ToLower(k) == "content" {
			m.Content = v
		}
		// properly parse Date field
		if strings.ToLower(k) == "date" {
			t, err := time.Parse("Jan 2, 2006 at 3:04pm (MST)", v)
			if err != nil {
				return err
			}
			m.Date = t
		}
	}
	return nil
}

// MessageFromJsonFile initializes a message from a Json filename
// TODO: Refactor to work from a []bytes or iobuffer to generalize
func MessageFromJsonFile(filename string) Message {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var m Message
	err = json.Unmarshal(byteValue, &m)
	if err != nil {
		log.Panic(err)
	}
	return m
}

// formatContent adds a tab and space to properly format line breaks in the Content field when printing
func formatContent(content string) string {
	// TODO: for extremely large JSON it would be better to do this with a buffer
	ret := ""
	for i, line := range strings.Split(strings.TrimSuffix(content, "\n"), "\n") {
		if i == 0 {
			ret += fmt.Sprintf("%s\n", line)
		} else {
			ret += fmt.Sprintf("\t %s\n", line)

		}
	}
	return ret
}

func (m *Message) String() string {
	return fmt.Sprintf("Title: %s\nTo: %s\nFrom: %s\nDate: %s\nContent: %s",
		m.Title,
		m.To,
		m.From,
		m.Date,
		formatContent(m.Content),
	)
}
