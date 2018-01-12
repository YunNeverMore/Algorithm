class Solution {
public:
    // // O(nlogn) sort version
    // int arrayPairSum(vector<int>& nums) {
    //     int total = 0;
    //     std::sort(nums.begin(), nums.end());
    //     for (int i = 0; i < nums.size(); i++) {
    //         if (i % 2 == 0) total += nums[i];
    //     }
    //     return total;
    // }
    
    // O(n) bucket sort
    int arrayPairSum(vector<int>& nums) {
        vector<int> bucket(20001, 0);
        for (int num : nums) {
            bucket[num + 10000]++;
        }
        int total = 0, count = 0;
        for (int i = 0; i < 20001; i++) {
            total += (bucket[i] + (count % 2 == 0)) / 2 * (i - 10000);
            count += bucket[i];
        }
        return total;
    }
};
