# init.sh 脚本执行前置配置
## init.sh 脚本参数调整
1. module 调整为预期的项目名
2. default_dir 调整为预期的 aigo 项目本地目录
3. output_dir 调整为预期的项目输出目录

## 修改 `aigo/config/config.json` 配置文件
1. 修改 `aigo/config/config.json` 配置文件中的 `output` 配置下的 `project_name` 与 `dir`
2. 修改 `aigo/config/config.json` 配置文件中的 `drivers` 配置，添加需要的驱动以及对应的鉴权信息 
3. 修改 `aigo/config/config.json` 配置文件中的 `tables` 配置，添加需要生成的表名，留空则默认生成指定数据库下的所有表

# init.sh 脚本执行流程简介
1. 将 `aigo/templates` 目录下的文件复制到指定输出目录下
2. 编译 `main.go` 文件，生成可执行文件 
3. 运行 main.go 可行性文件，生成文件 
4. 执行 `go mod tidy` 命令，下载依赖包 
5. 执行文本替换指令，将指定输出目录下所有 `gen-templates` 替换成指定的项目名称 
6. 删除指定输出目录下所有 .tpl 后缀的文件