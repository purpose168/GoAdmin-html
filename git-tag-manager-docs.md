# Git标签管理工具操作说明

## 1. 脚本简介

`git-tag-manager.sh` 是一个功能完整的Git标签管理脚本，旨在简化Git标签的创建、管理和操作流程。该脚本提供了直观的命令行界面，支持Git标签管理的各种常用操作，适合团队协作环境中使用。

### 1.1 主要特点

- **用户友好**：提供清晰的彩色提示和错误信息
- **功能完整**：覆盖Git标签管理的所有常用操作
- **安全检查**：自动验证Git环境和仓库状态
- **易于扩展**：模块化设计，便于添加新功能
- **详细帮助**：内置完整的使用说明和示例

## 2. 功能列表

| 功能 | 命令 | 描述 |
|------|------|------|
| 创建轻量级标签 | `create-lightweight` | 创建简单的轻量级标签 |
| 创建附注标签 | `create-annotated` | 创建包含元数据的附注标签 |
| 为历史提交创建标签 | `create-history` | 为指定的历史提交创建标签 |
| 列出标签 | `list` | 列出标签（支持模式筛选） |
| 查看标签详情 | `show` | 查看标签的完整信息 |
| 推送单个标签 | `push-single` | 将单个标签推送到远程仓库 |
| 推送所有标签 | `push-all` | 将所有本地标签推送到远程仓库 |
| 检出标签 | `checkout` | 检出特定标签（分离头指针状态） |
| 基于标签创建分支 | `branch-from-tag` | 基于标签创建新分支 |
| 删除本地标签 | `delete-local` | 删除本地标签 |
| 删除远程标签 | `delete-remote` | 删除远程仓库中的标签 |
| 显示帮助 | `help` | 显示使用帮助信息 |

## 3. 安装步骤

### 3.1 克隆仓库（如果尚未克隆）

```bash
git clone https://github.com/purpose168/GoAdmin-html.git
cd GoAdmin-html
```

### 3.2 赋予脚本执行权限

```bash
chmod +x git-tag-manager.sh
```

### 3.3 验证安装

```bash
./git-tag-manager.sh help
```

如果安装成功，将显示完整的帮助信息。

## 4. 使用方法

### 4.1 基本语法

```bash
./git-tag-manager.sh <功能> [参数]
```

### 4.2 显示帮助信息

```bash
./git-tag-manager.sh help
```

## 5. 功能详细说明

### 5.1 创建轻量级标签

**功能**：创建一个指向当前提交的轻量级标签。

**语法**：
```bash
./git-tag-manager.sh create-lightweight <tag-name>
```

**示例**：
```bash
./git-tag-manager.sh create-lightweight v1.0.0-light
```

**输出**：
```
成功: 成功创建轻量级标签: v1.0.0-light
```

### 5.2 创建附注标签

**功能**：创建一个包含元数据（创建者、日期、标签消息）的附注标签。

**语法**：
```bash
./git-tag-manager.sh create-annotated <tag-name> <message>
```

**示例**：
```bash
./git-tag-manager.sh create-annotated v1.0.0 "第一个稳定版本发布"
```

**输出**：
```
成功: 成功创建附注标签: v1.0.0
```

### 5.3 为历史提交创建标签

**功能**：为指定的历史提交创建附注标签。

**语法**：
```bash
./git-tag-manager.sh create-history <tag-name> <commit-sha> <message>
```

**示例**：
```bash
./git-tag-manager.sh create-history v0.9.0 dfd107d "为历史提交创建标签"
```

**输出**：
```
成功: 成功为历史提交 dfd107d 创建标签: v0.9.0
```

### 5.4 列出标签

**功能**：列出所有本地标签，支持模式筛选。

**语法**：
```bash
./git-tag-manager.sh list [pattern]
```

**示例1**：列出所有标签
```bash
./git-tag-manager.sh list
```

**输出1**：
```
v0.9.0
v1.0.0
v1.0.0-light
```

**示例2**：按模式筛选标签
```bash
./git-tag-manager.sh list v1.*
```

**输出2**：
```
v1.0.0
v1.0.0-light
```

### 5.5 查看标签详情

**功能**：查看标签的完整信息，包括标签元数据和关联的提交信息。

**语法**：
```bash
./git-tag-manager.sh show <tag-name>
```

**示例**：
```bash
./git-tag-manager.sh show v1.0.0
```

**输出**：
```
tag v1.0.0
Tagger: User Name <user@example.com>
Date:   Sun Jan 5 12:00:00 2026 +0800

第一个稳定版本发布

commit dfd107dc52dc55923c049da7a26e2f98b4201356 (tag: v1.0.0)
Author: User Name <user@example.com>
Date:   Sun Jan 5 12:00:00 2026 +0800

    提交信息
```

### 5.6 推送单个标签

**功能**：将单个标签推送到远程仓库。

**语法**：
```bash
./git-tag-manager.sh push-single <tag-name>
```

**示例**：
```bash
./git-tag-manager.sh push-single v1.0.0
```

**输出**：
```
成功: 成功推送标签 v1.0.0 到远程
```

### 5.7 推送所有标签

**功能**：将所有本地标签推送到远程仓库。

**语法**：
```bash
./git-tag-manager.sh push-all
```

