package main

// Not deployed yet — see the note on regionDef.
var regionRU = regionDef{
	Lang: "ru",

	Greeting: `Привет и добро пожаловать в sh-development 👋
Меня зовут Сергей Шумилов 😎

Снизу слева - кнопка с моими приложениями

Мой Telegram - @sergey_showmelove
Мой сайт - sh-development.ru

Если у вас возникнут вопросы или предложения - you are welcome! 🙂`,

	Text: pageText{
		Title:    "auth-center",
		SignedAs: "Вы вошли как:",
		Checking: "Проверяем…",

		NoSession:   "в ссылке нет сессии входа — начните со страницы входа",
		UnknownLink: "ссылка входа не распознана",

		JS: jsText{
			Next:       "ДАЛЕЕ",
			Back:       "НАЗАД",
			Failed:     "Ошибка. Попробуйте снова",
			NoAnswer:   "сервер не ответил за %d с",
			NoSDK:      "telegram-web-app.js не загрузился — страница должна открываться внутри Telegram",
			NoInitData: "нет initData — страница открыта не по ссылке входа",
		},
	},
}
