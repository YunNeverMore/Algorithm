class Solution {
public:
  
    int totalFruit(vector<int>& tree) {
      int count = 0, res = 0;
      unordered_map<int, int> dict;
      for (int i = 0, j = 0; i < tree.size(); i++) {
        count += dict[tree[i]]++ == 0;
        while (count > 2) {
          count -= --dict[tree[j++]] == 0;
        }
        res = max(res, i - j + 1);
      }
      return res;
    }
};
