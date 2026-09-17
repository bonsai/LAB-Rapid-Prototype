package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"
  "net/http"
  "os"
  "path/filepath"
  "strings"
)

type Message struct { Role string `json:"role"`; Content string `json:"content"` }
type Request struct { Model string `json:"model"`; Messages []Message `json:"messages"` }
type Response struct { Choices []struct { Message Message `json:"message"` } `json:"choices"` }

func main() {
  input := os.Getenv("PROTOTYPE_INPUT"); if input == "" { input = "input.md" }
  md, err := os.ReadFile(input); if err != nil { panic(err) }
  base := strings.TrimRight(os.Getenv("LLM_BASE_URL"), "/"); if base == "" { base = "https://api.openai.com/v1" }
  key := os.Getenv("LLM_API_KEY"); if key == "" { panic("LLM_API_KEY is required") }
  model := os.Getenv("LLM_MODEL"); if model == "" { model = "gpt-5-mini" }
  system := "You are the Rapid-Prototype Agent. Convert the supplied Markdown into an implementation-oriented prototype specification. Return Markdown only. Include Goal, Inputs, Outputs, API/contract, implementation plan, and acceptance criteria. Do not claim code was built."
  result, err := call(base, key, model, system, string(md)); if err != nil { panic(err) }
  if err := os.MkdirAll("artifacts", 0755); err != nil { panic(err) }
  name := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input))
  out := filepath.Join("artifacts", name+".prototype.md")
  if err := os.WriteFile(out, []byte(result+"\n"), 0644); err != nil { panic(err) }
  fmt.Println(out)
}

func call(base, key, model, system, user string) (string, error) {
  payload, _ := json.Marshal(Request{Model:model, Messages:[]Message{{"system",system},{"user",user}}})
  req, err := http.NewRequest(http.MethodPost, base+"/chat/completions", bytes.NewReader(payload)); if err != nil { return "", err }
  req.Header.Set("Authorization", "Bearer "+key); req.Header.Set("Content-Type", "application/json")
  resp, err := http.DefaultClient.Do(req); if err != nil { return "", err }; defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body); if err != nil { return "", err }
  if resp.StatusCode >= 300 { return "", fmt.Errorf("LLM request failed: %s: %s", resp.Status, body) }
  var decoded Response; if err := json.Unmarshal(body, &decoded); err != nil { return "", err }
  if len(decoded.Choices) == 0 { return "", fmt.Errorf("LLM returned no choices") }
  return decoded.Choices[0].Message.Content, nil
}
