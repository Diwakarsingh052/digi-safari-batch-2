package calc

import (
	"dummyProj/result"
)

func Bar() {
	//in case of cyclic deps, take the common part out in a new package
	result.Result = 10
}
