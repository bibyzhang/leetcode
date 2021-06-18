import java.lang.System;

public class DepthFirstSearch {
    private static int vertex; //顶点个数
    private static Graph g;//图的存储

    static boolean found = false;

    public static void main(String[] args){
        vertex = 8;
        g = new Graph(vertex);
        g.addEdge(0,1);
        g.addEdge(0,3);
        g.addEdge(1,4);
        g.addEdge(1,2);
        g.addEdge(2,5);
        g.addEdge(3,4);
        g.addEdge(4,5);
        g.addEdge(4,6);
        g.addEdge(5,7);
        g.addEdge(6,7);
        DFSraverse(0, 6);
    }

    public static void DFSraverse(int s, int t) {
        found = false;
        boolean[] visited = new boolean[vertex];//记录访问状态
        int[] prev = new int[vertex];//记录搜索路劲
        for (int i=0; i<vertex; ++i) {
            prev[i] = -1;
        }
        dfs(s, t, visited, prev);
        print(prev, s, t);
    }

    public static void dfs(int w, int t, boolean[] visited, int[] prev) {
        if(found==true) return;
        visited[w] = true;//记录节点已经被访问
        if(w==t) {//起始顶点等于终止顶点
            found = true;
            return;
        }
        //对每个顶点的链表进行遍历(邻接表存储方式，每个顶点存储指向顶点的链表，无向链表一个边要存储两次)
        for (int i=0; i<g.Adjacency[w].size(); ++i) {
            int q = g.Adjacency[w].get(i);
            if(!visited[q]) {
                prev[q] = w;
                dfs(q, t, visited, prev);
            }
        }
    }

    private static void print(int[] prev, int s, int t) { // 递归打印s->t的路径
        if (prev[t] != -1 && t != s) {
            print(prev, s, prev[t]);
        }
        System.out.print(t + " ");
    }
}