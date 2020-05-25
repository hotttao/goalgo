# 编程思想
## 1. 原理
基础的数据结构与算法中，有几块非常难懂，贪心，分治，回溯和动态规划这四个编程思想应该算是"名列前茅"了。本文希望通过详细解答几个经典示例来帮助大家搞懂他们。在进入实战之前，我们先来看看他们的区别。

### 1.1 动态规划
动态规划适合解决的问题可以概括为“一个模型三个特征”。
1. 一个模型: **多阶段决策最优解模型**。动态规划通常被用来解决最优问题，而解决问题的过程，需要经历多个决策阶段。每个决策阶段都对应着一组状态。然后我们寻找一组决策序列，经过这组决策序列，能够产生最终期望求解的最优值。
2. 三个特征:
    - 最优子结构: 我们可以通过子问题的最优解，推导出问题的最优解
    - 无后效性: 某阶段状态一旦确定，就不受之后阶段的决策影响
    - 重复子问题: 不同的决策序列，到达某个相同的阶段时，可能会产生重复的状态

了解了动态，接下我们就以动态为标杆，看看其他编程思想有什么不同

### 1.2 分治
贪心、回溯、动态规划可以归为一类，都可以抽象成我们今天讲的那个多阶段决策最优解模型，而分治算法解决的问题尽管大部分也是最优解问题，但是，大部分都不能抽象成多阶段决策模型。

在重复子问题这一点上，动态规划和分治算法的区分非常明显。分治算法要求分割成的子问题，不能有重复子问题，而动态规划正好相反，动态规划之所以高效，就是因为回溯算法实现中存在大量的重复子问题。

### 1.3 贪心
它能解决的问题需要满足三个条件，最优子结构、无后效性和贪心选择性。**贪心选择性**的意思是，通过局部最优的选择，能产生全局的最优选择。每一个阶段，我们都选择当前看起来最优的决策，所有阶段的决策完成之后，最终由这些局部最优解构成全局最优解。贪心算法实际上是动态规划算法的一种特殊情况。

### 1.4 回溯
回溯算法是个“万金油”。基本上能用的动态规划、贪心解决的问题，我们都可以用回溯算法解决。回溯算法相当于穷举搜索。不过，回溯算法的时间复杂度非常高，是指数级别的，只能用来解决小规模数据的问题。

它们之间的区别讲完了，接下来我们就来看看如何用它们来解决我们的编程问题。

## 2. 实战
### 2.1 最少加油次数
[leecode 871](https://leetcode-cn.com/problems/minimum-number-of-refueling-stops/)

#### 回溯
```python
class Solution(object):
    def minRefuelStops(self, target, startFuel, stations):
        """
        :type target: int
        :type startFuel: int
        :type stations: List[List[int]]
        :rtype: int
        """
        self.min_station = len(stations) + 1
        if startFuel >= target:
            return 0
        self.refuel_stop(target, stations, 0, startFuel, 0)
        self.min_station = self.min_station \
            if self.min_station <= len(stations) else -1
        return self.min_station


    def refuel_stop(self, target, station, i, fuel, c):
        """
        :param target: 距离目标地点还有多远
        :param station:
        :param i: 表示到达第 i 个加油站，i 从 0 开始计数
        :param fuel: 能走多远
        :param c: 加油次数
        :return:
        """
        # 已经到最有一个加油站了
        if i == len(station):
            if fuel >= target:
                self.min_station = min(c, self.min_station)
            return

        # 剩下的油走不到下一站了，也无法到达 target
        if fuel < min(station[i][0], target):
            return

        # 已经足够达到终点
        if fuel >= target:
            self.min_station = min(c, self.min_station)
            return
        # 没到达终点，油够到下一站
        # 不加油
        self.refuel_stop(target, station, i + 1, fuel, c)
        # 加油
        if fuel >= station[i][0]:
            self.refuel_stop(target, station, i + 1, fuel + station[i][1], c + 1)
```

#### 动态规划
使用动态规划需要我们稍微专变一下思路: 我们需要计算出第 i 次油能走的最远距离dp[i]，然后找出满足 dp[i] >= target 的最小 i。如果我们画出上面的递归调用图:

我们如果仔细思考上面的回溯代码，就可以发现，第 i + 1 次走的最远距离为第 t 次 (0 <= t <= i)能走的最远距离加上第 i 个加油站能添加的油。

```python
class Solution(object):
    def minRefuelStops(self, target, startFuel, stations):
        dp = [startFuel] + [0] * len(stations)
        for i, (location, capacity) in enumerate(stations):
            for t in xrange(i, -1, -1):
                if dp[t] >= location:
                    dp[t+1] = max(dp[t+1], dp[t] + capacity)

        for i, d in enumerate(dp):
            if d >= target: return i
        return -1
```