package partitions

import (
	"fmt"
	"testing"
)

func TestPartitions(t *testing.T) {
	partitions, num_partitions := partition(5, 3, []int{})
	fmt.Printf("\npartitions: %v\n", partitions)
	fmt.Printf("\nnum_partitions: %v\n", num_partitions)

}
