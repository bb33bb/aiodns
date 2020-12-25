package main

import "C"
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/miekg/dns"
)

const (
	TYPE_REST int = iota
	TYPE_ADDR
	TYPE_LIST
	TYPE_CDNS
	TYPE_ODNS
	TYPE_METH
)

var (
	client = &dns.Client{Net: "tcp"}

	ListenAddr = ":53"
	ChinaDNS   = "223.5.5.5:53"
	OtherDNS   = "1.1.1.1:53"

	ChinaList = make([]string, 0)

	ServeMux  *dns.ServeMux
	TCPSocket *dns.Server
	UDPSocket *dns.Server
)

//export aiodns_dial
func aiodns_dial(name int, value *C.char) bool {
	switch name {
	case TYPE_REST:
		{
			ChinaDNS = "223.5.5.5:53"
			OtherDNS = "1.1.1.1:53"
			ChinaList = make([]string, 0)

			if TCPSocket != nil {
				_ = TCPSocket.Shutdown()

				TCPSocket = nil
			}

			if UDPSocket != nil {
				_ = UDPSocket.Shutdown()

				UDPSocket = nil
			}

			ServeMux = nil

			fmt.Println("[aiodns][aiodns_dial] TYPE_REST")
		}
	case TYPE_ADDR:
		ListenAddr = C.GoString(value)

		fmt.Printf("[aiodns][aiodns_dial] TYPE_ADDR => %s\n", C.GoString(value))
	case TYPE_LIST:
		{
			fd, err := os.Open(C.GoString(value))
			if err != nil {
				return false
			}
			defer fd.Close()

			br := bufio.NewReader(fd)
			for {
				data, _, err := br.ReadLine()
				if err != nil {
					if err != io.EOF {
						fmt.Printf("[aiodns][aiodns_dial][br.ReadLine] %v\n", err)
						return false
					}

					break
				}

				text := strings.TrimSpace(string(data))
				if text != "" {
					ChinaList = append(ChinaList, text)
				}
			}

			fmt.Printf("[aiodns][aiodns_dial] TYPE_LIST => %s\n", C.GoString(value))
		}
	case TYPE_CDNS:
		ChinaDNS = C.GoString(value)

		fmt.Printf("[aiodns][aiodns_dial] TYPE_CDNS => %s\n", C.GoString(value))
	case TYPE_ODNS:
		OtherDNS = C.GoString(value)

		fmt.Printf("[aiodns][aiodns_dial] TYPE_ODNS => %s\n", C.GoString(value))
	case TYPE_METH:
		client = &dns.Client{Net: strings.ToLower(C.GoString(value))}

		fmt.Printf("[aiodns][aiodns_dial] TYPE_METH => %s\n", C.GoString(value))
	default:
		return false
	}

	return true
}

//export aiodns_init
func aiodns_init() bool {
	ServeMux = dns.NewServeMux()
	ServeMux.HandleFunc("in-addr.arpa.", handleServerName)

	for i := 0; i < len(ChinaList); i++ {
		ServeMux.HandleFunc(ChinaList[i], handleChina)
	}

	ServeMux.HandleFunc(".", handleOther)

	TCPSocket = &dns.Server{Net: "tcp", Addr: ListenAddr, Handler: ServeMux}
	UDPSocket = &dns.Server{Net: "udp", Addr: ListenAddr, Handler: ServeMux}

	go func() { _ = TCPSocket.ListenAndServe() }()
	go func() { _ = UDPSocket.ListenAndServe() }()

	return true
}

//export aiodns_free
func aiodns_free() {
	ChinaList = nil

	if TCPSocket != nil {
		_ = TCPSocket.Shutdown()

		TCPSocket = nil
	}

	if UDPSocket != nil {
		_ = UDPSocket.Shutdown()

		UDPSocket = nil
	}

	ServeMux = nil
}

func handleServerName(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)

	for i := 0; i < len(r.Question); i++ {
		mx, err := dns.NewRR(fmt.Sprintf("%s PTR aioCloud", r.Question[i].Name))
		if err != nil {
			log.Println(err)
			return
		}

		m.Answer = append(m.Answer, mx)
	}

	_ = w.WriteMsg(m)
}

func handleChina(w dns.ResponseWriter, r *dns.Msg) {
	m, _, err := client.Exchange(r, ChinaDNS)
	if err != nil {
		return
	}

	_ = w.WriteMsg(m)
}

func handleOther(w dns.ResponseWriter, r *dns.Msg) {
	m, _, err := client.Exchange(r, OtherDNS)
	if err != nil {
		return
	}

	_ = w.WriteMsg(m)
}

func main() {
}
