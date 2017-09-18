// This program simulates a toy robot moving atop a 5x5 grid surface, according to a set of given commands.

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// the program reads in robot movement commands from an input file
	// ensure input filename is entered on command line
	if len(os.Args) <= 1 {
		fmt.Println("Error: please enter the robot operations command file name on command-line (e.g. > go run REA-Robot.go commands.txt)")
		return
	}

	// capture input file name from command line and open for reading
	inFileName := os.Args[1]
	file, err := os.Open(inFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		return
	}
	defer file.Close()

	// setup scanner to read through command file
	scanner := bufio.NewScanner(file)

	// create new robot 🤖
	c_3PO := createRobot()

	// iterate through the commands
	for scanner.Scan() {
		line := scanner.Text()

		err := c_3PO.Action(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	}
}

// ----------------- end main() --------------------

// type struct to represent a robot
type robot struct {
	X      int                    // x-coordinate of current position
	Y      int                    // y-coordinate of current position
	F      string                 // current facing direction (NORTH, SOUTH, EAST or WEST)
	Placed bool                   // indicator whether the robot is placed on the table or not
	CmdMap map[string]interface{} // Map structure mapping command tokens (from input) to action methods (of robot)

	NORTH, EAST, SOUTH, WEST         string // Constant literals holding direction names
	PLACE, MOVE, LEFT, RIGHT, REPORT string // Constant literals holding command names

	// maximum values of X- and Y-coordinatess the robot can travel - minimum assumed to be 0
	XLIMIT int
	YLIMIT int

	// Verbose mode flag - if set the robot would print out extra information about its activities
	Verbose bool
}

// builder function to create and return a reference to a new robot instance
func createRobot() *robot {
	var r robot
	r.X = 0
	r.Y = 0
	r.F = "NONE"
	r.Placed = false

	r.PLACE = "PLACE"
	r.MOVE = "MOVE"
	r.LEFT = "LEFT"
	r.RIGHT = "RIGHT"
	r.REPORT = "REPORT"

	r.CmdMap = map[string]interface{}{
		r.PLACE:  r.place,
		r.MOVE:   r.move,
		r.LEFT:   r.left,
		r.RIGHT:  r.right,
		r.REPORT: r.report,
	}

	r.NORTH = "NORTH"
	r.EAST = "EAST"
	r.SOUTH = "SOUTH"
	r.WEST = "WEST"

	// table is 5x5
	r.XLIMIT = 5
	r.YLIMIT = 5

	// verbose activity mode off by default
	// if set to true, will print out detailed activity information to console
	r.Verbose = false

	// return pointer to this new  struct instance
	return &r
}

// --------------------- methods of robot type ------------------------
// high-level, exported method to take a command text line read from the input file and invoke the corresponding robot action
func (r *robot) Action(commandline string) error {
	// if the input line starts with a #, ignore as comment (and ignore blank lines)
	commentMatcher := regexp.MustCompile(`^#`)
	if commentMatcher.MatchString(strings.TrimSpace(commandline)) || strings.TrimSpace(commandline) == "" {
		return nil
	}

	// split the command string by whitespace delimiter to obtain command name (and, if applicable, any arguments)
	line_split := strings.Split(commandline, " ")
	command := strings.TrimSpace(line_split[0])

	// check if the command is one with parameters (e.g. PLACE X Y F) or single (e.g. MOVE)
	if len(line_split) > 1 {
		// parameterized command e.g. PLACE X Y F
		if command == r.PLACE {
			place_coords := strings.Split(line_split[1], ",")

			if len(place_coords) == 3 {
				place_x, e1 := strconv.Atoi(strings.TrimSpace(place_coords[0]))
				place_y, e2 := strconv.Atoi(strings.TrimSpace(place_coords[1]))
				place_facing := strings.TrimSpace(place_coords[2])

				if e1 == nil && e2 == nil && place_facing != "" {
					err := r.CmdMap[command].(func(int, int, string) error)(place_x, place_y, place_facing)
					return err
				} else {
					return fmt.Errorf("Invalid <%s> parameters (%s)", command, line_split[1])
				}
			} else {
				return fmt.Errorf("Invalid <%s> parameters (%s)", command, line_split[1])
			}
		} else {
			return fmt.Errorf("Invalid command <%s>", commandline)
		}
	} else {
		// standalone command e.g. MOVE

		// NOTE: this check is required to prevent program crash in the event of malformed input such as
		// a parameterized command (e.g. PLACE issued without parameters). If this safeguard is not present,
		// this will cause an attempted invocation of a method with incorrect signature for the command name in that case.
		if command == r.MOVE || command == r.LEFT || command == r.RIGHT || command == r.REPORT {
			err := r.CmdMap[command].(func() error)()
			return err
		} else {
			return fmt.Errorf("Invalid command <%s>", command)
		}
	}
}

