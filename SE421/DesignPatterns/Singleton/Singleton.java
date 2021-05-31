
public class Singleton {
	private static Singleton instance;
	
	private Singleton() {
		
	}

	public static Singleton getInstance() {
		// TODO Auto-generated method stub
		if (instance == null) {
			instance = new Singleton();
			return instance;
		}
		else {
		return instance;
		}
	}

}
