from typing import List


class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        maxSoFar = 0
        current = 0

        for i in range(k):
            maxSoFar += nums[i]

        hi = k
        lo = 0
        current = maxSoFar
        while hi < len(nums):
            current -= nums[lo]
            current += nums[hi]

            if current > maxSoFar:
                maxSoFar = current
            lo += 1
            hi += 1

        return maxSoFar/k


s = Solution()
inputArr = [1, 12, -5, -6, 50, 3]
print(s.findMaxAverage(inputArr, 4))
