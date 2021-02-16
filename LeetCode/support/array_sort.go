package support

/*
十大经典排序算法：
https://mp.weixin.qq.com/s/Ab5vW2MbqkmwJXMYrfrTpg
*/

/*
冒泡排序
*/
func Bubbling(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			a := nums[j-1]
			b := nums[j]
			if a > b {
				nums[j] = a
				nums[j-1] = b
			}
		}
	}
}

/*
插入排序
【稳定】
*/
func InsertSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			a := nums[j-1]
			b := nums[j]
			if a > b {
				nums[j] = a
				nums[j-1] = b
			} else {
				break
			}
		}
	}
}

/*
选择排序
和插入排序相反 选择是从未排序区间选择一个最小值放在已排区间末尾
因为选择排序 必定要遍历整个未排区间，而插入排序在理想情况下，只需要遍历已排区间的1个元素，因此选择排序常常被插入排序替代
【不稳定】
*/
func SelectSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		minIndex := i
		for j := i; j < len(nums); j++ {
			if nums[j] < min {
				minIndex = j
				min = nums[j]
			}
		}
		nums[minIndex] = nums[i]
		nums[i] = min
	}
}

/**
快速排序
快排可以简单理解为：每排序一次把pivot移动到 它应处的 排序后的位置
B站教学视频： https://b23.tv/Lqlpet
【不稳定】
18 7 9 5 3 46 2 4 43
*/
func FastSort(nums []int) {
	fast_sort_segment(nums, 0, len(nums)-1)
}

/**
三个参数代表 一个数组片段
*/
func fast_sort_segment(nums []int, start int, end int) {
	size := end - start
	if size < 1 {
		return
	}
	pivotIndex := start
	pivot := nums[pivotIndex]
	l, r := start, end
	// 是否正在移动左指针
	isLeft := false
	for l != r {
		// 从右指针 左右指针交规前进，当右指针没有比pivot大时 下次循环前进左指针； 左指针前进规则同理
		if isLeft {
			// 当正在移动左指针时 如果左指针元素大于pivot  则交换
			//【无法保证相同元素的顺序】
			if nums[l] > pivot {
				nums[pivotIndex] = nums[l]
				pivotIndex = l
				isLeft = false
			}
		} else {
			// 移动右指针时，如果右指针对应的值比中轴小，则拿当前右指针对应的数据和中轴pivot交换;否则 下次循环仍然前进右指针，右指针前进1格
			if nums[r] < pivot {
				nums[pivotIndex] = nums[r]
				pivotIndex = r
				isLeft = true
			}
		}
		if isLeft {
			l++
		} else {
			r--
		}
	}
	// 最终得到中轴的位置，此时pivotIndex左边序列都比pivot小，右边都比pivot大
	nums[pivotIndex] = pivot

	// 对pivot左右两边的序列再排序
	fast_sort_segment(nums, start, pivotIndex-1)
	fast_sort_segment(nums, pivotIndex+1, end)
}

/*
二路快排
二路快排和快排的区别在于：快排每工作一次，把轴放在正确的排序位；而二路快排，则是每工作一次，把轴和轴相同的元素一块放在正确的排序位，这样就不需要再为重复元素排序了
图文教学： https://www.freesion.com/article/5881911751/
【不稳定】
*/
func FastSort2Road(nums []int) {

}

func fastSort2road_segment(nums []int, start int, end int) {
	if end-start < 1 {
		return
	}
	// todo

}

/*
希尔排序
【原地、不稳定】
*/
func ShellSort(nums []int) {
	// 希尔排序
	if nums == nil || len(nums) < 2 {
		return
	}
	l := len(nums)
	// 分组长度
	// 第一个循环是有来多级分组的 每次按照上一个分组步长/2 直到步长等于1
	for gl := l / 2; gl > 0; gl /= 2 {
		// 第二个循环 是遍历所有分组的 如从0到
		for k := 0; k < gl; k++ {
			// 第三个和第四个循环是对每个分组做插入排序用的
			for i := gl + k; i < l; i += gl {
				// 这个子循环内 要进行插入排序
				for j := i; j >= gl; j -= gl {
					if nums[j] < nums[j-gl] {
						b := nums[j]
						nums[j] = nums[j-gl]
						nums[j-gl] = b
					} else {
						break
					}
				}
			}
		}
	}
}

/**
归并排序
图解归并：https://www.cnblogs.com/chengxiao/p/6194356.html
【稳定、非原地】
*/
func MergeSort(nums []int) {
	// 归并排序
	if nums == nil || len(nums) < 2 {
		return
	}
	merge_sort_split(nums, 0, len(nums)-1)
}

func merge_sort_split(nums []int, start int, end int) {
	l := end + 1 - start
	if l < 2 {
		return
	} else if l == 2 {
		if nums[start] > nums[end] {
			tem := nums[start]
			nums[start] = nums[end]
			nums[end] = tem
		}
		return
	}
	i := l / 2
	merge_sort_split(nums, start, i)
	merge_sort_split(nums, i+1, end)

	// 合并切片结果
	merge_sort_merge(nums, start, i, end)
}

/*
每一次split切割数组 完成以后都要把结果合并一次
在合并的时候 从start-> i  以及i->end 两边都已经有序了
*/
func merge_sort_merge(nums []int, start int, i int, end int) {
	l := end + 1 - start
	if l <= 2 {
		return
	}
	// 两个有序数组的合并，start->i 和i+1->end
	temarr := make([]int, l)
	k := 0
	a := start
	b := i + 1
	for ; a <= i; a++ {
		av := nums[a]
		for b <= end {
			bv := nums[b]
			if av <= bv {
				temarr[k] = av
				k++
				if a <= i {
					// a没有遍历完的情况下 才打断去a中找更小的值，如果a已经遍历完了 要把二遍历完
					break
				} else {
					b++
				}
			} else {
				temarr[k] = bv
				b++
				k++
			}
		}

	}
	k = 0
	for i := start; i <= end; i++ {
		nums[i] = temarr[k]
		k++
	}

}
