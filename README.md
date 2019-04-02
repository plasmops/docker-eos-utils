# EOS Utils

## keygen

### Generate ordinal keys

#### Given number of keys

```bash
./keygen cleos --keys=3
```

#### Write to a file

```bash
./keygen sync --cleos-keys=3 --out-file=keys
```

Note that operation creates and updates file in case if some entries has been removed. For example:

```
keygen ≫ starting task key
keygen.key ≫ "0": [5KGe8v32SX3vtVGLWMpf4f3aPYYHNgiJDqbKWm16DMY7aH6SZQe, EOS6CWSQGo6AMnFGLAHdyZTsu3RXHApacfVtggwQ2Wcfz2UTe2YWc]
keygen.key ≫ "1": [5KE7unDio3kb83cmmADf8mBrs5UL24Y9Zw5pZTe2MrzBg3MLUQA, EOS8G7u9SdRgshHDQK5h9TyUrmC51Zw9c4XLMfmb5PhGrx61kiVfy]
keygen.key ≫ "2": [5KETArwbqBoioYhDpKPTiTpao1Ftm6SFL4CCHn68PasYwjWucEJ, EOS5HEBUbr6JkSEKyD8HMvg9CJhwbLF3sK5DiojR7ngtxVAMWrnC4]
keygen ≫ starting task out
keygen.out ≫ "0":
keygen.out ≫ - 5J7tVrfw2tJfpQkmzJuqCcDFesJpQjT8AzYbyXebkpFphkk4ajG
keygen.out ≫ - EOS6ekWgkzaMcZ8bSsejB4e4jYKN1JH3MUdq4W7ZdoYXn7bfjbAE9
keygen.out ≫ "2":
keygen.out ≫ - 5JyzrjXvfyHbmNo36K6TW64wDRVSzJiPSYFJT2Lz2fyfrE5iasY
keygen.out ≫ - EOS6byTw6Mo2312vxet5eRNosoPMSUwMJhKo5MVsXDUCyYVCmSwXn
keygen ≫ starting task sync
--- keys	2019-03-28 14:38:19.286695864 +0300
+++ .21739-keys-update	2019-03-28 14:38:22.734782063 +0300
@@ -1,6 +1,9 @@
"0":
- 5J7tVrfw2tJfpQkmzJuqCcDFesJpQjT8AzYbyXebkpFphkk4ajG
- EOS6ekWgkzaMcZ8bSsejB4e4jYKN1JH3MUdq4W7ZdoYXn7bfjbAE9
+"1":
+- 5KE7unDio3kb83cmmADf8mBrs5UL24Y9Zw5pZTe2MrzBg3MLUQA
+- EOS8G7u9SdRgshHDQK5h9TyUrmC51Zw9c4XLMfmb5PhGrx61kiVfy
"2":
- 5JyzrjXvfyHbmNo36K6TW64wDRVSzJiPSYFJT2Lz2fyfrE5iasY
- EOS6byTw6Mo2312vxet5eRNosoPMSUwMJhKo5MVsXDUCyYVCmSwXn
```

### Named keys

To generate named keys, first create a YAML file containing a map as bellow:

```yaml
# testnet-map.yaml
accounts:
  - eosio.bpay
  - eosio.msig
  - eosio.names
  - eosio.ram
  - eosio.ramfee
  - eosio.saving
  - eosio.stake
  - eosio.system
  - eosio.token
  - eosio.vpay
  - arsp
  - audp
  - brlp
```

Now use keygen:

```bash
./keygen sync --cleos-image=registry.plasma-bank.com/blockchain/eos --cleos-map=./testnet-account-ids.yaml --out-file=test
```


## eosinit

### Create: producer accounts

This operation 


```bash
## You can define custom env
# export CONTRACTS=/contracts
# export EOSURL="http://127.0.0.1:8888"


## dryrun
./eosinit create producer-accounts --dryrun=true --from=testnet.yaml

## run
./eosinit create producer-accounts --from=testnet.yaml
```

### Acttion: activate producers setprods

```bash
./eosinit action setprods --schedule-from=testnet.yaml
```
