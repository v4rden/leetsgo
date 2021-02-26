package MergeSortedArray

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	m-- // first array(values only) index
	n-- // second array index
	resultIndex := len(nums1) - 1

	for ; n >= 0 && m >= 0; resultIndex-- {
		if nums1[m] > nums2[n] {
			nums1[resultIndex] = nums1[m]
			m--
		} else {
			nums1[resultIndex] = nums2[n]
			n--
		}
	}

	for n >= 0 {
		nums1[resultIndex] = nums2[n]
		n--
		resultIndex--
	}
}
