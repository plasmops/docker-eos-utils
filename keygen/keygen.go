package main
import "github.com/mumoshu/variant/pkg/run"
func main() {
    run.YAML(`
#!/usr/bin/env variant

tasks:
  ## Generate a key
  key:
    description: Generates eos key-pair(s)
    parameters:
    - name: count
      description: Number of key-pairs to create
      type: integer
      default: 1
    runner:
      image: eosio/eos
      command: bash
      args: [-c]

    ## Generate a JSON object with a given number of keypairs
    script: |
      {{- range $i, $e := until (.count) }}
      pair=$(cleos create key --to-console | sed 's/.*: //' | sed ':a;N;$!ba;s/\n/, /g')
      printf '"{{ $i }}": [%s]\n' "$pair"
      {{- end }}

  out:
    description: Cats a keypairs yaml file
    parameters:
    - name: file
      description: Path to a keypairs yaml file
      type: string
      required: true
    script: |
      if [ -f "{{ .file }}" ]; then
        cat {{ .file }}
      else
        echo "{}"
      fi

  ## Sync a keyfile
  sync:
    description: Creates a keypair yaml file
    options:
    - name: key
      type: object
      required: true
    - name: out
      type: object
      required: true
    script: |
      set -e
      {{- $update := merge .key .out }}
      cat <<EOF > .$$-{{ .file }}-update
      {{ $update | toYaml -}}
      EOF

      if [ -f "{{ .file }}" ]; then
        ## sync changes
        diff -u {{ .file }} .$$-{{ .file }}-update || \
          cat .$$-{{ .file }}-update > {{ .file }}
      else
        ## create new keysfile
        cat .$$-{{ .file }}-update > {{ .file }}
        echo "... Created new file - {{ .file }}"
      fi

      rm -f .$$-{{ .file }}-update
`)
}