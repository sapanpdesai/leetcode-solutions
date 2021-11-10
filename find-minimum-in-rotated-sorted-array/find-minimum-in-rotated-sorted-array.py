class Solution(object):
    def findMin(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        l = 0
        r = n - 1
        
        if len(nums) == 1:
            return nums[0]
        
        while l <= r:
            mid = (l+r)//2
            if nums[mid] > nums[mid+1]:
                return nums[mid+1]
            if nums[mid-1] > nums[mid]:
                return nums[mid]
            if nums[mid] > nums[r]:
                l = mid
            else:
                r = mid
        return 