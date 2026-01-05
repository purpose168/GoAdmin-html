#!/bin/bash

# ============================================================================
# Git提交助手脚本
# ============================================================================
# 功能描述：
#   该脚本用于自动化Git提交流程，支持约定式提交规范（Conventional Commits）
#   主要功能包括：
#   1. 检查Git环境和工作区状态
#   2. 交互式输入提交信息（类型、作用域、描述）
#   3. 验证提交信息格式
#   4. 自动执行git add和git commit命令
#   5. 提供清晰的执行反馈和错误处理
#
# 使用方法：
#   ./git-commit.sh
#
# 依赖要求：
#   - Git命令行工具
#   - Bash shell环境
# ============================================================================

# 约定式提交类型列表
# 说明：这些是符合Conventional Commits规范的标准提交类型
# feat: 新功能（feature）
# fix: 修复bug
# docs: 文档变更
# style: 代码格式调整（不影响代码运行的变更）
# refactor: 重构（既不是新增功能，也不是修复bug的代码变更）
# test: 添加测试或修正测试
# chore: 构建过程或辅助工具的变动
# build: 影响构建系统或外部依赖的变更
# ci: CI配置文件和脚本的变更
# perf: 性能优化
# revert: 回退之前的提交
VALID_TYPES=("feat" "fix" "docs" "style" "refactor" "test" "chore" "build" "ci" "perf" "revert")

# 颜色定义
# 说明：使用ANSI转义码为终端输出添加颜色，提高可读性
# \033[0;32m: 绿色（用于成功消息）
# \033[0;31m: 红色（用于错误消息）
# \033[1;33m: 黄色（用于警告消息）
# \033[0;34m: 蓝色（用于信息消息）
# \033[0m: 重置颜色（恢复默认颜色）
GREEN="\033[0;32m"
RED="\033[0;31m"
YELLOW="\033[1;33m"
BLUE="\033[0;34m"
NC="\033[0m" # No Color (重置颜色)

# ============================================================================
# 工具函数
# ============================================================================

# 打印带有颜色的消息
# 参数：
#   $1 (color): 颜色代码（如GREEN、RED等）
#   $2 (message): 要打印的消息内容
# 返回值：无
# 说明：该函数用于统一管理终端输出的颜色格式
print_message() {
    local color=$1      # 颜色代码
    local message=$2    # 消息内容
    echo -e "${color}${message}${NC}"  # -e参数启用转义序列
}

# 打印错误消息并退出
# 参数：
#   $1 (message): 错误消息内容
# 返回值：无（函数会调用exit 1退出脚本）
# 说明：该函数用于在发生错误时显示错误信息并终止脚本执行
print_error_and_exit() {
    local message=$1    # 错误消息内容
    print_message "${RED}" "错误: ${message}"  # 使用红色显示错误
    exit 1  # 退出脚本，返回状态码1表示错误
}

# 打印成功消息
# 参数：
#   $1 (message): 成功消息内容
# 返回值：无
# 说明：该函数用于显示操作成功的消息
print_success() {
    local message=$1    # 成功消息内容
    print_message "${GREEN}" "成功: ${message}"  # 使用绿色显示成功消息
}

# 打印警告消息
# 参数：
#   $1 (message): 警告消息内容
# 返回值：无
# 说明：该函数用于显示警告信息，提示用户注意某些情况
print_warning() {
    local message=$1    # 警告消息内容
    print_message "${YELLOW}" "警告: ${message}"  # 使用黄色显示警告
}

# 打印信息消息
# 参数：
#   $1 (message): 信息消息内容
# 返回值：无
# 说明：该函数用于显示一般性的信息提示
print_info() {
    local message=$1    # 信息消息内容
    print_message "${BLUE}" "信息: ${message}"  # 使用蓝色显示信息
}

# ============================================================================
# 环境检查函数
# ============================================================================

