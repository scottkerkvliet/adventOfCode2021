using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AoC2021 {
    public static class Day20 {

        static string[] input = Read.File("image.txt");

        public static void Solve() {

            string imageEnhancementAlgorithm = input[0];
            Dictionary<string, bool> outputImagePixels = new Dictionary<string, bool>();
            List<string> nextOutputImageOn = new List<string>();
            List<string> nextOutputImageOff = new List<string>();

            int minX = -1, minY = -1;
            int maxX = input[2].Length + 1, maxY = input.Length - 2 + 1;
            int iter = 0;
            for (int line = 2; line < input.Length; line++) {
                for (int i = 0; i < input[line].Length; i++) {
                    if(input[line][i] == '#') {
                        string position = i + "," + (line - 2);
                        outputImagePixels.Add(position, true);
                    }
                    else {
                        string position = i + "," + (line - 2);
                        outputImagePixels.Add(position, false);
                    }
                }
            }

            // cycle through each point on the image, though
            // each iteration we increase outter bounds by 1
            while(iter < 50) {
                for (int y = minY; y < maxY; y++) {
                    for (int x = minX; x < maxX; x++) {
                        // setup our string to make a binary number
                        string currentPixelPos = x + "," + y;
                        string binaryString = "";
                        int imageEnhancementIndex;
                        // our sliding 3x3 window
                        for (int wY = y - 1; wY <= y + 1; wY++) {
                            for (int wX = x - 1; wX <= x + 1; wX++) {
                                if(iter % 2 == 0) { // if we are odd, infinite lit, even infinite off.
                                    if(outputImagePixels.ContainsKey(wX + "," + wY)) {
                                        if (outputImagePixels[wX + "," + wY]) {
                                            binaryString += '1';
                                        }
                                        else {
                                            binaryString += '0';
                                        }
                                    }
                                    else {
                                        binaryString += '0';
                                    }  
                                }
                                else {
                                    if(outputImagePixels.ContainsKey(wX + "," + wY)) {
                                        if (outputImagePixels[wX + "," + wY]) {
                                            binaryString += '1';
                                        }
                                        else{
                                            binaryString += '0';
                                        }
                                    }
                                    else {
                                        binaryString += '1';
                                    }
                                }
                            }
                        }
                        // now that we have the binary, we can use that number to determine the pixel
                        imageEnhancementIndex = Convert.ToInt32(binaryString, 2);

                        if (imageEnhancementAlgorithm[imageEnhancementIndex] == '#') {
                            nextOutputImageOn.Add(currentPixelPos);
                        }
                        else {
                            nextOutputImageOff.Add(currentPixelPos);
                        }
                    }
                }
                // Adjust outward and update the image
                outputImagePixels.Clear();
                foreach (string s in nextOutputImageOn) {
                    outputImagePixels.Add(s, true);
                }
                foreach (string s in nextOutputImageOff) {
                    outputImagePixels.Add(s, false);
                }
                nextOutputImageOn.Clear();
                nextOutputImageOff.Clear();
                minX--;
                minY--;
                maxX++;
                maxY++;
                iter++;

                if(iter == 2) {
                    int onCountP1 = 0;
                    foreach (KeyValuePair<string, bool> kvp in outputImagePixels) {
                        if (kvp.Value)
                            onCountP1++;
                    }
                    Console.WriteLine(onCountP1);
                }

            }
            int onCount = 0;
            foreach(KeyValuePair<string,bool> kvp in outputImagePixels) {
                if (kvp.Value)
                    onCount++;
            }
            Console.WriteLine(onCount);
        }
    }
}
