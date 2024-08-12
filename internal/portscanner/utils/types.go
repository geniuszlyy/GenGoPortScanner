package utils

// Структура для хранения результата сканирования
type OutputResult struct {
	Target      string
	Version     string
	Players     string
	Description string
}

// Полный ответ от сервера
type FullResponse struct {
	Players struct {
		Max    int `json:"max"`
		Online int `json:"online"`
	}

	Version struct {
		Name string `json:"name"`
	}

	Description string
}

// Ответ сервера
type Response struct {
	Players struct {
		Online int `json:"online"`
		Max    int `json:"max"`
	} `json:"players"`

	Version struct {
		Name string `json:"name"`
	} `json:"version"`
}

// Ответ MOTD
type ResponseMOTD struct {
	Description struct {
		Text string `json:"text"`
	}
}
