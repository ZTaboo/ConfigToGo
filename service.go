package main

import (
	"bytes"
	"changeme/model"
	"changeme/utils"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
	"syscall"
	"text/template"
)

//go:embed all:template
var jsonTemplate embed.FS

func JsonToStruct(jsonData string, name string) (string, error) {
	var buf bytes.Buffer
	var ress = new([]model.ResData)
	t := template.Must(template.ParseFS(jsonTemplate, "template/*.tmpl"))
	switch name {
	case "json":
		res, err := parseJson(jsonData, name)
		ress = res
		if err != nil {
			return "", err
		}
		break
	case "yaml":
		res, err := parseYaml(jsonData, name)
		ress = res
		if err != nil {
			return "", err
		}
		break
	case "toml":
		res, err := parseToml(jsonData, name)
		ress = res
		if err != nil {
			return "", err
		}
		break
	case "hcl":
		res, err := parseHcl(jsonData, name)
		ress = res
		if err != nil {
			return "", err
		}
		break
	case "env":
		res, err := parseEnv(jsonData, name)
		ress = res
		if err != nil {
			return "", err
		}
		break
	case "ini":
		res, err := parseIni(jsonData, name)
		ress = res
		if err != nil {
			return "", err
		}
		break
	default:
		break
	}

	err := t.ExecuteTemplate(&buf, "language", ress)
	if err != nil {
		return "", err
	}
	temp, err := os.CreateTemp("", "zero")
	if err != nil {
		return "", err
	}
	temp.Write(buf.Bytes())
	_ = temp.Close()
	return printOutStr(temp.Name()), nil
}
func printOutStr(data string) string {
	fmtPath := runtime.GOROOT() + "/bin/gofmt"
	cmd := exec.Command(fmtPath, data)
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error:", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			fmt.Println("delete error:", err)
		}
	}(data)
	//strings.Replace(string(output), "\t", "    ", -1)
	return string(output)
}

func parseJson(str string, name string) (*[]model.ResData, error) {
	var data interface{}
	json.Unmarshal([]byte(str), &data)
	res, err := isMap(data)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	var resData = new([]model.ResData)
	result := rangeMap(res, "ZeroModel", resData, 0, name)
	return result, nil
}

func parseYaml(str string, name string) (*[]model.ResData, error) {
	var data interface{}
	yaml.Unmarshal([]byte(str), &data)
	res, err := isMap(data)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	var resData = new([]model.ResData)
	result := rangeMap(res, "ZeroModel", resData, 0, name)
	return result, nil
}

func parseToml(str string, name string) (*[]model.ResData, error) {
	var data interface{}
	// 解析字符串流
	viper.SetConfigType("toml")
	err := viper.ReadConfig(bytes.NewBuffer([]byte(str)))
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&data)
	if err != nil {
		return nil, err
	}
	res, err := isMap(data)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	var resData = new([]model.ResData)
	result := rangeMap(res, "ZeroModel", resData, 0, name)
	return result, nil
}

func parseHcl(str string, name string) (*[]model.ResData, error) {
	var data interface{}
	// 解析字符串流
	viper.SetConfigType("hcl")
	err := viper.ReadConfig(bytes.NewBuffer([]byte(str)))
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&data)
	if err != nil {
		return nil, err
	}
	res, err := isMap(data)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	var resData = new([]model.ResData)
	result := rangeMap(res, "ZeroModel", resData, 0, name)
	return result, nil
}

func parseEnv(str string, name string) (*[]model.ResData, error) {
	var data interface{}
	// 解析字符串流
	viper.SetConfigType("env")
	err := viper.ReadConfig(bytes.NewBuffer([]byte(str)))
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&data)
	if err != nil {
		return nil, err
	}
	res, err := isMap(data)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	var resData = new([]model.ResData)
	result := rangeMap(res, "ZeroModel", resData, 0, name)
	return result, nil
}

func parseIni(str string, name string) (*[]model.ResData, error) {
	var data interface{}
	// 解析字符串流
	viper.SetConfigType("ini")
	err := viper.ReadConfig(bytes.NewBuffer([]byte(str)))
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&data)
	if err != nil {
		return nil, err
	}
	res, err := isMap(data)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	var resData = new([]model.ResData)
	result := rangeMap(res, "ZeroModel", resData, 0, name)
	return result, nil
}

// 将序列化后的对象解析为数组结构体
func rangeMap(data any, structName string, listData *[]model.ResData, cont int, name string) *[]model.ResData {
	*listData = append(*listData, model.ResData{
		Name: strings.ToTitle(structName[:1]) + structName[1:],
	})
	var paramData []model.Param
	for k, v := range data.(map[string]interface{}) {
		switch v.(type) {
		case map[string]interface{}:
			if reflect.ValueOf(v).Type() == nil {
				break
			} else {
				rangeMap(v, k, listData, len(*listData), name)
				tmpData := utils.Case(k)
				paramData = append(paramData, model.Param{
					Name:    tmpData,
					Type:    tmpData,
					Tag:     strings.ToLower(tmpData[:1]) + tmpData[1:],
					TagType: name,
				})
			}
			break
		case []interface{}:
			switch v.([]interface{})[0].(type) {
			case map[string]interface{}:
				if reflect.ValueOf(v).Type() == nil {
					break
				} else {
					rangeMap(v.([]interface{})[0], k, listData, len(*listData), name)
					tmpData := utils.Case(k)
					paramData = append(paramData, model.Param{
						Name:    tmpData,
						Type:    "[]" + tmpData,
						Tag:     strings.ToLower(tmpData[:1]) + tmpData[1:],
						TagType: name,
					})
				}
				break
			default:
				if reflect.ValueOf(v).Type() == nil {
					break
				} else {
					tmpData := utils.Case(k)
					paramData = append(paramData, model.Param{
						Name:    tmpData,
						Type:    "[]" + utils.GetType(v.([]interface{})[0]),
						Tag:     strings.ToLower(tmpData[:1]) + tmpData[1:],
						TagType: name,
					})
				}
				break

			}

			break
		case []map[string]interface{}:
			fmt.Println("数组对象参数")
			tmpData := utils.Case(k)
			paramData = append(paramData, model.Param{
				Name:    tmpData,
				Type:    "[]" + tmpData,
				Tag:     strings.ToLower(tmpData[:1]) + tmpData[1:],
				TagType: name,
			})
			rangeMap(v.([]map[string]interface{})[0], k, listData, len(*listData), name)
			break
		default:
			if v == nil {
				break
			} else {
				tmpData := utils.Case(k)
				paramData = append(paramData, model.Param{
					Name:    tmpData,
					Type:    utils.GetType(v),
					Tag:     strings.ToLower(tmpData[:1]) + tmpData[1:],
					TagType: name,
				})
			}
			break
		}
	}
	(*listData)[cont].Param = &paramData
	return listData
}

// 判断json格式是否正确,如果是数组对象返回第一个json对象
func isMap(data any) (any, error) {
	switch data.(type) {
	case map[string]interface{}:
		return data, nil
	case []interface{}:
		switch data.([]interface{})[0].(type) {
		case map[string]interface{}:
			fmt.Println("数组对象")
			return data.([]interface{})[0], nil
		default:
			return nil, errors.New("类型错误")
		}
		break
	default:
		return nil, errors.New("类型错误")
	}
	return nil, errors.New("类型错误")
}
