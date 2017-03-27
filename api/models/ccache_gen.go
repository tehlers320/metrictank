package models

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *CCacheDelete) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Patterns":
			var zbai uint32
			zbai, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Patterns) >= int(zbai) {
				z.Patterns = (z.Patterns)[:zbai]
			} else {
				z.Patterns = make([]string, zbai)
			}
			for zxvk := range z.Patterns {
				z.Patterns[zxvk], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "OrgId":
			z.OrgId, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Propagate":
			z.Propagate, err = dc.ReadBool()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *CCacheDelete) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Patterns"
	err = en.Append(0x83, 0xa8, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Patterns)))
	if err != nil {
		return
	}
	for zxvk := range z.Patterns {
		err = en.WriteString(z.Patterns[zxvk])
		if err != nil {
			return
		}
	}
	// write "OrgId"
	err = en.Append(0xa5, 0x4f, 0x72, 0x67, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.OrgId)
	if err != nil {
		return
	}
	// write "Propagate"
	err = en.Append(0xa9, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Propagate)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CCacheDelete) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Patterns"
	o = append(o, 0x83, 0xa8, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Patterns)))
	for zxvk := range z.Patterns {
		o = msgp.AppendString(o, z.Patterns[zxvk])
	}
	// string "OrgId"
	o = append(o, 0xa5, 0x4f, 0x72, 0x67, 0x49, 0x64)
	o = msgp.AppendInt(o, z.OrgId)
	// string "Propagate"
	o = append(o, 0xa9, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65)
	o = msgp.AppendBool(o, z.Propagate)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CCacheDelete) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Patterns":
			var zajw uint32
			zajw, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Patterns) >= int(zajw) {
				z.Patterns = (z.Patterns)[:zajw]
			} else {
				z.Patterns = make([]string, zajw)
			}
			for zxvk := range z.Patterns {
				z.Patterns[zxvk], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "OrgId":
			z.OrgId, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Propagate":
			z.Propagate, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CCacheDelete) Msgsize() (s int) {
	s = 1 + 9 + msgp.ArrayHeaderSize
	for zxvk := range z.Patterns {
		s += msgp.StringPrefixSize + len(z.Patterns[zxvk])
	}
	s += 6 + msgp.IntSize + 10 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *CCacheDeleteResp) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zwht uint32
	zwht, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zwht > 0 {
		zwht--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "PeerErrors":
			z.PeerErrors, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "DeletedSeries":
			z.DeletedSeries, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "DeletedArchives":
			z.DeletedArchives, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z CCacheDeleteResp) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "PeerErrors"
	err = en.Append(0x83, 0xaa, 0x50, 0x65, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.PeerErrors)
	if err != nil {
		return
	}
	// write "DeletedSeries"
	err = en.Append(0xad, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.DeletedSeries)
	if err != nil {
		return
	}
	// write "DeletedArchives"
	err = en.Append(0xaf, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.DeletedArchives)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z CCacheDeleteResp) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "PeerErrors"
	o = append(o, 0x83, 0xaa, 0x50, 0x65, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73)
	o = msgp.AppendInt(o, z.PeerErrors)
	// string "DeletedSeries"
	o = append(o, 0xad, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73)
	o = msgp.AppendInt(o, z.DeletedSeries)
	// string "DeletedArchives"
	o = append(o, 0xaf, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x73)
	o = msgp.AppendInt(o, z.DeletedArchives)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CCacheDeleteResp) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zhct uint32
	zhct, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zhct > 0 {
		zhct--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "PeerErrors":
			z.PeerErrors, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "DeletedSeries":
			z.DeletedSeries, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "DeletedArchives":
			z.DeletedArchives, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z CCacheDeleteResp) Msgsize() (s int) {
	s = 1 + 11 + msgp.IntSize + 14 + msgp.IntSize + 16 + msgp.IntSize
	return
}
