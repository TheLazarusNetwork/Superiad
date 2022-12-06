# virtuacoin_wallet

Self custodial wallet, covering all primary functions.
Wallet provides these features

- Native Asset for the Configured Network
    - [ ] transfer
    - [ ] checkbalance
    - [ ] fetchwallet
    - [ ] signMessage
    - [ ] verifySignature
- ERC20
    - [ ] transfer
    - [ ] checkbalance
- ERC721
    - [ ] approve
    - [ ] approveAll
    - [ ] transfer
    - [ ] checkbalance
    - [ ] isowner

## Redis for development
```bash
    podman run --name="virtuacoin-wallet-redis" --rm -d -p 6379:6379 redis -c log_statement=all
```