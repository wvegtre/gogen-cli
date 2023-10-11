package convertor

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
)

func SnakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) == 1 {
			parts[i] = strings.ToUpper(part[:1])
		} else {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, "")
}

func LowFirstLetter(s string) string {
	if len(s) == 1 {
		return strings.ToLower(s)
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func ParseTPLAndWrite(projectName, tplFile, writeFile string, data interface{}) error {
	// 解析模板文件
	tpl, err := template.ParseFiles(tplFile)
	if err != nil {
		return errors.WithStack(err)
	}
	// 创建一个新文件
	file, err := os.OpenFile(writeFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.WithStack(err)
	}
	defer file.Close()
	// 将数据写入文件
	err = tpl.Execute(file, data)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func DeleteFileWithSuffixName(file, suffix string) error {
	err := filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == suffix {
			fmt.Printf("Deleting %s\n", path)
			if err := os.Remove(path); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func ReplaceFileDir(dir, projectName string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, f := range files {
		if f.IsDir() {
			err = ReplaceFileDir(filepath.Join(dir, f.Name()), projectName)
			if err != nil {
				return errors.WithStack(err)
			}
			continue
		}
		if !strings.HasSuffix(f.Name(), ".go") {
			log.Println(aurora.Yellow("ReplaceFileDir skip file " + dir + "/" + f.Name()))
			continue
		}
		filePath := filepath.Join(dir, f.Name())
		replaceMap := map[string]string{
			"&#34;":         "\"",
			"&lt;":          "<",
			"gen-templates": projectName,
		}
		for k, v := range replaceMap {
			err = replaceFileContent(filePath, k, v)
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

func replaceFileContent(writeFile string, oldKeyword, newKeyword string) error {
	file, err := os.OpenFile(writeFile, os.O_RDWR, 0644)
	if err != nil {
		return errors.WithStack(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, oldKeyword, newKeyword)
	}
	err = file.Truncate(0)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return errors.WithStack(err)
	}
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err = fmt.Fprintln(writer, line)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	err = writer.Flush()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func fieldTypeMapping(fieldType string) string {
	m := map[string]string{
		"datetime":  "time.Time",
		"timestamp": "int64",
		"text":      "string",
		"varchar":   "string",
		"decimal":   "float64",
		"long":      "int64",
		"int":       "int",
	}
	if v, ok := m[fieldType]; ok {
		return v
	}
	return fieldType
}
