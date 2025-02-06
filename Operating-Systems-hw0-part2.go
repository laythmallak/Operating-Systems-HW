package main
import "fmt"
import "errors"

// Stack struct
type Stack struct {
	values []int // Slice to store the numbers in the stack
	size   int   // Max size of the stack
}

// Push func adds a number to the top of the stack.
// Checks for stack overflow. Appends the number to the stack.
func (s *Stack) Push(val int) error {
	// Check if the stack is full (stack overflow)
	if len(s.values) >= s.size {
		return errors.New("stack overflow") // Return an error if the stack is full
	}
	// Append the item to the slice (top of the stack)
	s.values = append(s.values, val)
	return nil // Return nil to end the function
}

// Pop func removes and returns the top number from the stack.
// Checks for stack underflow. Pops the top number from the stack.
func (s *Stack) Pop() (int, error) {
	// Check if the stack is empty (stack underflow)
	if len(s.values) == 0 {
		return 0, errors.New("stack underflow") // Return an error if the stack is empty
	}
	// Get the top number from the slice
	val := s.values[len(s.values)-1]
	// Remove the top number from the slice
	s.values = s.values[:len(s.values)-1]
	return val, nil // Return the popped number and nil to end function
}

func main() {
	// Create a stack with a maximum size of 100
	stack := Stack{size: 100}

	// Push some numbers onto the stack using for loop
	for i := 1; i <= 5; i++ {
		err := stack.Push(i) // Push the number onto the stack
		if err != nil {
			fmt.Println(err) // Print an error message if the stack is full (stack overflow)
			return
		}
		fmt.Printf("Pushed: %d\n", i) // Print the number that was pushed
	}

	// Pop numbers from the stack using for loop
	for i := 1; i <= 5; i++ {
		val, err := stack.Pop() // Pop the top number from the stack
		if err != nil {
			fmt.Println(err) // Print an error message if the stack is empty (stack underflow)
			return
		}
		fmt.Printf("Popped: %d\n", val) // Print the number that was popped
	}
}