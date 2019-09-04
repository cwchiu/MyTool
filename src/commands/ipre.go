package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	mdns "github.com/miekg/dns"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

// https://github.com/m3ng9i/IP-resolver

type DnsAddr struct {
	Name    string
	Address string
}

type DnsAddrs []DnsAddr

type Answer struct {
	DnsAddr
	IP    []string
	Error error
}

type Answers []Answer

type AnswerJson struct {
	DnsAddr
	IP    []string
	Error string
}

type AnswersJson []AnswerJson

type ReadConfigError struct {
	Path   string
	Errmsg string
	Exit   bool
}

func (e *ReadConfigError) Error() string {
	return fmt.Sprintf("Configuration file '%s' load error: %s", e.Path, e.Errmsg)
}

var appname string

/* from qlibgo/dns

Get a domain's IPs from a specific name server.

Parameters:
    domain      the domain you want to query
    nameserver  name server's IP address
    port        53 in general
    net         tcp or udp
    timeout     in seconds, can be omitted

Here's an example：
    r, e := ARecords("www.example.com", "8.8.8.8", 53, "tcp")
    if e != nil {
        fmt.Println(e)
    } else {
        fmt.Println(r)
    }
*/
func ARecords(domain, nameserver string, port uint16, net string, timeout ...uint8) ([]string, error) {
	var result []string

	if net != "tcp" && net != "udp" {
		return result, errors.New("The Parameter 'net' should only be 'tcp' or 'udp'.")
	}

	msg := new(mdns.Msg)
	msg.SetQuestion(mdns.Fqdn(domain), mdns.TypeA)

	var client *mdns.Client
	if len(timeout) > 0 {
		tm := time.Duration(timeout[0]) * time.Second
		client = &mdns.Client{Net: net, DialTimeout: tm, ReadTimeout: tm, WriteTimeout: tm}
	} else {
		client = &mdns.Client{Net: net}
	}

	r, _, err := client.Exchange(msg, fmt.Sprintf("%s:%d", nameserver, port))
	if err != nil {
		return result, err
	}

	for _, i := range r.Answer {
		if t, ok := i.(*mdns.A); ok {
			result = append(result, t.A.String())
		}
	}

	return result, nil
}

/*
Use goroutines to query one domain with multiple name servers.

Parameters:
    dns     name server configuration
    domain  the domain you want to query
    net     tcp or udp
*/
func query(dns DnsAddrs, domain string, net string) Answers {
	var wg sync.WaitGroup
	answers := make(Answers, len(dns))
	for j, i := range dns {
		wg.Add(1)
		go func(n int, d DnsAddr) {
			defer wg.Done()
			var answer Answer
			answer.DnsAddr = d
			ip, err := ARecords(domain, d.Address, 53, net, 3)
			if err != nil {
				answer.Error = err
			} else {
				if len(ip) == 0 {
					answer.Error = errors.New("No result")
				} else {
					answer.IP = ip
				}
			}
			answers[n] = answer
		}(j, i)
	}

	wg.Wait()
	return answers
}

// Get all the IPs from the query results.
func (a Answers) allIP() []string {

	var ips []string
	i := make(map[string]bool)

	for _, item := range a {
		for _, ip := range item.IP {
			i[ip] = true
		}
	}

	for key, _ := range i {
		ips = append(ips, key)
	}

	sort.Strings(ips)
	return ips

}

func in(ip string, ips []string) bool {
	for _, i := range ips {
		if i == ip {
			return true
		}
	}
	return false
}

