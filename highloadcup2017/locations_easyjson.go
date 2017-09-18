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

func easyjsonE9bf8de2DecodeHighloadcup(in *jlexer.Lexer, out *LocationAvg) {
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
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "avg":
			out.Avg = float64(in.Float64())
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
func easyjsonE9bf8de2EncodeHighloadcup(out *jwriter.Writer, in LocationAvg) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"avg\":")
	out.Float64(float64(in.Avg))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LocationAvg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9bf8de2EncodeHighloadcup(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LocationAvg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9bf8de2EncodeHighloadcup(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LocationAvg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9bf8de2DecodeHighloadcup(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LocationAvg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9bf8de2DecodeHighloadcup(l, v)
}
func easyjsonE9bf8de2DecodeHighloadcup1(in *jlexer.Lexer, out *Location) {
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
		case "place":
			out.Place = string(in.String())
		case "country":
			out.Country = string(in.String())
		case "city":
			out.City = string(in.String())
		case "distance":
			out.Distance = int(in.Int())
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
func easyjsonE9bf8de2EncodeHighloadcup1(out *jwriter.Writer, in Location) {
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
	out.RawString("\"place\":")
	out.String(string(in.Place))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"country\":")
	out.String(string(in.Country))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"city\":")
	out.String(string(in.City))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"distance\":")
	out.Int(int(in.Distance))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Location) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9bf8de2EncodeHighloadcup1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Location) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9bf8de2EncodeHighloadcup1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Location) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9bf8de2DecodeHighloadcup1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Location) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9bf8de2DecodeHighloadcup1(l, v)
}