#!/bin/bash
set -ex

# *************************************************************
# Description: K8s安装部署工具Helm的定制仓库更新所需脚本，
#              1. 安装helm，使用其相关的repo操作工具
#              2. git pull，把./src下的最新的chart打包，copy到./release
#              2. 为./release发布目录生成索引index.yaml
#
# Author:      Hchnr
# Date:        2018-12-29
#
# *************************************************************

# *******************************************
# RESOLVE chart to be updated
base_dir=$(cd `dirname $0`; pwd)
mkdir -p release

if [ $# -ne 1 ]
then
    echo "Example: ./update.sh /update"
    exit
fi

chart_name=all

# *******************************************
# PULL charts source code
cd $base_dir
repo_dir=$base_dir/charts
if [ -d $repo_dir ]; then
    cd $repo_dir
    old_commit=`git log | sed -n '1p'| awk '{print $2}'`
    git pull
    new_commit=`git log | sed -n '1p'| awk '{print $2}'`
else
    git clone git@github.com:jiweil/charts.git
    chart_name=init
fi

# *******************************************
# [SOURCE chart] => [TGZ chart]
cd $repo_dir/src
if [ "all" == $chart_name ]; then
    chart_dirs=`ls`
    chart_dirs=($chart_dirs)
    change_dirs=`git diff --name-only $old_commit $new_commit`
    change_dirs=($change_dirs)
    i=0
    for change_dir in ${change_dirs[@]}; do
        change_dir=${change_dir#src/}
        change_dir=${change_dir%%/*}
        change_dirs[$i]=$change_dir
        let i=i+1
    done
    change_dirs=($(awk -vRS=' ' '!a[$1]++' <<< ${change_dirs[@]}))
    echo "change_dirs: " ${#change_dirs[@]} ${change_dirs[@]}
    for change_dir in ${change_dirs[@]}; do
        if echo ${chart_dirs[@]} | grep $change_dir &>/dev/null; then
            helm package $change_dir || true
        fi
    done
elif [ "init" == $chart_name ]; then
    chart_dirs=`ls`
    chart_dirs=($chart_dirs)
    for chart_dir in ${chart_dirs[@]}; do
        helm package $chart_dir || true
    done
fi

# *******************************************
# MV .tgz files to release
chart_tars=`ls | grep tgz`
chart_tars=($chart_tars)
for chart_tar in ${chart_tars[@]}; do
    mv $chart_tar "$base_dir/release/"
done

# *******************************************
# CREATE index.yaml for chart repo
cd $base_dir/release
helm repo index . --url=https://chart.shannonai.com:4443/release/
mv index.yaml ../release/ || true

