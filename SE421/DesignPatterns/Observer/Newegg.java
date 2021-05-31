public class Newegg implements Observer {
    private Double graphicCardCost;
    private Subject nvidia;

    public Double getGraphicCardCost() {
        return this.graphicCardCost;
    }

    public void setGraphicCardCost(Double graphicCardCost) {
        this.graphicCardCost = graphicCardCost;
    }

    public Newegg(Subject nvidia) {
        this.nvidia = nvidia;

    }

    @Override
    public void update(Double graphicCardCost) {
        this.graphicCardCost = graphicCardCost;
    }

}
