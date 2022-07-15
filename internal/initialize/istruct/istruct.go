package istruct

import (
	"GhortLinks/internal/initialize/icommon"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type IStruct struct {
	*Base    `mapstructure:"base"`
	*Program `mapstructure:"program"`
	*Http    `mapstructure:"serve"`
	*Journal `mapstructure:"journal"`
	*Redis   `mapstructure:"redis"`
	Database []interface{} `mapstructure:"database"`
}

// Base 基础配置信息
type Base struct {
	RuntimeMode  string `mapstructure:"runtime_mode"`
	TimeLocation string `mapstructure:"time_location"`
	MachineID    int64  `mapstructure:"machine_id"`
	SnowRuntime  string `mapstructure:"snow_runtime"`
}

// Program 项目配置信息
type Program struct {
	ProgramName    string `mapstructure:"program_name"`
	ProgramVersion string `mapstructure:"program_version"`
}

// Http http配置信息
type Http struct {
	HttpListenPort     string `mapstructure:"http_listen_port"`
	HttpReadsTimeout   int    `mapstructure:"http_reads_timeout"`
	HttpWriteTimeout   int    `mapstructure:"http_write_timeout"`
	HttpMaxHeaderBytes int    `mapstructure:"http_max_header_bytes"`
}

// Journal 项目日志配置信息
type Journal struct {
	JournalPrintMode   string `mapstructure:"journal_print_mode"`
	JournalRecordLevel string `mapstructure:"journal_record_level"`
	JournalDebugPath   string `mapstructure:"journal_debug_path"`
	JournalMaxIoSize   int    `mapstructure:"journal_max_io_size"`
	JournalEachMaxAge  int    `mapstructure:"journal_each_max_age"`
	JournalMaxBackups  int    `mapstructure:"journal_max_backups"`
}

// Database 数据库配置信息
type Database struct {
	DbDriverName      string `mapstructure:"db_driver_name"`
	DbSourceStr       string `mapstructure:"db_source_str"`
	DbMaxOpen         int    `mapstructure:"db_max_open"`
	DbMaxIdle         int    `mapstructure:"db_max_idle"`
	DbMaxConnLifetime int    `mapstructure:"db_max_conn_lifetime"`
}

// Redis 缓存配置信息
type Redis struct {
	RedisHost         string `mapstructure:"redis_host"`
	RedisConnTimeout  int    `mapstructure:"redis_conn_timeout"`
	RedisPassword     string `mapstructure:"redis_password"`
	RedisDbUse        int    `mapstructure:"redis_db_use"`
	RedisReadTimeout  int    `mapstructure:"redis_read_timeout"`
	RedisWriteTimeout int    `mapstructure:"redis_write_timeout"`
	RedisMaxOpen      int    `mapstructure:"redis_max_open"`
	RedisMaxIdle      int    `mapstructure:"redis_max_idle"`
}

// Conf 全局变量,用于保存所有的配置信息
var Conf = new(IStruct)

// InitializeConfig 初始化配置信息
func InitializeConfig() (err error) {
	viper.SetConfigName(fmt.Sprintf("config_%s", icommon.CURRENT_RUNTIME)) // 指定配置文件名称,不携带后缀名
	viper.SetConfigType("yaml")                                            // 指定配置文件类型
	viper.AddConfigPath(icommon.CONFIG_PATH)                               // 指定查找配置文件的路径 (相对路径)
	// 读取配置文件信息
	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err: %v\n", err)
		return err
	}
	// 将读取到的配置文件内容反序列化到 Conf 变量中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal() failed, err: %v\n", err)
		return err
	}
	// 监听配置文件是否修改
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件内容进行了修改 ...")
		// 讲字符串解码到相应的数据结构
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal() failed, err: %v\n", err)
		}
	})
	return
}
