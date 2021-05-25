/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    /** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
    Solution(ListNode* head) {
      head_ = head;
    }

    /** Returns a random node's value. */
    int getRandom() {
      ListNode* p = head_, *selected = p;
      int k = 0;
      while (p) {
        if (rand() % ++k == 0) {
          selected = p;
        }
        p = p->next;
      }
      return selected->val;
    }
private:
    ListNode* head_;
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution obj = new Solution(head);
 * int param_1 = obj.getRandom();
 */
