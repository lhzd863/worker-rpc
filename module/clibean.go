package module

type MetaCliConf struct {
	Version string `yaml:"version"`
	Name    string `yaml:"name"`
	Ip      string `yaml:"ip"`
	Port    string `yaml:"port"`
	JobConf string `yaml:"jobconf"`
}

type MetaJobConf struct {
	Version string                      `yaml:"version"`
	Name    string                      `yaml:"name"`
	Attr    map[string]*MetaJobAttrConf `yaml:"attr"`
}

type MetaJobAttrConf struct {
	Sys       string        `yaml:"sys"`
	Job       string        `yaml:"job"`
	Context   string        `yaml:"context"`
	Cmd       []interface{} `yaml:"cmd"`
	Parameter []interface{} `yaml:"parameter"`
	Timeout   int64         `yaml:"timeout"`
	Retry     int8          `yaml:"retry"`
	Ip        string        `yaml:"ip"`
	Port      string        `yaml:"port"`
        Describle string        `yaml:"describle"`
        Enable    string        `yaml:"enable"`
}


