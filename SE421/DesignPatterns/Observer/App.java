
public class App {
    public static void main(String[] args) {
        Nvidia nv = new Nvidia();
        Amazon user1 = new Amazon(nv);
        Amazon user2 = new Amazon(nv);
        Newegg user3 = new Newegg(nv);
        Newegg user4 = new Newegg(nv);
        nv.register(user1);
        nv.register(user2);
        nv.register(user3);
        nv.register(user4);
        System.out.println(user1.getGraphicCardCost());
        System.out.println(user2.getGraphicCardCost());
        System.out.println(user3.getGraphicCardCost());
        System.out.println(user4.getGraphicCardCost());
        nv.setGraphicCardCost(50.4);
        System.out.println(user1.getGraphicCardCost());
        System.out.println(user2.getGraphicCardCost());
        System.out.println(user3.getGraphicCardCost());
        System.out.println(user4.getGraphicCardCost());
    }
}
