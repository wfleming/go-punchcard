package punchcard

import (
	"testing"
	"time"
)

func TestParseLogLineGoodInLine(t *testing.T) {
	parsedTime, _ := time.Parse(time.RFC3339, "2015-02-10T15:30:10Z")
	line := "IN\t2015-02-10T15:30:10Z\tcoding"
	logline, err := parseLogLine(line)
	if err != nil {
		t.Error("Unexpected parsing error", err)
	} else if logline.project != "coding" {
		t.Error("LogLine has wrong project name")
	} else if !logline.time.Equal(parsedTime) {
		t.Error("LogLine has wrong time")
	} else if IN != logline.action {
		t.Error("LogLine has wrong action")
	}
}

func TestParseLogLineGoodOutLine(t *testing.T) {
	parsedTime, _ := time.Parse(time.RFC3339, "2015-02-10T15:30:10Z")
	line := "OUT\t2015-02-10T15:30:10Z"
	logline, err := parseLogLine(line)
	if err != nil {
		t.Error("Unexpected parsing error", err)
	} else if logline.project != "" {
		t.Error("LogLine has wrong project name")
	} else if !logline.time.Equal(parsedTime) {
		t.Error("LogLine has wrong time")
	} else if OUT != logline.action {
		t.Error("LogLine has wrong action")
	}
}

func TestParseLogLineBadOutLine(t *testing.T) {
	line := "OUT\t2015-02-10T15:30:10Z\tfoo bar"
	if _, err := parseLogLine(line); err == nil {
		t.Error("Parsing should have failed")
	}
}

func TestParseLogLineBadInLine1(t *testing.T) {
	line := "IN\t2015-02-10T15:30:10Z"
	if _, err := parseLogLine(line); err == nil {
		t.Error("Parsing should have failed")
	}
}

func TestParseLogLineBadInLine2(t *testing.T) {
	line := "FOO\t2015-02-10T15:30:10Z\tfoo bar"
	if _, err := parseLogLine(line); err == nil {
		t.Error("Parsing should have failed")
	}
}

func TestParseLogLineBadInLine3(t *testing.T) {
	line := "IN\tbad-time\tfoo bar"
	if _, err := parseLogLine(line); err == nil {
		t.Error("Parsing should have failed")
	}
}

func TestLogLineStringInLine(t *testing.T) {
	parsedTime, _ := time.Parse(time.RFC3339, "2015-02-10T15:30:10Z")
	logline := LogLine{IN, parsedTime, "project name"}
	strline := logline.String()
	if "IN\t2015-02-10T15:30:10Z\tproject name" != strline {
		t.Error("Not expected line string", strline)
	}
}

func TestLogLineStringOutLine(t *testing.T) {
	parsedTime, _ := time.Parse(time.RFC3339, "2015-02-10T15:30:10Z")
	logline := LogLine{OUT, parsedTime, ""}
	strline := logline.String()
	if "OUT\t2015-02-10T15:30:10Z" != strline {
		t.Error("Not expected line string", strline)
	}
}

func TestLogLineStringBadLine(t *testing.T) {
	parsedTime, _ := time.Parse(time.RFC3339, "2015-02-10T15:30:10Z")
	logline := LogLine{-1, parsedTime, "project name"}
	strline := logline.String()
	if "INVALID_LOG_LINE_ACTION" != strline {
		t.Error("Not expected line string", strline)
	}
}
