package main

import "C"
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aiocloud/aiodns/tools"
	"github.com/miekg/dns"
)

var (
	client = &dns.Client{Net: "tcp"}

	ChinaCON = ""
	ChinaDNS = "119.29.29.29:53"
	OtherDNS = "1.1.1.1:53"

	ChinaList []string

	ServeMux  *dns.ServeMux
	TCPSocket *dns.Server
	UDPSocket *dns.Server
)

//export aiodns_dial
func aiodns_dial(chinacon *C.char, chinadns *C.char, otherdns *C.char) bool {
	ChinaCON = tools.String(C.GoString(chinacon))
	ChinaDNS = tools.String(C.GoString(chinadns))
	OtherDNS = tools.String(C.GoString(otherdns))
	ChinaList = make([]string, 0)

	fd, err := os.Open(ChinaCON)
	if err != nil {
		fmt.Printf("[aiodns][aiodns_dial][os.Open] %v\n", err)
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

	fmt.Printf("[aiodns][aiodns_dial] ChinaCON => %s\n", ChinaCON)
	fmt.Printf("[aiodns][aiodns_dial] ChinaDNS => %s\n", ChinaDNS)
	fmt.Printf("[aiodns][aiodns_dial] OtherDNS => %s\n", OtherDNS)
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

	TCPSocket = &dns.Server{Net: "tcp", Addr: ":53", Handler: ServeMux}
	UDPSocket = &dns.Server{Net: "udp", Addr: ":53", Handler: ServeMux}

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
