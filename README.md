# 📸 Image Similarity Search with Go

Welcome to **ImageSimilaritySearchGo**! This project helps you find the top similar images in a dataset using color histograms and Go’s concurrent processing power! 🚀

## 📋 Project Structure

```
ImageSimilaritySearchGo/
├── Src/                       # Main similarity search code
│   └── main.go                # Main Go program for similarity search
├── Query/                     # Directory for query images
│   └── query.jpg              # Example query image
└── TimeMeasurement/           # Timing experiment code
    └── time_experiment.go     # Independent timing experiment code
```

## 🌟 Features
- **Efficient Similarity Search**: Finds the most similar images to a given query image using color histograms.
- **Parallel Processing**: Speeds up the process with customizable goroutines.
- **Timing Experiment**: Measure how different numbers of goroutines (`k`) affect the execution time of the similarity search.

## 🔧 Setup & Run Instructions

### Running the Similarity Search

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/YourUsername/ImageSimilaritySearchGo.git
   cd ImageSimilaritySearchGo
   ```

2. **Organize Your Images**:
   - Add your **query image** to the `Query` directory (e.g., `Query/query.jpg`).
   - Ensure your dataset images (e.g., `.jpg` files) are in a separate directory (e.g., `imageDataset2_15_20`).

3. **Run the Similarity Search**:
   From within the `Src` directory, use the following command:
   ```bash
   go run similaritySearch.go ../Query/query.jpg ../imageDataset2_15_20 <k>
   ```
   Replace:
   - `query.jpg` with the query image filename you want to use.
   - `imageDataset2_15_20` with your dataset folder.
   - `<k>` with the number of goroutines to use (e.g., `80`).

4. **Example Run**:
   ```bash
   go run similaritySearch.go ../Query/q00.jpg ../imageDataset2_15_20 80
   ```
   Expected output:
   ```
   Finding similarity with K=80
   Top 5 similar images:
   1: 1144.jpg - Score: 1.000000
   2: 3806.jpg - Score: 0.704005
   3: 3756.jpg - Score: 0.660608
   4: 3714.jpg - Score: 0.659687
   5: 3668.jpg - Score: 0.643304
   ```

### Running the Timing Experiment

To observe how increasing the number of goroutines (`k`) affects the time taken to run the similarity search:

1. Navigate to the `TimeMeasurement` directory:
   ```bash
   cd TimeMeasurement
   ```

2. Run the experiment with an initial value for `k`. For example:
   ```bash
   go run time_experiment.go ../Query/q00.jpg ../imageDataset2_15_20 5
   ```
   This will run the similarity search with `k` values increasing as multiples of the initial value (e.g., 5, 10, 20, 40), and it will output the time taken for each.

3. **Sample Output**:
<img width="730" alt="Time Measurement" src="https://github.com/user-attachments/assets/c31821c0-8fe2-4c3a-a773-f5f53f4f6d8a">


   As shown, increasing `k` generally decreases the time taken to complete the algorithm, demonstrating the performance gains from parallel processing in Go using Go routines 

## 💡 Tips
- 🖼️ Ensure all images in the dataset are in `.jpg` format.
- ⚙️ Experiment with different `k` values to see how parallel processing impacts performance on your machine.

## 🛠 Requirements
- [Go Language](https://golang.org/dl/) (version 1.16 or higher recommended)

---

Happy searching! 😊
