//
// Created by ChenYuZhao on 2020/7/1.
//
package bpf

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Define global slices for storing all plugins
var PluginServices []*PluginService

func init() {
	// Read the name of the plug-in in the directory
	files, _ := ioutil.ReadDir("./plugins")
	fmt.Println(files)
	for _, f := range files {
		// Register plugins
		file,_ := os.Open("./plugins/"+f.Name())
		RegisterPluginService(f.Name(),file,"")
	}

	//OutputPluginService()
}

type PluginService struct {
	F *os.File
	Name string
	Info string
}

// Register plugins
func RegisterPluginService(name string, f *os.File, info string) {
	PluginServices = append(PluginServices, &PluginService{
		F : f,
		Name : name,
		Info : info,
	})
}

// Outputs the names of all plugins
func OutputPluginService() {
	for _,plugin := range PluginServices {
		fmt.Println(plugin.Name)
		fmt.Println(plugin.Info)
		fmt.Println(plugin.F)
	}
}

// Save this plugin to database
func (p *PluginService) Save2DB() {

}


/*
type DirectorPy struct {
	//生成python文件的接口
	pyBuilder PyBuilder
}

type DirectorC struct {
	//生成C文件的接口
	cBuilder CBuilder
}

func NewDirectorPy(b PyBuilder) DirectorPy {
	return DirectorPy{
		pyBuilder: b,
	}
}

func NewDirectorC(c CBuilder) DirectorC {
	return DirectorC{
		cBuilder: c,
	}
}

//生成python文件的步骤，接口和实现分离
func (d *DirectorPy) ConstructPy() {
	d.pyBuilder.AddCommonPyFileFront()
	d.pyBuilder.AddPrivatePyFile()
	d.pyBuilder.AddCommonPyFileEnd()
}

//生成C文件的步骤，接口和实现分离
func (d *DirectorC) ConstructC() {
	d.cBuilder.AddCommonCFileFront()
	d.cBuilder.AddPrivateCFile()
	d.cBuilder.AddCommonCFileEnd()
}
*/