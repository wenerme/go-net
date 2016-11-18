package sdp

import (
	"bufio"
	"strings"
	"errors"
	"strconv"
	"github.com/Sirupsen/logrus"
	"io"
)

// An SDP parser MUST completely ignore any session description that contains a type letter that it does not understand.
// An SDP parser MUST ignore any attribute it doesn't understand.
// The sequence CRLF (0x0d0a) is used to end a record, although parsers SHOULD be tolerant and also accept records terminated with a single newline character.
// If the "a=charset" attribute is not present, these octet strings MUST be interpreted as containing ISO-10646 characters in UTF-8 encoding (the presence of the "a=charset" attribute may force some fields to be interpreted differently).
func Parse(r io.Reader) (*T, error) {
	s := bufio.NewScanner(r)
	t := &T{Attribute:Attr(make(map[string]string))}
	var err error
	var media *Media
	for s.Scan() {
		parts := strings.SplitN(s.Text(), "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			if len(key) != 1 {
				return t, errors.New("SDP only allows 1-character variables")
			}
			val := parts[1]
			switch key {
			// version
			case "v":
				ver, err := strconv.Atoi(val)
				if err != nil {
					return t, err
				}
				t.Version = ver
			case "z":
				t.TimeZoneAdjustments = val
			// owner/creator and session identifier
			case "o":
				// o=<username> <session id> <version> <network type> <address type> <address>
				//t.Origin = val
				t.Origin, err = ParseOrigin(val)
				if err != nil {
					return nil, err
				}
			// session name
			case "s":
				t.SessionName = val
			// session information
			case "i":
				if media != nil {
					media.Title = val
				} else {
					t.SessionInformation = val
				}
			// URI of description
			case "u":
				t.URI = val
			// email address
			case "e":
				t.Email = val
			// phone number
			case "p":
				t.Phone = val
			// connection information - not required if included in all media
			case "c":
				// TODO: parse this
				if media != nil {
					media.ConnectionInformation = val
				} else {
					t.ConnectionInformation = val
				}
			case "b":
				// TODO: parse this
				if media != nil {
					media.BandwidthInformation = append(media.BandwidthInformation, val)
				} else {
					t.BandwidthInformation = append(t.BandwidthInformation, val)
				}
			case "k":
				if media != nil {
					media.EncryptionKey = val
				} else {
					t.EncryptionKey = val
				}
			case "m":
				m := Media{Attribute:Attr(make(map[string]string))}
				media = &m
				err = ParseMedia(val, media)
				if err != nil {
					return nil, err
				}
				t.Media = append(t.Media, m)
			case "a":
				// a=<attribute>
				// a=<attribute>:<value>
				splits := strings.SplitN(val, ":", 2)
				attrName := splits[0]
				attrVal := ""
				if len(splits) > 1 {
					attrVal = splits[1]
				}
				if media != nil {
					media.Attribute[attrName] = attrVal
				} else {
					t.Attribute[attrName] = attrVal
				}

			// TODO
			//case "t":
			//case "r":

			default:
				logrus.WithField("attr", key).WithField("val", val).Debug("Ignored")
			}
		} else if len(s.Text()) == 0 {
			logrus.WithField("line", s.Text()).Info("Invalid")
		}
	}
	return t, nil
}

func ParseMedia(s string, m *Media) (err error) {
	splits := strings.SplitN(s, " ", 4)
	if len(splits) != 4 {
		err = errors.New("Invalid media string")
		return
	}
	m.Type = splits[0]
	parts := strings.SplitN(splits[1], "/", 2)
	m.Port, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}
	if len(parts) > 1 {
		m.PortNumber, err = strconv.Atoi(parts[1])
		if err != nil {
			return
		}
	}
	m.Proto = splits[2]
	m.Format = splits[3]
	return
}

func ParseOrigin(s string) (o Origin, err error) {
	// o=<username> <sess-id> <sess-version> <nettype> <addrtype> <unicast-address>
	splits := strings.Split(s, " ")
	if len(splits) != 6 {
		err = errors.New("Invalid Origin string")
		return
	}
	o.Username = splits[0]
	o.SessionID, err = strconv.Atoi(splits[1])
	if err != nil {
		return
	}
	o.SessionVersion, err = strconv.Atoi(splits[2])
	if err != nil {
		return
	}
	o.NetType = splits[3]
	o.AddressType = splits[4]
	o.UnicastAddress = splits[5]
	return
}
