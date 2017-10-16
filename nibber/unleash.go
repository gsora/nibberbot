package nibber

// List of the string -> emoji associations.
//
// Just add another tuple and OrderedSubstitution will make it available.
var Emojis = map[string]string{
	"b":       "\xF0\x9F\x85\xB1",
	"g":       "\xF0\x9F\x85\xB1",
	"a":       "\xF0\x9F\x85\xB0",
	"o":       "\xF0\x9F\x85\xBE",
	"p":       "\xF0\x9F\x85\xBF",
	"ab":      "\xF0\x9F\x86\x8E",
	"cl":      "\xF0\x9F\x86\x91",
	"c":       "\xC2\xA9",
	"r":       "\xC2\xAE",
	"i":       "\xE2\x84\xB9",
	"x":       "\xE2\x9D\x8C",
	"?":       "\xE2\x9D\x93",
	"!":       "\xE2\x9D\x97",
	"docker":  "\xF0\x9F\x90\xB3",
	"1":       "\x31\xE2\x83\xA3",
	"2":       "\x32\xE2\x83\xA3",
	"3":       "\x33\xE2\x83\xA3",
	"4":       "\x34\xE2\x83\xA3",
	"5":       "\x35\xE2\x83\xA3",
	"6":       "\x36\xE2\x83\xA3",
	"7":       "\x37\xE2\x83\xA3",
	"8":       "\x38\xE2\x83\xA3",
	"9":       "\x39\xE2\x83\xA3",
	"0":       "\x30\xE2\x83\xA3",
	"!!":      "\xE2\x80\xBC",
	"!?":      "\xE2\x81\x89",
	"peroni":  "\xF0\x9F\x8D\xBA",
	"2peroni": "\xF0\x9F\x8D\xBB",
	"allah":   "\xF0\x9F\x91\xB3 \xF0\x9F\x92\xA3",
	"ngul":    "\xF0\x9F\x86\x96 \xF0\x9F\x86\x92",
	"sos":     "\xF0\x9F\x86\x98",
	"ok":      "\xF0\x9F\x86\x97",
	"nsfw":    "\xF0\x9F\x94\x9E"
}
