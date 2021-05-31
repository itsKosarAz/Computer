import java.util.*;
import java.io.IOException;
import java.io.PrintWriter;
public class SE421_Lab1 {

	public static void main(String[] args) throws IOException {
		double num1;
		double num2;
		
		Scanner input = new Scanner(System.in);
		
		
		System.out.printf("Enter a decimal number: ");
		num1 = input.nextDouble();
		
		System.out.printf("Enter a decimal number: ");
		num2 = input.nextDouble();
		
		System.out.printf("%f + %f = %f\n", num1, num2, num1 + num2);
		
		System.out.printf("Enter a decimal number: ");
		num1 = input.nextDouble();
		
		System.out.printf("Enter a decimal number: ");
		num2 = input.nextDouble();
		
		System.out.printf("%f * %f = %f\n", num1, num2, num1 * num2);
		
		System.out.printf("Enter a decimal number: ");
		num1 = input.nextDouble();
		String num1_text = String.valueOf(num1);
		try (PrintWriter out = new PrintWriter("num1.txt")) {
			out.write(num1_text);
		}
		
		System.out.printf("Enter a decimal number: ");
		num2 = input.nextDouble();
		String num2_text = String.valueOf(num1);
		try (PrintWriter out = new PrintWriter("num2.txt")) {
			out.write(num2_text);
		}
		
		String result_text = String.valueOf(num1 / num2);
		try (PrintWriter out = new PrintWriter("result.txt")) {
			out.write(result_text);
		}
		System.out.printf("%f / %f = %f\n", num1, num2, num1 / num2);
	}

}
