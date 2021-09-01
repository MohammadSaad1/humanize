package da

import (
	"github.com/imbue11235/humanize/language"
)

// Language ...
var Language = language.Map{
	"slice": language.Map{
		"connector": "og",
		"rest":      "èn anden|%d andre",
	},
	"time": language.Map{
		"now":    "lige nu",
		"future": "om %s",
		"past":   "%s siden",
		"approximation": language.Map{
			"second": "et sekund|%d sekunder",
			"minute": "et minut|%d minutter",
			"hour":   "en time|%d timer",
			"day":    "en dag|%d dage",
			"week":   "en uge|%d uger",
			"month":  "en måned|%d måneder",
			"year":   "et år|%d år",
			"decade": "et årti|%d årtier",
			"long":   "lang tid",
		},
		"precision": language.Map{
			"second": "1 sekund|%d sekunder",
			"minute": "1 minut|%d minutter",
			"hour":   "1 time|%d timer",
			"day":    "1 dag|%d dage",
			"month":  "1 måned|%d måneder",
			"year":   "1 år|%d år",
		},
	},
	"int": language.Map{
		"thousand":    "1 tusind|%d tusinde",
		"million":     "1 million|%d millioner",
		"billion":     "1 milliard|%d milliarder",
		"trillion":    "1 billion|%d billioner",
		"quadrillion": "1 billiard|%d billiarder",
		"quintillion": "1 trillion|%d trillioner",
		"sextillion":  "1 trilliard|%d trilliarder",
		"septillion":  "1 kvadrillion|%d kvadrillioner",
		"octillion":   "1 kvadrilliard|%d kvadrilliarder",
	},
}
