package mosestokenizer

import (
	"regexp"
	"strings"
)

type regex struct {
	regex *regexp.Regexp
	replace string
}

var normalizer = map[string][]regex {
	"whitespace": {
		// Extra Whitespace
		{regexp.MustCompile(`\r`), ``},
		{regexp.MustCompile(`\(`), ` (`},
		{regexp.MustCompile(`\)`), `) `},
		{regexp.MustCompile(` +`), ` `},
		{regexp.MustCompile(`\) ([.!:?;,])`), `)$1`},
		{regexp.MustCompile(`\( `), `(`},
		{regexp.MustCompile(` \)`), `)`},
		{regexp.MustCompile(`(\d) %`), `$1%`},
		{regexp.MustCompile(` ([:;])`), `$1`},
	},
	"unicode": {
		{regexp.MustCompile("`"), `'`},
		{regexp.MustCompile(`''`), ` " `},
		{regexp.MustCompile(`[„“”]`), `"`},
		{regexp.MustCompile(`–`), `-`},
		{regexp.MustCompile(`—`), ` - `},
		{regexp.MustCompile(` +`), ` `},
		{regexp.MustCompile(`´`), `'`},
		{regexp.MustCompile(`([a-zA-Z])[‘’]([a-zA-Z])`), `$1'$2`},
		{regexp.MustCompile(`[‘‚’]|''|´´`), `"`},
		{regexp.MustCompile(`…`), `...`},
	},
	"french_quotes": {
		{regexp.MustCompile(` « |« |«| » | »|»`), `"`},
	},
	"pseudo_spaces": {
		{regexp.MustCompile(` ([%:?!;])`), `$1`},
		{regexp.MustCompile(`([nº|,]) `), `$1 `},
		{regexp.MustCompile(` (ºC|cm)`), ` $1`},
	},
	"en_quote_comma": {
		{regexp.MustCompile(`"([,.]+)`), `$1"`},
	},
	"de_es_fr_quote_comma": {
		{regexp.MustCompile(`,"`), `",`},
		{regexp.MustCompile(`(\.+)"(\s*[^<])`), `"$1$2`},
	},
}

func Normalize(text string, lang string) string {
	options := []string{"whitespace", "unicode", "french_quotes", "pseudo_spaces"}
	if lang == "en" {
		options = append(options, "en_quote_comma")
	} else if lang == "de" || lang == "es" || lang == "fr" {
		options = append(options, "de_es_fr_quote_comma")
	}
	for _, option := range options {
		for _, queue := range normalizer[option] {
			text = queue.regex.ReplaceAllString(text, queue.replace)
		}
	}
	text = removeAsciiJunk(text)
	return strings.TrimSpace(text)
}

func removeAsciiJunk(text string) string {
	var result strings.Builder
	for _, r := range text {
		if r >= 32 {
			result.WriteRune(r)
		}
	}
	return result.String()
}
