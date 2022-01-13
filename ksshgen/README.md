# ksshgen

suspend execution for an interval of time, maybe shorter or longer.

```bash
go install bitbucket.org/ai69/keiki/ksshgen@latest
```

Bug:

1. passphase is not encrypted for ed25519

2. 2048 + dsa is broken to preview
3. generated dsa is not working on git or ssh

https://stackoverflow.com/questions/70389802/generate-dsa-keys-for-openssh-in-golang
