package internal

import (
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	changer Changer
}

func NewWatch(changer Changer) *Watcher {
	return &Watcher{
		changer: changer,
	}
}

func (w *Watcher) Start(path string) (err error) {
	if watcher, nwe := fsnotify.NewWatcher(); nil != nwe { // 创建监听对象
		err = nwe
	} else if ae := watcher.Add(path); nil != ae { // 添加要监听的文件或目录
		err = ae
	} else {
		go w.watch(watcher)
	}

	return
}

func (w *Watcher) watch(watcher *fsnotify.Watcher) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if event.Has(fsnotify.Write) && ok {
				w.changer.Wrote()
			}
		}
	}
}
