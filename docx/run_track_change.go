package docx

import (
	"github.com/gomutex/godocx/wml/ctypes"
)

type RunTrackChange struct {
	root *RootDoc               // root is the root document to which this run track change belongs.
	ct   *ctypes.RunTrackChange // ct is the underlying run track change element from the wml/ctypes package.
}

func newRunTrackChange(root *RootDoc, ct *ctypes.RunTrackChange) *RunTrackChange {
	return &RunTrackChange{root: root, ct: ct}
}

// AddRun adds a new Run to the RunTrackChange and returns it.
//
// Returns:
//   - *Run: The newly added Run instance.
func (rtc *RunTrackChange) AddRun() *Run {
	run := &ctypes.Run{}
	rtc.ct.Children = append(rtc.ct.Children, ctypes.RunTrackChangeChild{Run: run})
	return newRun(rtc.root, run)
}
