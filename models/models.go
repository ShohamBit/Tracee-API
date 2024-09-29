package models

// Holds the server info
type ServerInfo struct {
	IP   string
	Port string
}

// Address returns the full server address as IP:Port
func (s *ServerInfo) Address() string {
	return s.IP + ":" + s.Port
}

var (
	DefaultIP   = "localhost"
	DefaultPort = "4466"
)
