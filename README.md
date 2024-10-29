Here’s a README with emojis for your project, which explains the setup, usage, and features in a concise and engaging way:

---

# 📸 Image Similarity Search with Go

Welcome to **ImageSimilaritySearchGo**! This project helps you find the top similar images in a dataset using color histograms and Go’s concurrent processing power! 🚀

## 📋 Project Structure

```
ImageSimilaritySearchGo/
├── Src/              # Main source code directory
│   └── main.go       # Main Go program for similarity search
└── Query/            # Place your query images here
    └── query.jpg     # Example query image
```

## 🌟 Features
- **Efficient Similarity Search**: Finds the most similar images to a given query image using color histograms.
- **Parallel Processing**: Speeds up the process with customizable goroutines (set by `k`).
- **Flexible Directory Setup**: Simply add your query images in the `Query` folder and the dataset in any accessible directory.

## 🔧 Setup & Run Instructions

1. **Clone the Repository**: 
   ```bash
   git clone https://github.com/YourUsername/ImageSimilaritySearchGo.git
   cd ImageSimilaritySearchGo
   ```

2. **Organize Your Images**:
   - Add your **query image** to the `Query` directory (e.g., `Query/query.jpg`).
   - Ensure your dataset images (e.g., `.jpg` files) are in a separate directory (e.g., `imageDataset2_15_20`).

3. **Run the Program**:
   From within the `Src` directory, use the following command:
   ```bash
   go run main.go ../Query/query.jpg ../imageDataset2_15_20 <k>
   ```
   Replace:
   - `query.jpg` with your query image filename. For example you can use q00.jpg for query image 0 or qu01.jpg for query image 1
   - `imageDataset2_15_20` with your dataset folder.
   - `<k>` with the number of goroutines to use (e.g., `80`).

4. **Example Run**:
   ```bash
   go run similaritySearch.go ../Query/q00.jpg ../imageDataset2_15_20 80
   ```
   Output using k=80 go routines:
   <img width="793" alt="Screenshot 2024-10-28 212518" src="https://github.com/user-attachments/assets/12fab938-4975-4441-b5ee-d97219cc6b8f">


## 📚 How It Works
The program:
1. Computes a **color histogram** for each image in the dataset.
2. **Compares** histograms using intersection similarity.
3. **Outputs** the top 5 most similar images based on similarity scores. 

## 💡 Tips
- 🖼️ Ensure all images in the dataset are `.jpg` format.
- ⚙️ Experiment with `k` to balance between processing speed and system load.

## 🛠 Requirements
- [Go Language](https://golang.org/dl/) (version 1.16 or higher recommended)

---

Happy searching! 😊
