package entity

type InputTextInput struct {
	Text string
}

type InputTextRespond struct {
	Status int8
	TextID string
}

type SearchTextInput struct {
	SearchText string
}

type SearchTextRespond struct {
	Status int8
	TextID []string
}

type GetTextRespond struct {
	Status int8
	Text   string
}
