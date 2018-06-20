package data

import (
	"testing"
	"time"
)

func TestCompareAuthorizationTime(t *testing.T) {
	beg := time.Now().Unix() -3000
	end := time.Now().Unix() + 3000

	valid := CompareAuthorizationTime(beg,end)
	if valid ==false {
		t.Fail()
	}

	beg = time.Now().Unix() + 200
	end = time.Now().Unix() + 3000
	valid = CompareAuthorizationTime(beg, end)
	if valid == true {
		t.Fail()
	}

}
