package reflectutil

import (
	"reflect"
	"unicode"
)

type GetTagDefaultStrategy interface {
	Acquire(sf reflect.StructField) string
}

type FirstLowerCaseFieldNameStrategy struct{}

func (s *FirstLowerCaseFieldNameStrategy) Acquire(sf reflect.StructField) string {
	nameRune := []rune(sf.Name)
	nameRune[0] = unicode.ToLower(nameRune[0])
	return string(nameRune)
}

type FieldNameStrategy struct{}

func (s *FieldNameStrategy) Acquire(sf reflect.StructField) string {
	return sf.Name
}

type GetTagNameConfig struct {
	TagKey          string
	DefaultStrategy GetTagDefaultStrategy
}

type GetTagNameOption func(config *GetTagNameConfig)

func WithTagKey(tagKey string) GetTagNameOption {
	return func(config *GetTagNameConfig) {
		config.TagKey = tagKey
	}
}

func WithDefaultStrategy(defaultStrategy GetTagDefaultStrategy) GetTagNameOption {
	return func(config *GetTagNameConfig) {
		config.DefaultStrategy = defaultStrategy
	}
}

func NewDefaultGetTagNameConfig(opts ...GetTagNameOption) *GetTagNameConfig {
	config := &GetTagNameConfig{
		DefaultStrategy: &FirstLowerCaseFieldNameStrategy{},
	}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

func GetTagNames(structT interface{}, opts ...GetTagNameOption) []string {
	config := NewDefaultGetTagNameConfig(opts...)
	t := reflect.TypeOf(structT)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	ret := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagName := field.Tag.Get(config.TagKey)
		if len(tagName) == 0 {
			tagName = config.DefaultStrategy.Acquire(field)
		}
		ret = append(ret, tagName)
	}
	return ret
}
