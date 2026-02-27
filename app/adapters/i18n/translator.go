package i18n

import (
	_ "embed"
	"encoding/json"
	"strings"
)

//go:embed locales/en.json
var enBytes []byte

//go:embed locales/th.json
var thBytes []byte

type bundle struct {
	en map[string]string
	th map[string]string
}

type Translator struct {
	b bundle
}

func New() *Translator {
	var en map[string]string
	var th map[string]string
	_ = json.Unmarshal(enBytes, &en)
	_ = json.Unmarshal(thBytes, &th)
	return &Translator{b: bundle{en: en, th: th}}
}

func (t *Translator) Translate(lang, code string, args ...string) string {
	m := t.b.en
	if strings.HasPrefix(strings.ToLower(lang), "th") {
		m = t.b.th
	}
	msg, ok := m[code]
	if !ok {
		msg = code
	}
	for i, a := range args {
		msg = strings.ReplaceAll(msg, "{"+itoa(i)+"}", a)
	}
	return msg
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	d := []byte{}
	for i > 0 {
		d = append([]byte{byte('0' + i%10)}, d...)
		i /= 10
	}
	return string(d)
}