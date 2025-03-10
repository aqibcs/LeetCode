def removeElement(nums, val):
    k = 0
    for i in range(len(nums)):
        if nums[i] != val:
            nums[k] = nums[i]
            k += 1
    return k 

# Example Usage:
nums1 = [3, 2, 2, 3]
val1 = 3
result1 = removeElement(nums1, val1)
print(result1, nums1[:result1])

nums2 = [0, 1, 2, 2, 3, 0, 4, 2]
val2 = 2
result2 = removeElement(nums2, val2)
print(result2, nums2[:result2])
