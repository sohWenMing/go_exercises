package partitions

/*
	I want to be able to:
	get all the possible partitions
	get the number of partitions

	a partition should be expressed as a slice of integers
*/

type part_slice []int

func partition(
	number int,
	max_num int,
	cur_part_slice part_slice) (partitions []part_slice, num_partitions int) {
	num_partitions = 0
	if number == 0 {
		num_partitions += 1
		partitions = append(partitions, cur_part_slice)
		return
	}
	if number < 0 {
		return
	}
	if max_num < 1 {
		return
	}
	// deal first with accounting for subtracting max num from number
	parts_1, num_parts_1 := partition(number-max_num, max_num, append(cur_part_slice, max_num))
	parts_2, num_parts_2 := partition(number, max_num-1, cur_part_slice)
	partitions = append(partitions, parts_1...)
	partitions = append(partitions, parts_2...)
	num_partitions = num_partitions + num_parts_1 + num_parts_2
	return
}
