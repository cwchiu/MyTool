package clamav

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	// "io/ioutil"
	"io"
	"net"
	"os"
	"strings"
)

type ClamAV struct {
	Entry string
}

func (self ClamAV) Request(command string) (string, error) {
	conn, err := net.Dial("tcp", self.Entry)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	conn.Write([]byte(command))
	// buf, err := ioutil.ReadAll(conn)
	// buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 16) // using small tmo buffer for demonstrating
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				// fmt.Println("read error:", err)
                panic(err)
			}
			break
		}
		// fmt.Println("got", n, "bytes.")
		// buf = append(buf, tmp[:n]...)
		fmt.Print(string(tmp[:n]))
	}
	// fmt.Println("total size:", len(buf))
	if err != nil {
		return "", err
	}
	// return string(buf), nil
	return "", nil
}

func (self ClamAV) Version() (string, error) {
	return self.Request("VERSION")
}

func (self ClamAV) Reload() (string, error) {
	return self.Request("RELOAD")
}

func (self ClamAV) Shutdown() (string, error) {
	return self.Request("SHUTDOWN")
}

func (self ClamAV) Ping() (string, error) {
	return self.Request("PING")
}

func (self ClamAV) Scan(fn string) (string, error) {
	ret, err := self.Request(fmt.Sprintf("SCAN %s", fn))
	if err != nil {
		return "", err
	}

	if strings.Contains(ret, "failed: No such file or directory") {
		return "", errors.New("No such file or directory")
	}
	return strings.TrimSpace(strings.SplitN(ret, ":", 2)[1]), nil
}

func CreateClamAV(entry string) (*ClamAV, error) {
	return &ClamAV{
		Entry: entry,
	}, nil
}

type ClamAVEnviron struct {
	Server string
}

func GetClamAVEnviron() (*ClamAVEnviron, error) {
	server := os.Getenv("CLAMAV_SERVER")
	if server == "" {
		panic("required CLAMAV_SERVER enviroment variable")
	}
	return &ClamAVEnviron{
		Server: server,
	}, nil
}

func isScanCommand(command string) bool {
	switch command {
	case
		"SCAN",
		"MULTISCAN",
		"ALLMATCHSCAN",
		"CONTSCAN":
		return true
	}
	return false
}

func isSupportCommand(command string) bool {
	switch command {
	case
		"PING",
		"VERSION",
		"RELOAD",
		"SHUTDOWN":
		return true
	}
	return false
}

type ScanCommandHandler func(cmd *cobra.Command, args []string)

func createScanCommandHandler(cmd string) ScanCommandHandler {
	return func(_ *cobra.Command, args []string) {
		if len(args) < 1 {
			panic("required <file>")
		}

		env, _ := GetClamAVEnviron()
		inst, err := CreateClamAV(env.Server)
		if err != nil {
			panic(err)
		}
		scan_result, err := inst.Request(fmt.Sprintf("%s %s", cmd, args[0]))
		if err != nil {
			panic(err)
		}
		fmt.Println(scan_result)
	}
}

func createNoArgsCommandHandler(cmd string) ScanCommandHandler {
	return func(_ *cobra.Command, args []string) {
		env, _ := GetClamAVEnviron()
		inst, err := CreateClamAV(env.Server)
		if err != nil {
			panic(err)
		}
		scan_result, err := inst.Request(cmd)
		if err != nil {
			panic(err)
		}
		fmt.Println(scan_result)
	}
}
