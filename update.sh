#!/bin/bash

BINARIES="keygen producers"

for binary in $BINARIES; do
mkdir -p build/$binary

cat <<EOF > build/$binary/main.go
package main
import "github.com/mumoshu/variant/pkg/run"
func main() {
    run.YAML(\`
$(cat $binary)
\`)
}
EOF

done
