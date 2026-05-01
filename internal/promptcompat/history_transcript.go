package promptcompat

import (
	"strings"
)

const CurrentInputContextFilename = "history.txt"

func BuildOpenAIHistoryTranscript(messages []any) string {
	return buildOpenAIInjectedFileTranscript(messages)
}

func BuildOpenAICurrentUserInputTranscript(text string) string {
	if strings.TrimSpace(text) == "" {
		return ""
	}
	return BuildOpenAICurrentInputContextTranscript([]any{
		map[string]any{"role": "user", "content": text},
	})
}

func BuildOpenAICurrentInputContextTranscript(messages []any) string {
	return buildOpenAIInjectedFileTranscript(messages)
}

func buildOpenAIInjectedFileTranscript(messages []any) string {
	normalized := NormalizeOpenAIMessagesForPrompt(messages, "")
	if len(normalized) == 0 {
		return ""
	}
	var b strings.Builder
	for _, m := range normalized {
		role, _ := m["role"].(string)
		content, _ := m["content"].(string)
		content = strings.TrimSpace(content)
		if content == "" {
			continue
		}
		switch role {
		case "system":
			b.WriteString("System: ")
			b.WriteString(content)
			b.WriteString("\n\n")
		case "assistant":
			b.WriteString("Assistant: ")
			b.WriteString(content)
			b.WriteString("\n\n")
		case "tool":
			b.WriteString("Tool: ")
			b.WriteString(content)
			b.WriteString("\n\n")
		default:
			b.WriteString("User: ")
			b.WriteString(content)
			b.WriteString("\n\n")
		}
	}
	return strings.TrimSpace(b.String())
}
