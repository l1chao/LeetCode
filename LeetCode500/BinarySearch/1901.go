package main
 
func findPeakGrid(mat [][]int) []int { 
	left, right, mid := 0, 0, 0
	for i := range mat {
		left, right = 0, len(mat[0])-1
		for left <= right {
			mid = (left + right) / 2
			if mat[i][mid] > mat[i][mid-1] && mat[i][mid] > mat[i][mid+1] {
				if mat[i][mid] 
			}  
		}
	}
}
