package algorithms

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	resLen := len(nums1) + len(nums2)
	res := make([]int, resLen)
	nums1Len, i := len(nums1), 0
	nums2Len, j := len(nums2), 0

	for k := 0; k < resLen; k++ {
		if i == nums1Len || (i < nums1Len && j < nums2Len && nums1[i] > nums2[j]) {
			res[k] = nums2[j]
			j++
			continue
		}

		if j == nums2Len || (i < nums1Len && j < nums2Len && nums1[i] <= nums2[j]) {
			res[k] = nums1[i]
			i++
			continue
		}
	}
	return findMedian(res)
}

func findMedian(nums []int) float64 {
	numsLen := len(nums)
	if numsLen == 0 {
		panic("无法获取中位数")
	}

	if numsLen%2 == 0 {
		return float64((nums[numsLen/2] + nums[numsLen/2-1])) / 2.0
	}
	return float64(nums[numsLen/2])
}

func FindMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	isOne := (l%2 != 0)
	mid := l/2 - 1
	if isOne {
		mid = l / 2
	}
	p, q := 0, 0
	var cur int
	res := float64(0)
	for p < len(nums1) && q < len(nums2) {
		if nums1[p] < nums2[q] {
			cur = nums1[p]
			p++
		} else {
			cur = nums2[q]
			q++
		}
		if isOne {
			if (p + q - 1) == mid {
				res = float64(cur) * 2
			}
		} else {
			if (p+q-1) == mid || (p+q-1) == mid+1 {
				res += float64(cur)
			}
		}
	}

	for p < len(nums1) {
		cur = nums1[p]
		p++
		if isOne {
			if (p + q - 1) == mid {
				res = float64(cur) * 2
			}
		} else {
			if (p+q-1) == mid || (p+q-1) == mid+1 {
				res += float64(cur)
			}
		}
	}

	for q < len(nums2) {
		cur = nums2[q]
		q++
		if isOne {
			if (p + q - 1) == mid {
				res = float64(cur) * 2
			}
		} else {
			if (p+q-1) == mid || (p+q-1) == mid+1 {
				res += float64(cur)
			}
		}
	}
	return res / 2
}
