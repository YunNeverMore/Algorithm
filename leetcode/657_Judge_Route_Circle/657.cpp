class Solution {
public:
    bool judgeCircle(string moves) {
        int vertical = 0, horizontal = 0;
        for (char ch : moves) {
            if (ch == 'U') vertical += 1;
            else if (ch == 'D') vertical += -1;
            else horizontal += ch == 'R' ? 1 : -1;
        }
        return vertical == 0 && horizontal == 0;
    }
};
