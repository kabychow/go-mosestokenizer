package mosestokenizer

import (
	"github.com/khaibin/go-mosestokenizer/nonbreaking_prefix"
	"github.com/khaibin/go-mosestokenizer/perluniprops"
	"regexp"
	"strings"
	"unicode"
)

var alpha = perluniprops.ALPHA
var num = perluniprops.NUM
var alnum = perluniprops.ALNUM

var rPadNotAlnum = regex{regexp.MustCompile("([^" + alnum + " .'`,-])"), ` $1 `}
var rTokenEndsWithPeriod = regexp.MustCompile(`^(\S+)\.$`)

var rReplaceMultidots = []regex{
	{regexp.MustCompile(`\.([.]+)`), ` DOTMULTI$1`},
	{regexp.MustCompile(`DOTMULTI\.([^.])`), `DOTDOTMULTI $1`},
	{regexp.MustCompile(`DOTMULTI\.`), `DOTDOTMULTI`},
}

var rRestoreMultidots = []regex{
	{regexp.MustCompile(`DOTDOTMULTI`), `DOTMULTI.`},
	{regexp.MustCompile(`DOTMULTI`), `.`},
}

var rCommaSeparate = []regex{
	{regexp.MustCompile("([^" + perluniprops.NUM + "])[,]"), `$1 , `},
	{regexp.MustCompile("[,]([^" + num + "])"), ` , $1`},
	{regexp.MustCompile("([" + num + "])[,]$"), `$1 , `},
}

var rApostrophe = map[string][]regex{
	"en": {
		{regexp.MustCompile("([^" + alpha + "])[']([^" + alpha + "])"), `$1 ' $2`},
		{regexp.MustCompile("([^" + alnum + "])[']([" + alpha + "])"), `$1 ' $2`},
		{regexp.MustCompile("([" + alpha + "])[']([^" + alpha + "])"), `$1 ' $2`},
		{regexp.MustCompile("([" + alpha + "])[']([" + alpha + "])"), `$1 '$2`},
		{regexp.MustCompile("([" + num + "])[']([s])"), `$1 '$2`},
	},
	"fr_it": {
		{regexp.MustCompile("([^" + alpha + "])[']([^" + alpha + "])"), `$1 ' $2`},
		{regexp.MustCompile("([^" + alpha + "])[']([" + alpha + "])"), `$1 ' $2`},
		{regexp.MustCompile("([" + alpha + "])[']([^" + alpha + "])"), `$1 ' $2`},
		{regexp.MustCompile("([" + alpha + "])[']([" + alpha + "])"), `$1 '$2`},
	},
	"other": {
		{regexp.MustCompile(`'`), ` ' `},
	},
}

func TokenizeAsString(text, lang string) string {
	text = Normalize(text, lang)
	text = rPadNotAlnum.regex.ReplaceAllString(text, rPadNotAlnum.replace)
	text = replaceMultidots(text)
	text = commaSeparate(text)
	text = apostrophe(text, lang)
	text = handleNonBreakingPrefixes(text, lang)
	text = restoreMultidots(text)
	return text
}

func Tokenize(text, lang string) []string {
	return strings.Fields(TokenizeAsString(text, lang))
}

func replaceMultidots(text string) string {
	for i := 0; i < len(rReplaceMultidots); i++ {
		r := &rReplaceMultidots[i]
		text = r.regex.ReplaceAllString(text, r.replace)
		if i == len(rReplaceMultidots) - 1 && r.regex.MatchString(text) {
			i = 0
		}
	}
	return text
}

func restoreMultidots(text string) string {
	for i := 0; i < len(rRestoreMultidots); i++ {
		r := &rRestoreMultidots[i]
		text = r.regex.ReplaceAllString(text, r.replace)
		if i == 0 && r.regex.MatchString(text) {
			i = -1
		}
	}
	return text
}

func commaSeparate(text string) string {
	for _, r := range rCommaSeparate {
		text = r.regex.ReplaceAllString(text, r.replace)
	}
	return text
}

func apostrophe(text, lang string) string {
	var qApostrophe []regex
	if lang == "en" {
		qApostrophe = rApostrophe["en"]
	} else if lang == "fr" || lang == "it" {
		qApostrophe = rApostrophe["fr_it"]
	} else {
		qApostrophe = rApostrophe["other"]
	}
	for _, q := range qApostrophe {
		text = q.regex.ReplaceAllString(text, q.replace)
	}
	return text
}

func handleNonBreakingPrefixes(text, lang string) string {
	tokens := strings.Fields(text)
	for i, token := range tokens {
		isLastToken := i == len(tokens) - 1
		matches := rTokenEndsWithPeriod.FindStringSubmatch(token)
		if len(matches) >= 2 {
			prefix := matches[1]
			if !(strings.ContainsRune(prefix, '.') && strings.ContainsAny(prefix, alpha)) &&
				!nonbreaking_prefix.Find(prefix, lang) && !(!isLastToken && unicode.IsLower(rune(tokens[i+1][0]))) &&
				!(nonbreaking_prefix.FindNumeric(prefix, lang) && !isLastToken &&
					unicode.IsNumber(rune(tokens[i+1][0]))) {
				tokens[i] = prefix + " ."
			}
		}
	}
	return strings.Join(tokens, " ")
}