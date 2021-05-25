class Solution {
public:
    int findKthLargest2(vector<int>& nums, int k) {
      if (!k || nums.size() < k) return 0;
      priority_queue<int, vector<int>, std::greater<int>> pq;
      for (auto num : nums) {
        if (pq.size() < k) {
          pq.push(num);
        } else if (pq.top() < num) {
          pq.pop();
          pq.push(num);
        }
      }
      return pq.empty() ? 0: pq.top();
    }

    int twoWayPartition(vector<int>& nums, int start, int end) {
      int i = start + 1, j = end, target = nums[start];
      while (i <= j) {
        if (nums[i] < target && nums[j] >= target) {
          swap(nums[i++], nums[j--]);
        } else {
          if (nums[i] >= target) i++;
          if (nums[j] < target) j--;
        }
      }
      swap(nums[start], nums[j]);
      return j;
    }

    int threeWayPartition(vector<int>& nums, int start, int end) {
      int i = start, k = i, j = end, target = nums[start];
      while (i <= j) {
        if (nums[i] < target) {
          swap(nums[i], nums[j--]);
        } else if (nums[i] > target) {
          swap(nums[i++], nums[k++]);
        } else i++;
      }
      return j;
    }

    int findKthLargest(vector<int>& nums, int k) {
      int low = 0, high = nums.size() - 1;
      while (low < high) {
        int mid = threeWayPartition(nums, low, high);
        if (mid < k - 1) {
          low = mid + 1;
        } else if (mid > k - 1){
          high = mid - 1;
        } else {
          return nums[mid];
        }
      }
      return nums[low];
    }
};
