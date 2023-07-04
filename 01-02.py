# Solution for Maximum depth of binary tree
# Link: https://leetcode.com/problems/maximum-depth-of-binary-tree/

from collections import deque
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None, depth=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        queue = deque()
        visited = set()

        if not root:
            return 0

        queue.append((root, 1))
        max_depth = 1

        while queue:
            (u, u_depth) = queue.popleft()

            max_depth = max(max_depth, u_depth)

            if u.left:
                queue.append((u.left, u_depth + 1))
            if u.right:
                queue.append((u.right, u_depth + 1))

        return max_depth
