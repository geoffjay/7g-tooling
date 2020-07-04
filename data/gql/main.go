//go:generate go-bindata -ignore=\.go -pkg=gql -o=bindata.go ./...
package gql

import "bytes"

func GetQueryList() []string {
	return AssetNames()
}

func GetQuery(name string) string {
	buf := bytes.Buffer{}
	b := MustAsset(name)
	buf.Write(b)

	// Add a newline if the file does not end in a newline.
	if len(b) > 0 && b[len(b)-1] != '\n' {
		buf.WriteByte('\n')
	}

	return buf.String()
}
