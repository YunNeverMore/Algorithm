class Solution {
public:
    bool is_integer(string str) {
        if (!str.empty() && str[0] == '-') str = str.substr(1);
        for (char ch : str) {
            if (!isdigit(ch)) return false;
        }
        return true;
    }

    int calPoints(vector<string>& ops) {
        int total = 0;
        vector<int> st;
        for (string& str : ops) {
            int cur_num = 0;
            if (str == "C" && st.size() > 0) {
                total -= st.back();
                st.pop_back();
                continue;
            } else if (str == "+") {
                cur_num = st[st.size() - 1] + st[st.size() - 2];
            } else if (str == "D") {
                cur_num = st.back() * 2;
            } else if (is_integer(str)) {
                cur_num = stoi(str);
            }
            st.push_back(cur_num);
            total += cur_num;
        }
        return total;
    }
};
