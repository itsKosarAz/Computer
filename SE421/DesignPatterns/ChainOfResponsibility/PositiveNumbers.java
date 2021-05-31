
public class PositiveNumbers implements Numbers {
	private Numbers next;

	public static void main(String[] args) {
		// TODO Auto-generated method stub

	}

	@Override
	public void processNumbers(int num) {
		// TODO Auto-generated method stub
		if (num > 0) {
			System.out.println("The number is positive");
		} else {
			if (next != null) {
				next.processNumbers(num);
			}
		}
	}

	@Override
	public void setNumbers(Numbers next) {
		// TODO Auto-generated method stub
		this.next = next;
	}

}
