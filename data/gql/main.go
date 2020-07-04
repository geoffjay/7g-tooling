//go:generate go-bindata -ignore=\.go -pkg=gql -o=bindata.go ./...
package gql

import (
	"bytes"
	"fmt"
)

func GetQueryList() []string {
	return AssetNames()
}

func GetQuery(query string) string {
	buf := bytes.Buffer{}
	name := fmt.Sprintf("%s.gql", query)
	b := MustAsset(name)
	buf.Write(b)

	// Add a newline if the file does not end in a newline.
	if len(b) > 0 && b[len(b)-1] != '\n' {
		buf.WriteByte('\n')
	}

	return buf.String()
}
