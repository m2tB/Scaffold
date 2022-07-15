package icommon

import (
	"database/sql"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"net"
	"time"
)

var (
	SYSTEM_MAP  = []string{"console", "service"}
	RUNTIME_MAP = []string{"testing", "develop", "product"}

	CURRENT_SYSTEM_IS_SERVICE = false
	CURRENT_RUNTIME           = ""

	CURRENT_LOCAL_IP      = net.ParseIP("127.0.0.1")
	CURRENT_TIME_LOCATION *time.Location

	CURRENT_DEFAULT_DB   *sql.DB
	CURRENT_DEFAULT_GORM *gorm.DB

	CURRENT_SNOW_GENERAL *snowflake.Node

	DEFAULT_LOCAL = "zh"

	CURRENT_REDIS_SESSION_SECRET = "secret"
)

const (
	// CONFIG_PATH - 配置文件路径
	CONFIG_PATH = "../conf/"

	RUNTIME_MODE_TESTING = "testing"
	RUNTIME_MODE_DEVELOP = "develop"
	RUNTIME_MODE_PRODUCT = "product"

	JOURNAL_PRINT_MODE_FILE    = "file"
	JOURNAL_PRINT_MODE_CONSOLE = "console"

	CURRENT_TIME_FORMAT = "2006-01-02 15:04:05"

	CURRENT_TRACE_ID = "traceId"

	CURRENT_TRANSLATOR = "translator"
	CURRENT_VALIDATOR  = "validator"

	DEFAULT_TOKEN_OR_JWT_EXPIRES = 60 * 60
)
