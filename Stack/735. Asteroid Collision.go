package Stack

func asteroidCollision(asteroids []int) []int {
	stack := []int{} // Stack to hold surviving asteroids

	for _, a := range asteroids {
		// Check for collisions: while there are asteroids in the stack and
		// the current asteroid is moving left (negative) and the last asteroid
		// in the stack is moving right (positive)
		for len(stack) > 0 && a < 0 && stack[len(stack)-1] > 0 {
			// Compare the sizes of the colliding asteroids
			if stack[len(stack)-1] < -a {
				// The last asteroid in the stack is smaller, so it explodes
				stack = stack[:len(stack)-1] // Remove the last asteroid
				continue                     // Check again for collision with the new last asteroid
			} else if stack[len(stack)-1] == -a {
				// Both asteroids are equal, so they both explode
				stack = stack[:len(stack)-1] // Remove the last asteroid
			}
			// If the last asteroid in the stack is larger, the current asteroid is destroyed
			break
		}

		// If the current asteroid is positive or if no collision occurs, append it to the stack
		if len(stack) == 0 || a > 0 {
			stack = append(stack, a)
		}
	}

	return stack // Return the surviving asteroids
}
