import java.util.LinkedList;

//存储结构无向图
public class Graph {
    private int vertex; //顶点个数
    public LinkedList<Integer> Adjacency[];//邻接表

    /**
     * 邻接表方式存储
     * @param vertex int 顶点个数
     */
    public Graph(int vertex) {
        this.vertex = vertex;
        Adjacency = new LinkedList[vertex];
        for (int i=0; i<vertex; ++i) {
            Adjacency[i] = new LinkedList<>();
        }
    }

    /**
     * 无向图的一条边存两次
     * @param s int 起始顶点
     * @param t int 终止顶点
     */
    public void addEdge(int s, int t) {
        Adjacency[s].add(t);
        Adjacency[t].add(s);
    }
}