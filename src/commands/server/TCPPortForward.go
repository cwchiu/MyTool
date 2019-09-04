package server

import (
	"github.com/spf13/cobra"
	"log"
    "net"
    "io"
)

// https://github.com/kintoandar/fwd/blob/master/fwd.go

func getLocalAddrs() ([]net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	var list []net.IP
	for _, addr := range addrs {
		v := addr.(*net.IPNet)
		if v.IP.To4() != nil {
			list = append(list, v.IP)
		}
	}
	return list, nil
}

func fwd(src net.Conn, remote string, proto string) {
	dst, err := net.Dial(proto, remote)
	errHandler(err)
	go func() {
		_, err = io.Copy(src, dst)
		errPrinter(err)
	}()
	go func() {
		_, err = io.Copy(dst, src)
		errPrinter(err)
	}()
}

func errHandler(err error) {
	if err != nil {
        panic(err)
	}
}

// TODO: merge error handling functions
func errPrinter(err error) {
	if err != nil {
		log.Fatalf("[Error] %s\n", err.Error())
	}
}

func tcpStart(from string, to string) {
	proto := "tcp"

	localAddress, err := net.ResolveTCPAddr(proto, from)
	errHandler(err)

	remoteAddress, err := net.ResolveTCPAddr(proto, to)
	errHandler(err)

	listener, err := net.ListenTCP(proto, localAddress)
	errHandler(err)

	defer listener.Close()

	log.Printf("Forwarding %s traffic from '%v' to '%v'\n", proto, localAddress, remoteAddress)
	log.Println("<CTRL+C> to exit")

	for {
		src, err := listener.Accept()
		errHandler(err)
		log.Printf("New connection established from '%v'\n", src.RemoteAddr())
		go fwd(src, to, proto)
	}
}

func udpStart(from string, to string) {
	proto := "udp"

	localAddress, err := net.ResolveUDPAddr(proto, from)
	errHandler(err)

	remoteAddress, err := net.ResolveUDPAddr(proto, to)
	errHandler(err)

	listener, err := net.ListenUDP(proto, localAddress)
	errHandler(err)
	defer listener.Close()

	dst, err := net.DialUDP(proto, nil, remoteAddress)
	errHandler(err)
	defer dst.Close()

	log.Printf("Forwarding %s traffic from '%v' to '%v'\n", proto, localAddress, remoteAddress)
	log.Println("<CTRL+C> to exit")

	buf := make([]byte, 512)
	for {
		rnum, err := listener.Read(buf[0:])
		errHandler(err)

		_, err = dst.Write(buf[:rnum])
		errHandler(err)

		log.Printf("%d bytes forwared\n", rnum)
	}
}

func SetupTcpPortForwardCommand(rootCmd *cobra.Command) {

	var from string
	var to string
	var udp bool

	cmd := &cobra.Command{
		Use:   "fwd",
		Short: "TCP Port Forward",
		Run: func(cmd *cobra.Command, args []string) {

            if udp {
				udpStart(from, to)
			} else {
				tcpStart(from, to)
			}

		},
	}
    flags := cmd.Flags()
	flags.StringVarP(&from, "from", "f", "127.0.0.1:8000", "source HOST:PORT")
	flags.StringVarP(&to, "to", "t", "127.0.0.1:80", "destination HOST:PORT")
	flags.BoolVarP(&udp, "udp", "u", false, "UDP模式")

	rootCmd.AddCommand(cmd)
}
