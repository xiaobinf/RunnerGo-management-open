package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"kp-management/internal/pkg/conf"
)

func MustInitConf() {
	viper.SetConfigFile("../../configs/dev.yaml")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&conf.Conf); err != nil {
		panic(fmt.Errorf("unmarshal error config file: %w", err))
	}

	fmt.Println("config initialized")
}

func main() {
	MustInitConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", conf.Conf.MySQL.Username,
		conf.Conf.MySQL.Password, conf.Conf.MySQL.Host, conf.Conf.MySQL.Port, conf.Conf.MySQL.DBName, conf.Conf.MySQL.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		//OutPath: "./internal/pkg/dal/query",
		OutPath: "../../internal/pkg/dal/query",
	})

	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("target"),
		g.GenerateModel("operation"),
		g.GenerateModel("user"),
		g.GenerateModel("team"),
		g.GenerateModel("team_env"),
		g.GenerateModel("team_env_service"),
		g.GenerateModel("user_team"),
		g.GenerateModel("setting"),
		g.GenerateModel("sms_log"),
		g.GenerateModel("report_machine"),
		g.GenerateModel("variable"),
		g.GenerateModel("variable_import"),
		g.GenerateModel("team_user_queue"),
		g.GenerateModel("machine"),
		g.GenerateModel("preinstall_conf"),
		g.GenerateModel("auto_plan"),
		g.GenerateModel("auto_plan_email"),
		g.GenerateModel("auto_plan_timed_task_conf"),
		g.GenerateModel("auto_plan_task_conf"),
		g.GenerateModel("auto_plan_task_conf"),
		g.GenerateModel("auto_plan_report"),
		g.GenerateModel("stress_plan"),
		g.GenerateModel("stress_plan_task_conf"),
		g.GenerateModel("stress_plan_report"),
		g.GenerateModel("stress_plan_email"),
		g.GenerateModel("stress_plan_timed_task_conf"),
		g.GenerateModel("target_debug_log"),
		g.GenerateModel("user_collect_info"),
		g.GenerateModel("public_function"),
	)

	g.Execute()
}
