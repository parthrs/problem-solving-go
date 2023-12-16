package stacks

func willCollide(i, j int) (bool, []int) {
	if i > 0 && j > 0 || i < 0 && j < 0 || i < 0 && j > 0 {
		return false, []int{i, j}
	} else if i+j == 0 {
		return true, []int{}
	} else if i+j > 0 {
		return true, []int{i}
	} else {
		return true, []int{j}
	}
}

func AsteroidCollision(asteroids []int) []int {
	if len(asteroids) == 0 {
		return []int{}
	}
	stack := make([]int, 0, len(asteroids))
	for _, ast := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, ast)
			continue
		}
		collision, survivors := willCollide(stack[len(stack)-1], ast)
		if !collision {
			stack = append(stack, ast)
			continue
		}
		stack = stack[:len(stack)-1]
		for collision && len(survivors) > 0 && len(stack) > 0 {
			collision, survivors = willCollide(stack[len(stack)-1], survivors[0])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, survivors...)
	}
	return stack
}
