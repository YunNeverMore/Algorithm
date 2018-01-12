class Solution {
public:
    vector<string> findWords(vector<string>& words) {
        vector<string> rows = {"qwertyuiop", "asdfghjkl", "zxcvbnm"};
        vector<int> dict(26, 0);
        for (int i = 0; i < rows.size(); i++) {
            for (char ch : rows[i]) {
                dict[ch - 'a'] = 1 << i;
            }
        }
        vector<string> res;
        for (string& word : words) {
            int bit = 7;
            for (char ch : word) {
                bit &= dict[tolower(ch) - 'a'];
                if (bit == 0) break;
            }
            if (bit != 0 && bit != 7) {
                res.push_back(word);
            }
        }
        return res;
    }
};
