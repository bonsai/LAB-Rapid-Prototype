# POC: POST /requests を GitHub Issue に接続する

Source: GitHub Issue #1

## Goal
`POST /requests` をFactory Labの正式な郵便受けとして実装する。

## Flow
```text
LLM / CLI / MCP / EXT
        ↓
 POST /requests
        ↓
 GitHub Issue
        ↓
 input/issue-*.md
        ↓
 AW
        ↓
 prototyper.go
        ↓
 artifact
```

## Acceptance Criteria
- [ ] Go HTTP APIを追加
- [ ] `POST /requests` をOpenAPI contractに準拠させる
- [ ] GitHub APIでIssueを作成できる
- [ ] `prototype` ラベルを付与する
- [ ] API responseにIssue番号/URLを返す
- [ ] テストを追加する
- [ ] GitHub Actionsでbuild/testできる

## Principle
**人間はIssueを書く。機械はPOSTする。**

Issueをrequestの記録、POSTをmachine intakeとして扱う。
