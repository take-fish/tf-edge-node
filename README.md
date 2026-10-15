# tf-edge-node

FishMatrix（摸鱼矩阵）边缘节点服务 —— 基于 Go + Gin，单二进制分发到每个成员机器运行。

- 角色：本机业务服务的唯一网络进出口，后续承接 mTLS 隧道、P2P 转发、本地缓存
- 端口：`8090`（可用 `PORT` 环境变量覆盖）
- 健康检查：`GET /healthz`

## 本地开发

```bash
go run .            # 启动开发服务
go test ./...       # 运行测试
golangci-lint run   # 静态检查（需安装 golangci-lint）
```

## 构建

```bash
go build -trimpath -ldflags "-s -w -X main.version=v0.1.0" -o bin/tf-edge-node .
```

## 分支模型

- `master`：保护分支，只接受 PR（CI 必须通过），合并后自动发版
- `test`：联调分支，可直接推送
- 开发流：`feature/*` → PR 到 `test` → 联调 → PR 到 `master` → 自动发版
