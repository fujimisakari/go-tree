package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		OutputHelp()
		return
	}

	arg := args[1]
	switch arg {
	case "sample":
		OutputSampleTree()
		return
	case "help":
		OutputHelp()
		return
	default:
		yamlPath := arg
		if _, err := os.Stat(yamlPath); err != nil {
			fmt.Println("YamlFile not found:", yamlPath)
			return
		}
		OutputGoTree(yamlPath)
	}
}

func OutputHelp() {
	helpMsg := `Usage:

    go-tree [yaml-path|sample]


The commands are:

    sample    output sample tree`

	fmt.Println(helpMsg)
}

func OutputSampleTree() {
	sampleTree := `
root-dir: "Sample"
go-tree:
  - fizz
  - buzz
  - dir1:
    - comp1-1
    - comp1-2
    - comp1-3
    - comp1-4
  - dir2:
    - comp2-1
    - comp2-2
    - comp2-3
    - comp2-4
    - dir3:
      - comp3-1
      - comp3-2
      - dir4:
        - comp4-1
        - comp4-2
        - comp4-3
      - comp3-3
      - comp3-4
    - comp2-5
    - comp2-6
  - dir5:
    - comp4-1
    - comp4-2
    - comp4-3
    - comp4-4
  - foo
  - bar
`
	fmt.Println(sampleTree)
}

func OutputGoTree(yamlPath string) {
	treeData := getTreeDataBy(yamlPath)
	rootDir := "."
	if v, ok := treeData["root-dir"]; ok {
		rootDir = v.(string)
	}
	isLastLoopFlags := []bool{}
	output(rootDir, isLastLoopFlags)
	goTreeRecursion(treeData["go-tree"], isLastLoopFlags)
}

func getTreeDataBy(yamlPath string) map[interface{}]interface{} {
	buf, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		panic(err)
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}

	return m
}

func goTreeRecursion(treeData interface{}, isLastLoopFlags []bool) {

	if isKind(treeData, reflect.String) {
		output(treeData.(string), isLastLoopFlags)
	}

	if isKind(treeData, reflect.Slice) {
		dataList := treeData.([]interface{})
		maxCount := len(dataList)
		for i := 0; i < maxCount; i++ {
			isLastLoop := false
			if i == (maxCount - 1) {
				isLastLoop = true
			}
			flags := append(isLastLoopFlags, []bool{isLastLoop}...)
			goTreeRecursion(dataList[i], flags)
		}
	}

	if isKind(treeData, reflect.Map) {
		for directory, dataList := range treeData.(map[interface{}]interface{}) {
			goTreeRecursion(directory.(string), isLastLoopFlags)
			goTreeRecursion(dataList, isLastLoopFlags)
		}
	}
}

func isKind(what interface{}, kind reflect.Kind) bool {
	return reflect.ValueOf(what).Kind() == kind
}

func output(data string, isLastLoopFlags []bool) {
	pathLine := ""
	maxCount := len(isLastLoopFlags)
	for i := 0; i < maxCount; i++ {
		isLast := isLastLoopFlags[i]
		if i == (maxCount - 1) {
			if isLast {
				pathLine += "└── "
			} else {
				pathLine += "├── "
			}
		} else {
			if isLast {
				pathLine += "    "
			} else {
				pathLine += "│   "
			}
		}
	}
	pathLine += data
	fmt.Println(pathLine)
}
