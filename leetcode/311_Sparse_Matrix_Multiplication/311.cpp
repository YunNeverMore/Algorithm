class Solution {
public:
    int rowMultiCol(vector<pair<int, int>>& row, vector<pair<int, int>>& col) {
      int i = 0, j = 0, total = 0;
      while (i < row.size() && j < col.size()) {
        if (row[i].first == col[j].first) {
          total += row[i++].second * col[j++].second;
        } else if (row[i].first < col[j].first) {
          i++;
        } else {
          j++;
        }
      }
      return total;
    }

    vector<vector<int>> multiply2(vector<vector<int>>& A, vector<vector<int>>& B) {
      if (B.empty()) return {};
      vector<vector<pair<int, int>>> processedA(A.size()), processedB(B[0].size());
      for (int i = 0; i < A.size(); i++) {
        for (int j = 0; j < A[i].size(); j++) {
          if (A[i][j]) processedA[i].push_back({j, A[i][j]});
        }
      }
      for (int i = 0; i < B.size(); i++) {
        for (int j = 0; j < B[i].size(); j++) {
          if (B[i][j]) processedB[j].push_back({i, B[i][j]});
        }
      }
      vector<vector<int>> multipied(processedA.size(), vector<int>(processedB.size(), 0));
      for (int i = 0; i < processedA.size(); i++) {
        for (int k = 0; k < processedB.size(); k++) {
          multipied[i][k] = rowMultiCol(processedA[i], processedB[k]);
        }
      }
      return multipied;
    }


   vector<vector<int>> multiply(vector<vector<int>>& A, vector<vector<int>>& B) {
     if (B.empty()) return {};
     int m = A.size(), n = A[0].size(), nB = B[0].size();
     vector<vector<int>> C(m, vector<int>(nB, 0));
     for (int i = 0; i < m; i++) {
       for (int j = 0; j < n; j++) {
         if (!A[i][j]) continue;
         for (int k = 0; k < nB; k++) {
           if (B[j][k]) C[i][k] += A[i][j] * B[j][k];
         }
       }
     }
     return C;
   }
};
