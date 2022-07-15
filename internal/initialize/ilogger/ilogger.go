package ilogger

import (
	"GhortLinks/internal/initialize/icommon"
	"GhortLinks/internal/initialize/istruct"
	"GhortLinks/utils/istring"
	"GhortLinks/utils/izap/color"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var lg *zap.Logger

// InitializeLogger 初始化Logger - 等待引入viper配置
func InitializeLogger(config *istruct.IStruct) (err error) {
	recordPath := config.JournalDebugPath
	authWriteSyncer := getLogWriter(
		recordPath,
		config.JournalMaxIoSize,
		config.JournalMaxBackups,
		config.JournalEachMaxAge,
	)
	encoder := getProdEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText(istring.String2Byte(config.JournalRecordLevel))
	if err != nil {
		return
	}

	// 根据环境参数处理日志输出位置
	var core zapcore.Core
	if config.RuntimeMode == icommon.RUNTIME_MODE_PRODUCT || config.JournalPrintMode == icommon.JOURNAL_PRINT_MODE_FILE {
		// 生产模式或指定时,将日志输出到日志文件中
		core = zapcore.NewCore(encoder, authWriteSyncer, l)
	} else {
		// 开发|测试 | 未指定模式时,将日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(getDevEncoder())
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	}

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	return
}

// getProdEncoder 生产环境日志输出配置
func getProdEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = iEncodeTime
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = iEncodeLevel
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = iEncodeCaller
	encoderConfig.ConsoleSeparator = "  "
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getDevEncoder 开发环境日志输出配置
func getDevEncoder() zapcore.EncoderConfig {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = iEncodeTime
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = iEncodeLevel
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = iEncodeCaller
	encoderConfig.ConsoleSeparator = "  "
	return encoderConfig
}

// getLogWriter 使用lumberjack对日志进行切割和保存
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// iEncodeLevel 自定义日志等级格式显示
func iEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(color.PrintColorLevelForZap(level))
}

// iEncodeTime 自定义时间格式显示
func iEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[M2TB] " + t.Format(icommon.CURRENT_TIME_FORMAT) + "  |")
}

// iEncodeCaller 自定义行号显示
func iEncodeCaller(_ zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("|")
}
