package sdp

import "strconv"

// a=rtpmap:<payload type> <encoding name>/<clock rate> [/<encoding parameters>
func (self Attr)GetRTPMap() {

}

// a=fmtp:<format> <format specific parameters>
//
// This attribute allows parameters that are specific to a
// particular format to be conveyed in a way that SDP does not
// have to understand them.  The format must be one of the formats
// specified for the media.  Format-specific parameters may be any
// set of parameters required to be conveyed by SDP and given
// unchanged to the media tool that will use this format.  At most
// one instance of this attribute is allowed for each format.
func (self Attr)GetFormatParameters() {

}

//a=quality:<quality>
//
//This gives a suggestion for the quality of the encoding as an
//integer value.  The intention of the quality attribute for
//video is to specify a non-default trade-off between frame-rate
//and still-image quality.  For video, the value is in the range
//0 to 10, with the following suggested meaning:
//
//10 - the best still-image quality the compression scheme can give.
//5  - the default behaviour given no quality suggestion.
//0  - the worst still-image quality the codec designer thinks is still usable.
func (self Attr)GetQuality() int {
	i, _ := strconv.Atoi(self["quality"])
	return i
}

//a=framerate:<frame rate>
//
//This gives the maximum video frame rate in frames/sec.  It is
//intended as a recommendation for the encoding of video data.
//Decimal representations of fractional values using the notation
//"<integer>.<fraction>" are allowed.  It is a media-level
//attribute, defined only for video media, and it is not
//dependent on charset.
func (self Attr)GetFrameRate() {

}

//a=recvonly
//
//This specifies that the tools should be started in receive-only
//mode where applicable.  It can be either a session- or media-
//level attribute, and it is not dependent on charset.  Note that
//recvonly applies to the media only, not to any associated
//control protocol (e.g., an RTP-based system in recvonly mode
//SHOULD still send RTCP packets).
func (self Attr)IsRecvOnly() {

}

//a=sendrecv
//
//This specifies that the tools should be started in send and
//receive mode.  This is necessary for interactive conferences
//with tools that default to receive-only mode.  It can be either
//a session or media-level attribute, and it is not dependent on
//charset.
//If none of the attributes "sendonly", "recvonly", "inactive",
//and "sendrecv" is present, "sendrecv" SHOULD be assumed as the
//default for sessions that are not of the conference type
//"broadcast" or "H332" (see below).
func (self Attr)IsSendRecv() {

}

//a=sendonly
//
//This specifies that the tools should be started in send-only
//mode.  An example may be where a different unicast address is
//to be used for a traffic destination than for a traffic source.
//In such a case, two media descriptions may be used, one
//sendonly and one recvonly.  It can be either a session- or
//media-level attribute, but would normally only be used as a
//media attribute.  It is not dependent on charset.  Note that
//sendonly applies only to the media, and any associated control
//protocol (e.g., RTCP) SHOULD still be received and processed as
//normal.
func (self Attr)IsSendOnly() {

}

//a=inactive
//
//This specifies that the tools should be started in inactive
//mode.  This is necessary for interactive conferences where
//users can put other users on hold.  No media is sent over an
//inactive media stream.  Note that an RTP-based system SHOULD
//still send RTCP, even if started inactive.  It can be either a
//session or media-level attribute, and it is not dependent on
//charset.
func (self Attr)IsInactive() {

}
