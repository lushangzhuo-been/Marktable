package i18n

import (
	"context"
	"embed"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

type I18n struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
	current   language.Tag
}

type Config struct {
	DefaultLanguage language.Tag
	Format          string
	EmbedFiles      *embed.FS
}

func New(cfg Config) (*I18n, error) {
	bundle := i18n.NewBundle(cfg.DefaultLanguage)

	switch cfg.Format {
	case "toml":
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	default:
		return nil, fmt.Errorf("unsupported format: %s", cfg.Format)
	}

	if cfg.EmbedFiles != nil {
		files, err := cfg.EmbedFiles.ReadDir(".")
		if err != nil {
			return nil, fmt.Errorf("failed to read embed files: %w", err)
		}

		for _, file := range files {
			data, err := cfg.EmbedFiles.ReadFile(file.Name())
			if err != nil {
				return nil, fmt.Errorf("failed to read embed file %s: %w", file.Name(), err)
			}

			if _, err := bundle.ParseMessageFileBytes(data, file.Name()); err != nil {
				return nil, fmt.Errorf("failed to parse embed file %s: %w", file.Name(), err)
			}
		}
	}

	return &I18n{
		bundle:  bundle,
		current: cfg.DefaultLanguage,
	}, nil
}

func (i *I18n) SetLanguage(lang string) {
	tag, err := language.Parse(lang)
	if err != nil {
		tag = i.bundle.LanguageTags()[0]
	}
	i.current = tag
	i.localizer = i18n.NewLocalizer(i.bundle, tag.String())
}

func (i *I18n) T(messageID string, data ...map[string]interface{}) string {
	if i.localizer == nil {
		i.localizer = i18n.NewLocalizer(i.bundle, i.current.String())
	}

	var templateData map[string]interface{}
	if len(data) > 0 {
		templateData = data[0]
	}

	msg, err := i.localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
	if err != nil {
		return messageID
	}
	return msg
}

type contextKey struct{}

var i18nKey = contextKey{}

func NewContext(ctx context.Context, i *I18n) context.Context {
	return context.WithValue(ctx, i18nKey, i)
}

func FromContext(ctx context.Context) *I18n {
	if val, ok := ctx.Value(i18nKey).(*I18n); ok {
		return val
	}
	return nil
}
