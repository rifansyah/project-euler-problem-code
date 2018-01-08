import java.math.BigInteger;

/**
 *
 * @author rifansyah
 */
public class Problem20 {

    public static void main(String[] args) {
        String factorial = getFactorial(100);
        int sumOfFactorial = getSumOfFactorialDigits(factorial);
        
        System.out.print("Sum of factorial digits : " + sumOfFactorial);
    }
    
    public static String getFactorial(int num) {
        BigInteger result = BigInteger.valueOf(1);
        
        for(int i = 1; i <= num; i++) {
            result = result.multiply(BigInteger.valueOf(i));
        }
        return result.toString();
    }
    
    public static int getSumOfFactorialDigits(String num) {
        int result = 0;
        for(int i = 0; i < num.length(); i++) {
            result += Integer.parseInt("" + num.charAt(i));
        }
        return result;
    }
}

// question :
/*
n! means n × (n − 1) × ... × 3 × 2 × 1
For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
Find the sum of the digits in the number 100!
*/