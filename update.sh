#!/bin/bash

BINARIES="keygen producers"

for binary in $BINARIES; do

cat <<EOF > $binary.go
package main
import "github.com/mumoshu/variant/pkg/run"
func main() {
    run.YAML(\`
$(cat $binary)
\`)
}
EOF

done
