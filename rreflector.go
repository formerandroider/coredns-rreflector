package rreflector

import (
	"context"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
	"strings"
)

var log = clog.NewWithPlugin("rreflector")

type rreflectorHandler struct {
	host string
}

func newRReflectorHandler() *rreflectorHandler {
	return &rreflectorHandler{host: ""}
}

func (h *rreflectorHandler) ServeDNS(_ context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{
		Req: r,
		W:   w,
	}

	name := state.Name()
	log.Infof("parsing qname %s", name)
	parts := strings.Split(name[:len(name)-2], ".")
	arpaHost := parts[len(parts)-2]
	log.Debugf("arpa host: %s", name)
	addrParts := parts[:len(parts)-2]
	log.Debugf("address parts: %v", addrParts)

	a := &dns.Msg{}
	a.SetReply(r)
	a.Authoritative = true

	var joiner string
	switch arpaHost {
	case "ip6":
		joiner = ":"
		log.Debug("detected ipv6 lookup")
	case "in-addr":
		joiner = "."
		log.Debug("detected ipv4 lookup")
	default:
		log.Error("rreflector plugin called for non-rDNS lookup")
		a.Rcode = dns.RcodeNameError
		w.WriteMsg(a)
		return dns.RcodeNameError, nil
	}

	var reversedAddrParts []string
	for i := len(addrParts) - 1; i >= 0; i-- {
		reversedAddrParts = append(reversedAddrParts, addrParts[i])
	}

	log.Infof("original address: %s", strings.Join(reversedAddrParts, joiner))

	ptr := new(dns.PTR)
	ptr.Hdr = dns.RR_Header{
		Name:   state.QName(),
		Rrtype: dns.TypePTR,
		Class:  state.QClass(),
	}
	ptr.Ptr = strings.Join(reversedAddrParts, joiner) + h.host + "."

	a.Answer = []dns.RR{ptr}
	a.Rcode = dns.RcodeSuccess

	w.WriteMsg(a)

	return dns.RcodeSuccess, nil
}

func (h *rreflectorHandler) Name() string {
	return "rreflector"
}
