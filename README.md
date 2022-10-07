# Superiad

A custodial wallet, covering all primary functions.
This custodial wallet provides these features

- ERC721
    - approve
    - approveAll
    - transfer
    - checkbalance
    - isowner
- ERC20
    - transfer
    - checkbalance
- Native
    - transfer
    - checkbalance
    - fetchwallet
    - signMessage
    - verifySignature

This system also supports locking the wallet using which the user can lock his funds.

## Postgres for development
```bash
 docker run --name="superiad" --rm -d -p 5432:5432 \
      -e POSTGRES_PASSWORD=superiad \
      -e POSTGRES_USER=superiad \
      -e POSTGRES_DB=superiad \
      postgres -c log_statement=all
```