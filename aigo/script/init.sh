#!/bin/bash

# 指定默认目录和文件/目录列表
module="chocolate"
default_dir="/Users/hb.li/Documents/royce/star-royce/gogen-cli/aigo"
target_dir="/Users/hb.li/Documents/royce/wvegtre/$module"

# 判断默认目录是否存在，如果不存在就报错
if [ ! -d "$default_dir" ]; then
    echo "Error: $default_dir : No such file or directory"
    exit 1
fi

# 判断指定目录是否存在，如果不存在就创建它
if [ ! -d "$target_dir" ]; then
  mkdir -p "$target_dir" && echo "Created $target_dir"
fi

# 初始化基础项目目录结构
cp -r "$default_dir"/templates/* "$target_dir/" && echo "Init $target_dir"

echo "Generating code files!"
./run.sh "$default_dir" "$target_dir" "$module"

echo "$module"" project init successfully! You can view it in ""$target_dir"