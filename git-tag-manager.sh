#!/bin/bash

# Git标签管理工具
# 功能：
# 1. 创建轻量级标签
# 2. 创建附注标签
# 3. 为历史提交创建标签
# 4. 查看标签列表
# 5. 查看标签详细信息
# 6. 推送单个标签到远程
# 7. 推送所有标签到远程
# 8. 检出特定标签
# 9. 基于标签创建新分支
# 10. 删除本地标签
# 11. 删除远程标签

# 颜色定义
GREEN="\033[0;32m"
RED="\033[0;31m"
YELLOW="\033[1;33m"
BLUE="\033[0;34m"
NC="\033[0m" # No Color

# 打印带有颜色的消息
print_message() {
    local color=$1
    local message=$2
    echo -e "${color}${message}${NC}"
}

# 打印成功消息
print_success() {
    local message=$1
    print_message "${GREEN}" "成功: ${message}"
}

# 打印错误消息
print_error() {
    local message=$1
    print_message "${RED}" "错误: ${message}"
}

# 打印警告消息
print_warning() {
    local message=$1
    print_message "${YELLOW}" "警告: ${message}"
}

# 打印信息消息
print_info() {
    local message=$1
    print_message "${BLUE}" "信息: ${message}"
}

# 检查Git是否可用
check_git_available() {
    if ! command -v git &> /dev/null; then
        print_error "Git未安装，请先安装Git"
        exit 1
    fi
}

# 检查是否在Git仓库中
check_in_git_repo() {
    if ! git rev-parse --is-inside-work-tree &> /dev/null; then
        print_error "当前目录不是Git仓库"
        exit 1
    fi
}

# 显示使用帮助
show_help() {
    echo "Git标签管理工具使用说明："
    echo ""
    echo "使用方式："
    echo "  ./git-tag-manager.sh <功能> [参数]"
    echo ""
    echo "功能列表："
    echo "  create-lightweight <tag-name>          创建轻量级标签"
    echo "  create-annotated <tag-name> <message>   创建附注标签"
    echo "  create-history <tag-name> <commit-sha> <message>  为历史提交创建标签"
    echo "  list [pattern]                         列出标签（可带模式筛选）"
    echo "  show <tag-name>                        查看标签详细信息"
    echo "  push-single <tag-name>                 推送单个标签到远程"
    echo "  push-all                               推送所有标签到远程"
    echo "  checkout <tag-name>                    检出特定标签（分离头指针）"
    echo "  branch-from-tag <branch-name> <tag-name>  基于标签创建新分支"
    echo "  delete-local <tag-name>                删除本地标签"
    echo "  delete-remote <tag-name>               删除远程标签"
    echo "  help                                   显示此帮助信息"
    echo ""
    echo "示例："
    echo "  ./git-tag-manager.sh create-annotated v1.0.0 \"第一个稳定版本\""
    echo "  ./git-tag-manager.sh list v1.*"
    echo "  ./git-tag-manager.sh push-single v1.0.0"
    echo "  ./git-tag-manager.sh delete-local v0.9.0"
}

# 创建轻量级标签
create_lightweight() {
    local tag_name=$1
    if [ -z "$tag_name" ]; then
        print_error "请提供标签名称"
        show_help
        exit 1
    fi
    
    git tag "$tag_name"
    if [ $? -eq 0 ]; then
        print_success "成功创建轻量级标签: $tag_name"
    else
        print_error "创建轻量级标签失败"
        exit 1
    fi
}

# 创建附注标签
create_annotated() {
    local tag_name=$1
    local message=$2
    if [ -z "$tag_name" ] || [ -z "$message" ]; then
        print_error "请提供标签名称和标签消息"
        show_help
        exit 1
    fi
    
    git tag -a "$tag_name" -m "$message"
    if [ $? -eq 0 ]; then
        print_success "成功创建附注标签: $tag_name"
    else
        print_error "创建附注标签失败"
        exit 1
    fi
}

# 为历史提交创建标签
create_history() {
    local tag_name=$1
    local commit_sha=$2
    local message=$3
    if [ -z "$tag_name" ] || [ -z "$commit_sha" ] || [ -z "$message" ]; then
        print_error "请提供标签名称、提交哈希值和标签消息"
        show_help
        exit 1
    fi
    
    git tag -a "$tag_name" "$commit_sha" -m "$message"
    if [ $? -eq 0 ]; then
        print_success "成功为历史提交 $commit_sha 创建标签: $tag_name"
    else
        print_error "为历史提交创建标签失败"
        exit 1
    fi
}

