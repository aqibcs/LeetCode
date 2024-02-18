#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> numMap;

        for (int i = 0; i < nums.size(); ++i) {
            int complement = target - nums[i];

            // Check if the complement is already in the map
            if (numMap.find(complement) != numMap.end()) {
                // Return the indices of the two numbers
                return {numMap[complement], i};
            }

            // If complement is not in the map, add the current number and its index
            numMap[nums[i]] = i;
        }

        // No solution found, return an empty vector
        return {};
    }
};

int main() {
    Solution solution;

    // Example 1
    vector<int> nums1 = {2, 7, 11, 15};
    int target1 = 9;
    vector<int> result1 = solution.twoSum(nums1, target1);

    cout << "Example 1: ";
    for (int index : result1) {
        cout << index << " ";
    }
    cout << endl;

    // Example 2
    vector<int> nums2 = {3, 2, 4};
    int target2 = 6;
    vector<int> result2 = solution.twoSum(nums2, target2);

    cout << "Example 2: ";
    for (int index : result2) {
        cout << index << " ";
    }
    cout << endl;

    // Example 3
    vector<int> nums3 = {3, 3};
    int target3 = 6;
    vector<int> result3 = solution.twoSum(nums3, target3);

    cout << "Example 3: ";
    for (int index : result3) {
        cout << index << " ";
    }
    cout << endl;

    return 0;
}
