# sputnik

Sputnik network Blockchain Environment

## Sputnik app-chain binaries installation (sputnikd)

```
go: go version go1.22.9 linux/amd64
name: sputnik
```

## Prerequisites

### Install go

```
sudo rm -rvf /usr/local/go/
wget https://golang.org/dl/go1.22.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz
rm go1.22.4.linux-amd64.tar.gz
```

### Put PATH to ~/.profile

```
nano .profile
```

```
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
```

### Use source ~/.profile

```
source .profile
```

### Check go

```
go version
```

### Install packages

```
sudo apt-get update
sudo apt-get upgrade
sudo apt install mc htop screen git gcc make
```

## Binary building

### Clone source from repo

```
git clone https://github.com/olimdzhon/sputnik.git
```

### Open source directory

```
cd sputnik
```

### Build binary

```
make install
```

## Network launch

### Init

```bash:
sputnikd init "<moniker-name>" --chain-id dvs-4.4
```

### Set minimum-gas-prices = "" in app.toml to minimum-gas-prices = "0.25usignal"

```
nano ~/.sputnik/config/app.toml
```

### Generate keys

```bash:
# To create new keypair - make sure you save the mnemonics!
sputnikd keys add <key-name>
```

or

```
# Restore existing odin wallet with mnemonic seed phrase.
# You will be prompted to enter mnemonic seed.
sputnikd keys add <key-name> --recover
```

or

```
# Add keys using ledger
sputnikd keys show <key-name> --ledger
```

Check your key:

```
# Query the keystore for your public address
sputnikd keys show <key-name> -a
```

### Create account to genesis

```
sputnikd genesis add-genesis-account <key-name> 10000000usignal
```

### Create GenTX

```
# Create the gentx.
# Note, your gentx will be rejected if you use any amount greater than 1000000usignal.
# Make sure that all participants built their gentx files without typos.

sputnikd genesis gentx <key-name> 1000000usignal \
  --pubkey=$(sputnikd tendermint show-validator) \
  --chain-id=test-core-1 \
  --moniker="my-moniker" \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01"
```

### Add all accounts to genesis

```
# Add account addresses of all participants before generating genesis.
# (whose Gentx files you're using to generate genesis)
sputnikd genesis add-genesis-account <account-address> 1000000usignal
```

### Generate genesis

```
sputnikd genesis collect-gentxs
```

### Start network

```
sputnikd start
```
