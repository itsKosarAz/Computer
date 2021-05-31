
public class Clients {
	private Service currentService;
	private String clientName;
	
	public Clients(Service currentService, String clientName) {
		this.currentService = currentService;
		this.clientName = clientName;
	}
	
	public void setService(Service currentService) {
		this.currentService = currentService;
	}

}
