// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package protocol

import (
	json "encoding/json"
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

func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol(in *jlexer.Lexer, out *ValidationResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "accepted":
			out.Accepted = bool(in.Bool())
		case "message":
			if in.IsNull() {
				in.Skip()
				out.Message = nil
			} else {
				if out.Message == nil {
					out.Message = new(string)
				}
				*out.Message = string(in.String())
			}
		case "code":
			if in.IsNull() {
				in.Skip()
				out.Code = nil
			} else {
				if out.Code == nil {
					out.Code = new(uint16)
				}
				*out.Code = uint16(in.Uint16())
			}
		case "mutated_object":
			if m, ok := out.MutatedObject.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.MutatedObject.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.MutatedObject = in.Interface()
			}
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
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol(out *jwriter.Writer, in ValidationResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"accepted\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Accepted))
	}
	if in.Message != nil {
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(*in.Message))
	}
	if in.Code != nil {
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.Uint16(uint16(*in.Code))
	}
	if in.MutatedObject != nil {
		const prefix string = ",\"mutated_object\":"
		out.RawString(prefix)
		if m, ok := in.MutatedObject.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.MutatedObject.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.MutatedObject))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ValidationResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ValidationResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ValidationResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ValidationResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol1(in *jlexer.Lexer, out *ValidationRequest) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "request":
			(out.Request).UnmarshalEasyJSON(in)
		case "settings":
			(out.Settings).UnmarshalEasyJSON(in)
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
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol1(out *jwriter.Writer, in ValidationRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"request\":"
		out.RawString(prefix[1:])
		(in.Request).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"settings\":"
		out.RawString(prefix)
		(in.Settings).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ValidationRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ValidationRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ValidationRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ValidationRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol1(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol2(in *jlexer.Lexer, out *UserInfo) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "username":
			out.Username = string(in.String())
		case "groups":
			if in.IsNull() {
				in.Skip()
				out.Groups = nil
			} else {
				in.Delim('[')
				if out.Groups == nil {
					if !in.IsDelim(']') {
						out.Groups = make([]string, 0, 4)
					} else {
						out.Groups = []string{}
					}
				} else {
					out.Groups = (out.Groups)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Groups = append(out.Groups, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "extra":
			(out.Extra).UnmarshalEasyJSON(in)
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
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol2(out *jwriter.Writer, in UserInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix[1:])
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"groups\":"
		out.RawString(prefix)
		if in.Groups == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Groups {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if (in.Extra).IsDefined() {
		const prefix string = ",\"extra\":"
		out.RawString(prefix)
		(in.Extra).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol2(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol3(in *jlexer.Lexer, out *SettingsValidationResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "valid":
			out.Valid = bool(in.Bool())
		case "message":
			if in.IsNull() {
				in.Skip()
				out.Message = nil
			} else {
				if out.Message == nil {
					out.Message = new(string)
				}
				*out.Message = string(in.String())
			}
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
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol3(out *jwriter.Writer, in SettingsValidationResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"valid\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Valid))
	}
	if in.Message != nil {
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(*in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SettingsValidationResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SettingsValidationResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SettingsValidationResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SettingsValidationResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol3(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol4(in *jlexer.Lexer, out *KubernetesAdmissionRequest) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "uid":
			out.Uid = string(in.String())
		case "kind":
			(out.Kind).UnmarshalEasyJSON(in)
		case "resource":
			(out.Resource).UnmarshalEasyJSON(in)
		case "subResource":
			out.SubResource = string(in.String())
		case "requestKind":
			(out.RequestKind).UnmarshalEasyJSON(in)
		case "requestResource":
			(out.RequestResource).UnmarshalEasyJSON(in)
		case "requestSubResource":
			out.RequestSubResource = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "namespace":
			out.Namespace = string(in.String())
		case "operation":
			out.Operation = string(in.String())
		case "userInfo":
			(out.UserInfo).UnmarshalEasyJSON(in)
		case "object":
			(out.Object).UnmarshalEasyJSON(in)
		case "oldObject":
			(out.OldObject).UnmarshalEasyJSON(in)
		case "dryRun":
			out.DryRun = bool(in.Bool())
		case "options":
			(out.Options).UnmarshalEasyJSON(in)
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
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol4(out *jwriter.Writer, in KubernetesAdmissionRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"uid\":"
		out.RawString(prefix[1:])
		out.String(string(in.Uid))
	}
	{
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		(in.Kind).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"resource\":"
		out.RawString(prefix)
		(in.Resource).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"subResource\":"
		out.RawString(prefix)
		out.String(string(in.SubResource))
	}
	{
		const prefix string = ",\"requestKind\":"
		out.RawString(prefix)
		(in.RequestKind).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"requestResource\":"
		out.RawString(prefix)
		(in.RequestResource).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"requestSubResource\":"
		out.RawString(prefix)
		out.String(string(in.RequestSubResource))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"namespace\":"
		out.RawString(prefix)
		out.String(string(in.Namespace))
	}
	{
		const prefix string = ",\"operation\":"
		out.RawString(prefix)
		out.String(string(in.Operation))
	}
	{
		const prefix string = ",\"userInfo\":"
		out.RawString(prefix)
		(in.UserInfo).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"object\":"
		out.RawString(prefix)
		(in.Object).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"oldObject\":"
		out.RawString(prefix)
		(in.OldObject).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"dryRun\":"
		out.RawString(prefix)
		out.Bool(bool(in.DryRun))
	}
	{
		const prefix string = ",\"options\":"
		out.RawString(prefix)
		(in.Options).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v KubernetesAdmissionRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v KubernetesAdmissionRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *KubernetesAdmissionRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *KubernetesAdmissionRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol4(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol5(in *jlexer.Lexer, out *GroupVersionResource) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "group":
			out.Group = string(in.String())
		case "version":
			out.Version = string(in.String())
		case "kind":
			out.Kind = string(in.String())
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
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol5(out *jwriter.Writer, in GroupVersionResource) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"group\":"
		out.RawString(prefix[1:])
		out.String(string(in.Group))
	}
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GroupVersionResource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GroupVersionResource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GroupVersionResource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GroupVersionResource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol5(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol6(in *jlexer.Lexer, out *GroupVersionKind) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "group":
			out.Group = string(in.String())
		case "version":
			out.Version = string(in.String())
		case "kind":
			out.Kind = string(in.String())
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
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol6(out *jwriter.Writer, in GroupVersionKind) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"group\":"
		out.RawString(prefix[1:])
		out.String(string(in.Group))
	}
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GroupVersionKind) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GroupVersionKind) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoProtocol6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GroupVersionKind) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GroupVersionKind) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoProtocol6(l, v)
}