// Output the query results.
func (a Answers) output() {

	allip := a.allIP()

	resultNum := len(allip)
	if resultNum == 0 {
		resultNum = 1 // leave room for displaying error
	}

	/*
	   First line is name servers's names, the second line is theirs IPs. Example:
	   DNS1         DNS2         DNS3         DNS4
	   1.1.1.1      2.2.2.2      3.3.3.3      4.4.4.4
	*/
	head := make([]string, len(a)*2)

	/*
	   A domain's IPs queried from different name servers. Example:
	   11.11.11.11  Timout       -            -
	   11.11.11.12  -            11.11.11.12  -
	   -            -            11.11.11.13  -
	   -            -            -            11.11.11.14
	*/
	ip := make([]string, len(a)*resultNum)

	// Fill ip with "-"
	for i, _ := range ip {
		ip[i] = "-"
	}

	// Fill head and ip
	for i, item := range a {
		head[i] = item.Name
		head[i+len(a)] = item.Address

		if item.Error == nil {
			for j := 0; j < len(allip); j++ {
				if in(allip[j], item.IP) {
					ip[j*len(a)+i] = allip[j]
				}
			}
		} else {
			ip[i] = errToString(item.Error)
		}
	}

	// Output variable "head" and "ip" to stdout.
	// The max length of IP string is 15, plus two space is 17, that's the reason I use "17" in Printf.

	for i, item := range head {
		fmt.Printf("%-17.16s", item)
		if (i+1)%len(a) == 0 {
			fmt.Println()
		}
	}

	fmt.Println(strings.Repeat("-", 17*len(a)))

	for i, item := range ip {
		fmt.Printf("%-17s", item)
		if (i+1)%len(a) == 0 {
			fmt.Println()
		}
	}
}

// Output all IPs resolved from all nameserver and ignore errors.
func (a Answers) outputNormal() {

	allip := a.allIP()

	for _, i := range allip {
		fmt.Println(i)
	}
}

// Output Json with full error message.
func (a Answers) outputJson() {

	aj := make(AnswersJson, len(a))
	for j, item := range a {
		aj[j].DnsAddr = item.DnsAddr
		aj[j].IP = item.IP
		if item.Error != nil {
			aj[j].Error = item.Error.Error()
		}
	}

	b, err := json.Marshal(aj)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred when generating json: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(b))
}

// Convert error message to a short form.
func errToString(err error) string {
	if err == nil {
		return ""
	}

	var f = func(str, substr string) bool {
		return strings.Contains(strings.ToLower(str), substr)
	}

	s := err.Error()

	if f(s, "timeout") {
		// Errors like "dial tcp 8.8.8.8:53: ConnectEx tcp: i/o timeout"
		//   or "WSARecv tcp 192.168.0.1:3586: i/o timeout"
		return "Timeout"

	} else if f(s, "refused the network connection") {
		// Errors like "dial tcp 8.8.8.8:53: ConnectEx tcp:
		//   The remote system refused the network connection."
		// When this error came up after you send a tcp query, this maybe means
		//   the nameserver doesn't support tcp protocol, so choose udp instead.
		return "Conn refused"

	} else if f(s, "no service is operating") {
		// Errors like "WSARecv udp 192.168.0.1:1573: No service is operating at
		//   the destination network endpoint on the remote system."
		return "NS invalid"

	} else if f(s, "forcibly closed by the remote host") {
		// Errors like "WSARecv udp 192.168.0.1:1590: An existing connection
		//   was forcibly closed by the remote host.
		return "NS invalid"

	} else if f(s, "no result") {
		return "No result"

	}

	// Deal with other error message
	return "Connect error"

}

func getDefaultConfig() (DnsAddrs, error) {
	conf := []DnsAddr{
		DnsAddr{Name: "AliDNS", Address: "223.5.5.5"},
		DnsAddr{Name: "114DNS", Address: "114.114.114.114"},
		DnsAddr{Name: "Google", Address: "8.8.8.8"},
		DnsAddr{Name: "IBM", Address: "9.9.9.9"},
		DnsAddr{Name: "OpenDNS", Address: "208.67.222.222"},
		DnsAddr{Name: "Hinet", Address: "168.95.192.1"},
	}
	return conf, nil
}

func init() {
	var format string
	cmd := &cobra.Command{Use: "ipre <domain>", Short: "IP反查",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <domain>")
			}

			net := "tcp"
			conf, _ := getDefaultConfig()
			result := query(conf, args[0], net)
			if format == "std" {
				result.output()
			} else if format == "json" {
				result.outputJson()
			} else {
				result.outputNormal()
			}
		},
	}
	cmd.Flags().StringVarP(&format, "format", "f", "ip", "std, json, ip")

	rootCmd.AddCommand(cmd)
}
