package main

import "testing"

func TestCountAfterDays(t *testing.T) {
	t.Run("3", func(t *testing.T) {
		ct := countAfterDays(3, 3)
		if ct != 1 {
			t.Errorf("expected 1 but was %d", ct)
		}
	})
	t.Run("4", func(t *testing.T) {
		ct := countAfterDays(3, 4)
		if ct != 2 {
			t.Errorf("expected 2 but was %d", ct)
		}
	})
	t.Run("10", func(t *testing.T) {
		ct := countAfterDays(3, 10)
		if ct != 2 {
			t.Errorf("expected 2 but was %d", ct)
		}
	})
	t.Run("11", func(t *testing.T) {
		ct := countAfterDays(3, 11)
		if ct != 3 {
			t.Errorf("expected 3 but was %d", ct)
		}
	})
	t.Run("13", func(t *testing.T) {
		ct := countAfterDays(3, 13)
		if ct != 4 {
			t.Errorf("expected 4 but was %d", ct)
		}
	})
	t.Run("18", func(t *testing.T) {
		ct := countAfterDays(3, 18)
		if ct != 5 {
			t.Errorf("expected 5 but was %d", ct)
		}
	})


	//t.Run("Can spawn new fish", func(t *testing.T) {
	//	ct := countAfterDays(3, 4)
	//	if ct != 2 {
	//		t.Errorf("expected 2 but was %d", ct)
	//	}
	//})
	//t.Run("Children can spawn new fish", func(t *testing.T) {
	//	ct := countAfterDays(3, 15)
	//	if ct != 4 {
	//		t.Errorf("expected 4 but was %d", ct)
	//	}
	//})
}

func TestPart1(t *testing.T) {
	initialStates := []int{3, 4, 3, 1, 2}
	t.Run("After 18 days", func(t *testing.T) {
		ct := countAllAfterDays(initialStates, 18)
		if ct != 26 {
			t.Errorf("Expected 26 lanternfish after 18 days but was %d", ct)
		}
	})
	t.Run("After 80 days", func(t *testing.T) {
		ct := countAllAfterDays(initialStates, 80)
		if ct != 5934 {
			t.Errorf("Expected 5934 lanternfish after 80 days but was %d", ct)
		}
	})
	t.Run("After 256 days", func(t *testing.T) {
		ct := countAllAfterDays(initialStates, 256)
		if ct != 26984457539 {
			t.Errorf("Expected 26984457539 lanternfish after 256 days but was %d", ct)
		}
	})
}

//0		3					1
//1		2					0
//2		1					6,8
//3		0					5,7
//4		6,8				4,6
//5		5,7				3,5
//6		4,6				2,4
//7		3,5				1,3
//8		2,4				0,2
//9		1,3				6,1,8
//10	0,2				5,0,7
//11	6,1,8				4,8,6,8
//12	5,0,7				3,7,5,7
//13	4,8,6,8			2,6,4,6
//14	3,7,5,7			1,5,3,5
//15	2,6,4,6			0,4,2,4
//16	1,5,3,5			6,3,1,3,8
//17	0,4,2,4			5,2,0,2,7
//18	6,3,1,3,8			4,1,8,1,6,8

//18 days
//1:6, 2:5, 3:5:, 4:4, 5:4
//3,4,3,1,2 -> 5+4+5+6+5 = 25
