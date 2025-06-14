package response

import (
	example "github.com/imbossa/3G/docs/proto/example"
	"github.com/imbossa/3G/internal/entity"
)

// NewTranslationHistory -.
func NewTranslationHistory(translationHistory entity.TranslationHistory) *example.GetHistoryResponse {
	history := make([]*example.TranslationHistory, len(translationHistory.History))

	for i, h := range translationHistory.History {
		history[i] = &example.TranslationHistory{
			Source:      h.Source,
			Destination: h.Destination,
			Original:    h.Original,
			Translation: h.Translation,
		}
	}

	return &example.GetHistoryResponse{History: history}
}
