package router

import "os"
import "io/ioutil"
import "encoding/json"

type RunPara struct {
	RunPort string `json:"run_port"`
	RunFile string `json:"file_path"`
}

var Run RunPara

func Init() {
	name := "./config/run.json"
	jsonFile, err := os.Open( name )
	if err != nil {
		panic("打开文件错误，请查看:" + name )
	}
	defer jsonFile.Close()
	data, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal( data, &Run )
}
