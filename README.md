# Zypher Network tools in Golang

## 1. POP

### 1.1 verify

You can use this, before your agent send request to LLM, to verify your system prompt. And then it will return you true or false.
```
pop.Verify(context.Background(), []byte(your system prompt), apiKey)
```



