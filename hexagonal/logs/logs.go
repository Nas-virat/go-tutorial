package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger


func init(){
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //change time format
	config.EncoderConfig.StacktraceKey = ""

	var err error
	//zap.AddCallerSkip(1) is options skip the call 
	// make log not to call on "logs/logs.go:26" but to the right line 
	log,err = config.Build(zap.AddCallerSkip(1))
	if err != nil{
		panic(err)
	}
}

func Info(message string,fields ...zap.Field){
	log.Info(message, fields...)
}

func Debug(message string,fields ...zap.Field){
	log.Debug(message, fields...)
}

// interface{} is similar to object 
func Error(message interface{},fields ...zap.Field){
	// check type of message
	switch v := message.(type){
	case error: // error type
		log.Error(v.Error(), fields...)
	case string: // string type
		log.Error(v, fields...)
	}

}