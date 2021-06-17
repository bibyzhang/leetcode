import java.lang.System;

public class DepthFirstSearch {
    private static int vertex; //顶点个数
    private static Graph g;

    static boolean found = false;

    public static void main(String[] args){
        vertex = 8;
        g = new Graph(vertex);
        g.addEdge(1-1,2-1);
        g.addEdge(1-1,4-1);
        g.addEdge(2-1,5-1);
        g.addEdge(2-1,3-1);
        g.addEdge(3-1,6-1);
        g.addEdge(4-1,5-1);
        g.addEdge(5-1,6-1);
        g.addEdge(5-1,7-1);
        g.addEdge(6-1,8-1);
        g.addEdge(7-1,8-1);
        DFSraverse(1-1, 7-1);
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
//        System.out.print(g.Adjacency[0].size());
//        System.out.print(g.bbbsize[0]);
//        System.out.print(g.Adjacency[w].size);
//        System.out.print(g.link_size(w));
//        print(g.Adjacency[w], w, t);
        //对每个顶点的链表进行遍历
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