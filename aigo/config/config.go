package config

type GenConfig struct {
	Output        Output           `json:"output"`
	Drivers       GenConfigDrivers `json:"drivers"`
	Tables        []string         `json:"tables"`
	TemplatePaths *TemplatePaths   `json:"template_paths"`
}

type GenConfigDrivers struct {
	Mysql *DBAuth `json:"mysql"`
}

type DBAuth struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	IP       string `json:"ip" validate:"required"`
	Port     string `json:"port" validate:"required"`
	Database string `json:"database" validate:"required"`
}

type Output struct {
	Dir         string `json:"dir"`
	ProjectName string `json:"project_name"`
}

type TemplatePaths struct {
	Basic            string `json:"basic"`
	InitAPIRouter    string `json:"init_api_router"`
	AddV1APIRouter   string `json:"add_v1_api_router"`
	APIParameter     string `json:"api_parameter"`
	APIRouters       string `json:"api_routers"`
	InternalService  string `json:"internal_service"`
	InternalDatabase string `json:"internal_database"`
}
