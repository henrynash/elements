package main

import (
	"fmt"
)

func main() {

	// This is an array; it is a fixed length list. In this case, a list of 3 strings "a", "b", and "c"
	// We rarely use this in Antha.
	var array = [3]string{"a", "b", "c"}

	fmt.Println("An array of length 3: ", array)

	// More commonly we'll use a slice instead since the length is not restricted up front.
	// This is a slice; it is a list which has no fixed length. In this case, a list of 3 strings "x", "y", and "z"
	var slice = []string{"x", "y", "z"}

	// we can also create an empty slice and add entries to the slice after.
	var slice2 []string

	fmt.Println("An empty slice:", slice2)

	// We add entries using append
	slice2 = append(slice2, "d", "e", "f")

	fmt.Println("A slice after using append:", slice2)

	// Alternatively we can make a slice with an initial length using make
	var slice3 = make([]string, 3)

	fmt.Println("A slice of length 3:", slice3)

	// if we define the length initially in this way instead of using append
	// we can directly set the value for a position in the slice
	// in this case we'll set position 1 to "g"

	slice3[1] = "g"

	fmt.Println("A slice of length 3 after updating position 1:", slice3)

	// You'll notice the first position was not updated, but the second was.
	// The reason for this is that in Antha, Golang and most programming languages we start counting from 0 and not from 1.
	slice3[0] = "h"

	fmt.Println("A slice of length 3 after updating position 0:", slice3)

	slice3[2] = "i"

	fmt.Println("A slice of length 3 after updating position 2:", slice3)

	// if we were now to try to set position 3  in the same way (slice3[3] = "j") an error would occur since the slice has only 3 positions.
	// To add an additional value we'll need to go back to append.

	slice3 = append(slice3, "j")

	fmt.Println("An updated slice, now of length 4 after appending a new value:", slice3)

	// A common operation with slices is to use a for loop to iteratively go through each position.
	// A for loop will continue until the statement following it is no longer true.

	var counter int

	// if the {} follow directly from for, the loop will continue indefinitely unless a break occurs
	for {
		// To prevent an infinite loop we'll use the counter variable
		// in each loop we'll increase the counter by 1 and break when the counter is equal to the length of the slice3
		if counter >= len(slice3) {
			break
		}

		fmt.Println("before change: ", counter, slice3[counter])

		// for each loop we'll concatenate the value of slice3 to the corresponding position of slice
		slice3[counter] = slice3[counter] + slice[counter]

		fmt.Println("after change: ", counter, slice3[counter])

		counter = counter + 1
	}

	// A more common short hand for this places all three conditions of the loop on the same line as for separated by ;
	// we initialise a variable called counter to 0.
	// Continue as long as counter is less than length of slice3.
	// Each iteration of the loop we increase counter by 1 (counter++ is short hand for counter = counter + 1)
	for counter := 0; counter < len(slice3); counter++ {

		slice3[counter] = slice[counter]

		fmt.Println(counter, slice3[counter])
	}

	// The simplest way to loop through a slice is using the range keyword.
	// here we're setting a variable called index which is an integer that will correspond to the position in the slice
	// each iteration of the loop, index will increase until all positions have been evaluated.
	for index := range slice {
		fmt.Println("This is the easiest way to loop through a slice: using range and the index:", slice, index, slice[index])
	}

	// if we add a second output from range, that will be an alias for the instance of slice at the position index
	for index, specificInstanceOfSliceAtIndex := range slice {
		fmt.Println("This is the easiest way to loop through a slice: using range and the index and the instance of that slice at the index: ", slice, index, specificInstanceOfSliceAtIndex, slice[index])
	}

	// if we only want to use the second output the first must be replaced by an _
	for _, specificInstanceOfSliceAtIndex := range slice {
		fmt.Println("This is the easiest way to loop through a slice: using range and the instance of that slice at the index", slice, specificInstanceOfSliceAtIndex)
	}

}
