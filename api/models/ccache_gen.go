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
	// map header, size 2
	// write "Patterns"
	err = en.Append(0x82, 0xa8, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73)
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
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CCacheDelete) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Patterns"
	o = append(o, 0x82, 0xa8, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Patterns)))
	for zxvk := range z.Patterns {
		o = msgp.AppendString(o, z.Patterns[zxvk])
	}
	// string "OrgId"
	o = append(o, 0xa5, 0x4f, 0x72, 0x67, 0x49, 0x64)
	o = msgp.AppendInt(o, z.OrgId)
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
	s += 6 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *CCacheDeleteResp) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zhct uint32
	zhct, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zhct > 0 {
		zhct--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Nodes":
			var zcua uint32
			zcua, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Nodes) >= int(zcua) {
				z.Nodes = (z.Nodes)[:zcua]
			} else {
				z.Nodes = make([]string, zcua)
			}
			for zwht := range z.Nodes {
				z.Nodes[zwht], err = dc.ReadString()
				if err != nil {
					return
				}
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
func (z *CCacheDeleteResp) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Nodes"
	err = en.Append(0x81, 0xa5, 0x4e, 0x6f, 0x64, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Nodes)))
	if err != nil {
		return
	}
	for zwht := range z.Nodes {
		err = en.WriteString(z.Nodes[zwht])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CCacheDeleteResp) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Nodes"
	o = append(o, 0x81, 0xa5, 0x4e, 0x6f, 0x64, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Nodes)))
	for zwht := range z.Nodes {
		o = msgp.AppendString(o, z.Nodes[zwht])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CCacheDeleteResp) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zxhx uint32
	zxhx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Nodes":
			var zlqf uint32
			zlqf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Nodes) >= int(zlqf) {
				z.Nodes = (z.Nodes)[:zlqf]
			} else {
				z.Nodes = make([]string, zlqf)
			}
			for zwht := range z.Nodes {
				z.Nodes[zwht], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
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
func (z *CCacheDeleteResp) Msgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zwht := range z.Nodes {
		s += msgp.StringPrefixSize + len(z.Nodes[zwht])
	}
	return
}
