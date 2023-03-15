package dbconverter

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"echo-shopping/tools/array"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

//map for converting mysql type to golang types
var fieldTypeMapping = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int",
	"smallint":           "int",
	"mediumint":          "int",
	"bigint":             "int",
	"int unsigned":       "int",
	"integer unsigned":   "int",
	"tinyint unsigned":   "int",
	"smallint unsigned":  "int",
	"mediumint unsigned": "int",
	"bigint unsigned":    "int",
	"bit":                "int",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "string", // time.Time
	"datetime":           "string", // time.Time
	"timestamp":          "string", // time.Time
	"time":               "string", // time.Time
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

type MySQLConverter struct {
	dbClient *sql.DB
	config   *mySQLConverterConfig
}

func NewMySQLConverter(dc DBConfig, ops ...MySQLConfigOption) *MySQLConverter {
	defaultConfig := newDefaultConfig()
	defaultConfig.DBConfig = dc
	defaultConfig.DBConfig.charset = "utf8"
	// run option for set custom value
	for _, op := range ops {
		op(defaultConfig)
	}
	return &MySQLConverter{
		config: defaultConfig,
	}
}

func newDefaultConfig() *mySQLConverterConfig {
	c := &mySQLConverterConfig{}
	c.fieldConfig = fieldConfig{
		EnableJsonTag: true,
	}
	c.fileConfig = fileConfig{
		AllInOneFile:        true,
		SaveDir:             "./",
		SaveFileDefaultName: "model.go",
	}
	return c
}

type mySQLConverterConfig struct {
	// you can set package name, if null, set by fileConfig.SaveFileDefaultName
	//packageName string
	// if nil, gen for all tables
	tables []string
	// some configs for struct filed
	fieldConfig
	// some configs for write file
	fileConfig
	// need auth config for dbClient connect
	DBConfig
}

type fieldConfig struct {
	EnableJsonTag bool // if add json tag on struct field, default true.
}

type fileConfig struct {
	AllInOneFile        bool
	SaveDir             string
	SaveFilePrefix      string
	SaveFileDefaultName string // default model.go, but invalid when AllInOneFile=false
}

type DBConfig struct {
	// {user}:{pwd}@tcp({ip}:{port})/{dbClient}?charset=utf8"
	UserName string `validate:"required"`
	Password string `validate:"required"`
	IP       string `validate:"required"`
	Port     string `validate:"required"`
	Database string `validate:"required"`
	charset  string // default utf8
}

type StructContentDetail struct {
	TableName  string
	StructName string
	Content    string
}

