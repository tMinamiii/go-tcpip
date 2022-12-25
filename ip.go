package tcpip

// IPHeader
//
// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |Version|  IHL  |Type of Service|          Total Length         |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |         Identification        |Flags|      Fragment Offset    |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |  Time to Live |    Protocol   |         Header Checksum       |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                       Source Address                          |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                    Destination Address                        |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                    Options                    |    Padding    |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
type IPHeader struct {
	Version              byte   // 1byte
	InternetHeaderLength byte   // 1byte
	ServiceType          []byte // 2byte
	TotalLength          []byte // 2byte
	Identification       []byte // 2byte
	Flags                []byte // 3bit + 13bit offset
	TTL                  byte   // 1byte
	Protocol             byte   // 1byte
	HeaderCheckSum       []byte // 2byte
	SourceAddress        []byte // 4byte
	DestinationAddress   []byte // 4byte
	Options              []byte // 3byte
	Padding              byte   // byte
}

func NewIPHeader(sourceIP, dstIP []byte, protocol string) IPHeader {
	ip := IPHeader{
		Version:              byte(0x04), // 4
		InternetHeaderLength: byte(0x05), // 5
		ServiceType:          []byte{0x00},
		TotalLength:          []byte{0x00, 0x00},
	}
	return ip
}
