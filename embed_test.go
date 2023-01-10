package godiscord

import "testing"

func TestEmbed(t *testing.T) {
	webhookURL := "http://localhost"

	embed := NewEmbed("some title", "some description stuff")
	embed.SetColor(`#2ECC71`)
	embed.SendToWebhook(webhookURL)
}
