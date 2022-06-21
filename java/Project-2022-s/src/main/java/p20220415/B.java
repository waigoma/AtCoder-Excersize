package p20220415;

import java.util.Scanner;

public class B {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String n = sc.nextLine();

        String[] ss = n.split(" ");

        int ni = Integer.parseInt(ss[0]);
        if (!(1 <= ni &&  ni <= 100000)) return;

        int border = Integer.parseInt(ss[1]);
        String s = sc.nextLine();
        String[] ss2 = s.split(" ");

        int total = 0;
        int t = 1;
        for (String str : ss2) {
            int i = Integer.parseInt(str);
            if (i < border) total++;
        }

        System.out.println(total);
    }
}
