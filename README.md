# EOS Utils


## Producers

### Create accounts

```bash
## You can define custom env
# export CONTRACTS=/contracts
# export EOSURL="http://127.0.0.1:8888"


## dryrun
./producers accounts --basename=eoschart --genesis=true --dryrun=true --from=testnet.yaml

## run
./producers accounts --basename=eoschart --genesis=true --from=testnet.yaml
```

### Register active producers

```bash
./producers register --schedule-genesis=true --schedule-basename=eoschart --schedule-from=testnet.yaml
```
