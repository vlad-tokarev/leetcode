package longcons

var (
	// we use min int32 as sentinel value to mark digits which are already inside some sequences
	// For example
	inside = int32(-2147483648)
	// we use max int32 as sentinel value to mark digits that are alone. They do not belong any sequences
	alone = int32(2147483647)
)

func longestConsecutive(nums []int) int {
	var nums32 = make([]int32, len(nums))
	for i, num := range nums {
		nums32[i] = int32(num)
	}
	return longestConsecutive32(nums32)
}

func longestConsecutive32(nums []int32) int {

	if len(nums) == 0 {
		return 0
	}

	var (
		seqDescriptor int32
		ok            bool
	)

	// m is a map, where key is number, and value is next or previous value in sequence
	m := make(map[int32]int32)

	for _, num := range nums {

		if _, ok = m[num]; ok {
			continue // handled
		}

		// first, lets check either the number can be joined on the right side [. . seqDescriptor ] + [num]
		{

			m[num] = alone // by default number is alone (is not joined to sequences yet)

			seqDescriptor, ok = m[num-1]
			if !ok { // we can't join 9 because 8 is not found is numbers space
				goto L
			}

			// if 8 is alone, or 8 is upper boundary, we can joint 9 as upper boundary to this sequence
			if seqDescriptor == alone {
				m[num-1] = num // 9 becomes neighbor
				m[num] = num - 1
				goto L
			}

			// let's imagine that num is equal to 9
			// in this case if 8 marked as value that already belongs to some sequence and
			// exist inside this sequence, it means that AT LEAST this sequence is [7, 8, 9]
			// it means that 9 is duplicate, and we can skip it
			//
			// If 8 is marked as bottom boundary, that at LEAST the sequence is [8, 9]
			// It means that 9 is duplicate, and we can skip it
			if seqDescriptor == inside || seqDescriptor > num-1 {
				m[num] = inside
				continue
			}

			if seqDescriptor < num-1 {
				m[num-1] = inside
				m[num] = seqDescriptor
				m[seqDescriptor] = num
			}
		}

	L:
		{ // handle [9] + [10, ...]
			seqDescriptor, ok = m[num+1]
			if !ok { // we can't join 9 because 10 is not found is numbers space
				continue
			}

			if seqDescriptor == inside || seqDescriptor < num+1 {
				m[num] = inside
				continue
			}

			if seqDescriptor == alone {
				if m[num] == alone {
					m[num] = num + 1
					m[num+1] = num
				} else {
					m[num+1] = m[num]
					m[m[num]] = num + 1
					m[num] = inside
				}
				continue
			}
			if seqDescriptor > num+1 {
				m[num+1] = inside
				if m[num] == alone {
					m[num] = seqDescriptor
					m[seqDescriptor] = num
				} else {
					m[m[num]] = seqDescriptor
					m[seqDescriptor] = m[num]
					m[num] = inside
				}
			}
		}

	}

	var (
		l      int32
		maxLen int32 = 1
	)
	for i, descriptor := range m {
		l = length(i, descriptor)
		if maxLen < l {
			maxLen = l
		}
	}

	return int(maxLen)

}

func length(i, descriptor int32) int32 {
	if descriptor == alone {
		return 1
	}
	if descriptor == inside {
		return 0
	}
	if descriptor > i {
		return descriptor - i + 1
	}

	return i - descriptor + 1
}
