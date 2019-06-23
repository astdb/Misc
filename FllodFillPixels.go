/*
An image is represented by a 2-D array of integers, each integer representing the pixel value of the image (from 0 to 65535).
Given a coordinate (sr, sc) representing the starting pixel (row and column) of the flood fill, and a pixel value newColor, "flood fill" the image.
To perform a "flood fill", consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color as the starting pixel), and so on. Replace the color of all of the aforementioned pixels with the newColor.
At the end, return the modified image.
Example 1:
Input: 
image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1, sc = 1, newColor = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]
Explanation: 
From the center of the image (with position (sr, sc) = (1, 1)), all pixels connected 
by a path of the same color as the starting pixel are colored with the new color.
Note the bottom corner is not colored 2, because it is not 4-directionally connected
to the starting pixel.

Note:
The length of image and image[0] will be in the range [1, 50].
The given starting pixel will satisfy 0 <= sr < image.length and 0 <= sc < image[0].length.
The value of each color in image[i][j] and newColor will be an integer in [0, 65535].

*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    	// check which pixels are on north/south/east/west to the designated start
	// floodfill ones with the same color as the starter
	// make note of the floodfilled pixels and consider their neighbors 
	// to continue floodfilling

	// make note of starting pixel color
var starterColor int
	if len(image) > sr && len(image[sr]) > sc {
		starterColor := image[sr][sc]
	} else {
		log.Fatal("Invalid starting pixel")
	}

	pixelsToFill := [][]int{{sr, sc}}

	for len(pixelsToFill) > 0 {
		// for each pixel, 
		// - check it's color, and fill if it's starterColor
		// - if filled, remove it from pixelsToFill and add its north/south/east/west neighbors to pixelsToFill
		
		ssr := pixelsToFill[len(pixelsToFill)-1][0]
ssc := pixelsToFill[len(pixelsToFill)-1][1]

		// north
		if (ssr-1) >= 0 {
			north := image[ssr-1][ssc]	
		if north == starterColor {
			// fill
			image[ssr-1][ssc] = newColor

			// remove
			pixelsToFill = pixelsToFill[:len(pixelsToFill)-1]

			// add neighbors
			if ssr-2 >= 0 && image[ssr-2][ssc] == starterColor {
				pixelsToFill = append(pixelsToFill, {image[ssr-2][ssc]})	// north
			}

			if (ssc+1) < len(image[ssr-1]) && image[ssr-1][ssc+1] == starterColor {
				pixelsToFill = append(pixelsToFill, {image[ssr-1][ssc+1]})	// east
			}

			// south was already filled

			if (ssc-1) >= 0 && image[ssr-1][ssc-1] == starterColor {
				pixelsToFill = append(pixelsToFill, {image[ssr-1][ssc-1]})	// west
			}
			
		}
		}

		// east
		if (ssc+1) < len(image[ssr]) {
			east := image[ssr][ssc+1]

			if east == starterColor {
				// fill
			image[ssr][ssc+1] = newColor

			// remove from to-be-filled list
			pixelsToFill = pixelsToFill[:len(pixelsToFill)-1]

			// add eligible neighbors to to-be-filled list
			// north
			if (ssr+1) < len(image) && image[ssr+1][ssc+1] == starterColor {
				pixelsToFill = append(pixelsToFill, {image[ssr+1][ssc+1]})
			}

			// east

			// south


			// west
			}
		}
		
	}	
	
}

