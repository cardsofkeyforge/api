package database

import (
	"errors"
	"fmt"
	"strings"
)

type FilterBuilder struct {
	filterSlice []string
	valuesSlice []interface{}
}

func (fb *FilterBuilder) Eq(attribute string, val interface{}) *FilterBuilder {
	if val != "" && val != 0 {
		fb.filterSlice = append(fb.filterSlice, fmt.Sprintf("'%s' %s ?", attribute, "="))
		fb.valuesSlice = append(fb.valuesSlice, val)
	}
	return fb
}

func (fb *FilterBuilder) Ne(attribute string, val interface{}) *FilterBuilder {
	if val != "" && val != 0 {
		fb.filterSlice = append(fb.filterSlice, fmt.Sprintf("'%s' %s ?", attribute, "!="))
		fb.valuesSlice = append(fb.valuesSlice, val)
	}
	return fb
}

func (fb *FilterBuilder) Gt(attribute string, val interface{}) *FilterBuilder {
	if val != "" && val != 0 {
		fb.filterSlice = append(fb.filterSlice, fmt.Sprintf("'%s' %s ?", attribute, ">"))
		fb.valuesSlice = append(fb.valuesSlice, val)
	}
	return fb
}

func (fb *FilterBuilder) Lt(attribute string, val interface{}) *FilterBuilder {
	if val != "" && val != 0 {
		fb.filterSlice = append(fb.filterSlice, fmt.Sprintf("'%s' %s ?", attribute, "<"))
		fb.valuesSlice = append(fb.valuesSlice, val)
	}
	return fb
}

func (fb *FilterBuilder) Ge(attribute string, val interface{}) *FilterBuilder {
	if val != "" && val != 0 {
		fb.filterSlice = append(fb.filterSlice, fmt.Sprintf("'%s' %s ?", attribute, ">="))
		fb.valuesSlice = append(fb.valuesSlice, val)
	}
	return fb
}

func (fb *FilterBuilder) Le(attribute string, val interface{}) *FilterBuilder {
	fb.filterSlice = append(fb.filterSlice, fmt.Sprintf("'%s' %s ?", attribute, "<="))
	fb.valuesSlice = append(fb.valuesSlice, val)
	return fb
}

func (fb *FilterBuilder) Bt(attribute string, val ...interface{}) *FilterBuilder {
	fb.filterSlice = append(fb.filterSlice, fmt.Sprintf("'%s' %s ? AND ?", attribute, "BETWEEN"))
	fb.valuesSlice = append(fb.valuesSlice, val)
	return fb
}

func (fb *FilterBuilder) In(attribute string, val []string) *FilterBuilder {
	if len(val) > 0 {
		fb.filterSlice = append(fb.filterSlice, fmt.Sprintf("'%s' %s ?", attribute, strings.Join(val, ",")))
		fb.valuesSlice = append(fb.valuesSlice, val)
	}
	return fb
}

func (fb *FilterBuilder) And() *FilterBuilder {
	if len(fb.filterSlice) > 0 {
		fb.filterSlice = append(fb.filterSlice, fmt.Sprintf(" %s ", "AND"))
	}
	return fb
}

func (fb *FilterBuilder) Or() *FilterBuilder {
	if len(fb.filterSlice) > 0 {
		fb.filterSlice = append(fb.filterSlice, fmt.Sprintf(" %s ", "OR"))
	}
	return fb
}

func (fb *FilterBuilder) Contains(attribute string, val interface{}) *FilterBuilder {
	if val != nil && val != "" {
		fb.filterSlice = append(fb.filterSlice, fmt.Sprintf("contains('%s', ?)", attribute))
		fb.valuesSlice = append(fb.valuesSlice, val)
	}
	return fb
}

func (fb *FilterBuilder) BeginsWith(attribute string, val interface{}) *FilterBuilder {
	if val != nil && val != "" {
		fb.filterSlice = append(fb.filterSlice, fmt.Sprintf("begins_with('%s', ?)", attribute))
		fb.valuesSlice = append(fb.valuesSlice, val)
	}
	return fb
}

func (fb *FilterBuilder) Build() (expression *string, values *[]interface{}, err error) {
	if len(fb.filterSlice) > 0 {
		last := fb.filterSlice[len(fb.filterSlice)-1]

		if strings.Contains(last, "AND") || strings.Contains(last, "OR") {
			fb.filterSlice[len(fb.filterSlice)-1] = "" // erasing reference
			fb.filterSlice = fb.filterSlice[:len(fb.filterSlice)-1]
		}

		join := strings.Join(fb.filterSlice, " ")
		expression = &join
		values = &fb.valuesSlice
	} else {
		err = errors.New("no values to build")
	}
	return
}
