# mime-fold

解析邮件头块：展开折行、按名字归一、按缺省层补齐缺失的头。只用标准库。

```
go test ./...
go vet ./...
go run ./cmd/mimefold --sample=fold
go run ./cmd/mimefold --sample=case
go run ./cmd/mimefold --sample=multi
go run ./cmd/mimefold --sample=default
go run ./cmd/mimefold --sample=work
```

## 口径

- **折行**：以空格或制表符开头的行是上一行值的续行，拼接时用单个空格连接。
- **名字**：头部名大小写不敏感，解析与查找都统一按小写。
- **重复**：可重复的头（如 `received`、`x-` 开头的自定义头）按出现顺序保留全部值。
- **缺省层**：缺省头只在目标里没有同名头时补上，不覆盖已有值。

## 不变量

- 每个头的值都已展开折行，名字都是小写。
- 查找结果按原始出现顺序返回。
- `scanned` 不随头数与查找次数乘积放大：两千个头查两千次，`scanned` 不超过 6000。

## 输出契约

场景打印一行 JSON，键随场景而定，常见键为 `values`（按序的值列表）与 `count`。
`--sample=work` 打印 `{"headers": N, "scanned": N}`。
