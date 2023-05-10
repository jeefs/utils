package utils

//根据块大小切分数组
func SliceChunk(s []int,chunkSize int) [][]int {
	res := make([][]int,0)
        for i := 0; i < len(s); i += chunkSize {
                left := i
                right := i + chunkSize
                var chunk []int
                if right > len(s) {
                        chunk = s[left:]
                } else {
                        chunk = s[left:right]
                }
	        res  = append(res,chunk)	

        }
	return res
}

//判断数组中(不重复)值是否连续
func IsContinuous(a []int) bool {
	min := -1
	max := -1
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] != 0 {
			if min > a[i] || min == -1 {
				min = a[i]
			}
			if max < a[i] || max == -1 {
				max = a[i]
			}
		}
	}
	if max-min > n-1 {
		return false
	} else {
		return true
	}
}
