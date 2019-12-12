/**
PUG
Copyright (C) 2019  Yesterday17

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
	"os"
)

var l = Logger{w: os.Stdout, ew: os.Stderr}

func Info(message string) {
	l.Info(message)
}

func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func Warn(message string) {
	l.Warn(message)
}

func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func Error(message string) {
	l.Error(message)
}

func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func Fatal(message string) {
	l.Fatal(message)
}

func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}
