package database

type Config struct {
	DSN          string `json:"-"`
	MaxOpenConns int    `json:"max-open-conns"`
	MaxIdleConns int    `json:"max-idle-conns"`
	MaxIdleTime  string `json:"max-idle-time"`
	Timeout      int    `json:"timeout"`
}