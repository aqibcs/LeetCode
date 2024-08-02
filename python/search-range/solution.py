class Solution:

    def searchRange(self, nums: List[int], target: int) -> List[int]:

        def find_first(nums, target):
            left, right = 0, len(nums) - 1
            first_position = -1

            while left <= right:
                mid = (left + right) // 2
                if nums[mid] == target:
                    first_position = mid
                    right = mid - 1
                elif nums[mid] < target:
                    left = mid + 1
                else:
                    right = mid - 1
            return first_position

        def find_last(nums, target):
            left, right = 0, len(nums) - 1
            last_position = -1

            while left <= right:
                mid = (left + right) // 2
                if nums[mid] == target:
                    last_position = mid
                    left = mid + 1
                elif nums[mid] < target:
                    left = mid + 1
                else:
                    right = mid - 1
            return last_position

        first_position = find_first(nums, target)
        last_position = find_last(nums, target)

        return [first_position, last_position]
