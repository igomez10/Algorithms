package meetingplanner

import "testing"

func MeetingPlanner(slotsA [][]int, slotsB [][]int, dur int) []int {
	idxA := 0
	idxB := 0

	for idxA < len(slotsA) && idxB < len(slotsB) {
		latestStartTime := slotsA[idxA][0]
		if latestStartTime < slotsB[idxB][0] {
			latestStartTime = slotsB[idxB][0]
		}

		earliestEndTime := slotsB[idxB][1]
		if slotsA[idxA][1] < earliestEndTime {
			earliestEndTime = slotsA[idxA][1]
		}

		maximumDur := earliestEndTime - latestStartTime
		if maximumDur >= dur {
			return []int{latestStartTime, latestStartTime + dur}
		}

		// where is earliestEndTime and increase it
		if slotsA[idxA][1] < slotsB[idxB][1] {
			idxA++
			continue
		}

		idxB++
	}

	return []int{}
}

func TestMeetingPlanner(t *testing.T) {

	tests := []struct {
		name   string
		slotsA [][]int
		slotsB [][]int
		dur    int
		want   []int
	}{
		{
			name:   "Example 1",
			slotsA: [][]int{{10, 50}, {60, 120}, {140, 210}},
			slotsB: [][]int{{0, 15}, {60, 70}},
			dur:    8,
			want:   []int{60, 68},
		},
		{
			name:   "Example 2",
			slotsA: [][]int{{10, 50}, {60, 120}, {140, 210}},
			slotsB: [][]int{{0, 15}, {60, 70}},
			dur:    12,
			want:   []int{},
		},
		{
			name:   "No overlap",
			slotsA: [][]int{{10, 50}, {60, 120}, {140, 210}},
			slotsB: [][]int{{0, 15}, {60, 70}},
			dur:    30,
			want:   []int{},
		},
		{
			name:   "Exact match",
			slotsA: [][]int{{10, 50}, {60, 120}, {140, 210}},
			slotsB: [][]int{{0, 15}, {60, 70}},
			dur:    10,
			want:   []int{60, 70},
		},
		{
			name:   "One large option",
			slotsA: [][]int{{1, 99}},
			slotsB: [][]int{{2, 5}},
			dur:    10,
			want:   []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MeetingPlanner(tt.slotsA, tt.slotsB, tt.dur)
			// Compare slices safely, including empty results
			if len(got) != len(tt.want) {
				t.Errorf("MeetingPlanner() length = %d, want %d (got=%v, want=%v)", len(got), len(tt.want), got, tt.want)
				return
			}
			if len(got) == 2 {
				if got[0] != tt.want[0] || got[1] != tt.want[1] {
					t.Errorf("MeetingPlanner() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
