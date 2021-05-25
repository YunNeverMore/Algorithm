class Solution {
public:
    // dp[i][j] = min(nums[i] + dp[i+1][j], nums[j] + dp[i][j - 1])
    bool PredictTheWinner(vector<int>& nums) {
      vector<int> dp(nums.size(), 0);
      for (int i = nums.size() - 1; i >= 0; i--) {
        for (int j = i + 1; j < nums.size(); j++) {
          dp[j] = max(nums[i] - dp[j], nums[j] - (j == i ? 0 : dp[j - 1]));
        }
      }
      return dp.back() >= 0;
    }
};