func (c *MySQLConverter) Run() (map[string][]StructContentDetail, error) {
	err := validator.New().Struct(c.config)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 链接mysql, 获取db对象
	err = c.dialMysql()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	log.Println("get mysql connect succeed.")
	// 获取表和字段的shcema
	tableColumns, err := c.getColumns()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 组装struct
	//structContentMap := make([]StructContentDetail, 0)
	groupMap := make(map[string][]StructContentDetail, 0)
	for tableRealName, item := range tableColumns {
		groupName, ok := _tableGroups[tableRealName]
		if !ok {
			groupName = "undefined"
		}
		structName := c.genStructNameByTableName(tableRealName)
		// 按分组存放好，每个组在一个 package 下
		groupMap[groupName] = append(groupMap[groupName], StructContentDetail{
			TableName:  tableRealName,
			StructName: structName,
			Content:    c.convertStructContent(tableRealName, structName, item),
		})
	}
	// 写入
	for groupName, contents := range groupMap {
		err = c.output(groupName, "", contents)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return groupMap, nil
}

func (c *MySQLConverter) dialMysql() error {
	if c.dbClient != nil {
		return nil
	}
	// {user}:{pwd}@tcp({ip}:{port})/{dbClient}?charset=utf8
	dsn := `%s:%s@tcp(%s:%s)/%s?charset=%s`
	dsn = fmt.Sprintf(dsn, c.config.UserName, c.config.Password, c.config.IP, c.config.Port, c.config.Database, c.config.charset)
	client, err := sql.Open("mysql", dsn)
	if err != nil {
		return errors.WithStack(err)
	}
	err = client.Ping()
	if err != nil {
		return errors.WithStack(err)
	}
	c.dbClient = client
	return nil
}

func (c *MySQLConverter) output(packageName string, fileName string, structContentMap []StructContentDetail) error {
	// 不分多个文件存储，直接 save 后 return
	basePath := c.config.SaveDir + packageName
	err := os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		return errors.WithStack(err)
	}
	if c.config.AllInOneFile {
		// 是否指定保存路径
		savePath := c.getSavePath(basePath, c.config.SaveFileDefaultName)
		var structContent string
		for _, v := range structContentMap {
			structContent += v.Content
		}
		return c.write2File(packageName, structContent, savePath)
	}
	// 分多文件存储
	for _, v := range structContentMap {
		// 格式化路径
		savePath := c.getSavePath(basePath, v.TableName)
		if err := c.write2File(packageName, v.Content, savePath); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (c *MySQLConverter) getSavePath(savePath, fileName string) string {
	if c.config.SaveFilePrefix != "" {
		savePath += "/" + c.config.SaveFilePrefix
	}
	savePath += "/" + fileName + ".go"
	return savePath
}

func (c *MySQLConverter) write2File(packageName, structContent, savePath string) error {
	// 如果有引入 time.Time, 则需要引入 time 包
	var moreContent string
	moreContent += "package " + packageName + "\n"
	if strings.Contains(structContent, "time.Time") {
		moreContent += "import \"time\"\n\n"
	}
	filePath := fmt.Sprintf("%s", savePath)
	f, err := os.Create(filePath)
	if err != nil {
		log.Println("Can not write file")
		return errors.WithStack(err)
	}
	defer f.Close()

	_, err = f.WriteString(moreContent + structContent)
	if err != nil {
		return errors.WithStack(err)
	}

	log.Println("write to field succeed. path: ", filePath)
	//
	//cmd := exec.Command("gofmt", "-w", filePath)
	//err = cmd.Run()
	//if err != nil {
	//	return errors.WithStack(err)
	//}
	return nil
}

func (c *MySQLConverter) convertStructContent(tableRealName, structName string, item []column) string {
	depth := 1
	var structContent string
	structContent += "type " + structName + " struct {\n"
	structContent += "gorm.Model\n"
	for _, v := range item {
		if c.isGormCommonFields(v.Tag) {
			continue
		}
		// 字段注释
		var comment string
		if v.ColumnComment != "" {
			comment = fmt.Sprintf(" // %s", v.ColumnComment)
		}
		structContent += fmt.Sprintf("%s%s %s %s%s\n",
			tab(depth), v.ColumnName, v.Type, v.Tag, comment)
	}
	structContent += tab(depth-1) + "}\n\n"
	structContent += c.addTableNameFunc(tableRealName, structName) + "\n\n"
	log.Println("convert to struct succeed. table: ", tableRealName, ", struct: ", structName)
	return structContent
}

func (c *MySQLConverter) isGormCommonFields(tag string) bool {
	set := array.NewStringSet([]string{
		"Id", "CreatedAt", "UpdatedAt", "DeletedAt",
	})
	return set.Contain(tag)
}

func (c *MySQLConverter) addTableNameFunc(tableName, structName string) string {
	temp := `
	func (%s) TableName() string {
		return "%s"
	}`
	return fmt.Sprintf(temp, structName, tableName)
}

func (c *MySQLConverter) genStructNameByTableName(tableName string) string {
	// user_auth -> UserAuth
	arr := strings.Split(tableName, "_")
	var structName string
	for _, v := range arr {
		vv := v
		if len(vv) == 1 {
			structName += strings.ToUpper(vv)
		} else {
			structName += strings.ToUpper(vv[:1]) + vv[1:]
		}
	}
	return structName + "Model"
}

type column struct {
	ColumnName    string
	Type          string
	Nullable      string
	TableName     string
	ColumnComment string
	Tag           string
}

// Function for fetching schema definition of passed table
func (c *MySQLConverter) getColumns(table ...string) (tableColumns map[string][]column, err error) {
	tableColumns = make(map[string][]column)
	// sql
	var sqlStr = `SELECT COLUMN_NAME,DATA_TYPE,IS_NULLABLE,TABLE_NAME,COLUMN_COMMENT
		FROM information_schema.COLUMNS
		WHERE table_schema = DATABASE()`
	if len(c.config.tables) > 0 {
		var whereTable string
		// AND TABLE_NAME in('t1','t2','t3')
		for _, v := range c.config.tables {
			whereTable += "'" + v + "'" + ","
		}
		sqlStr += fmt.Sprintf(" AND TABLE_NAME in ('%s')", whereTable[:len(whereTable)-1])
	}
	// sql排序
	sqlStr += " order by TABLE_NAME asc, ORDINAL_POSITION asc"

	rows, err := c.dbClient.Query(sqlStr)
	if err != nil {
		log.Println("Error reading table information: ", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		col := column{}
		err = rows.Scan(&col.ColumnName, &col.Type, &col.Nullable, &col.TableName, &col.ColumnComment)

		if err != nil {
			log.Println(err.Error())
			return
		}
		// ColumnName 转换成小写作为 tag
		tag := strings.ToLower(col.ColumnName)
		// add gorm tag
		col.Tag = "`" + fmt.Sprintf("%s:\"%s\"", "gorm", tag)
		if c.config.EnableJsonTag {
			col.Tag += fmt.Sprintf(" json:\"%s\"", tag)
		}
		col.Tag += "`"
		// 处理成首字母大写，后边struct 生成时需要
		col.ColumnName = c.camelCase(col.ColumnName)
		col.Type = fieldTypeMapping[col.Type]
		tableColumns[col.TableName] = append(tableColumns[col.TableName], col)
	}
	return
}

func (c *MySQLConverter) camelCase(str string) string {
	var text string
	//for _, p := range strings.Split(name, "_") {
	for _, p := range strings.Split(str, "_") {
		// 字段首字母大写的同时, 是否要把其他字母转换为小写
		switch len(p) {
		case 0:
		case 1:
			text += strings.ToUpper(p[0:1])
		default:
			text += strings.ToUpper(p[0:1]) + p[1:]
		}
	}
	return text
}

func tab(depth int) string {
	return strings.Repeat("\t", depth)
}
