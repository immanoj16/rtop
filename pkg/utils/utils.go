package utils

import (
	"log"
	"os/user"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

const (
	defaultPort  = 22
	maxPortRange = 65536
	minPortRange = 0
)

// ParseUserAndHost parses the argument passed to rtop (username@12.12.121.12:23)
func ParseUserAndHost(args []string, currentUser *user.User) (username string, host string) {
	if len(args) == 1 {
		argHost := args[0]
		if i := strings.Index(argHost, "@"); i != -1 {
			username = argHost[:i]
			if i+1 >= len(argHost) {
				pflag.Usage()
			}
			host = argHost[i+1:]
		} else {
			host = argHost
		}
	} else {
		pflag.Usage()
	}
	if len(username) == 0 {
		username = currentUser.Username
	}
	return
}

// ParseHostAndPort parses the host and port address (213.23.23.23:2343)
func ParseHostAndPort(addr string) (host string, port int) {
	if p := strings.Split(addr, ":"); len(p) == 2 {
		host = p[0]
		var err error
		if port, err = strconv.Atoi(p[1]); err != nil {
			log.Printf("bad port: %v", err)
			pflag.Usage()
		}
		if port <= minPortRange || port >= maxPortRange {
			log.Printf("bad port: %v", err)
			pflag.Usage()
		}
	} else {
		host = addr
	}

	if port == 0 {
		port = defaultPort
	}
	return
}
