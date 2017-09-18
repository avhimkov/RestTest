// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package main

import (
	json "encoding/json"
	"errors"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonEada991cDecodeHighloadcup(in *jlexer.Lexer, out *Visit) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.AddError(errors.New("Field is null"))
			return
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "location":
			out.Location = int(in.Int())
		case "user":
			out.User = int(in.Int())
		case "visited_at":
			out.VisitedAt = int(in.Int())
		case "mark":
			out.Mark = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonEada991cEncodeHighloadcup(out *jwriter.Writer, in Visit) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"id\":")
	out.Int(int(in.Id))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"location\":")
	out.Int(int(in.Location))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"user\":")
	out.Int(int(in.User))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"visited_at\":")
	out.Int(int(in.VisitedAt))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mark\":")
	out.Int(int(in.Mark))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Visit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEada991cEncodeHighloadcup(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Visit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEada991cEncodeHighloadcup(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Visit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEada991cDecodeHighloadcup(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Visit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEada991cDecodeHighloadcup(l, v)
}
