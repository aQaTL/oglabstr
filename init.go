package oglabstr

import "runtime"

func init() {
	runtime.LockOSThread()
}
