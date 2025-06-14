package main

func searchMatrix(matrix [][]int, target int) bool {
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}

	left, right := 0, len(matrix)-1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	col := left - 1
	left, right = 0, len(matrix[0])-1
	for left <= right {
		mid = (left + right) / 2
		if matrix[col][mid] == target {
			return true
		} else if matrix[col][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func searchMatrix1(matrix [][]int, target int) bool {
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}

	left, right := 0, len(matrix)-1
	mid := 0

	for left <= right {
		mid = (left + right) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	row := left - 1
	left, right = 0, len(matrix[0])-1
	for left <= right {
		mid = (left + right) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
