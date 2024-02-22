package gocva

import "strings"

type ClassBuilder struct {
	builder strings.Builder
}

func (b *ClassBuilder) Add(class string) {
	trimmedClass := strings.TrimSpace(class)

	if len(trimmedClass) == 0 {
		return
	}

	b.builder.WriteString(trimmedClass + " ")
}

func (b *ClassBuilder) String() string {
	return strings.TrimSpace(b.builder.String())
}
