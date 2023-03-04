package channels

import (
	"github.com/jiyeyuran/go-eventemitter"
)

var HavenCloud = eventemitter.NewEventEmitter(eventemitter.WithMaxListeners(300))
