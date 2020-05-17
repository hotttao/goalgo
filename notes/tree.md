# 树
## 1. 树的抽象
我们都知道树是一种数组组织形式，通过限定树中数据的组织方式，我们可以得到很多树的变种。因此要想学好树，我们就要从最基本的树开始，逐一去了解每种特殊的树的数据组织方式以及他们能提供的操作。

### 1.1 数的组织方式
我们所说的数据组织方式本质上应该包括两个方面:
1. 底层数据的存储方式: 包括链表和数组，因此就有了链表实现的树，与数组实现的树
2. 父子节点的排列方式: 排列方式包括如下几个层次:
    - 每个节点的子节点个数
    - 父子节点之间的大小顺序
    - 每一个节点的子节点之间的大小顺序
    - 树为否为完全二叉树

通常链表是实现树的通用方式，而数组实现的树通常仅限于完全二叉树。父子节点的排列方式决定了树的搜索属性，决定了每种树的特定用途。

### 1.2 树的抽象层次
下面是树的一个类层次结构，接下来我们会一一介绍下面的各种树。
```bash
Tree
|------- BinaryTree
|            | 
|            |------------ ArraryBinaryTree
|            | 
|            |------------ LinkedBinaryTree
```
## 2. 树
### 2,1 Tree
Tree 最普通的树，可以有任意的分叉和孩子数，支持如下操作:
1. `root()`: 返回树的根节点
2. `parent(p)`: 返回节点 p 的父节点
3. `children(p)`: 返回节点 p 的子节点
4. `is_root(p)`: 判断节点 p 是否为根节点
5. `is_leaf(p)`: 判断节点 p 是否是叶节点
5. `is_empty()`: 判断树是否为空
6. `depth(p)`: 计算节点 p 的深度
7. `height(p)`: 计算节点 p 的高度
8. `preorder(p)`: 先序遍历
9. `postorder(p)`: 后序遍历
10. `breadthfirst()`: 层序遍历，又称广度优先遍历

说明: 树的遍历按照父节点被访问的次序分为:
1. 先序遍历: 先访问树的父节点，在访问子节点
2. 后序遍历: 先访问树的子节点，在访问父节点
3. 层序遍历: 按树的层遍历所有节点

```python
class Tree(object):
    
    def depth(self, p):
        """
        :param p:
        :return: 返回 p 节点的深度
        """
        if self.is_root(p):
            return 0
        else:
            return 1 + self.depth(self.parent(p))

    def height(self, p=None):
        """
        :return: 返回树的高度
        """
        p = p or self.root()
        return self._height(p)

    def _height(self, p):
        if self.is_leaf(p):
            return 0
        else:
            return 1 + max(self._height(c) for c in self.children(p))
    def preorder(self):
        """
        :return: 树的前序遍历
        """
        if not self.is_empty():
            for p in self._subtree_preorder(self.root()):
                yield p

    def _subtree_preorder(self, p):
        yield p
        for i in self.children(p):
            for other in self._subtree_preorder(i):
                yield other

    def postorder(self):
        """
        :return: 后序遍历
        """
        if not self.is_empty():
            for p in self._subtree_postorder(self.root()):
                yield p

    def _subtree_postorder(self, p):
        for i in self.children(p):
            for other in self._subtree_preorder(i):
                yield other
        yield p

    def breadthfirst(self):
        """
        :return: 中序遍历
        """
        if not self.is_empty():
            queue = deque()
            queue.append(self.root())
            while len(queue) > 0:
                p = queue.popleft()
                yield p
                for c in self.children(p):
                    queue.append(c)
```

### 2.2 BinaryTree
BinaryTree 二叉树是每个节点最多只有两个分叉的树，他在 Tree 的基础上增加了如下几个操作:
1. `left(p)`: 返回节点 p 的左子节点
2. `right(p)`: 返回节点 p 的右子节点
3. `sibling(p)`: 返回节点 p 的兄弟节点
4. `inorder(p)`: 中序遍历
    - 中序遍历是二叉树特有的遍历方式
    - 节点的访问次序是左子节点->父节点->右子节点