func (r *robot) place(x int, y int, facing string) error {
	if x >= 0 && x <= r.XLIMIT && y >= 0 && y <= r.YLIMIT && (facing == r.NORTH || facing == r.EAST || facing == r.SOUTH || facing == r.WEST) {
		r.X = x
		r.Y = y
		r.F = facing
		r.Placed = true
		
		if r.Verbose {
			fmt.Printf("Placed on (%d,%d), facing %s\n", x, y, facing)
		}

		return nil
	}

	return fmt.Errorf("Robot placement attempted with invalid coordinates (X: %d, Y: %d, FACING: %s)", x, y, facing)
}

func (r *robot) move() error {
	// move robot one unit in the facing direction, with regards to table limitations and returns an error if fails to do so
	if !r.Placed {
		// if robot not placed on table, ignore move command
		return fmt.Errorf("move() called on unplaced robot.")
	}

	if r.F == r.NORTH && r.Y < r.YLIMIT {
		// move 1 unit NORTH
		r.Y++

		if r.Verbose {
			fmt.Printf("Moving one unit North..\n")
		}

		return nil
	}

	if r.F == r.EAST && r.X < r.XLIMIT {
		// move 1 unit EAST
		r.X++

		if r.Verbose {
			fmt.Printf("Moving one unit East..\n")
		}

		return nil
	}

	if r.F == r.SOUTH && r.Y > 0 {
		// move 1 unit SOUTH
		r.Y--

		if r.Verbose {
			fmt.Printf("Moving one unit South..\n")
		}

		return nil
	}

	if r.F == r.WEST && r.X > 0 {
		// move 1 unit WEST
		r.X--

		if r.Verbose {
			fmt.Printf("Moving one unit West..\n")
		}

		return nil
	}

	// if program execution reaches this point it would mean that the robot was placed on table but has a facing of neither defined value (NORTH, EAST, SOUTH or WEST) - return error
	return fmt.Errorf("move() called on robot with undefined initial facing or facing towards edge of table (X: %d, Y: %d, FACING: %s)", r.X, r.Y, r.F)
}

func (r *robot) left() error {
	// rotate robot 90degrees counter-clockwise and returns an error if failed to do so

	if !r.Placed {
		// if robot not placed on table, ignore left command
		return fmt.Errorf("left() called on unplaced robot.")
	}

	if r.F == r.NORTH {
		r.F = r.WEST

		if r.Verbose {
			fmt.Printf("Turning to face West..\n")
		}

		return nil
	}

	if r.F == r.EAST {
		r.F = r.NORTH

		if r.Verbose {
			fmt.Printf("Turning to face North..\n")
		}

		return nil
	}

	if r.F == r.SOUTH {
		r.F = r.EAST

		if r.Verbose {
			fmt.Printf("Turning to face East..\n")
		}

		return nil
	}

	if r.F == r.WEST {
		r.F = r.SOUTH

		if r.Verbose {
			fmt.Printf("Turning to face South..\n")
		}

		return nil
	}

	// if program execution reaches this point it would mean that the robot was placed on table but has a facing of neither defined value (NORTH, EAST, SOUTH or WEST) - return error
	return fmt.Errorf("left() called on robot with possibly undefined initial facing (X: %d, Y: %d, FACING: %s)", r.X, r.Y, r.F)
}

func (r *robot) right() error {
	// rotate robot 90degrees clockwise and returns an error if filed to do so

	if !r.Placed {
		// if robot not placed on table, ignore right command
		return fmt.Errorf("right() called on unplaced robot.")
	}

	if r.F == r.NORTH {
		r.F = r.EAST

		if r.Verbose {
			fmt.Printf("Turning to face East..\n")
		}

		return nil
	}

	if r.F == r.EAST {
		r.F = r.SOUTH

		if r.Verbose {
			fmt.Printf("Turning to face South..\n")
		}

		return nil
	}

	if r.F == r.SOUTH {
		r.F = r.WEST

		if r.Verbose {
			fmt.Printf("Turning to face West..\n")
		}

		return nil
	}

	if r.F == r.WEST {
		r.F = r.NORTH

		if r.Verbose {
			fmt.Printf("Turning to face North..\n")
		}

		return nil
	}

	// if program execution reaches this point it would mean that the robot was placed on table but has a facing of neither defined value (NORTH, EAST, SOUTH or WEST) - return error
	return fmt.Errorf("right() called on robot with possibly undefined initial facing (X: %d, Y: %d, FACING: %s)", r.X, r.Y, r.F)
}

func (r *robot) report() error {
	// reports the robot's location along with an error if unplaced
	var err error
	err = nil
	if !r.Placed {
		// return error if robot is not placed
		err = fmt.Errorf("report() called on unplaced robot.")
		return err
	}

	// return r.X, r.Y, r.F, err
	fmt.Printf("%d, %d, %s\n", r.X, r.Y, r.F)
	return err
}

// ----------------- end methods of robot type ----------------------
