package goaway

var profanities = []string{
	"anal",
	"anus",
	"arrse",
	"arse",
	"assfucker",
	"asses",
	"assfucker",
	"assfukka",
	"asshole",
	"assholes",
	"asswhole",
	"ballbag",
	"balls",
	"ballsack",
	"bastard",
	"beastial",
	"beastiality",
	"bellend",
	"bestial",
	"bestiality",
	"biatch",
	"bitch",
	"bitcher",
	"bitchers",
	"bitches",
	"bitchin",
	"bitching",
	"bloody",
	"blowjob",
	"blowjobs",
	"boiolas",
	"bollock",
	"bollok",
	"boner",
	"boob",
	"boobs",
	"booobs",
	"boooobs",
	"booooobs",
	"booooooobs",
	"breasts",
	"buceta",
	"bugger",
	"bum",
	"bunnyfucker",
	"butt",
	"butthole",
	"buttmunch",
	"buttplug",
	"cock",
	"cocksucker",
	"carpetmuncher",
	"cawk",
	"chink",
	"cipa",
	"clit",
	"clitoris",
	"clits",
	"cnut",
	"cock",
	"cockface",
	"cockhead",
	"cockmunch",
	"cockmuncher",
	"cocks",
	"cocksuck",
	"cocksucked",
	"cocksucker",
	"cocksucking",
	"cocksucks",
	"cocksuka",
	"cocksukka",
	"cokmuncher",
	"coksucka",
	"coon",
	"crap",
	"cummer",
	"cumming",
	"cums",
	"cumshot",
	"cunilingus",
	"cunillingus",
	"cunnilingus",
	"cunt",
	"cuntlick",
	"cuntlicker",
	"cuntlicking",
	"cunts",
	"cyalis",
	"cyberfuc",
	"cyberfuck",
	"cyberfucked",
	"cyberfucker",
	"cyberfuckers",
	"cyberfucking",
	"damn",
	"dick",
	"dickhead",
	"dildo",
	"dildos",
	"dink",
	"dinks",
	"dirsa",
	"dlck",
	"dogfucker",
	"doggin",
	"dogging",
	"donkeyribber",
	"doosh",
	"duche",
	"dyke",
	"ejaculate",
	"ejaculated",
	"ejaculates",
	"ejaculating",
	"ejaculatings",
	"ejaculation",
	"ejakulate",
	"fuck",
	"fag",
	"fagging",
	"faggitt",
	"faggot",
	"faggs",
	"fagot",
	"fagots",
	"fags",
	"fanny",
	"fannyflaps",
	"fannyfucker",
	"fanyy",
	"fatass",
	"fcuk",
	"fcuker",
	"fcuking",
	"feck",
	"fecker",
	"felching",
	"fellate",
	"fellatio",
	"fingerfuck",
	"fingerfucked",
	"fingerfucker",
	"fingerfuckers",
	"fingerfucking",
	"fingerfucks",
	"fistfuck",
	"fistfucked",
	"fistfucker",
	"fistfuckers",
	"fistfucking",
	"fistfuckings",
	"fistfucks",
	"flange",
	"fook",
	"fooker",
	"fuck",
	"fucka",
	"fucked",
	"fucker",
	"fuckers",
	"fuckhead",
	"fuckheads",
	"fuckin",
	"fucking",
	"fuckings",
	"fuckingshitmotherfucker",
	"fuckme",
	"fucks",
	"fuckwhit",
	"fuckwit",
	"fudgepacker",
	"fuk",
	"fuker",
	"fukker",
	"fukkin",
	"fuks",
	"fukwhit",
	"fukwit",
	"fux",
	"fuxor",
	"gangbang",
	"gangbanged",
	"gangbangs",
	"gaylord",
	"gaysex",
	"goatse",
	"God",
	"goddam",
	"goddamned",
	"goddamn",
	"goddamned",
	"hardcoresex",
	"hell",
	"heshe",
	"hoar",
	"hoare",
	"hoer",
	"homo",
	"hore",
	"horniest",
	"horny",
	"hotsex",
	"incest",
	"jackoff",
	"jackoff",
	"jap",
	"jerk",
	"jerkoff",
	"jism",
	"jiz",
	"jizm",
	"jizz",
	"kawk",
	"knobead",
	"knobed",
	"knobend",
	"knobhead",
	"knobjocky",
	"knobjokey",
	"kock",
	"kondum",
	"kondums",
	"kum",
	"kummer",
	"kumming",
	"kums",
	"kunilingus",
	"labia",
	"masochist",
	"masterbate",
	"masterbat",
	"masterbate",
	"masterbation",
	"masterbations",
	"masturbate",
	"materbation",
	"mofo",
	"mothafuck",
	"mothafucka",
	"mothafuckas",
	"mothafuckaz",
	"mothafucked",
	"mothafucker",
	"mothafuckers",
	"mothafuckin",
	"mothafucking",
	"mothafuckings",
	"mothafucks",
	"motherfuck",
	"motherfucked",
	"motherfucker",
	"motherfuckers",
	"motherfuckin",
	"motherfucking",
	"motherfuckings",
	"motherfuckka",
	"motherfucks",
	"muff",
	"mutha",
	"muthafecker",
	"muthafuckker",
	"muther",
	"mutherfucker",
	"nazi",
	"nigga",
	"niggah",
	"niggas",
	"niggaz",
	"nigger",
	"niggers",
	"nob",
	"nobjokey",
	"nobhead",
	"nobjocky",
	"nobjokey",
	"numbnuts",
	"nutsack",
	"orgasim",
	"orgasims",
	"orgasm",
	"orgasms",
	"pawn",
	"pecker",
	"penis",
	"penisfucker",
	"phonesex",
	"phuck",
	"phuk",
	"phuked",
	"phuking",
	"phukked",
	"phukking",
	"phuks",
	"phuq",
	"pigfucker",
	"pimpis",
	"piss",
	"pissed",
	"pisser",
	"pissers",
	"pisses",
	"pissflaps",
	"pissin",
	"pissing",
	"pissoff",
	"poop",
	"porn",
	"porno",
	"pornos",
	"prick",
	"pricks",
	"pron",
	"pube",
	"pusse",
	"pussi",
	"pussies",
	"pussy",
	"pussys",
	"queer",
	"rectum",
	"retard",
	"rimjob",
	"rimjaw",
	"rimming",
	"schlong",
	"screwing",
	"scroat",
	"scrote",
	"scrotum",
	"semen",
	"shag",
	"shagger",
	"shaggin",
	"shagging",
	"shemale",
	"shit",
	"shitdick",
	"shite",
	"shited",
	"shitey",
	"shitfuck",
	"shitfull",
	"shithead",
	"shiting",
	"shitings",
	"shits",
	"shitted",
	"shitter",
	"shitters",
	"shitting",
	"shittings",
	"shitty",
	"skank",
	"slut",
	"sluts",
	"smegma",
	"smut",
	"snatch",
	"sonofabitch",
	"spac",
	"spunk",
	"suckmy",
	"teets",
	"teez",
	"testical",
	"testicle",
	"tit",
	"titfuck",
	"tits",
	"titt",
	"tittiefucker",
	"titties",
	"tittyfuck",
	"tittywank",
	"titwank",
	"tosser",
	"turd",
	"twat",
	"twathead",
	"twatty",
	"twunt",
	"twunter",
	"vagina",
	"viagra",
	"vulva",
	"woose",
	"wang",
	"wank",
	"wanker",
	"wanky",
	"whoar",
	"whore",
	"willies",
	"willy",
}
