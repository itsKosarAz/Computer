import java.util.ArrayList;

public class Nvidia implements Subject {
    private Double graphicCardCost;
    private ArrayList<Observer> Users;

    public ArrayList<Observer> getUsers() {
        return this.Users;
    }

    public Nvidia() {
        Users = new ArrayList<Observer>();
    }

    public Double getGraphicCardCost() {
        return this.graphicCardCost;
    }

    public void setGraphicCardCost(Double graphicCardCost) {
        this.graphicCardCost = graphicCardCost;
        notifyObserver();
    }

    @Override
    public void register(Observer o) {
        this.Users.add(o);
    }

    @Override
    public void unregister(Observer o) {
        int index = this.Users.indexOf(o);
        System.out.println("User" + index + 1 + "is removed");
        this.Users.remove(index);
    }

    @Override
    public void notifyObserver() {
        for (Observer o : Users) {
            o.update(graphicCardCost);
        }
    }

}
