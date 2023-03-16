package main

type GenConfig struct {
	Output  Output  `json:"output"`
	Drivers Drivers `json:"drivers"`
}

type Drivers struct {
	Mysql Mysql `json:"mysql"`
}

type Mysql struct {
	UserName string `json:"user_name"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
	DB       string `json:"db"`
	Tables   string `json:"tables"`
	Charset  string `json:"charset"`
}

type Output struct {
	Dir            string `json:"dir"`
	GroupInOneFile bool   `json:"group_in_one_file"`
}
