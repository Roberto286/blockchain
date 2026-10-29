package shared

import "sync"

var Mutex = &sync.Mutex{}

var Announcements = make(chan string)

var validators = make(map[string]int)
