  
public class Amazon implements Observer {
    private Double graphicCardCost;

    public Double getGraphicCardCost() {
        return this.graphicCardCost;
    }

    public void setGraphicCardCost(Double graphicCardCost) {
        this.graphicCardCost = graphicCardCost;
    }

    private Subject nvidia;

    public Amazon(Subject nvidia) {
        this.nvidia = nvidia;
    }

    @Override
    public void update(Double graphicCardCost) {
        this.graphicCardCost = graphicCardCost;
    }

}
