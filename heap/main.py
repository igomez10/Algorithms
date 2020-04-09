import heapq
# k = 4
# nums = [3, 2, 3, 1, 2, 4, 5, 5, 6]


class Solution:
    def findKthLargest(self, nums: [int], k: int) -> int:
        # print(list(reversed(sorted(set(nums)))))
        heap = []
        heapq.heapify(heap)

        for i in range(len(nums)):
            if len(heap) > k:
                del(heap[len(heap)-1])
            heapq.heappush(heap, -1 * nums[i])

        largest = heap[-1]
        return largest


s = Solution()

# nums = [3, 2, 1, 5, 6, 4]
# k = 2

nums = [3, 2, 3, 1, 2, 4, 5, 5, 6]
k = 4
kth = s.findKthLargest(nums, k)
if kth != 4:
    print("1.fail", 4, kth)


nums = [3, 2, 1, 5, 6, 4]
k = 2
kth = s.findKthLargest(nums, k)
if kth != 5:
    print("2.fail", 5, kth)


nums = [3, 2, 1, 5, 6, 4]
k = 2
