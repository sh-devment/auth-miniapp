package main

var regionCOM = regionDef{
	Lang: "en",

	Greeting: `Hi and welcome to sh-development 👋
My name is Sergey Shumilov 😎

Bottom left - the button with my apps

My Telegram - @sergey_showmelove
My website - sh-development.com

If you have any questions or ideas - you are welcome! 🙂`,

	Text: pageText{
		Title:    "auth-center",
		SignedAs: "Signed in as:",
		Checking: "Checking…",

		NoSession:   "the link carries no login session — start from the login page",
		UnknownLink: "login link not recognised",

		JS: jsText{
			Next:       "NEXT",
			Back:       "BACK",
			Failed:     "Error. Please try again",
			NoAnswer:   "the server did not answer within %d s",
			NoSDK:      "telegram-web-app.js failed to load — this page must be opened inside Telegram",
			NoInitData: "no initData — the page was not opened from a login link",
		},
	},
}
