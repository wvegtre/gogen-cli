#!/bin/bash

# 指定默认目录和文件/目录列表
default_dir="/Users/hb.li/Documents/royce/star-royce/gogen-cli/codegen"
target_dir="/Users/hb.li/Documents/royce/wvegtre/easymoney"
items=("demo")
module="easymoney"

# 判断默认目录是否存在，如果不存在就报错
if [ ! -d "$default_dir" ]; then
    echo "Error: $default_dir : No such file or directory"
    exit 1
fi

# 判断指定目录是否存在，如果不存在就创建它
if [ ! -d "$target_dir" ]; then
  mkdir -p "$target_dir"
fi

# 遍历文件/目录列表，将每个文件或目录复制到指定目录下
for item in "${items[@]}"; do
  # 判断默认目录是否存在，如果不存在就报错
  if [ ! -d "$default_dir/$item" ]; then
      echo "Error: $default_dir/$item : No such file or directory"
      exit 1
  fi
  cp -r "$default_dir/$item" "$target_dir/"
done

# 对复制后的文件夹内的所有文件执行文本替换操作
find "$target_dir" -type f -exec sed -i "s#demo_moudle/#$module/#g" {} +

echo "Copied from demo and replaced module successfully!"

./run.sh
echo "Customize code auto gen successfully!"

(
  project_path="$target_dir/$module/cmd/app"
  cd "$project_path" || exit

  echo "Run go mod for new project in""$project_path"
  go mod
)

# shellcheck disable=SC2027
echo "$module"" project init successfully! You can view it in ""$target_dir"
