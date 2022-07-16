package strutil

import "strings"

type JoinReq struct {
	Elems     []string
	Sep       string
	Prefix    string
	Suffix    string
	OmitEmpty bool
	sb        *strings.Builder
}

func (r *JoinReq) initStringBuilder() {
	r.sb = &strings.Builder{}
	n := len(r.Sep) * (len(r.Elems) - 1)
	for i := 0; i < len(r.Elems); i++ {
		n += len(r.Elems[i])
	}
	r.sb.Grow(n)
}

func (r *JoinReq) writeString(s string, writeSep bool) {
	if r.OmitEmpty && s == "" {
		return
	}
	if writeSep {
		r.sb.WriteString(r.Sep)
	}
	r.sb.WriteString(s)
}

func NewJoinReq(elems []string) *JoinReq {
	return &JoinReq{Elems: elems}
}

func (r *JoinReq) SetSep(sep string) *JoinReq {
	r.Sep = sep
	return r
}

func (r *JoinReq) SetPrefix(prefix string) *JoinReq {
	r.Prefix = prefix
	return r
}

func (r *JoinReq) SetSuffix(suffix string) *JoinReq {
	r.Suffix = suffix
	return r
}

func (r *JoinReq) SetOmitEmpty(omitEmpty bool) *JoinReq {
	r.OmitEmpty = omitEmpty
	return r
}

func (r *JoinReq) Join() string {
	switch len(r.Elems) {
	case 0:
		return ""
	case 1:
		return r.Elems[0]
	}
	r.initStringBuilder()
	r.writeString(r.Prefix, false)
	r.writeString(r.Elems[0], false)
	for _, s := range r.Elems[1:] {
		r.writeString(s, true)
	}
	r.writeString(r.Suffix, false)
	return r.sb.String()
}
