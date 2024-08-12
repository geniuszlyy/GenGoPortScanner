package utils

import "fmt"

// Форматирует результат в JSON
func FormatJSON(t OutputResult) string {
	return fmt.Sprintf(`{"target": "%s", "version": "%s", "players": "%s", "description": "%s"}`,
		t.Target, t.Version, t.Players, t.Description)
}

// Форматирует результат в CSV
func FormatCSV(t OutputResult) string {
	return fmt.Sprintf(`"%s","%s","%s","%s"`, t.Target, t.Version, t.Players, t.Description)
}

// Форматирует результат в QuBo
func FormatQubo(t OutputResult) string {
	return fmt.Sprintf("(%s)(%s)(%s)(%s)", t.Target, t.Players, t.Version, t.Description)
}
