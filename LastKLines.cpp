
// problem: print out the last k lines of the input file

// solution: implement a k-sized circular array and insert each line as its read, if required replacing the oldest line. 
// oldest-line must be memorized in a counter variable. that's the insert index to the buffer e.g. buffer[oldest] = newline 

// reading a text file
#include <stdlib.h>
#include <iostream>
#include <fstream>
#include <string>
using namespace std;

int main (int argc, char* argv[]) {
  if (argc < 3 ){
    cerr << "Please enter input file name and number of lines to read from end of file.\n";
    return 1;
  }

  // read filename and number of lines to read from command line
  string inputFile = argv[1];
  int numLines = atoi(argv[2]);

  if (numLines <= 0) {
    cerr << "Please enter a valid number of lines to read from the end of file.\n";
    return 1;
  }

  string line;  // placeholder for current line read  
  ifstream myfile(inputFile.c_str()); // input filestream to read the file

  // declare rotating array buffer (complete with dynamic 'next' index)
  string lastKLines[numLines];
  int lastKLines_next = 0;      // index to enter next new line read
  bool rotated = false;         // flag indicating if the lines in buffer started getting overwritten

  int i = 0;                    // line counter
  if (myfile.is_open()) {
    while ( getline (myfile,line) ) {
      i++;
      // cout << line << '\n';

      // store line in buffer
      lastKLines[lastKLines_next] = line;

      // update 'next' index appropriately
      lastKLines_next++;
      if(lastKLines_next >= numLines) {
        lastKLines_next = 0;
        rotated = true;
      }
    }

    // done reading file: close file input tstream
    myfile.close();
  } else {
    // error opening file input stream
    cout << "Unable to open file";
  }
  
  // index of the oldest line in buffer (therefore the first to print)
  int j = lastKLines_next;
  int k = lastKLines_next - 1;  // index of the newest line read to buffer (therefore last to print)

  if (k < 0) {
    k = numLines - 1;
  }

  while(true){
    if(j >= numLines){
      j = 0;
    }

    if(!rotated && j > k) {
      // don't print anything - top half of buffer would be empty if not 'rotated'

    } else {
      // print line in buffer
      cout << lastKLines[j] << "\n";
    }

    if( j == lastKLines_next-1 ){
      break;
    }

    j++;
  }
  
  return 0;
}
