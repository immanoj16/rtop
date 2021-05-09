package utils

import (
	"fmt"
	"os/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseUserAndHost(t *testing.T) {
	usernameArg := "admin"
	hostArg := "12.12.32.12"
	args := []string{fmt.Sprintf("%s@%s", usernameArg, hostArg)}
	username, host := ParseUserAndHost(args, nil)
	require.Equal(t, usernameArg, username)
	require.Equal(t, hostArg, host)

	hostArg = "12.12.23.23:22"
	args = []string{fmt.Sprintf("%s@%s", usernameArg, hostArg)}
	username, host = ParseUserAndHost(args, nil)
	require.Equal(t, usernameArg, username)
	require.Equal(t, hostArg, host)

	currentUser, _ := user.Current()
	args = []string{fmt.Sprintf("%s", hostArg)}
	username, host = ParseUserAndHost(args, currentUser)
	require.Equal(t, currentUser.Username, username)
	require.Equal(t, hostArg, host)
}

func TestParseHostAndPort(t *testing.T) {
	addr := "12.12.32.12"
	host, port := ParseHostAndPort(addr)
	require.Equal(t, host, addr)
	require.Equal(t, port, 22)

	addrPort := 2344
	host, port = ParseHostAndPort(fmt.Sprintf("%s:%d", addr, addrPort))
	require.Equal(t, host, addr)
	require.Equal(t, port, addrPort)
}
