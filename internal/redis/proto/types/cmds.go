package types

type CMD string

const (
	CMD_PING CMD = "PING"
	CMD_ECHO CMD = "ECHO"
	CMD_SET  CMD = "SET"
	CMD_GET  CMD = "GET"
)

var CmdStringMap = map[string]CMD{
	"PING": CMD_PING,
	"ECHO": CMD_ECHO,
	"SET":  CMD_SET,
	"GET":  CMD_GET,
}

var CmdArgsMap = map[CMD]int{
	CMD_PING: 0,
	CMD_ECHO: 1,
	CMD_SET:  2,
	CMD_GET:  1,
}
