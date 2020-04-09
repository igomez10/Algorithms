# There are n servers numbered from 0 to n-1 connected by undirected server-to-server connections forming a
# network where connections[i] = [a, b] represents a connection between servers a and b.
# Any server can reach any other server directly or indirectly through the network.
# A critical connection is a connection that, if removed, will make some server unable to reach some other
# server.
# Return all critical connections in the network in any order.


# Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
# Output: [[1,3]]
# Explanation: [[3,1]] is also accepted.

# IDEA: check removing each connection and
# run dfs from the first connection, if the number of reached nodes is different than n then we return

import collections


class Solution:

    def parseConnections(self, connections, graph):

        for i in range(len(connections)):
            if connections[i][0] not in graph:
                graph[connections[i][0]] = [connections[i][1]]

            if connections[i][1] not in graph:
                graph[connections[i][1]] = [connections[i][0]]

            if connections[i][0] not in graph[connections[i][1]]:
                graph[connections[i][1]].append(connections[i][0])

            if connections[i][1] not in graph[connections[i][0]]:
                graph[connections[i][0]].append(connections[i][1])

    def bfs(self, graph, root, mc):  # missingConnection = [x,y]
        queue = collections.deque()
        queue.append(root)
        seen = {}
        while len(queue) > 0:
            curr = queue.pop()
            seen[curr] = True  # visited node
            # add all neighbors to queue
            for i in range(len(graph[curr])):
                if graph[curr][i] in mc and curr in mc:
                    pass

                elif graph[curr][i] not in seen:
                    # ADD NODE
                    queue.append(graph[curr][i])
                    seen[graph[curr][i]] = False

        return len(seen)

    def criticalConnections(self, n, connections):
        graph = {}
        self.parseConnections(connections, graph)

        expectedNodes = self.bfs(graph, connections[0][0], [-1, -1])

        criticalConnections = []
        for i in range(len(connections)):
            receivedNodes = self.bfs(graph, connections[0][0], connections[i])
            if receivedNodes != expectedNodes:
                criticalConnections.append(connections[i])
        return criticalConnections
        # graph= {
        #
        # "4" : [1,2,5,7,8]
        # "5" : [4]
        # }


s = Solution()
n = 4,
connections = [[0, 1], [1, 2], [2, 0], [1, 3]]

print(s.criticalConnections(4, connections))