**输出**：
```
成功: 成功推送所有标签到远程
```

### 5.8 检出标签

**功能**：检出特定标签，进入分离头指针状态。

**语法**：
```bash
./git-tag-manager.sh checkout <tag-name>
```

**示例**：
```bash
./git-tag-manager.sh checkout v1.0.0
```

**输出**：
```
成功: 成功检出标签: v1.0.0
警告: 注意：当前处于分离头指针状态，建议基于标签创建分支进行开发
```

### 5.9 基于标签创建分支

**功能**：基于指定标签创建一个新分支，适合在标签基础上进行开发。

**语法**：
```bash
./git-tag-manager.sh branch-from-tag <branch-name> <tag-name>
```

**示例**：
```bash
./git-tag-manager.sh branch-from-tag feature/v1.0.1 v1.0.0
```

**输出**：
```
成功: 成功基于标签 v1.0.0 创建分支: feature/v1.0.1
```

### 5.10 删除本地标签

**功能**：删除本地仓库中的指定标签。

**语法**：
```bash
./git-tag-manager.sh delete-local <tag-name>
```

**示例**：
```bash
./git-tag-manager.sh delete-local v1.0.0-light
```

**输出**：
```
成功: 成功删除本地标签: v1.0.0-light
```

### 5.11 删除远程标签

**功能**：删除远程仓库中的指定标签。

**语法**：
```bash
./git-tag-manager.sh delete-remote <tag-name>
```

**示例**：
```bash
./git-tag-manager.sh delete-remote v0.9.0
```

**输出**：
```
成功: 成功删除远程标签: v0.9.0
```

## 6. 最佳实践

### 6.1 标签命名规范

- **语义化版本**：遵循 `major.minor.patch` 格式（如 `v1.2.3`）
- **预发布标签**：使用 `-alpha`、`-beta`、`-rc` 等后缀（如 `v2.0.0-beta`）
- **环境标签**：可添加环境后缀（如 `v1.0.0-prod`、`v1.0.0-staging`）
- **使用 `v` 前缀**：有助于区分标签和分支名

### 6.2 标签使用建议

1. **正式发布使用附注标签**：附注标签包含完整的元数据，便于追溯
2. **临时标记使用轻量级标签**：轻量级标签适合本地临时标记
3. **定期清理过期标签**：避免标签过多导致混乱
4. **推送标签前确认**：确保标签内容正确后再推送到远程
5. **使用GPG签名重要标签**：增强标签的安全性和可信度

### 6.3 团队协作建议

1. **统一标签策略**：团队内部约定标签命名规范和使用流程
2. **标签与CHANGELOG结合**：每次创建标签时更新CHANGELOG文件
3. **标签与CI/CD集成**：结合自动化流程自动创建和推送标签
4. **避免修改已推送标签**：已推送到远程的标签应视为不可变

## 7. 注意事项

1. **分离头指针状态**：直接检出标签会进入分离头指针状态，建议基于标签创建分支进行开发
2. **远程标签删除**：删除远程标签后，其他团队成员需要执行 `git fetch --prune --tags` 更新本地标签列表
3. **标签不可修改**：Git标签一旦创建，建议不要修改，如需更新应创建新标签
4. **权限问题**：删除远程标签需要相应的仓库权限
5. **网络连接**：推送和删除远程标签需要稳定的网络连接

## 8. 故障排除

### 8.1 常见错误及解决方法

| 错误信息 | 可能原因 | 解决方法 |
|----------|----------|----------|
| `Git未安装，请先安装Git` | 系统未安装Git | 安装Git：`sudo apt install git`（Ubuntu）或 `brew install git`（macOS） |
| `当前目录不是Git仓库` | 在非Git仓库目录执行命令 | 切换到Git仓库目录或初始化Git仓库：`git init` |
| `创建标签失败` | 标签名已存在或权限问题 | 检查标签名是否已存在，或使用不同的标签名 |
| `推送标签失败` | 网络问题或权限不足 | 检查网络连接，确保拥有推送权限 |
| `删除标签失败` | 标签不存在或权限不足 | 检查标签名是否正确，或确认拥有相应权限 |

### 8.2 获取帮助

如果遇到其他问题，可以：

1. 查看脚本帮助：`./git-tag-manager.sh help`
2. 检查Git状态：`git status`
3. 查看Git日志：`git log --oneline`
4. 检查远程标签：`git ls-remote --tags origin`

## 9. 更新日志

### v1.0.0 (2026-01-05)

- 初始版本发布
- 支持创建、查看、推送、检出和删除标签
- 提供完整的彩色提示和错误处理
- 包含详细的使用文档和示例

## 10. 贡献指南

欢迎提交问题和改进建议！如果您想为该脚本贡献代码，请：

1. Fork 仓库
2. 创建特性分支
3. 提交更改
4. 推送分支
5. 打开 Pull Request

## 11. 许可证

本脚本采用 MIT 许可证，详情请查看项目根目录下的 LICENSE 文件。

## 12. 联系方式

如有问题或建议，欢迎通过以下方式联系：

- 项目地址：https://github.com/purpose168/GoAdmin-html
- 问题反馈：https://github.com/purpose168/GoAdmin-html/issues

---

**文档更新日期**：2026-01-05
**文档版本**：v1.0.0
