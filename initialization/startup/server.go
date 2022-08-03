package startup

type Server struct {
	Port int `json:"port" yaml:"port"`
	JwtSecret string `json:"jwt_secret" yaml:"jwt_secret"`
	ImgSizeLimit string `json:"max_img_size" yaml:"max_img_size"`
	VideoSizeLimit string `json:"max_video_size" yaml:"max_video_size"`
	CronTaskRequired bool `json:"cron_tasks_required" yaml:"cron_tasks_required"`
	Migrate bool `json:"migration_required" yaml:"migration_required"`
}

var ServerGlobalSetting *Server