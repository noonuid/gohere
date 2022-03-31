# 回溯

思路
如果对回溯算法基础还不了解的话，我还特意录制了一期视频：带你学透回溯算法（理论篇） 可以结合题解和视频一起看，希望对大家理解回溯算法有所帮助。

此时我们已经学习了77.组合问题、 131.分割回文串和78.子集问题，接下来看一看排列问题。

相信这个排列问题就算是让你用for循环暴力把结果搜索出来，这个暴力也不是很好写。

所以正如我们在关于回溯算法，你该了解这些！所讲的为什么回溯法是暴力搜索，效率这么低，还要用它？

因为一些问题能暴力搜出来就已经很不错了！

我以[1,2,3]为例，抽象成树形结构如下：

![46.全排列](https://pic.leetcode-cn.com/1631607821-lOapRO-file_1631607821406)

回溯三部曲
递归函数参数
首先排列是有序的，也就是说[1,2] 和[2,1] 是两个集合，这和之前分析的子集以及组合所不同的地方。

可以看出元素1在[1,2]中已经使用过了，但是在[2,1]中还要在使用一次1，所以处理排列问题就不用使用startIndex了。

但排列问题需要一个used数组，标记已经选择的元素，如图橘黄色部分所示:

![46.全排列](https://pic.leetcode-cn.com/1631607821-lOapRO-file_1631607821406)

代码如下：

```
vector<vector<int>> result;
vector<int> path;
void backtracking (vector<int>& nums, vector<bool>& used)
```

递归终止条件

![46.全排列](https://pic.leetcode-cn.com/1631607821-lOapRO-file_1631607821406)


可以看出叶子节点，就是收割结果的地方。

那么什么时候，算是到达叶子节点呢？

当收集元素的数组path的大小达到和nums数组一样大的时候，说明找到了一个全排列，也表示到达了叶子节点。

代码如下：

// 此时说明找到了一组

```c++
if (path.size() == nums.size()) {
    result.push_back(path);
    return;
}
```

单层搜索的逻辑
这里和77.组合问题、131.切割问题和78.子集问题最大的不同就是for循环里不用startIndex了。

因为排列问题，每次都要从头开始搜索，例如元素1在[1,2]中已经使用过了，但是在[2,1]中还要再使用一次1。

而used数组，其实就是记录此时path里都有哪些元素使用了，一个排列里一个元素只能使用一次。

代码如下：

```c++
for (int i = 0; i < nums.size(); i++) {
    if (used[i] == true) continue; // path里已经收录的元素，直接跳过
    used[i] = true;
    path.push_back(nums[i]);
    backtracking(nums, used);
    path.pop_back();
    used[i] = false;
}
```


整体C++代码如下：

```c++
class Solution {
public:
    vector<vector<int>> result;
    vector<int> path;
    void backtracking (vector<int>& nums, vector<bool>& used) {
        // 此时说明找到了一组
        if (path.size() == nums.size()) {
            result.push_back(path);
            return;
        }
        for (int i = 0; i < nums.size(); i++) {
            if (used[i] == true) continue; // path里已经收录的元素，直接跳过
            used[i] = true;
            path.push_back(nums[i]);
            backtracking(nums, used);
            path.pop_back();
            used[i] = false;
        }
    }
    vector<vector<int>> permute(vector<int>& nums) {
        result.clear();
        path.clear();
        vector<bool> used(nums.size(), false);
        backtracking(nums, used);
        return result;
    }
};
```

总结
大家此时可以感受出排列问题的不同：

每层都是从0开始搜索而不是startIndex
需要used数组记录path里都放了哪些元素了
排列问题是回溯算法解决的经典题目，大家可以好好体会体会。

其他语言版本
Java：

```java
class Solution {

List<List<Integer>> result = new ArrayList<>();// 存放符合条件结果的集合
LinkedList<Integer> path = new LinkedList<>();// 用来存放符合条件结果
boolean[] used;
public List<List<Integer>> permute(int[] nums) {
    if (nums.length == 0){
        return result;
    }
    used = new boolean[nums.length];
    permuteHelper(nums);
    return result;
}

private void permuteHelper(int[] nums){
    if (path.size() == nums.length){
        result.add(new ArrayList<>(path));
        return;
    }
    for (int i = 0; i < nums.length; i++){
        if (used[i]){
            continue;
        }
        used[i] = true;
        path.add(nums[i]);
        permuteHelper(nums);
        path.removeLast();
        used[i] = false;
    }
}

}
```

// 解法2：通过判断path中是否存在数字，排除已经选择的数字

```java
class Solution {
    List<List<Integer>> result = new ArrayList<>();
    LinkedList<Integer> path = new LinkedList<>();
    public List<List<Integer>> permute(int[] nums) {
        if (nums.length == 0) return result;
        backtrack(nums, path);
        return result;
    }
    public void backtrack(int[] nums, LinkedList<Integer> path) {
        if (path.size() == nums.length) {
            result.add(new ArrayList<>(path));
        }
        for (int i =0; i < nums.length; i++) {
            // 如果path中已有，则跳过
            if (path.contains(nums[i])) {
                continue;
            } 
            path.add(nums[i]);
            backtrack(nums, path);
            path.removeLast();
        }
    }
}
```

Python：

```python
class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []  #存放符合条件结果的集合
        path = []  #用来存放符合条件的结果
        used = []  #用来存放已经用过的数字
        def backtrack(nums,used):
            if len(path) == len(nums):
                return res.append(path[:])  #此时说明找到了一组
            for i in range(0,len(nums)):
                if nums[i] in used:
                    continue  #used里已经收录的元素，直接跳过
                path.append(nums[i])
                used.append(nums[i])
                backtrack(nums,used)
                used.pop()
                path.pop()
        backtrack(nums,used)
        return res
```


Python（优化，不用used数组）：

```python
class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []  #存放符合条件结果的集合
        path = []  #用来存放符合条件的结果
        def backtrack(nums):
            if len(path) == len(nums):
                return res.append(path[:])  #此时说明找到了一组
            for i in range(0,len(nums)):
                if nums[i] in path:  #path里已经收录的元素，直接跳过
                    continue
                path.append(nums[i])
                backtrack(nums)  #递归
                path.pop()  #回溯
        backtrack(nums)
        return res
```


Go：

```go
var res [][]int
func permute(nums []int) [][]int {
	res = [][]int{}
	backTrack(nums,len(nums),[]int{})
	return res
}
func backTrack(nums []int,numsLen int,path []int)  {
	if len(nums)==0{
		p:=make([]int,len(path))
		copy(p,path)
		res = append(res,p)
	}
	for i:=0;i<numsLen;i++{
		cur:=nums[i]
		path = append(path,cur)
		nums = append(nums[:i],nums[i+1:]...)//直接使用切片
		backTrack(nums,len(nums),path)
		nums = append(nums[:i],append([]int{cur},nums[i:]...)...)//回溯的时候切片也要复原，元素位置不能变
		path = path[:len(path)-1]
	}
}
```

Javascript:

```javascript

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permute = function(nums) {
    const res = [], path = [];
    backtracking(nums, nums.length, []);
    return res;
    
    function backtracking(n, k, used) {
        if(path.length === k) {
            res.push(Array.from(path));
            return;
        }
        for (let i = 0; i < k; i++ ) {
            if(used[i]) continue;
            path.push(n[i]);
            used[i] = true; // 同支
            backtracking(n, k, used);
            path.pop();
            used[i] = false;
        }
    }
};
```

回溯算法力扣题目总结
按照如下顺序刷力扣上的题目，相信会帮你在学习回溯算法的路上少走很多弯路。

关于回溯算法，你该了解这些！
组合问题
77.组合
216.组合总和III
17.电话号码的字母组合
39.组合总和
40.组合总和II
分割问题
131.分割回文串
93.复原IP地址
子集问题
78.子集
90.子集II
排列问题
46.全排列
47.全排列II
棋盘问题
51.N皇后
37.解数独
其他
491.递增子序列
332.重新安排行程
回溯算法总结篇
大家好，我是程序员Carl，点击我的头像，查看力扣详细刷题攻略，你会发现相见恨晚！

如果感觉题解对你有帮助，不要吝啬给一个👍吧！

作者：carlsun-2
链接：https://leetcode-cn.com/problems/permutations/solution/dai-ma-sui-xiang-lu-dai-ni-xue-tou-hui-s-mfrp/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。