```python
class BinaryTree(Tree):
    def slide(self, p):
        """
        :param p:
        :return: 返回节点的兄弟节点
        """
        parent = self.parent(p)
        if parent is not None:
            left = self.left(parent)
            right = self.right(parent)
            if left == p:
                return right
            else:
                return left

    def children(self, p):
        """
        :param p:
        :return: 返回节点的所有子节点
        """
        left = self.left(p)
        if p is not None:
            yield left
        right = self.right(p)
        if right is not None:
            yield right
    
    def inorder(self):
        """
        :return: 中序遍历
        """
        if not self.is_empty():
            return self._subtree_inorder(self.root())

    def _subtree_inorder(self, p):
        left = self.left(p)
        if left is not None:
            for other in self._subtree_inorder(left):
                yield other
        yield p
        right = self.right(p)
        if right is not None:
            for other in self._subtree_inorder(right):
                yield other
```

## 3. 树遍历的抽象-欧拉图和模板方法
前面我们实现了树的四种遍历方式，以按照不同的顺序获取树中的元素。但是这提供的抽象能力还不够，更多时候，我们需要获取树遍历过程中的更多信息，比如当前位置的深度，或者从根节点到当前位置的完整路径，或者返回下一层信息到上一层。因此我们需要一个更通用的框架，即基于概念实现树的遍历--欧拉遍历。

什么是欧拉遍历，我们来看下面的伪代码:
```python
Algorithm eulertour(T, p)
    pre visit for p                         # 第一访问节点 p，前序遍历可执行的操作位于此处
    for each child c in T.children(p) do
        eulertour(T, c)
    post visit for p                        # 第二次访问节点 p，后续遍历可执行的操作位于此处
```

显然我们可以把 eulertour 定义为模板方法，让继承自 eulertour 的子类去实现需要的前序后续遍历需要执行的操作。下面是 eulertour 的具体实现:

```python
class EulerTour(object):
    # 一般欧拉遍历的实现
    def __init__(self, tree):
        self._tree = tree

    def execute(self):
        if not self._tree.is_empty():
            return self._tour(self._tree.root(), 0, [])

    def _tour(self, p, d, path):
        """
        :param p: 当前遍历的节点
        :param d: 节点所处的树的深度
        :param path: 从根到达当前节点的路径
        :return: 后续遍历的返回值
        """
        self._hook_previsit(p, d, path) # 先序遍历需实现的抽象方法
        path.append(0)                  # path 最后一个索引记录了，当前节点所有子节点的排序
        result = []
        for c in self._tree.children(p):
            result.append(self._tour(c, d + 1, path))
            path[-1] += 1
        value = self._hook_postvisit(p, d, path, result)
        path.pop()
        return value

    def _hook_previsit(self, p, d, path):
        pass

    def _hook_postvisit(self, p, d, path, result):
        """
        :param p: 当前遍历的节点
        :param d: 节点所处的树的深度
        :param path: 从根到达当前节点的路径
        :param result: 子节点后续遍历返回值的列表
        :return: 后续遍历的返回值
        """


class BinaryEulerTour(BinaryEulerTour):
    # 二叉树的欧拉遍历
    def __init__(self, tree):
        super(BinaryEulerTour, self).__init__(tree)

    def _tour(self, p, d, path):
        self._hook_previsit(p, d, path)
        result = [None, None]
        if self._tree.left(p):
            path.append(0)
            result[0] = self._tour(self._tree.left(p), d + 1, path)
            path.pop()
        
        self._hook_invisit(p, d, path)
        
        if self._tree.right(p):
            path.append(1)
            result[1] = self._tour(self._tree.right(p), d + 1, path)
            path.pop()

        value = self._hook_postvisit(p, d, path, result)
        return value
    
    def _hook_invisit(self, p, d, path):
        pass
```

利用 BinaryEulerTour 我们可以开发一个用于计算二叉树的图形布局的子类，该算法用以下两条规则为二叉树的每个节点指定 x，坐标：
1. x(p): 在节点 p 之前，中序遍历访问的节点数量
2. y(p): 是 T 中 p 的深度

```python
class BinaryLayout(BinaryEulerTour)
    def __init__(self, tree):
        super().__init__(tree)
        self._count = 0
    
    def _hook_invisit(self, p, d, path):
        p.element().setX(self._count)
        p.element().setY(d)
        self._count+=1
```
除了 BinaryEulerTour 提供的钩子函数外，我们以 _count 实例变量的形式引入了额外的状态，从而调整了 BinaryEulerTour 框架，扩展了框架提供的功能。