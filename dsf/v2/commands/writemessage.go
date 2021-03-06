package commands

import "github.com/Duet3D/dsf-go/dsf/v2/machine/messages"

// WriteMessage writes an arbitrary message.
// If neither OutputMessage nor LogMessage is true the message is
// written to the console output.
type WriteMessage struct {
	BaseCommand
	// Type of the message to write
	Type messages.MessageType
	// Content of the message to write
	Content string
	// OutputMessage on the console and via the object model
	OutputMessage bool
	// LogMessage writes the message to the log file (if applicable)
	LogMessage bool
}

// NewWriteMessage creates a new WriteMessage
func NewWriteMessage(mType messages.MessageType, content string, outputMessage, logMessage bool) *WriteMessage {
	return &WriteMessage{
		BaseCommand:   *NewBaseCommand("WriteMessage"),
		Type:          mType,
		Content:       content,
		OutputMessage: outputMessage,
		LogMessage:    logMessage,
	}
}
