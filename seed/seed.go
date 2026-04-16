package main

import (
    "log"

    "github.com/joho/godotenv"
    "github.com/Frientia/my-firebase-backend/config"
    "github.com/Frientia/my-firebase-backend/models"
)

func main() {
    godotenv.Load()
    config.InitDatabase()

    products := []models.Product{
        {
            Name:        "Atomic Habits",
            Price:       120000,
            Category:    "Self Development",
            Stock:       40,
            Description: "Buku tentang membangun kebiasaan kecil yang berdampak besar",
            ImageURL:    "https://ibb.co.com/xq4vXdLJ",
        },
        {
            Name:        "Rich Dad Poor Dad",
            Price:       100000,
            Category:    "Finance",
            Stock:       60,
            Description: "Panduan mengelola keuangan dan mindset orang kaya",
            ImageURL:    "https://ibb.co.com/JjMFB7Vm",
        },
        {
            Name:        "Laskar Pelangi",
            Price:       90000,
            Category:    "Novel",
            Stock:       80,
            Description: "Kisah inspiratif anak-anak Belitung",
            ImageURL:    "https://ibb.co.com/1fQm2M1V",
        },
        {
            Name:        "Clean Code",
            Price:       150000,
            Category:    "Programming",
            Stock:       25,
            Description: "Panduan menulis kode yang rapi dan maintainable",
            ImageURL:    "https://ibb.co.com/whhPqcrZ",
        },
        {
            Name:        "Harry Potter and the Sorcerer's Stone",
            Price:       130000,
            Category:    "Fantasy",
            Stock:       35,
            Description: "Petualangan Harry Potter di dunia sihir",
            ImageURL:    "https://ibb.co.com/7xS57WGr",
        },
        {
            Name:        "The Psychology of Money",
            Price:       110000,
            Category:    "Finance",
            Stock:       45,
            Description: "Pemahaman tentang perilaku manusia terhadap uang",
            ImageURL:    "https://ibb.co.com/QFSKfnMg",
        },
        {
            Name:        "Deep Work",
            Price:       125000,
            Category:    "Productivity",
            Stock:       30,
            Description: "Strategi fokus tanpa distraksi di dunia digital",
            ImageURL:    "https://ibb.co.com/fGtJpB9C",
        },
        {
            Name:        "Sapiens: A Brief History of Humankind",
            Price:       140000,
            Category:    "History",
            Stock:       20,
            Description: "Sejarah singkat umat manusia dari masa ke masa",
            ImageURL:    "https://ibb.co.com/dwQFGtnN",
        },
        {
            Name:        "The Alchemist",
            Price:       95000,
            Category:    "Novel",
            Stock:       70,
            Description: "Perjalanan menemukan makna hidup dan impian",
            ImageURL:    "https://ibb.co.com/qLZ6SDsS",
        },
        {
            Name:        "Flutter for Beginners",
            Price:       160000,
            Category:    "Programming",
            Stock:       15,
            Description: "Panduan lengkap belajar Flutter dari dasar",
            ImageURL:    "https://ibb.co.com/HfKbwgHg",
        },
    }

    for _, p := range products {
        config.DB.Create(&p)
    }

    log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