# 列出标签
list_tags() {
    local pattern=$1
    if [ -z "$pattern" ]; then
        git tag -l
    else
        git tag -l "$pattern"
    fi
    
    if [ $? -ne 0 ]; then
        print_error "列出标签失败"
        exit 1
    fi
}

# 查看标签详细信息
show_tag() {
    local tag_name=$1
    if [ -z "$tag_name" ]; then
        print_error "请提供标签名称"
        show_help
        exit 1
    fi
    
    git show "$tag_name"
    if [ $? -ne 0 ]; then
        print_error "查看标签详细信息失败"
        exit 1
    fi
}

# 推送单个标签到远程
push_single() {
    local tag_name=$1
    if [ -z "$tag_name" ]; then
        print_error "请提供标签名称"
        show_help
        exit 1
    fi
    
    git push origin "$tag_name"
    if [ $? -eq 0 ]; then
        print_success "成功推送标签 $tag_name 到远程"
    else
        print_error "推送标签失败"
        exit 1
    fi
}

# 推送所有标签到远程
push_all() {
    git push origin --tags
    if [ $? -eq 0 ]; then
        print_success "成功推送所有标签到远程"
    else
        print_error "推送所有标签失败"
        exit 1
    fi
}

# 检出特定标签
checkout_tag() {
    local tag_name=$1
    if [ -z "$tag_name" ]; then
        print_error "请提供标签名称"
        show_help
        exit 1
    fi
    
    git checkout "$tag_name"
    if [ $? -eq 0 ]; then
        print_success "成功检出标签: $tag_name"
        print_warning "注意：当前处于分离头指针状态，建议基于标签创建分支进行开发"
    else
        print_error "检出标签失败"
        exit 1
    fi
}

# 基于标签创建新分支
branch_from_tag() {
    local branch_name=$1
    local tag_name=$2
    if [ -z "$branch_name" ] || [ -z "$tag_name" ]; then
        print_error "请提供分支名称和标签名称"
        show_help
        exit 1
    fi
    
    git checkout -b "$branch_name" "$tag_name"
    if [ $? -eq 0 ]; then
        print_success "成功基于标签 $tag_name 创建分支: $branch_name"
    else
        print_error "基于标签创建分支失败"
        exit 1
    fi
}

# 删除本地标签
delete_local() {
    local tag_name=$1
    if [ -z "$tag_name" ]; then
        print_error "请提供标签名称"
        show_help
        exit 1
    fi
    
    git tag -d "$tag_name"
    if [ $? -eq 0 ]; then
        print_success "成功删除本地标签: $tag_name"
    else
        print_error "删除本地标签失败"
        exit 1
    fi
}

# 删除远程标签
delete_remote() {
    local tag_name=$1
    if [ -z "$tag_name" ]; then
        print_error "请提供标签名称"
        show_help
        exit 1
    fi
    
    git push origin --delete "$tag_name"
    if [ $? -eq 0 ]; then
        print_success "成功删除远程标签: $tag_name"
    else
        print_error "删除远程标签失败"
        exit 1
    fi
}

# 主函数
main() {
    # 检查Git环境
    check_git_available
    check_in_git_repo
    
    # 获取功能参数
    local function=$1
    shift
    
    # 根据功能执行相应操作
    case "$function" in
        create-lightweight)
            create_lightweight "$@"
            ;;
        create-annotated)
            create_annotated "$@"
            ;;
        create-history)
            create_history "$@"
            ;;
        list)
            list_tags "$@"
            ;;
        show)
            show_tag "$@"
            ;;
        push-single)
            push_single "$@"
            ;;
        push-all)
            push_all "$@"
            ;;
        checkout)
            checkout_tag "$@"
            ;;
        branch-from-tag)
            branch_from_tag "$@"
            ;;
        delete-local)
            delete_local "$@"
            ;;
        delete-remote)
            delete_remote "$@"
            ;;
        help)
            show_help
            ;;
        *)
            print_error "无效的功能：$function"
            show_help
            exit 1
            ;;
    esac
}

# 执行主函数
main "$@"
