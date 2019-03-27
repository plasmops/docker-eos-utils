package main
import "github.com/mumoshu/variant/pkg/run"
func main() {
    run.YAML(`
#!/usr/bin/env variant

tasks:

  accounts:
    description: Create and run producuer accounts script
    options:
    - name: from
      description: Path to the producers keys data file
      type: object
      required: true
    - name: dryrun
      description: Set to true to print the script only
      type: boolean
      default: false
    script: |
      cat <<EOF | tee /tmp/.$$.create.accounts
      set -e

      GENESIS=\${GENESIS:-no}
      BASENAME=\${BASENAME:-eosproducer}
      CONTRACTS=\${CONTRACTS:-/contracts}
      EOSURL=\${EOSURL:-http://127.0.0.1:8888}

      # keys substitution
      {{- range $i, $e := values (get "from") }}
      KEY_{{ $i }}="{{ first $e }}"
      PUBKEY_{{ $i }}="{{ last $e }}"
      {{- end }}

      # create wallet
      cleos -u \$EOSURL wallet create --file /tmp/wallet.txt

      # import genesis or a simple producers key
      cleos -u \$EOSURL wallet import --private-key \$KEY_0

      i=0      
      if [ "\$GENESIS" != "no" ]; then
        # set the base contract
        cleos -u \$EOSURL set contract eosio \${CONTRACTS}/eosio.bios
        i=1
      fi

      # create producer accounts
      while [ \$i -lt {{ len (keys (get "from" )) }} ]; do
        key=\$(eval echo '$'KEY_\$i)
        pubkey=\$(eval echo '$'PUBKEY_\$i)
        account=\$(eval echo '$'{BASENAME}\$((\$i + 1)))

        cleos -u \$EOSURL wallet import --private-key \$key
        cleos -u \$EOSURL create account eosio \$account \$pubkey
        
        i=\$((\$i + 1))
      done
      EOF
      {{ if not (get "dryrun") }}bash /tmp/.$$.create.accounts{{ end }}
      rs=$?
      rm -f /tmp/.$$.create.accounts && exit $rs

  schedule:
    description: Create producers schedule JSON
    options:
    - name: from
      description: Path to the producers keys data file
      type: object
      required: true
    - name: basename
      description: Producers basename (basename+ordinal=producer_name, ex. eosproducer5)
      type: string
      default: eosproducer
    - name: genesis
      description: Set to true to enable genesis schedule generation (i.e. the first producer_name=eosio)
      type: boolean
      default: false
    ## Notes:
    #  genesis=true - then the first producer name is set to eosio (convention).
    script: |
      cat <<EOF
      {
        "schedule": [
          {{- $items := values (get "from") -}}
          {{- range $i, $e := $items }}
          {
            "producer_name": "{{ if and (get "genesis") (eq $i 0) }}eosio{{ else }}{{ get "basename" }}{{ add1 $i }}{{ end }}",
            "block_signing_key": "{{ last $e }}"
          }{{ if lt (add1 $i) (len ($items)) }},{{ end }}
          {{- end }}
        ]
      }
      EOF

  register:
    description: Register producers
    options:
    - name: schedule
      type: object
      requred: true
    script: |
      EOSURL=${EOSURL:-http://127.0.0.1:8888}
      cleos -u $EOSURL push action eosio setprods '{{ get "schedule" | toJson }}' -p eosio@active 

  cleanup:
    description: Cleanup keosd instance and unlink wallet
    script: |
      keosd_pid=$(ps -ao pid,comm | grep keosd | sed -r 's/ *([0-9]+).*/\1/')
      kill -TERM $keosd_pid
      rm -f /tmp/wallet.txt ~/eosio-wallet
`)
}
