package constant

const (
	EmptyStr = ""
)

const (
	AdaptiveFormatTemplateStr = "%v"
)

const (
	JSONObjectLeftBoundary        = "{"
	JSONObjectLeftBoundarySymbol  = '{'
	JSONObjectRightBoundary       = "}"
	JSONObjectRightBoundarySymbol = '}'
	JSONArrayLeftBoundary         = "["
	JSONArrayLeftBoundarySymbol   = '['
	JSONArrayRightBoundary        = "]"
	JSONArrayRightBoundarySymbol  = ']'
	EmptyJSONObject               = JSONObjectLeftBoundary + JSONObjectRightBoundary
	EmptyJSONArray                = JSONArrayLeftBoundary + JSONArrayRightBoundary
)

const (
	EscapeSymbol = '\\'
	Escape       = "\\"
)

const (
	// CarriageReturn 回车
	CarriageReturn       = "\r"
	CarriageReturnSymbol = '\r'

	// LineFeed 换行
	LineFeed       = "\n"
	LineFeedSymbol = '\n'

	// WindowsLineSep Windows换行
	WindowsLineSep = CarriageReturn + LineFeed
	// MacOSLineSep MacOS换行
	MacOSLineSep = CarriageReturn
	// UnixLineSep Unix系统换行
	UnixLineSep = LineFeed
)
