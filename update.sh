#!/bin/bash

cat <<EOF > producers.go
package main
import "github.com/mumoshu/variant/pkg/run"
func main() {
    run.YAML(\`
$(cat producers)
\`)
}
EOF