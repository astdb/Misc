/*
Dr. Black has been murdered. Detective Jill must determine the murderer, crime scene, and weapon. There are six possible murderers (numbered 1 through 6, Professor Plum to Mrs. Peacock), 10 locations (1 through 10, ballroom to cellar), and six weapons (1 through 6, lead pipe to spanner). Detective Jill tries to guess the correct combination (there are 360 possibilities). Each guess is a theory. She asks her assistant, Jack, to confirm or refute each theory. When Jack refutes a theory, he reports that one of the guesses�murderer, location, or weapon�is wrong. The contestants are tasked with implementing a procedure that plays the role of Detective Jill. A brute-force program that tests all 360 theories earns a mere 50 points. An efficient program that tests no more than 20 theories earns an additional 50.
*/ 

import java.util.*; 

public class murder{ 

		char theoryStatus; 
		
		//create random theory for test purposes 
		Random randGen = new Random();
		int[] correctTheory = new int[3]; 
		correctTheory[0] = randGen.nextInt(6); 
		correctTheory[1] = randGen.nextInt(10); 
		correctTheory[2] = randGen.nextInt(6); 
		
		System.out.println("\nCorrect theory is: (" + correctTheory[0] + ", " + correctTheory[1] + ", " + correctTheory[2] + ")\n" );
		
		//	System.exit(0);
		
		System.out.print("Jill: (" + suspect + ", " + location + ", " + weapon + ")" );
		theoryStatus = check(suspect, location, weapon, correctTheory); 		
		
		while( true ) {
			if( theoryStatus == 'm' ) { 
				System.out.println(" | Incorrect: suspect");
				suspect++; 
			} else if( theoryStatus == 'l' ) { 
				System.out.println(" | Incorrect: location");
				location++; 
			} else if( theoryStatus == 'w' ) { 
				System.out.println(" | Incorrect: weapon");
				weapon++; 
			} else if( theoryStatus == 'c' ) { 
				System.out.println(" | *********** Correct! *********** ");
				break; 
			}
			
			System.out.print("\nJill: (" + suspect + ", " + location + ", " + weapon + ")" );
			theoryStatus = check(suspect, location, weapon, correctTheory); 	
		}
		
	}
	
	private static char check(int s, int l, int w, int[] ct){ 
		if( ct[0] != s ) { 
			return 'm';
		}
		
		if( ct[1] != l ) { 
			return 'l';
		}
		
		if( ct[2] != w ) { 
			return 'w';
		}
		
		return 'c';
	}
}
