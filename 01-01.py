# Solution for Find if Path Exists in Graph
# Link: https://leetcode.com/problems/find-if-path-exists-in-graph/

from collections import deque
class Solution:
    graph = None
    n = None

    def buildGraph(self, n: int, edges: List[List[int]]):
        self.n = n
        self.graph = [[] for i in range(n)]

        for (u, v) in edges:
            self.graph[u].append(v)
            self.graph[v].append(u)

    def validPath(self, n: int, edges: List[List[int]], source: int, destination: int) -> bool:
        self.buildGraph(n, edges)
        
        queue = deque()
        queue.append(source)
        visited = set()
        visited.add(source)

        while queue: # not empty
            u = queue.popleft()
            if u == destination: #reached the destination 
                return True

            for v in self.graph[u]:
                if v not in visited:
                    queue.append(v)
                    visited.add(v)

        return False
