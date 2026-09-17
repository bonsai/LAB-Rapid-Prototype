# test: OpenAPI契約を自動で証明するContract Test

Source: GitHub Issue #3

## 目的

`contract/openapi.yaml` に書いたAPI契約を、テストで自動的に証明する。

## 考え方

```text
OpenAPI
  ↓
契約
  ↓
Go API
  ↓
Contract Test
  ↓
PASS
  ↓
「このコードは契約を守った」という証拠
```

- OpenAPI = 契約書
- Go = 契約を実装する人
- Contract Test = 契約を守ったか確認するテスト
- GitHub Actions = 毎回確認して証拠を残す係

## 対象

`POST /requests`

### Request

- `title`: 必須
- `body`: 必須
- `source`: `human | llm | cli | mcp | ext`

### Response

- HTTP `201`
- JSON
- `accepted: true`
- 送信した `title` / `body` / `source` を返す
- GitHub Issueを作成した場合はIssue number / URLも返す

## Acceptance Criteria

- [ ] OpenAPIのrequest schemaとテスト内容が一致する
- [ ] `POST /requests` をテストする
- [ ] 必須項目を検証する
- [ ] `201` responseを検証する
- [ ] response schemaを検証する
- [ ] GitHub Actionsで `go test ./...` を実行する
- [ ] PASSしたCI runを契約証明のEvidenceとして扱える
- [ ] OpenAPI変更時に契約テストも更新する

## ゴール

**Contract → Test → Evidence** をLABの標準パターンにする。

> Contract is the promise.
> Test is the proof.
> GitHub Actions is the witness.
