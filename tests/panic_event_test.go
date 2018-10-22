package parse_test

import (
	"strings"
	"testing"

	"github.com/mfridman/tparse/parse"
)

func TestPanicEvent(t *testing.T) {

	// The input contained a test that panicked, we need to catch this.

	pkgs, err := parse.Do(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	for name, pkg := range pkgs {
		if pkg.Summary == nil {
			t.Fatalf("got nil summary for pkg %q; summary should never be nil", name)
		}

		if _, ok := pkg.HasPanic(); !ok {
			t.Fatalf("got has panic false for pkg %q; want ok to be true", name)
		}
	}

}

const input = `{"Time":"2018-10-21T22:15:24.47322-04:00","Action":"run","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus"}
{"Time":"2018-10-21T22:15:24.473515-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"=== RUN   TestStatus\n"}
{"Time":"2018-10-21T22:15:24.473542-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"=== PAUSE TestStatus\n"}
{"Time":"2018-10-21T22:15:24.47355-04:00","Action":"pause","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus"}
{"Time":"2018-10-21T22:15:24.473565-04:00","Action":"cont","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus"}
{"Time":"2018-10-21T22:15:24.473573-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"=== CONT  TestStatus\n"}
{"Time":"2018-10-21T22:15:24.473588-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"--- FAIL: TestStatus (0.00s)\n"}
{"Time":"2018-10-21T22:15:24.47549-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"panic: runtime error: invalid memory address or nil pointer dereference [recovered]\n"}
{"Time":"2018-10-21T22:15:24.475513-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"\tpanic: runtime error: invalid memory address or nil pointer dereference\n"}
{"Time":"2018-10-21T22:15:24.475532-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1112389]\n"}
{"Time":"2018-10-21T22:15:24.47554-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"\n"}
{"Time":"2018-10-21T22:15:24.475549-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"goroutine 18 [running]:\n"}
{"Time":"2018-10-21T22:15:24.475559-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"testing.tRunner.func1(0xc0000b6300)\n"}
{"Time":"2018-10-21T22:15:24.475567-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"\t/usr/local/go/src/testing/testing.go:792 +0x387\n"}
{"Time":"2018-10-21T22:15:24.475581-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"panic(0x1137980, 0x1262100)\n"}
{"Time":"2018-10-21T22:15:24.475651-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"\t/usr/local/go/src/runtime/panic.go:513 +0x1b9\n"}
{"Time":"2018-10-21T22:15:24.475682-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"github.com/mfridman/tparse/tests_test.TestStatus.func1(0x116177e, 0xe, 0x1185120, 0xc00006c820, 0x0, 0x0, 0x0, 0xc00002e6c0)\n"}
{"Time":"2018-10-21T22:15:24.475695-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"\t/Users/michael.fridman/go/src/github.com/mfridman/tparse/tests/status_test.go:26 +0x69\n"}
{"Time":"2018-10-21T22:15:24.475749-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"path/filepath.walk(0x116177e, 0xe, 0x1185120, 0xc00006c820, 0xc0000666a0, 0x0, 0x10)\n"}
{"Time":"2018-10-21T22:15:24.475773-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"\t/usr/local/go/src/path/filepath/path.go:362 +0xf6\n"}
{"Time":"2018-10-21T22:15:24.475781-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"path/filepath.Walk(0x116177e, 0xe, 0xc0000666a0, 0x1c338b20, 0xf815f)\n"}
{"Time":"2018-10-21T22:15:24.475788-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"\t/usr/local/go/src/path/filepath/path.go:404 +0x105\n"}
{"Time":"2018-10-21T22:15:24.475798-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"github.com/mfridman/tparse/tests_test.TestStatus(0xc0000b6300)\n"}
{"Time":"2018-10-21T22:15:24.475936-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"\t/Users/michael.fridman/go/src/github.com/mfridman/tparse/tests/status_test.go:19 +0x7e\n"}
{"Time":"2018-10-21T22:15:24.475945-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"testing.tRunner(0xc0000b6300, 0x116ab18)\n"}
{"Time":"2018-10-21T22:15:24.475952-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"\t/usr/local/go/src/testing/testing.go:827 +0xbf\n"}
{"Time":"2018-10-21T22:15:24.475959-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"created by testing.(*T).Run\n"}
{"Time":"2018-10-21T22:15:24.475975-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"\t/usr/local/go/src/testing/testing.go:878 +0x353\n"}
{"Time":"2018-10-21T22:15:24.476216-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Output":"FAIL\tgithub.com/mfridman/tparse/tests\t0.014s\n"}
{"Time":"2018-10-21T22:15:24.476261-04:00","Action":"fail","Package":"github.com/mfridman/tparse/tests","Test":"TestStatus","Elapsed":0.014}`
