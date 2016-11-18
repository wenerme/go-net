package sdp

import (
	"fmt"
)

const MediaType = "application/sdp"

// The connection ("c=") and attribute ("a=") information in the
// session-level section applies to all the media of that session unless
// overridden by connection information or an attribute of the same name
// in the media description.
//
// '*' means optional in below
type T struct {
	// v=  (protocol version)
	//
	// This memo defines version 0.  There is no minor version number.
	Version               int
	// o=  (originator and session identifier)
	//
	// o=<username> <sess-id> <sess-version> <nettype> <addrtype> <unicast-address>
	// In general, the "o=" field serves as a globally unique identifier for
	// this version of this session description, and the subfields excepting
	// the version taken together identify the session irrespective of any
	// modifications.
	//
	// For privacy reasons, it is sometimes desirable to obfuscate the
	// username and IP address of the session originator.  If this is a
	// concern, an arbitrary <username> and private <unicast-address> MAY be
	// chosen to populate the "o=" field, provided that these are selected
	// in a manner that does not affect the global uniqueness of the field.
	Origin                Origin
	// s=  (session name)
	SessionName           string
	// i=* (session information)
	SessionInformation    string
	// u=* (URI of description)
	URI                   string
	// e=* (email address)
	Email                 string
	// p=* (phone number)
	Phone                 string
	// c=* (connection information -- not required if included in all media)
	//
	//  c=<nettype> <addrtype> <connection-address>
	ConnectionInformation string
	// b=* (zero or more bandwidth information lines)
	//
	// b=<bwtype>:<bandwidth>
	BandwidthInformation  []string
	// One or more time descriptions ("t=" and "r=" lines; see below)
	Times                 []TimeDescription
	// z=* (time zone adjustments)
	//
	// z=<adjustment time> <offset> <adjustment time> <offset> ....
	TimeZoneAdjustments   string
	// k=* (encryption key)
	//
	// k=<method>:[<encryption key>]
	EncryptionKey         string
	// a=* (zero or more session attribute lines)
	//
	// a=<attribute>
	// a=<attribute>:<value>
	Attribute             Attr
	// Zero or more media descriptions
	Media                 []Media
}

//Time description
type TimeDescription struct {
	// t=  (time the session is active)
	//
	// t=<start-time> <stop-time>
	SessionActiveTime uint
	// r=* (zero or more repeat times)
	//
	// r=<repeat interval> <active duration> <offsets from start-time>
	RepeatTimes       []uint
}

// Media description
//
// m=  (media name and transport address)
//
// m=<media> <port> <proto> <fmt> ...
// m=<media> <port>/<number of ports> <proto> <fmt> ...
type Media struct {
	// <media> is the media type.  Currently defined media are "audio",
	// "video", "text", "application", and "message"
	Type                  string
	Port                  int
	PortNumber            int
	// * udp: denotes an unspecified protocol running over UDP.
	// * RTP/AVP: denotes RTP used under the RTP Profile for Audio and Video Conferences with Minimal Control running over UDP.
	// * RTP/SAVP: denotes the Secure Real-time Transport Protocol running over UDP
	Proto                 string
	Format                string

	// i=* (media title)
	Title                 string
	// c=* (connection information -- optional if included at session level)
	ConnectionInformation string
	// b=* (zero or more bandwidth information lines)
	BandwidthInformation  []string
	// k=* (encryption key)
	EncryptionKey         string
	// a=* (zero or more media attribute lines)
	Attribute             Attr
}

type Attr map[string]string

type ConnectionInformation struct {
	NetType string
}

// o=<username> <sess-id> <sess-version> <nettype> <addrtype> <unicast-address>
type Origin struct {
	//<username> is the user's login on the originating host, or it is "-"
	//	if the originating host does not support the concept of user IDs.
	//	The <username> MUST NOT contain spaces.
	Username       string

	// <sess-id> is a numeric string such that the tuple of <username>,
	// 	<sess-id>, <nettype>, <addrtype>, and <unicast-address> forms a
	// 	globally unique identifier for the session.  The method of
	// 	<sess-id> allocation is up to the creating tool, but it has been
	// 	suggested that a Network Time Protocol (NTP) format timestamp be
	// 	used to ensure uniqueness.
	SessionID      int

	// <sess-version> is a version number for this session description.  Its
	//	usage is up to the creating tool, so long as <sess-version> is
	//	increased when a modification is made to the session data.  Again,
	//	it is RECOMMENDED that an NTP format timestamp is used.
	SessionVersion int
	// <nettype> is a text string giving the type of network.  Initially
	//	"IN" is defined to have the meaning "Internet", but other values
	//	MAY be registered in the future (see Section 8).
	NetType        string
	// <addrtype> is a text string giving the type of the address that
	// 	follows.  Initially "IP4" and "IP6" are defined, but other values
	// 	MAY be registered in the future (see Section 8).
	AddressType    string
	// <unicast-address> is the address of the machine from which the
	// session was created.  For an address type of IP4, this is either
	// the fully qualified domain name of the machine or the dotted-
	// decimal representation of the IP version 4 address of the machine.
	// For an address type of IP6, this is either the fully qualified
	// domain name of the machine or the compressed textual
	// representation of the IP version 6 address of the machine.  For
	// both IP4 and IP6, the fully qualified domain name is the form that
	// SHOULD be given unless this is unavailable, in which case the
	// globally unique address MAY be substituted.  A local IP address
	// MUST NOT be used in any context where the SDP description might
	// leave the scope in which the address is meaningful (for example, a
	// local address MUST NOT be included in an application-level
	// referral that might leave the scope).
	UnicastAddress string
}

func (self Origin)String() string {
	// <username> <sess-id> <sess-version> <nettype> <addrtype> <unicast-address>
	return fmt.Sprintf("%v %v %v %v %v %v", self.Username, self.SessionID, self.SessionVersion, self.NetType, self.AddressType, self.UnicastAddress);
}
