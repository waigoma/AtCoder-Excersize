package p20220422;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        List<Integer> inputList = new ArrayList<>();

        while (scanner.hasNextLine()) {
            String line = scanner.nextLine();
            if (line.equals("")) break;
            inputList.add(Integer.parseInt(line));
        }

        int n = inputList.get(0) - 1;

        for (int i = 1; i < inputList.size(); i++) {
            int count = 1;
            int tmp = 0;
            for (int j = 1; j < inputList.size(); j++) {
                if (i == j) continue;
                if (inputList.get(j) > tmp) tmp = inputList.get(j);
                if (count++ > n) break;
            }
            System.out.println(tmp);
        }
    }
}
