package main

//
//import (
//	"net"
//	"net/http"
//	"net/http/httputil"
//	"os"
//	"runtime/debug"
//	"strings"
//	"time"
//
//	"github.com/gin-gonic/gin"
//	"github.com/natefinch/lumberjack"
//	"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
//)
//
//var logger *zap.Logger
//var sugarLogger *zap.SugaredLogger
//
////
////func main() {
////	InitLogger()
////	defer logger.Sync()
////	for i := 0; i < 1000000; i++ {
////		logger.Info("testing...")
////	}
////	simpleHttpGet("www.google.com")
////	simpleHttpGet("http://www.google.com")
////}
//
//func main() {
//	InitLogger()
//	//r := gin.Default()
//	r := gin.New()
//	r.Use(GinLogger(logger), GinRecovery(logger, true))
//	r.GET("/hello", func(c *gin.Context) {
//		c.String(http.StatusOK, "hello,world!")
//	})
//	r.Run()
//}
//
////func InitLogger() {
////	//logger, _ = zap.NewProduction()
////	logger, _ = zap.NewDevelopment()
////	sugarLogger = logger.Sugar()
////}
//func InitLogger() {
//	writeSyncer := getLogWriter()
//	encoder := getEncoder()
//	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
//
//	logger := zap.New(core, zap.AddCaller())
//	sugarLogger = logger.Sugar()
//}
//
//func getEncoder() zapcore.Encoder {
//	//	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
//	//下面是普通的格式
//	encoderConfig := zap.NewProductionEncoderConfig()
//	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
//	return zapcore.NewConsoleEncoder(encoderConfig)
//}
//
//func getLogWriter() zapcore.WriteSyncer {
//	//file,_:=os.Openfile("./test.log",os.O_CREATE|os.O_APPEND|os.O_RDWR)
//	lumberJackLogger := &lumberjack.Logger{
//		Filename:   "./test.log",
//		MaxSize:    1,     //M
//		MaxBackups: 5,     //最大备份数量
//		MaxAge:     30,    //最大备份天数
//		Compress:   false, //是否压缩
//	}
//	return zapcore.AddSync(lumberJackLogger)
//}
//func simpleHttpGet(url string) {
//	resp, err := http.Get(url)
//	if err != nil {
//		sugarLogger.Error(
//			"Error fetching url..",
//			zap.String("url", url),
//			zap.Error(err))
//	} else {
//		sugarLogger.Info("Success..",
//			zap.String("statusCode", resp.Status),
//			zap.String("url", url))
//		resp.Body.Close()
//	}
//}
//
//// GinLogger 接收gin框架默认的日志
//func GinLogger(logger *zap.Logger) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		start := time.Now()
//		path := c.Request.URL.Path
//		query := c.Request.URL.RawQuery
//		c.Next()
//
//		cost := time.Since(start)
//		logger.Info(path,
//			zap.Int("status", c.Writer.Status()),
//			zap.String("method", c.Request.Method),
//			zap.String("path", path),
//			zap.String("query", query),
//			zap.String("ip", c.ClientIP()),
//			zap.String("user-agent", c.Request.UserAgent()),
//			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
//			zap.Duration("cost", cost),
//		)
//	}
//}
//
//// GinRecovery recover掉项目可能出现的panic
//func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		defer func() {
//			if err := recover(); err != nil {
//				// Check for a broken connection, as it is not really a
//				// condition that warrants a panic stack trace.
//				var brokenPipe bool
//				if ne, ok := err.(*net.OpError); ok {
//					if se, ok := ne.Err.(*os.SyscallError); ok {
//						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
//							brokenPipe = true
//						}
//					}
//				}
//
//				httpRequest, _ := httputil.DumpRequest(c.Request, false)
//				if brokenPipe {
//					logger.Error(c.Request.URL.Path,
//						zap.Any("error", err),
//						zap.String("request", string(httpRequest)),
//					)
//					// If the connection is dead, we can't write a status to it.
//					c.Error(err.(error)) // nolint: errcheck
//					c.Abort()
//					return
//				}
//
//				if stack {
//					logger.Error("[Recovery from panic]",
//						zap.Any("error", err),
//						zap.String("request", string(httpRequest)),
//						zap.String("stack", string(debug.Stack())),
//					)
//				} else {
//					logger.Error("[Recovery from panic]",
//						zap.Any("error", err),
//						zap.String("request", string(httpRequest)),
//					)
//				}
//				c.AbortWithStatus(http.StatusInternalServerError)
//			}
//		}()
//		c.Next()
//	}
//}
