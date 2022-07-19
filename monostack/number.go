package monostack

// NextGreaterNumber 下一个更大的数
func NextGreaterNumber(numbers []int) []int {
	ans := make([]int, len(numbers))
	var stack []int
	for i, v := range numbers {
		for len(stack) != 0 && numbers[stack[len(stack)-1]] < v {
			ans[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// NextLessNumber 下一个更小的数
func NextLessNumber(numbers []int) []int {
	ans := make([]int, len(numbers))
	var stack []int
	for i, v := range numbers {
		for len(stack) != 0 && numbers[stack[len(stack)-1]] > v {
			ans[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// PrevGreaterNumber 前一个更大的数
func PrevGreaterNumber(numbers []int) []int {
	ans := make([]int, len(numbers))
	var stack []int
	for i := len(numbers) - 1; i >= 0; i-- {
		for len(stack) != 0 && numbers[stack[len(stack)-1]] < numbers[i] {
			ans[stack[len(stack)-1]] = numbers[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// PrevLessNumber 前一个更小的数
func PrevLessNumber(numbers []int) []int {
	ans := make([]int, len(numbers))
	var stack []int
	for i := len(numbers) - 1; i >= 0; i-- {
		for len(stack) != 0 && numbers[stack[len(stack)-1]] > numbers[i] {
			ans[stack[len(stack)-1]] = numbers[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

/*
   for (int i = 0; i < n; ++i) {
       while (!mono_stack.empty() && heights[mono_stack.top()] >= heights[i]) {
           mono_stack.pop();
       }
       left[i] = (mono_stack.empty() ? -1 : mono_stack.top());
       mono_stack.push(i);
   }
*/

func PrevLessNumberTemp(numbers []int) []int {
	ans := make([]int, len(numbers))
	var stack []int
	for i, v := range numbers {
		for len(stack) != 0 && numbers[stack[len(stack)-1]] >= v {
			stack = stack[:len(stack)-1]
		}
		ans[i] = -1
		if len(stack) != 0 {
			ans[i] = numbers[stack[len(stack)-1]]
		}
		stack = append(stack, i)
	}
	return ans
}
