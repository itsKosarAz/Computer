
public class Factory {

	public static Vehicles getVehicle(String name, int option) {
		if (option == 1) {
			return new Trucks(name);
		}
		else if (option == 2) {
			return new Trains(name);
		}
		else {
			return null;
		}
	}

}
