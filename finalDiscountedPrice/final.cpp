// 补充一种 c++ 的解法，毕竟是从 c++ 参考改过来的

class Solution {
public:
    vector<int> FindalDiscountedPrice(vector<int> &prices) {
        int n = prices.size();
        vector<int> res(n);
        stack<int> s;

        for(int i=0; i<n; i++) res[i] = prices[i];

        for(int i=0; i<n; i++) {
            while(!s.empty() && prices[s.top()] >= prices[i]) {
                int index = s.top();
                s.pop();
                res[index] = prices[index] - prices[i];

            }
            s.push(i);
        }
        return res;
    }
};