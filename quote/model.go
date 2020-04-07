package quote

type responseStruct struct {
	Contents contentStruct `json:"contents"`
}

type contentStruct struct {
	Quotes []quoteStruct `json:"quotes"`
}

type quoteStruct struct {
	Quote  string `json:"quote"`
	Author string `json:author`
}
