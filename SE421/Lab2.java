import java.util.*;
import java.io.IOException;
import java.io.PrintWriter;
public class SE421_Lab2 {
	
	public static void main(String[] args) throws IOException {
		double num1 = getnum1();
		double num2 = getnum2();
		double result = addition(num1, num2);
		print(result);
		
		num1 = getnum1();
		num2 = getnum2();
		result = multiplication(num1, num2);
		print(result);
		
		num1 = getnum1();
		txtnum1(num1);
		
		num2 = getnum2();
		txtnum2(num2);
		
		result = division(num1, num2);
		result(result);
		
		print(result);
	}
	
	public static double getnum1() {
		Scanner input = new Scanner(System.in);
		System.out.printf("Enter a decimal number: ");
		double num1 = input.nextDouble();
		return num1;
	}
	
	public static double getnum2() {
		Scanner input = new Scanner(System.in);
		System.out.printf("Enter a decimal number: ");
		double num2 = input.nextDouble();
		return num2;
	}
	
	public static double addition(double num1, double num2) {
		double result = num1 + num2;
		return result;
	}
	
	public static double multiplication(double num1, double num2) {
		double result = num1 * num2;
		return result;
	}
	
	public static double division(double num1, double num2) {
		double result = num1 / num2;
		return result;
	}
	
	public static void print(double result) {
		System.out.printf("%f\n", result);
	}
	
	public static void txtnum1(double num1) throws IOException {
		String num1_text = String.valueOf(num1);
		try (PrintWriter out = new PrintWriter("num1.txt")) {
			out.write(num1_text);
		}
	}
	
	public static void txtnum2(double num2) throws IOException {
		String num2_text = String.valueOf(num2);
		try (PrintWriter out = new PrintWriter("num2.txt")) {
			out.write(num2_text);
		}
	}
	
	public static void result(double result) throws IOException {
		String result_text = String.valueOf(result);
		try (PrintWriter out = new PrintWriter("result.txt")) {
			out.write(result_text);
		}
	}
}
