package logger

import (
	"fdms/src/infrastructure/config"
	"fdms/src/utils/date_utils"
	"fmt"
	"io"

	"strings"
	"time"

	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Ids struct {
	Key   string
	Value string
}

type FilteredWriter struct {
	w     zerolog.LevelWriter
	level zerolog.Level
}

func (w *FilteredWriter) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}
func (w *FilteredWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level == w.level {
		return w.w.WriteLevel(level, p)
	}
	return len(p), nil
}

var log zerolog.Logger

var name string

func loadConfigLog(asConsole bool) error {

	serverName, err := os.Hostname()
	if err != nil {
		return err
	}
	// parametrosDeMicroservicio := config.ObtenerParametrosDeMicroservicio()
	// name = parametrosDeMicroservicio.Nombre
	// ruta := config.Get().LogSettings.Ruta
	// id := strconv.Itoa(parametrosDeMicroservicio.Id)
	// _, err = os.Stat(filepath.Join(ruta+"MSWS", name, name))

	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		os.MkdirAll(filepath.Join(ruta+"MSWS",
	// 			string(parametrosDeMicroservicio.Tipo),
	// 			name+"-"+id),
	// 			0700)
	// 	} else {
	// 		return err
	// 	}
	// }

	// path := filepath.Join(ruta+"MSWS",
	// 	string(parametrosDeMicroservicio.Tipo),
	// 	name+"-"+id, name+"-"+id)
	// path := filepath.Join(ruta+"MSWS",
	// 	string(parametrosDeMicroservicio.Tipo),
	// 	name+"-"+id)

	path := "./"
	var writersTrace io.Writer
	var writersDebug io.Writer
	var writersInfo io.Writer
	var writersWarn io.Writer
	var writersError io.Writer
	var writersFatal io.Writer

	var writers []io.Writer

	var traceFilter FilteredWriter
	var debugFilter FilteredWriter
	var infoFilter FilteredWriter
	var warnFilter FilteredWriter
	var errorFilter FilteredWriter
	var fatalFilter FilteredWriter

	zerolog.TimeFieldFormat = date_utils.CompleteFormatDate

	if asConsole && config.Get().LogSettings.Console {
		if config.Get().LogSettings.BeutifyConsoleLog {
			output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
			output.FormatLevel = func(i interface{}) string {
				return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
			}
			writers = append(writers, output)
		} else {
			writers = append(writers, os.Stdout)
		}
	}
	switch config.Get().LogSettings.MinLevel {
	case "trace":
		{
			tracePath := filepath.Join(path, "TRACE")
			writersTrace = &lumberjack.Logger{
				Filename:   tracePath + "/Trace.log",                   // File name
				MaxSize:    config.Get().LogSettings.RotationMaxSizeMB, // Size in MB before file gets rotated
				MaxBackups: config.Get().LogSettings.MaxBackups,        // Max number of files kept before being overwritten
				MaxAge:     config.Get().LogSettings.MaxAgeDay,         // Max number of days to keep the files
				Compress:   config.Get().LogSettings.Compress,
			}
			traceWriter := zerolog.MultiLevelWriter(writersTrace)
			traceFilter = FilteredWriter{traceWriter, zerolog.TraceLevel}
			writers = append(writers, &traceFilter)
		}
		fallthrough
	case "debug":
		{
			debugPath := filepath.Join(path, "DEBUG")
			writersDebug = &lumberjack.Logger{
				Filename:   debugPath + "/Debug.log",                   // File name
				MaxSize:    config.Get().LogSettings.RotationMaxSizeMB, // Size in MB before file gets rotated
				MaxBackups: config.Get().LogSettings.MaxBackups,        // Max number of files kept before being overwritten
				MaxAge:     config.Get().LogSettings.MaxAgeDay,         // Max number of days to keep the files
				Compress:   config.Get().LogSettings.Compress,
			}
			debugWriter := zerolog.MultiLevelWriter(writersDebug)
			debugFilter = FilteredWriter{debugWriter, zerolog.DebugLevel}
			writers = append(writers, &debugFilter)

		}
		fallthrough
	case "info":
		{
			infoPath := filepath.Join(path, "INFO")
			writersInfo = &lumberjack.Logger{
				Filename:   infoPath + "/Info.log",
				MaxSize:    config.Get().LogSettings.RotationMaxSizeMB, // Size in MB before file gets rotated
				MaxBackups: config.Get().LogSettings.MaxBackups,        // Max number of files kept before being overwritten
				MaxAge:     config.Get().LogSettings.MaxAgeDay,         // Max number of days to keep the files
				Compress:   config.Get().LogSettings.Compress,
			}
			infoWritter := zerolog.MultiLevelWriter(writersInfo)
			infoFilter = FilteredWriter{infoWritter, zerolog.InfoLevel}
			writers = append(writers, &infoFilter)

		}
		fallthrough
	case "warn":
		{
			warnPath := filepath.Join(path, "WARN")
			writersWarn = &lumberjack.Logger{
				Filename:   warnPath + "/Warn.log",                     // File name
				MaxSize:    config.Get().LogSettings.RotationMaxSizeMB, // Size in MB before file gets rotated
				MaxBackups: config.Get().LogSettings.MaxBackups,        // Max number of files kept before being overwritten
				MaxAge:     config.Get().LogSettings.MaxAgeDay,         // Max number of days to keep the files
				Compress:   config.Get().LogSettings.Compress,
			}
			warnWritter := zerolog.MultiLevelWriter(writersWarn)
			warnFilter = FilteredWriter{warnWritter, zerolog.WarnLevel}
			writers = append(writers, &warnFilter)

		}
		fallthrough
	case "error":
		{
			errorPath := filepath.Join(path, "ERROR")

			writersError = &lumberjack.Logger{
				Filename:   errorPath + "/Error.log",                   // File name
				MaxSize:    config.Get().LogSettings.RotationMaxSizeMB, // Size in MB before file gets rotated
				MaxBackups: config.Get().LogSettings.MaxBackups,        // Max number of files kept before being overwritten
				MaxAge:     config.Get().LogSettings.MaxAgeDay,         // Max number of days to keep the files
				Compress:   config.Get().LogSettings.Compress,
			}
			errWritter := zerolog.MultiLevelWriter(writersError)
			errorFilter = FilteredWriter{errWritter, zerolog.ErrorLevel}
			writers = append(writers, &errorFilter)

		}
		fallthrough
	case "fatal":
		{
			fatalPath := filepath.Join(path, "FATAL")
			writersFatal = &lumberjack.Logger{
				Filename:   fatalPath + "/Fatal.log",                   // File name
				MaxSize:    config.Get().LogSettings.RotationMaxSizeMB, // Size in MB before file gets rotated
				MaxBackups: config.Get().LogSettings.MaxBackups,        // Max number of files kept before being overwritten
				MaxAge:     config.Get().LogSettings.MaxAgeDay,         // Max number of days to keep the files
				Compress:   config.Get().LogSettings.Compress,
			}
			fatalWriter := zerolog.MultiLevelWriter(writersFatal)
			fatalFilter = FilteredWriter{fatalWriter, zerolog.FatalLevel}
			writers = append(writers, &fatalFilter)

		}
	}
	w := zerolog.MultiLevelWriter(writers...)
	log = zerolog.New(w).With().Timestamp().Str(serverName, name).Logger()
	return nil
}

func Log() zerolog.Logger {
	return log
}

func Trace() *zerolog.Event {
	return log.Trace()
	//return &traceLog
}
func Debug() *zerolog.Event {
	return log.Debug()
	//return &debugLog
}
func Info() *zerolog.Event {
	return log.Info()
	//return &infoLog
}
func Warn() *zerolog.Event {
	return log.Warn()
	//return &warnLog
}
func Error() *zerolog.Event {
	return log.Error()
	//return &errorLog
}
func Fatal() *zerolog.Event {
	return log.Fatal()
	//return &fatalLog
}

func GetLogggerWithIdentifiers(identifiers map[string]string) zerolog.Logger {

	child := log.With()
	for key, val := range identifiers {
		child = child.Str(key, val)
	}

	return child.Logger()
}

func init() {
	fmt.Printf("LOGS INIT Numero de GoRutinas:%d Numero De GoThreads:%d \n",
		config.GetNumberOfGoRoutines(),
		config.GetNumbersOfThreads())
	err := loadConfigLog(true)
	if err != nil {
		panic(err)
	}
}
