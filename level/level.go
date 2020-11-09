package level

type LogLevel int

const (
	LevelInfo LogLevel = iota
	LevelWarning
	LevelError
)


func (ll LogLevel) String() string {
	return [...]string{"[INFO]", "[WARNING]", "[ERROR]"}[ll]
}

func (ll LogLevel) LongestLen() int {
	return 9
}