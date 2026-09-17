# Contract Proof

OpenAPI is the contract. Tests are the evidence that the implementation follows the contract.

```text
OpenAPI
  ↓
Contract Test
  ↓
Go API
  ↓
GitHub Actions
  ↓
PASS = contract evidence
```

## Proof principle

- `contract/openapi.yaml` defines what the API promises.
- Contract tests check the request and response shape.
- GitHub Actions runs the tests automatically.
- A passing workflow is recorded evidence for that commit.

The goal is not only to write a contract, but to continuously prove that the implementation still follows it.
