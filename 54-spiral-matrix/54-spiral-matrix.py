class Solution(object):
    def spiralOrder(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[int]
        """
        left, right = 0, len(matrix[0])
        top, bottom = 0, len(matrix)
        res = []
        
        while left < right and top < bottom:
            for i in range(left, right):
                res.append(matrix[top][i])
            top += 1
            
            for i in range(top, bottom):
                res.append(matrix[i][right-1])
            right -= 1
            
            if not(top < bottom and left < right):
                break
            
            for i in range(right-1, left, -1):
                res.append(matrix[bottom-1][i])
            bottom -= 1
            
            for i in range(bottom, top-1, -1):
                res.append(matrix[i][left])
            left += 1
        return res