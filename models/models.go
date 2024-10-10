package models

var (
	DefaultIP   = "localhost"
	DefaultPort = "4466"
)

// Holds the server info
type ServerInfo struct {
	IP   string
	Port string
}

type TestCase struct {
	Name           string
	Args           []string
	ExpectedOutput interface{}
}

// Address returns the full server address as IP:Port
func (s *ServerInfo) Address() string {
	return s.IP + ":" + s.Port
}
