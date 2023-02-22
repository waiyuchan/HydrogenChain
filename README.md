# HydrogenChain: 氢链

![](https://img.shields.io/badge/language-Golang-50b7e0.svg)
![](https://img.shields.io/badge/backend_frame-Gin-6db33f.svg)
![](https://img.shields.io/badge/deploy-Docker-blue.svg)
![](https://img.shields.io/badge/license-Apache_2.0-green.svg)

一款基于Golang实现的简易区块链服务

## 功能列表
- [ ] 获取某一条私有链的所有块
- [ ] 获取联盟链的所有块
- [ ] 向某一条私有链添加一个新的区块
- [ ] 向联盟链添加一个新的区块

## 区块链系统架构
```markdown
┌───────────────────────┐      ┌───────────────────────┐      ┌───────────────────────┐
│                       │      │                       │      │                       │
│      Node 1 (A)       │◀────▶│      Node 2 (B)       │◀────▶│      Node 3 (C)       │
│                       │      │                       │      │                       │
├───────────────────────┤      ├───────────────────────┤      ├───────────────────────┤
│                       │      │                       │      │                       │
│       API Server      │      │       API Server      │      │       API Server      │
│                       │      │                       │      │                       │
├───────────────┬───────┤      ├───────────────┬───────┤      ├───────────────┬───────┤
│               │       │      │               │       │      │               │       │
│  Blockchain   │  RPC  │      │  Blockchain   │  RPC  │      │  Blockchain   │  RPC  │
│   Database    │Client │      │   Database    │Client │      │   Database    │Client │
│               │       │      │               │       │      │               │       │
└───────────────┴───────┘      └───────────────┴───────┘      └───────────────┴───────┘

```