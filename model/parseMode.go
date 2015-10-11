package model

// ParseMode describes how a message should be parsed client-side
type ParseMode string

//ParseModes
const (
	ModeMarkdown = ParseMode("Markdown") // Parse as Markdown
	ModeDefault  = ParseMode("")         //Parse as text
)
