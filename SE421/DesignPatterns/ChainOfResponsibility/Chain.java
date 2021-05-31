
public class Chain {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Numbers negativeHandler = new NegativeNumbers();
		Numbers positiveHandler = new PositiveNumbers();
		Numbers zeroHandler = new Zero();
		
		negativeHandler.setNumbers(positiveHandler);
		positiveHandler.setNumbers(zeroHandler);
		negativeHandler.processNumbers(0);
	}

}
