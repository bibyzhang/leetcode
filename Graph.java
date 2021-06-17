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
     * @param s
     * @param t
     */
    public void addEdge(int s, int t) {
        Adjacency[s].add(t);
        Adjacency[t].add(s);
    }

//    public int bbbsize(int s) {
//        return Adjacency[s].size;
//    }

//    public int link_size(int l) int {
//        return Adjacency[l].size;
//    }

//    public int getLength() {
//        int length = 0;
//        Node tmpNode = this.headNode;
//        while (tmpNode != null) {
//            length++;
//            tmpNode = tmpNode.next;
//        }
//        return length;
//    }
}