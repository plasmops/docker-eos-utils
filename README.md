# EOS Utils


## Producers

### Create accounts

```bash
## You can define a custom path to contracts
# export CONTRACTS=/contracts

## dryrun
producers accounts --basename=eoschart --genesis=true --dryrun=true --from=testnet.yaml

## run
producers accounts --basename=eoschart --genesis=true --from=testnet.yaml
```

### Register active producers

```bash
producers register --schedule-genesis=true --schedule-basename=eoschart --schedule-from=testnet.yaml
```