# 检查Git是否可用
# 参数：无
# 返回值：无（如果Git不可用，会调用print_error_and_exit退出）
# 说明：该函数检查系统中是否安装了Git命令行工具
#       使用command -v命令检查git命令是否存在
#       &> /dev/null将输出重定向到/dev/null，不显示任何输出
check_git_available() {
    # command -v git检查git命令是否存在
    # !表示取反，如果git不存在则执行if块内的代码
    if ! command -v git &> /dev/null; then
        print_error_and_exit "Git未安装，请先安装Git"
    fi
}

# 检查是否在Git仓库中
# 参数：无
# 返回值：无（如果不在Git仓库中，会调用print_error_and_exit退出）
# 说明：该函数检查当前目录是否是一个Git仓库
#       使用git rev-parse --is-inside-work-tree命令检查
check_in_git_repo() {
    # git rev-parse --is-inside-work-tree检查是否在Git工作树中
    # 如果不在Git仓库中，命令会返回非零状态码
    if ! git rev-parse --is-inside-work-tree &> /dev/null; then
        print_error_and_exit "当前目录不是Git仓库"
    fi
}

# 检查工作区状态
# 参数：无
# 返回值：
#   0: 有未暂存的变更
#   1: 无变更（会调用print_error_and_exit退出）
# 说明：该函数检查工作区是否有未跟踪的文件或未暂存的修改
#       使用git status --porcelain命令获取简洁的状态输出
#       grep -q .检查输出是否为空（.表示匹配任何字符）
check_workspace_status() {
    print_info "检查工作区状态..."
    
    # git status --porcelain以简洁格式显示工作区状态
    # 输出格式：XY PATH
    #   X: 索引区的状态
    #   Y: 工作区的状态
    # grep -q .检查是否有任何输出（即是否有变更）
    if git status --porcelain | grep -q .; then
        print_info "发现未暂存的变更"
        return 0  # 有变更，返回0
    else
        print_error_and_exit "工作区无变更，无需提交"
    fi
}

# ============================================================================
# 验证函数
# ============================================================================

# 验证提交类型是否有效
# 参数：
#   $1 (type): 要验证的提交类型
# 返回值：
#   0: 类型有效
#   1: 类型无效
# 说明：该函数检查用户输入的提交类型是否在VALID_TYPES列表中
validate_type() {
    local type=$1  # 要验证的提交类型
    
    # 遍历所有有效的提交类型
    for valid_type in "${VALID_TYPES[@]}"; do
        # 如果找到匹配的类型，返回0（有效）
        if [ "$valid_type" = "$type" ]; then
            return 0
        fi
    done
    
    # 遍历完所有类型都没找到匹配，返回1（无效）
    return 1
}

# 验证提交信息格式
# 参数：
#   $1 (commit_msg): 要验证的提交信息
# 返回值：无（如果格式不符合规范，会调用print_error_and_exit退出）
# 说明：该函数验证提交信息是否符合约定式提交规范
#       格式要求：type(scope): description
#       - type: 小写字母
#       - scope: 可选，用括号括起来，可包含字母、数字、下划线、连字符
#       - description: 必需，以空格开头，首字母小写
validate_commit_message() {
    local commit_msg=$1  # 要验证的提交信息
    
    # 使用正则表达式验证提交信息格式
    # ^[a-z]+: 以小写字母开头（提交类型）
    # ([a-zA-Z0-9_-]+)?: 可选的作用域，包含字母、数字、下划线、连字符
    # [[:space:]]: 一个空格（冒号后的空格）
    # [a-zA-Z0-9]: 描述以字母或数字开头
    if [[ ! $commit_msg =~ ^[a-z]+([a-zA-Z0-9_-]+)?:[[:space:]][a-zA-Z0-9] ]]; then
        print_error_and_exit "提交信息不符合约定式提交规范\n示例: feat(login): add password reset functionality"
    fi
}

# ============================================================================
# 交互函数
# ============================================================================

