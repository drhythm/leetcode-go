package p2211

// https://leetcode.cn/problems/count-collisions-on-a-road
func countCollisions(directions string) (count int) {
	// Gemini Version two-pointer
	// left and right pointers
	// left, right := 0, len(directions) - 1
	// // Skip all Leftmost Car Moving Left(Safe)
	// for ; left <= right; left++ {
	// 	if directions[left] != 'L' {
	// 		break
	// 	}
	// }
	// // Skip all Rightmost Car Moving Right (Safe)
	// for ; right >= left; right-- {
	// 	if directions[right] != 'R' {
	// 		break
	// 	}
	// }
	// // Count each car is not stationary in between
	// for ; left <= right; left++ {
	// 	if directions[left] != 'S' {
	// 		count++
	// 	}
	// }
	// return

	// 'last' represents the state of the vehicle (or obstacle) immediately to the left.
	// Initialize with 'L' as a dummy state implies no collision from the left start.
	last := 'L'
	countR := 0 // Accumulator for consecutive cars moving Right

	for _, c := range directions {
		if c == 'L' {
			// Case 1: Current car moves Left
			if last == 'S' {
				// Crashes into a stationary car
				count++
			} else if last == 'R' {
				// Crashes into a stream of right-moving cars.
				// The current 'L' crashes (+1), and all previous 'R's crash (+countR).
				count += 1 + countR
				countR = 0
				last = 'S' // The pile-up becomes a stationary obstacle
			}
			// If last == 'L', it just drives away safely to the left.
		} else if c == 'S' {
			// Case 2: Current car is Stationary
			if last == 'R' {
				// All accumulated 'R' cars crash into this wall
				count += countR
				countR = 0
			}
			last = 'S'
		} else {
			// Case 3: Current car moves Right
			countR++
			last = 'R'
		}
	}
	return
}
