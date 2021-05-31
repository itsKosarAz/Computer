
public class Print {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Vehicles firstVehicle = Factory.getVehicle("ford", 1);
		
		firstVehicle.setName("chevy");
		
		System.out.println(firstVehicle.getName());
	}

}
