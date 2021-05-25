class Solution {
public:
    int sumSubarrayMins(vector<int>& A) {
      stack<pair<int, int>> incQue, empty;
      vector<int> leftLen(A.size(), 0);
      for (int i = 0; i < A.size(); i++) {
        while(!incQue.empty() && incQue.top().first >= A[i]) {
          incQue.pop();
        }
        leftLen[i] = incQue.empty() ? i + 1 : i - incQue.top().second;
        incQue.push({A[i], i});
      }

      incQue.swap(empty);
      long res = 0;
      for (int i = A.size() - 1; i >= 0; i--) {
        while(!incQue.empty() && incQue.top().first > A[i]) {
          incQue.pop();
        }
        int rightLen = incQue.empty() ? A.size() - i : incQue.top().second - i;
        res += leftLen[i] * rightLen * A[i];
        incQue.push({A[i], i});
      }
      long magicNum = pow(10, 9) + 7;
      return res % magicNum;
    }
};
