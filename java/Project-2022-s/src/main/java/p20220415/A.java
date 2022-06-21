package p20220415;

import java.util.Scanner;

public class A {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        scanner.close();

        if (n < 10) {
            System.out.println("000" + n);
        } else if (n < 100) {
            System.out.println("00" + n);
        } else if (n < 1000) {
            System.out.println("0" + n);
        } else {
            System.out.println(n);
        }
    }
}
