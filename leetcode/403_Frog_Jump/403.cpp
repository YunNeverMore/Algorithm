class Solution {
public:

    bool canCross2(vector<int>& stones) {
      if (stones.empty()) return false;
      vector<vector<bool>> dp(stones.size(), vector<bool>(stones.size(), false));
      dp[0][1] = true;
      unordered_map<int, int> indexs;
      for (int i = 0; i < stones.size(); i++)
        indexs[stones[i]] = i;

      for (int i = 1; i < stones.size(); i++) {
        for (int k = 1; k < stones.size(); k++) {
          for (int step = k - 1; step <= k + 1; step++) {
            if (step > 0 && indexs.count(stones[i] - step))
              dp[i][k] = dp[i][k] || dp[indexs[stones[i] - step]][step];
          }
          if (i == stones.size() - 1 && dp[i][k]) return true;
        }
      }
      return false;
    }

    bool canCross(vector<int>& stones) {
      if (stones.empty()) return false;
      unordered_map<int, unordered_set<int>> indexs;
      for (int i = 0; i < stones.size(); i++)
        indexs[stones[i]] = unordered_set<int>();
      indexs[0].insert(0);

      for (int i = 0; i < stones.size(); i++) {
        for (int k : indexs[stones[i]]) {
          for (int step = k - 1; step <= k + 1; step++) {
            if (step > 0 && indexs.count(stones[i] + step)) {
                indexs[stones[i] + step].insert(step);
            }
          }
        }
      }
      return !indexs[stones.back()].empty();
    }
};
