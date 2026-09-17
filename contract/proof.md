# Contract Proof

## Goal

Prove that the implementation follows `contract/openapi.yaml`.

## Evidence chain

```text
OpenAPI contract
      ↓
request validation
      ↓
Go implementation
      ↓
response validation
      ↓
go test ./...
      ↓
GitHub Actions
      ↓
PASS
```

## Acceptance Criteria

- [ ] OpenAPI defines the request schema.
- [ ] OpenAPI defines the response schema.
- [ ] Go tests verify required request fields.
- [ ] Go tests verify the response contains the promised Issue number and URL.
- [ ] GitHub Actions runs the contract tests on every change.
- [ ] A successful Actions run is treated as the evidence for that commit.

## Principle

**Contract is the promise. Test is the proof. GitHub Actions is the witness.**
