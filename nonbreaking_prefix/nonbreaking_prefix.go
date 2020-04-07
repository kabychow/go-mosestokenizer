package nonbreaking_prefix

import (
	"github.com/khaibin/go-mosestokenizer/nonbreaking_prefix/data"
	"strings"
)

var nonBreakingPrefix = map[string]map[string]interface{} {
	"ca": data.CA,
	"cs": data.CS,
	"de": data.DE,
	"el": data.EL,
	"en": data.EN,
	"es": data.ES,
	"fi": data.FI,
	"fr": data.FR,
	"ga": data.GA,
	"hu": data.HU,
	"is": data.IS,
	"it": data.IT,
	"lt": data.LT,
	"lv": data.LV,
	"nl": data.NL,
	"pl": data.PL,
	"pt": data.PT,
	"ro": data.RO,
	"ru": data.RU,
	"sk": data.SK,
	"sl": data.SL,
	"sv": data.SV,
	"ta": data.TA,
	"zh": data.ZH,
}

var nonBreakingPrefixNumeric = map[string]map[string]interface{} {
	"en": data.EN_NUMERIC,
	"ga": data.GA_NUMERIC,
	"hu": data.HU_NUMERIC,
	"is": data.IS_NUMERIC,
	"it": data.IT_NUMERIC,
	"lt": data.LT_NUMERIC,
	"lv": data.LV_NUMERIC,
	"nl": data.NL_NUMERIC,
	"pl": data.PL_NUMERIC,
	"pt": data.PT_NUMERIC,
	"sl": data.SL_NUMERIC,
	"ta": data.TA_NUMERIC,
	"zh": data.ZH_NUMERIC,
}

func Find(prefix, lang string) bool {
	return find(nonBreakingPrefix, prefix, lang)
}

func FindNumeric(prefix, lang string) bool {
	return find(nonBreakingPrefixNumeric, prefix, lang)
}

func find(from map[string]map[string]interface{}, prefix, lang string) bool {
	_, found := from[lang][strings.ToLower(prefix)]
	return found
}