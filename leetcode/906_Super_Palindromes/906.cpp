class Solution {
public:
    bool isPalidom(string str) {
      for (int i = 0, j = str.size() - 1; i < j; i++, j--) {
        if (str[i] != str[j]) return false;
      }
      return true;
    }

    void backtrack(string& cur, long l, long r, int& res) {
      if (cur.size() > 9) return;

      if (!cur.empty() && cur[0] != '0') {
        long curVal = stol(cur);
        long powVal = curVal * curVal;
        if (powVal > r) return;
        if (powVal >= l && isPalidom(to_string(powVal))) res++;
      }

      for (char ch = '0'; ch <= '9'; ch++) {
        cur = ch + cur + ch;
        backtrack(cur, l, r, res);
        cur.erase(cur.begin());
        cur.pop_back();
      }
    }

    int getDigits(long val) {
      return val == 0 ? 0 : 1 + getDigits(val / 10);
    }

    int superpalindromesInRange(string L, string R) {
      long l = stol(L), r = stol(R);
      // int ld = getDigits(sqrt(l)), rd = getDigits(sqrt(r));
      int res = 0;
      string str = "";
      backtrack(str, l, r, res);
      for (char ch = '0'; ch <= '9'; ch++) {
        str = string(1, ch);
        backtrack(str, l, r, res);
      }
      return res;
    }
};
