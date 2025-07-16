一開始就知道 maxSum 的合法範圍，不用再找一個「極小值」去初始化。

後面只要做兩步——sum += nums[i] - nums[i-k]、if sum>maxSum——程式簡潔、易讀。
