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

type Level int

const (
	LevelInfo  Level = 0
	LevelWarn  Level = 1
	LevelError Level = 2
	LevelFatal Level = 3
)

var (
	InfoColor  = White
	WarnColor  = Yellow
	ErrorColor = LightRed
	FatalColor = Red
)

func getLevelColor(level Level) color {
	switch level {
	case LevelInfo:
		return InfoColor
	case LevelWarn:
		return WarnColor
	case LevelError:
		return ErrorColor
	case LevelFatal:
		return FatalColor
	default:
		return Blue
	}
}
