// this file was generated by gomacro command: import "mime/quotedprintable"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"mime/quotedprintable"
)

func init() {
	Binds["mime/quotedprintable"] = map[string]Value{
		"NewReader":	ValueOf(quotedprintable.NewReader),
		"NewWriter":	ValueOf(quotedprintable.NewWriter),
	}
	Types["mime/quotedprintable"] = map[string]Type{
		"Reader":	TypeOf((*quotedprintable.Reader)(nil)).Elem(),
		"Writer":	TypeOf((*quotedprintable.Writer)(nil)).Elem(),
	}
	Proxies["mime/quotedprintable"] = map[string]Type{
	}
}