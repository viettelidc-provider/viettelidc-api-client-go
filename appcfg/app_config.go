package appcfg

var (
	appMode Mode = ModeDevelopment
)

type Mode int

const (
	ModeDevelopment Mode = iota
	ModeProduction
)

func (m Mode) IsProductionMode() bool {
	return m == ModeProduction
}

func IsProductionMode() bool {
	return appMode.IsProductionMode()
}
func SetMode(mode Mode) {
	appMode = mode
}

func (m Mode) String() string {
	switch m {
	case ModeDevelopment:
		return "dev"
	case ModeProduction:
		return "prod"
	default:
		return "unknown"
	}
}
