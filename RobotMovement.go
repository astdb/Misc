// This program simulates actions of a toy robot on to 5x5 tabletop according to a set of given commands.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// ensure input filename entered on command line
	if len(os.Args) <= 1 {
		fmt.Println("Error: please enter the robot operations command file name on command-line (e.g. > go run RobotMovement.go commands.txt)")
		return
	}

	// read commands file
	inFileName := os.Args[1]
	file, err := os.Open(inFileName)
	logErr(err)
	defer file.Close()

	// setup scanner for commands file input
	scanner := bufio.NewScanner(file)

	// create new robot
	c_3PO := createRobot()

	// iterate through the commands
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line != "" {
			// line_split will contain a slice of strings when the current line is 'exploded' by the comma delimiter
			line_split := strings.Split(line, " ")

			command := strings.TrimSpace(line_split[0])

			if command == "PLACE" {
				place_coords := strings.Split(line_split[1], ",")

				// parse x-coordinate of placement command
				place_x, err := strconv.Atoi(strings.TrimSpace(place_coords[0]))
				if err != nil {
					fmt.Fprintf(os.Stderr, "Placement X-coord argument error: %v\n", err)
					continue
				}

				// parse x-coordinate of placement command
				place_y, err := strconv.Atoi(strings.TrimSpace(place_coords[1]))
				if err != nil {
					fmt.Fprintf(os.Stderr, "Placement Y-coord argument error: %v\n", err)
					continue
				}

				// parse facing alignment of placement command
				place_facing := strings.TrimSpace(place_coords[2])

				// place robot and capture any placement error
				err = c_3PO.place(place_x, place_y, place_facing)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Placement error: %v\n", err)
				}

			}

			if command == "MOVE" {
				err := c_3PO.move()

				if err != nil {
					fmt.Fprintf(os.Stderr, "Movement error: %v\n", err)
				}
			}

			if command == "LEFT" {
				err := c_3PO.left()

				if err != nil {
					fmt.Fprintf(os.Stderr, "Rotation error: %v\n", err)
				}
			}

			if command == "RIGHT" {
				err := c_3PO.right()

				if err != nil {
					fmt.Fprintf(os.Stderr, "Rotation error: %v\n", err)
				}
			}

			if command == "REPORT" {
				current_x, current_y, current_facing, err := c_3PO.report()

				if err != nil {
					fmt.Fprintf(os.Stderr, "Report error: %v\n", err)
					continue
				}

				fmt.Printf("%d, %d, %s\n", current_x, current_y, current_facing)
			}
		}
	}
}

// ----------------- end main() --------------------

// type struct to represent a robot
type robot struct {
	X      int    // x-coordinate of current position
	Y      int    // y-coordinate of current position
	F      string // current facing direction (NORTH, SOUTH, EAST or WEST)
	Placed bool   // indicator whether the robot is placed on the table or not

	// maximum values of X- and Y-coordinatess the robot can travel - minimum assumed to be 0
	XLIMIT int
	YLIMIT int
}

// maker function create and return a reference to a robot instance
// func createRobot(x int, y int, facing string) *robot {
func createRobot() *robot {
	var r robot
	r.X = 0
	r.Y = 0
	r.F = "NONE"
	r.Placed = false

	// table is 5x5
	r.XLIMIT = 5
	r.YLIMIT = 5

	return &r
}

// ----------- methods of robot type (place, move, left, right, report) -----------
func (r *robot) place(x int, y int, facing string) error {
	if x >= 0 && x <= r.XLIMIT && y >= 0 && y <= r.YLIMIT && (facing == "NORTH" || facing == "EAST" || facing == "SOUTH" || facing == "WEST") {
		r.X = x
		r.Y = y
		r.F = facing
		r.Placed = true
		return nil
	}

	return fmt.Errorf("Robot placement attempted with invalid initial coordinates (X: %d, Y: %d, FACING: %s)", r.X, r.Y, r.F)
}

func (r *robot) move() error {
	// move robot one unit in the facing direction, with regards to table limitations
	if !r.Placed {
		// if robot not placed on table, ignore move command
		return fmt.Errorf("move() called on unplaced robot .")
	}

	if r.F == "NORTH" && r.Y < r.YLIMIT {
		// move 1 unit NORTH
		r.Y++
		return nil
	}

	if r.F == "EAST" && r.X < r.XLIMIT {
		// move 1 unit EAST
		r.X++
		return nil
	}

	if r.F == "SOUTH" && r.Y > 0 {
		// move 1 unit SOUTH
		r.Y--
		return nil
	}

	if r.F == "WEST" && r.X > 0 {
		// move 1 unit WEST
		r.X--
		return nil
	}

	return fmt.Errorf("move() called on robot with possibly undefined initial facing (X: %d, Y: %d, FACING: %s)", r.X, r.Y, r.F)
}

func (r *robot) left() error {
	// rotate robot 90degrees counter-clockwise

	if !r.Placed {
		// if robot not placed on table, ignore left command
		return fmt.Errorf("left() called on unplaced robot .")
	}

	if r.F == "NORTH" {
		r.F = "WEST"
		return nil
	}

	if r.F == "EAST" {
		r.F = "NORTH"
		return nil
	}

	if r.F == "SOUTH" {
		r.F = "EAST"
		return nil
	}

	if r.F == "WEST" {
		r.F = "SOUTH"
		return nil
	}

	return fmt.Errorf("left() called on robot with possibly undefined initial facing (X: %d, Y: %d, FACING: %s)", r.X, r.Y, r.F)
}

func (r *robot) right() error {
	// rotate robot 90degrees clockwise

	if !r.Placed {
		// if robot not placed on table, ignore right command
		return fmt.Errorf("left() called on unplaced robot .")
	}

	if r.F == "NORTH" {
		r.F = "EAST"
		return nil
	}

	if r.F == "EAST" {
		r.F = "SOUTH"
		return nil
	}

	if r.F == "SOUTH" {
		r.F = "WEST"
		return nil
	}

	if r.F == "WEST" {
		r.F = "NORTH"
		return nil
	}

	return fmt.Errorf("right() called on robot with possibly undefined initial facing (X: %d, Y: %d, FACING: %s)", r.X, r.Y, r.F)
}

func (r *robot) report() (int, int, string, error) {
	// fmt.Printf("X: %d, Y: %d, Facing: %s", r.X, r.Y, r.F);
	var err error
	err = nil
	if !r.Placed {
		// return error if robot is not placed
		err = fmt.Errorf("report() called on unplaced robot.")
	}

	return r.X, r.Y, r.F, err
	// return
}

// ----------------- end methods of robot type ----------------------

// utility function to evaluate and log errors encountered throughout program execution
func logErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
