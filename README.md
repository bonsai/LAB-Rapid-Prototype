# Rapid-Prototype

**MDをLLMから投げると、AWがGo Agentを起動してプロトタイプ成果品を生成する。**

## Flow

```text
LLM / Human
    ↓
  input/*.md
    ↓ git push
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

## Idea

これは量産工場ではなく、**試作のためのFactory Lab**。

MDを「原料」として投入し、Agentが要求をプロトタイプ仕様へ変換し、成果品をリポジトリへ戻す。

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

## AW

`input/*.md` のpushをトリガーに、GitHub Actionsが `go test` → `go build` → Agent実行 → artifact生成 → commit/push まで行う。
