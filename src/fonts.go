package src

import (
	_ "embed"
	"encoding/base64"
	"strings"
)

//go:embed fonts/inter-400.woff2
var inter400 []byte

//go:embed fonts/inter-500.woff2
var inter500 []byte

//go:embed fonts/inter-600.woff2
var inter600 []byte

const systemStack = "-apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif"

var embeddedFont = Font{
	FontFamily: "LeetCodeStats, " + systemStack,
	FontCSS:    buildFontCSS(),
}

var systemFont = Font{FontFamily: systemStack}

func buildFontCSS() string {
	var b strings.Builder
	b.WriteString("<style>")
	for _, face := range []struct {
		weight string
		data   []byte
	}{{"400", inter400}, {"500", inter500}, {"600", inter600}} {
		b.WriteString(`@font-face{font-family:LeetCodeStats;font-style:normal;font-weight:`)
		b.WriteString(face.weight)
		b.WriteString(`;src:url("data:font/woff2;base64,`)
		b.WriteString(base64.StdEncoding.EncodeToString(face.data))
		b.WriteString(`") format("woff2")}`)
	}
	b.WriteString("</style>")
	return b.String()
}

func GetFont(name string) Font {
	if name == "system" || name == "none" {
		return systemFont
	}
	return embeddedFont
}
