package other

// This should not be imported
import "github.com/google/martian/log"

func someFunc() {
	log.Debug("foo")
}
