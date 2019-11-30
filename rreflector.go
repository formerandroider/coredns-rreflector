package rreflector

import (
	"context"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("rreflector")

type RReflector struct {
}

func (p RReflector) ServeDNS(_ context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Debug("Received response")
	state := request.Request{
		Req: r,
		W:   w,
	}

	log.Debugf("QName: %s", state.QName())

	a := &dns.Msg{}
	a.SetReply(r)
	a.Authoritative = true

	var rr dns.RR
	rr.(*dns.PTR).Hdr = dns.RR_Header{
		Name:   state.QName(),
		Rrtype: dns.TypePTR,
		Class:  state.QClass(),
	}
	rr.(*dns.PTR).Ptr = ""

	a.Extra = []dns.RR{rr}

	if err := w.WriteMsg(a); err != nil {
		return dns.RcodeServerFailure, err
	}

	return 0, nil
}

func (p RReflector) Name() string {
	return "rreflector"
}
