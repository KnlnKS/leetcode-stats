package src

const systemStack = "-apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif"

func GetFont(name string) Font {
	return Font{FontFamily: systemStack}
}
