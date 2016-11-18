package sdp

import (
	"testing"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/bufio.v1"
)

func TestParse(t *testing.T) {
	s := `
v=0
o=- 936021522 936021522 IN IP4 184.72.239.149
s=BigBuckBunny_115k.mov
c=IN IP4 184.72.239.149
t=0 0
a=sdplang:en
a=range:npt=0- 596.48
a=control:*
m=audio 0 RTP/AVP 96
a=rtpmap:96 mpeg4-generic/12000/2
a=fmtp:96 profile-level-id=1;mode=AAC-hbr;sizelength=13;indexlength=3;indexdeltalength=3;config=1490
a=control:trackID=1
m=video 0 RTP/AVP 97
a=rtpmap:97 H264/90000
a=fmtp:97 packetization-mode=1;profile-level-id=42C01E;sprop-parameter-sets=Z0LAHtkDxWhAAAADAEAAAAwDxYuS,aMuMsg==
a=cliprect:0,0,160,240
a=framesize:97 240-160
a=framerate:24.0
a=control:trackID=2
`
	v, err := Parse(bufio.NewBufferString(s))
	if err != nil {
		panic(err)
	}
	spew.Dump(v)

}
