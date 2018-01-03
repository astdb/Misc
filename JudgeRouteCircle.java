public class JudgeRouteCircle {
    public static void main(String[] args) {
        System.out.println(judgeCircle("LL"));
    }

    public static boolean judgeCircle(String moves) {
        int X = 0;
        int Y = 0;

        for(int i = 0; i < moves.length(); i++) {
            char c = moves.charAt(i);            

            if(c == 'R') {
                X =X + 1;
            }

            if(c == 'L') {
                X = X - 1;
            }

            if(c == 'U') {
                Y = Y + 1;
            }

            if(c == 'D') {
                Y = Y - 1;
            }
        }

        if(X == 0 && Y == 0) {
            return true;
        }

        return false;
    }
}