# 获取提交信息
# 参数：无
# 返回值：输出符合规范的提交信息字符串
# 说明：该函数通过交互式提示获取用户输入的提交信息
#       包括提交类型、作用域（可选）和描述
get_commit_message() {
    # 显示提交信息格式说明
    print_info "请输入符合约定式提交规范的提交信息"
    print_info "格式: type(scope): description"
    print_info "有效类型: ${VALID_TYPES[*]}"
    print_info "scope为可选字段，用于指定变更范围"
    print_info "description为简短描述，首字母小写，不超过50个字符"
    
    local type scope description commit_msg  # 声明局部变量
    
    # 获取提交类型
    # 使用while循环持续提示，直到输入有效的类型
    while true; do
        read -p "请输入提交类型: " type  # 提示用户输入提交类型
        if validate_type "$type"; then  # 验证类型是否有效
            break  # 类型有效，退出循环
        else
            # 类型无效，显示警告并重新提示
            print_warning "无效的提交类型，请重新输入"
            print_info "有效类型: ${VALID_TYPES[*]}"
        fi
    done
    
    # 获取作用域（可选）
    read -p "请输入作用域(可选): " scope  # 提示用户输入作用域
    
    # 获取描述
    # 使用while循环持续提示，直到输入非空的描述
    while true; do
        read -p "请输入描述: " description  # 提示用户输入描述
        if [[ -n $description ]]; then  # 检查描述是否非空
            break  # 描述非空，退出循环
        else
            # 描述为空，显示警告并重新提示
            print_warning "描述不能为空，请重新输入"
        fi
    done
    
    # 构建提交信息
    # 如果提供了作用域，格式为type(scope): description
    # 如果没有提供作用域，格式为type: description
    if [[ -n $scope ]]; then
        commit_msg="$type($scope): $description"
    else
        commit_msg="$type: $description"
    fi
    
    # 验证构建的提交信息格式
    validate_commit_message "$commit_msg"
    
    # 输出提交信息（供调用者使用）
    echo "$commit_msg"
}

# ============================================================================
# Git操作函数
# ============================================================================

# 执行Git提交
# 参数：
#   $1 (commit_msg): 提交信息
# 返回值：
#   0: 提交成功
#   1: 提交失败（会调用print_error_and_exit退出）
# 说明：该函数执行git add和git commit命令完成提交
perform_git_commit() {
    local commit_msg=$1  # 提交信息
    
    # 执行git add命令，暂存所有变更
    print_info "执行git add命令..."
    if ! git add .; then  # git add .暂存当前目录及其子目录的所有变更
        print_error_and_exit "git add命令执行失败"
    fi
    
    # 执行git commit命令，提交变更
    print_info "执行git commit命令..."
    if git commit -m "$commit_msg"; then  # 使用指定的提交信息提交
        # 提交成功，显示成功消息
        print_success "提交成功！"
        print_success "提交信息: $commit_msg"
        return 0  # 返回0表示成功
    else
        # 提交失败，显示错误并退出
        print_error_and_exit "git commit命令执行失败"
    fi
}

# ============================================================================
# 主函数
# ============================================================================

# 主函数
# 说明：该函数是脚本的入口点，按顺序执行以下操作：
#   1. 检查Git环境
#   2. 检查工作区状态
#   3. 获取提交信息
#   4. 确认提交
#   5. 执行提交
main() {
    # 显示脚本标题
    print_info "=== Git提交助手 ==="
    
    # 第一步：检查Git环境
    # 检查Git是否安装
    check_git_available
    # 检查是否在Git仓库中
    check_in_git_repo
    
    # 第二步：检查工作区状态
    # 检查是否有未暂存的变更
    check_workspace_status
    
    # 第三步：获取提交信息
    # 通过交互式提示获取用户输入的提交信息
    local commit_msg=$(get_commit_message)
    
    # 第四步：确认提交
    # 提示用户确认是否要执行提交
    read -p "确认提交? (y/n): " confirm
    if [[ $confirm != [yY] ]]; then  # 如果用户输入的不是y或Y
        print_info "提交已取消"
        exit 0  # 退出脚本，返回状态码0表示正常退出
    fi
    
    # 第五步：执行提交
    # 执行git add和git commit命令
    perform_git_commit "$commit_msg"
}

# ============================================================================
# 脚本执行
# ============================================================================

# 执行主函数
# 说明：调用main函数启动脚本执行流程
main
