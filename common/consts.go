package common

const (
	BAD_PACKET_ERROR        = "Bad packet"
	VERSION_PREFIX          = "SSH-"
	LOCAL_VERSION           = VERSION_PREFIX + "2.0-gssh_0.1\r\n"
	PEER_VERSION_ERROR      = "Bad peer version"
	MAX_FULL_PACKET_SIZE    = 35000
	PACKET_LEN_BYTES_COUNT  = 4
	PADDING_LEN_BYTES_COUNT = 1
)
