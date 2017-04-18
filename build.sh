#!/bin/bash

# --------------------------------------------------
build_log="buildLog"
current_path=$(cd `dirname $0`; pwd)
shell_path=$current_path/tools/prebuild/shell
# --------------------------------------------------

# --------------------------------------------------
function echo_green
{
    local content=$@
    echo -e "\033[1;32m${content}\033[0m"
    return 0
}

function usage
{
    echo_green "
Usage:
    $0 [\$1]
Options:
    h|-h|help|-help   usage help
    init              初始化环境，包括环境变量设置、相关工具的安装。
                      如果本机没有go等开发环境建议首先选择此选项；否则不需要。
                      如果使用init参数，必须按照如下方法使用：
                      ./build.sh init && . ~/.bashrc
    buildd            编译整个项目(debug)。
    buildr            编译整个项目(release)。
    test              进行全部单元测试。
    package           进行项目打包。
                      目录结构为：
                      xxxx.tar.gz ─┬conf
                                   ├bin
                                   ├data
                                   └log
    "
    return 0
}
# --------------------------------------------------

# --------------------------------------------------
if [ $# -lt 1 ]
then
    usage
    exit -1
fi

mkdir -p $build_log
chmod +x $shell_path/*.sh 2>/dev/null

case $1 in
    h|help|-h|-help)
        usage
    ;;
    init)
        $shell_path/1_env_init.sh | tee $build_log/1_env_init.log
    ;;
    buildd)
        $shell_path/2_compile.sh debug | tee $build_log/2_compile.log
    ;;
	buildr)
        $shell_path/2_compile.sh release | tee $build_log/2_compile.log
    ;;
    test)
        $shell_path/3_unit_test.sh | tee $build_log/3_unit_test.log
    ;;
    package)
        $shell_path/4_package.sh | tee $build_log/4_package.log
    ;;
    *)
        usage
    ;;
esac

exit 0
# --------------------------------------------------
