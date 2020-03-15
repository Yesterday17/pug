/**
PUG
Copyright (C) 2019-2020  Yesterday17

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package log

import (
	"fmt"
	"io"
)

type Logger struct {
	logWriter   io.Writer
	errorWriter io.Writer

	WrappedLogWriter   io.Writer
	WrappedErrorWriter io.Writer
}

func (l *Logger) Log(level Level, message string) {
	_, _ = l.logWriter.Write([]byte(string(getLevelColor(level)) + message + string(endColor)))
}

func (l *Logger) Logf(level Level, format string, args ...interface{}) {
	l.Log(level, fmt.Sprintf(format, args...))
}

func (l *Logger) Info(message string) {
	l.Log(LevelInfo, message)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Logf(LevelInfo, format, args...)
}

func (l *Logger) Warn(message string) {
	l.Log(LevelWarn, message)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Logf(LevelWarn, format, args...)
}

func (l *Logger) Error(message string) {
	_, _ = l.errorWriter.Write([]byte(string(ErrorColor) + message + string(endColor)))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Fatal(message string) {
	_, _ = l.errorWriter.Write([]byte(string(FatalColor) + message + string(endColor)))
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}
