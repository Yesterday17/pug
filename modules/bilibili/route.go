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

package bilibili

type route struct {
	os      string
	query   string
	url     string
	profile string
}

var (
	upos = route{
		os:      "upos",
		query:   "os=upos&upcdn=ws",
		url:     "https://upos-sz-upcdnws.acgvideo.com/OK",
		profile: "ugcupos/bup", // if interactive it should be ugcupos/iv
	}
	kodo = route{
		os:      "kodo",
		query:   "os=kodo&bucket=bvcupcdnkodobm",
		url:     "https://up-na0.qbox.me/crossdomain.xml",
		profile: "ugcupos/bupfetch",
	}
)

func (m *Module) SelectRoute() {
	// TODO: Select route after ping
	m.Route = upos
}
