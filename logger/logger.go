package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

type LogConfig struct {
	FileName   string `json:"fileName"`   // FileName 日志文件位置
	MaxSize    int    `json:"maxSize"`    // MaxSize 进行切割之前，日志文件的最大大小(MB为单位)，默认为100MB
	MaxAge     int    `json:"maxAge"`     // MaxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数
	MaxBackups int    `json:"maxBackups"` // MaxBackups 是要保留的旧日志文件的最大数量。默认是保留所有旧的日志文件
}

// 负责设置 encoding 的日志格式
func getEncoder() zapcore.Encoder {
	// 获取一个指定的的EncoderConfig，进行自定义
	encodeConfig := zap.NewProductionEncoderConfig()
	// 设置每个日志条目使用的键。如果有任何键为空，则省略该条目的部分。
	// 序列化时间。eg: 2022-09-01T19:11:35.921+0800
	encodeConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	// "time":"2022-09-01T19:11:35.921+0800"
	encodeConfig.TimeKey = "time"
	// 将Level序列化为全大写字符串。例如，将info level序列化为INFO。
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 以 package/file:行 的格式 序列化调用程序，从完整路径中删除除最后一个目录外的所有目录。
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

// 负责日志写入的位置
func getLogWriter(filename string, maxsize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  // 文件位置
		MaxSize:    maxsize,   // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     maxAge,    // 保留旧文件的最大天数
		MaxBackups: maxBackup, // 保留旧文件的最大个数
		Compress:   false,     // 是否压缩/归档旧文件
	}

	// 每天分割一次日志
	go func() {
		for {
			// 使用北京时区
			beijingLocation, err := time.LoadLocation("Asia/Shanghai")
			if err != nil {
				return
			}
			// 获取当前时间（使用北京时区）
			now := time.Now().In(beijingLocation)
			//计算下一个零点
			next := now.Add(time.Hour * 60).In(beijingLocation)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, beijingLocation)
			t := time.NewTimer(next.Sub(now))
			<-t.C
			lumberJackLogger.Rotate()
		}
	}()

	return zapcore.AddSync(lumberJackLogger)
}

// InitLogger 初始化Logger
func InitLogger(lCfg LogConfig) {
	writeSyncer := getLogWriter(lCfg.FileName, lCfg.MaxSize, lCfg.MaxBackups, lCfg.MaxAge)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	logger = zap.New(core, zap.AddCaller())
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}
func WarnL(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}
func Error(fields ...zap.Field) {
	logger.Error("panic-error", fields...)
}

func Sync() error {
	return logger.Sync()
}
