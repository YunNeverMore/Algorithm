class Solution {
public:
    int findLessCount(vector<vector<int>>& matrix, int val) {
      int i = 0, j = matrix[0].size() - 1, res = 0;
      while (i < matrix.size() && j >= 0) {
        if (matrix[i][j] < val) {
          res += j + 1;
          i++;
        } else {
          j--;
        }
      }
      return res;
    }

    int kthSmallest(vector<vector<int>>& matrix, int k) {
      if (matrix.empty()) return 0;
      int low = matrix[0][0], high = matrix.back().back();
      while (low < high) {
        int mid = low + (high - low + 1) / 2;
        int count = findLessCount(matrix, mid);
        if (k - 1 >= count) {
          low = mid;
        } else {
          high = mid - 1;
        }
      }
      return low;
    }
};
