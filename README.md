# Rapid-Prototype

**MDをLLMから投げると、AWがGo Agentを起動してプロトタイプ成果品を生成する。**

## Factory Lab

これは量産工場ではなく、**試作のためのFactory Lab**。

```text
人間 → Issue → intake → MD → AW → Agent → artifact
LLM   → POST  → intake → MD → AW → Agent → artifact
```

- **Issue** = 人間の「これほしい」郵便受け
- **POST /requests** = 機械の正式な受付口
- **MD** = Agentへ渡す原料
- **AW** = 試作ライン
- **Agent** = `prototyper.go`
- **artifact** = 試作成果品

原則は **「人間はIssueを書く。機械はPOSTする。」**。

## Issue intake

`prototype` または `request` ラベル付きのIssueが開かれると、AWがIssue本文を `input/issue-<number>.md` に変換する。

```bash
gh api \
  repos/bonsai/LAB-Rapid-Prototype/issues \
  -f title='これほしい：Actions Viewer' \
  -f body='GitHub Actionsをブラウザから確認できるEXTがほしい'
```

人間の要望はIssueとして残り、機械からの投入も同じ受付モデルへ寄せる。

## API Contract

`contract/openapi.yaml` が正式な受付Contract。

```http
POST /requests
Content-Type: application/json
```

```json
{
  "title": "これほしい：Actions Viewer",
  "body": "GitHub Actionsをブラウザから確認できるEXTがほしい",
  "source": "llm"
}
```

`source` は `human | llm | cli | mcp | ext`。

## Flow

```text
Issue / POST /requests
          ↓
      input/*.md
          ↓
       git push
          ↓
   GitHub Actions (AW)
          ↓
     prototyper.go
          ↓
          LLM
          ↓
 artifacts/*.prototype.md
          ↓
       git push
          ↺
```

## LLM

OpenAI-compatible Chat Completions APIを使用する。

GitHub Actions Secrets:

- `LLM_API_KEY`

Variables（任意）:

- `LLM_BASE_URL` — 例: `https://api.openai.com/v1`
- `LLM_MODEL` — 例: `gpt-5-mini`

ローカルLLMやOpenAI互換サーバーにも切り替え可能。

## Local

```bash
go run ./cmd/prototyper
```

入力ファイルは `PROTOTYPE_INPUT` で変更できる。

```bash
PROTOTYPE_INPUT=input/example.md \
LLM_API_KEY=... \
go run ./cmd/prototyper
```

## Interfaces

```text
                 OpenAPI Contract
                       │
       ┌───────────────┼───────────────┐
       ↓               ↓               ↓
      API             CLI             MCP
       │                               │
       └───────────────┬───────────────┘
                       ↓
                     Agent
                       ↓
                    Artifact
                       ↑
                      EXT
```

APIを本体にして、CLI / MCP / EXTを入口として増やせる構造にする。
