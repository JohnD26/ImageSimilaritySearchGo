package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

type Histo struct {
	Name string
	H    []float64
}

func computeHistogram(imagePath string, depth int) (Histo, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return Histo{}, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return Histo{}, err
	}

	bounds := img.Bounds()
	histogram := make([]float64, 1<<(3*depth)) // Adjustment for reduced color space.

	var totalPixels int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			//  reduce the color depth and calculate histogram index.
			r = r >> (16 - depth) // shift for 8-bit depth images and then reducing  depth.
			g = g >> (16 - depth)
			b = b >> (16 - depth)
			index := (r << (2 * depth)) + (g << depth) + b
			histogram[index]++
			totalPixels++
		}
	}

	// Normalize the histogram.
	for i := range histogram {
		histogram[i] /= float64(totalPixels)
	}

	return Histo{Name: filepath.Base(imagePath), H: histogram}, nil
}

func compareHistograms(h1, h2 Histo) float64 {
	var intersection float64
	for i := range h1.H {
		intersection += min(h1.H[i], h2.H[i])
	}
	return intersection
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// divideDataset divides the dataset into K slices.
func divideDataset(files []os.DirEntry, k int) [][]os.DirEntry {
	var slices [][]os.DirEntry
	n := len(files)
	for i := 0; i < n; i += n / k {
		end := i + n/k
		if end > n {
			end = n
		}
		slices = append(slices, files[i:end])
	}
	return slices
}

// processSlice computes histograms for a slice of images.
func processSlice(slice []os.DirEntry, imageDatasetDirectory string, depth int, histogramsChan chan<- Histo, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, file := range slice {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".jpg" {
			filePath := filepath.Join(imageDatasetDirectory, file.Name())
			histo, err := computeHistogram(filePath, depth)
			if err != nil {
				fmt.Println("Error computing histogram for image:", file.Name(), err)
				continue
			}
			histogramsChan <- histo
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong input on terminal. Please write: go run similaritySearch.go <queryImageFilename> imageDataset2_15_20")
		os.Exit(1)
	}

	queryImageFilename := os.Args[1]
	imageDatasetDirectory := os.Args[2]

	queryHisto, err := computeHistogram(queryImageFilename, 3)
	if err != nil {
		fmt.Printf("Failed to compute histogram for query image: %v\n", err)
		return
	}

	files, err := os.ReadDir(imageDatasetDirectory)
	if err != nil {
		fmt.Printf("Failed to read dataset directory: %v\n", err)
		return
	}

	//BEGINNING
	//Specifying k, the number of Go Routines that we will be using to split the dataset
	k := 256
	fmt.Printf("Finding similarity with K=%d\n", k)
	var wg sync.WaitGroup
	histogramsChan := make(chan Histo, len(files))
	slices := divideDataset(files, k)

	for _, slice := range slices {
		wg.Add(1)
		go processSlice(slice, imageDatasetDirectory, 3, histogramsChan, &wg)
	}

	go func() {
		wg.Wait()
		close(histogramsChan)
	}()

	similarityScores := make(map[string]float64)
	for histo := range histogramsChan {
		score := compareHistograms(queryHisto, histo)
		similarityScores[histo.Name] = score
	}

	var pairs []nameScorePair
	for name, score := range similarityScores {
		pairs = append(pairs, nameScorePair{Name: name, Score: score})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Score > pairs[j].Score
	})

	fmt.Println("Top 5 similar images:")
	for i := 0; i < len(pairs) && i < 5; i++ {
		fmt.Printf("%d: %s - Score: %f\n", i+1, pairs[i].Name, pairs[i].Score)
	}

	//END

	/*MEASURING TIME*/
	//PRESS CONTROL SHIFT Z ON AN IDEA
	//WHILE MEASURING TIME WE MAKE SURE TO NOT USE ANY PRINT STATEMENTS AS PROF. LAGANIERE ONCE TOLD ME
	//PRINT STATEMENTS GO THROUGH THE OS TO DISPLAY INFORMATION AND GIVE YOU ERRONEOUS TIME MEASUREMENT
	/*k := 1048
	startTime := time.Now()
	var wg sync.WaitGroup
	histogramsChan := make(chan Histo, len(files))
	slices := divideDataset(files, k)

	for _, slice := range slices {
		wg.Add(1)
		go processSlice(slice, imageDatasetDirectory, 3, histogramsChan, &wg)
	}

	go func() {
		wg.Wait()
		close(histogramsChan)
	}()

	similarityScores := make(map[string]float64)
	for histo := range histogramsChan {
		score := compareHistograms(queryHisto, histo)
		similarityScores[histo.Name] = score
	}

	var pairs []nameScorePair
	for name, score := range similarityScores {
		pairs = append(pairs, nameScorePair{Name: name, Score: score})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Score > pairs[j].Score
	})

	// After measuring we print the execution time for this value of K
	endTime := time.Now()
	fmt.Printf("Execution time for K=%d: %v\n", k, endTime.Sub(startTime))*/

}

type nameScorePair struct {
	Name  string
	Score float64
}
