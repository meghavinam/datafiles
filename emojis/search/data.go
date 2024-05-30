package search

type Emoji struct {
	Emoji   string
	Label   string
	Tags    []string
	Emoicon string
}

var emjis = []Emoji{
	{Emoji: "smile", Label: "Smile emoji", Tags: []string{"face", "grin"}, Emoicon: ""},
	{Emoji: "cry", Label: "cry emoji", Tags: []string{"face", "grin", "eye"}, Emoicon: ""},
	{Emoji: "wonder", Label: "Wonder emoji", Tags: []string{"face", "grin", "mouth"}, Emoicon: ""},
}
