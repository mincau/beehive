/*
 *    Copyright (C) 2017 Timm Schäuble
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *		Timm Schäuble <tymmm1+gh@gmail.com>
 */

// Package simplepushbee is a Bee that is able to send push notifications to Android.
package simplepushbee

import (
	"github.com/muesli/beehive/bees"
	"github.com/simplepush/simplepush-go"
)

// SimplepushBee is a Bee that is able to send push notifications to Android.
type SimplepushBee struct {
	bees.Bee

	key      string
	password string
	salt     string
}

// Action triggers the action passed to it.
func (mod *SimplepushBee) Action(action bees.Action) []bees.Placeholder {
	outs := []bees.Placeholder{}

	switch action.Name {
	case "send":
		title := ""
		message := ""
		event := ""

		action.Options.Bind("title", &title)
		action.Options.Bind("message", &message)
		action.Options.Bind("event", &event)

		if mod.password != "" {
			simplepush.Send(simplepush.Message{mod.key, title, message, event, true, mod.password, mod.salt})
		} else {
			simplepush.Send(simplepush.Message{mod.key, title, message, event, false, "", ""})
		}

	default:
		panic("Unknown action triggered in " + mod.Name() + ": " + action.Name)
	}

	return outs
}

// ReloadOptions parses the config options and initializes the Bee.
func (mod *SimplepushBee) ReloadOptions(options bees.BeeOptions) {
	mod.SetOptions(options)

	options.Bind("key", &mod.key)
	options.Bind("password", &mod.password)
	options.Bind("salt", &mod.salt)
}
