package main

import "golang.org/x/crypto/cryptobyte"

type ClientHello struct{
	SNI string
}

func ParseClientHello(record [] byte)(c *ClientHello,ok bool){
	c=&ClientHello{}
	/*struct{
		ContentType type;
		ProtocolVersion legacy_record_version;
		uint16 length;
		opaque fragment[TLSPlaintext.length];
	} TLSPlaintext*/
	in:=cryptobyte.String(record)
	if !in.Skip(1) || ! in.Skip(2){
		return nil,false
	}
	var msg cryptobyte.String
	if ! in.ReadUint16LengthPrefixed(&msg) || !in.Empty(){
		return nil,false
	}



/*struct {
	NameType name_type;
	opaque HostName <1..2^16-1;
}ServerName; */



var nameType uint8
if !snl.ReadUint8(&nameType){
	return nil,false
}
var hostName cryptobyte.String
if !snl.ReadUint16LengthPrefixed(&hostName){
	return nil,false
}

if nameType !=0 /*host_name */{
	return nil,false
}
c.SNI= string(hostName)
}
}
return c,true
}