package main

import "testing"

func TestMeetingRooms(t *testing.T) {
	// [[0,30],[5,10],[15,20]]
	cases := map[int][][]int{
		// 2: {{0, 30}, {5, 10}, {15, 20}},
		2: {{1, 5}, {8, 9}, {8, 9}},
	}
	for k, v := range cases {
		ans := minMeetingRooms(v)
		if k != ans {
			t.Errorf("Expected %d for %+v but got %d", k, v, ans)
		}
	}

}
