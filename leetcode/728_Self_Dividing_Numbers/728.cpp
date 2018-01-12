class Solution {
public:
    vector<int> selfDividingNumbers(int left, int right) {
        vector<int> res;
        for (int num = left; num <= right; num++) {
            int remain = num;
            while (remain != 0) {
                int digit = remain % 10;
                if (digit == 0 || num % digit != 0) break;
                remain = remain / 10;
            }
            if (remain == 0) res.push_back(num);
        }
        return res;
    }
};
