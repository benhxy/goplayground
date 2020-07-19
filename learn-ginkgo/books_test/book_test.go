package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"learn-ginkgo/books"
)

var _ = Describe("Book", func() {
	var (
		shortBook books.Book
		longBook  books.Book
		book      books.Book
		json      string
		err       error
	)

	BeforeEach(func() {
		longBook = books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = books.Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}

		json = `{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
        }`
	})

	// Describe (given do something) - Context (when) - It (should)
	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("Novel"))
			})
		})

		GinkgoWriter.Write([]byte("Say something in verbose mode."))

		Context("With no more than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("Short story"))
			})
		})
	})

	// This execute just before each It()
	JustBeforeEach(func() {
		book, err = books.NewBookFromJSON(json)
	})

	Describe("Deserializing json", func() {
		Context("With a valid Book json", func() {
			It("should return a Book with correct values", func() {
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Pages).To(Equal(1488))
			})

			It("should not have error", func() {
				Expect(err).To(BeZero())
			})
		})

		Context("With an invalid Book json", func() {
			BeforeEach(func() {
				json = `{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":1488oops
                }`
			})

			It("should return a zero Book object", func() {
				Expect(book).To(BeZero())
			})

			It("should return error", func() {
				Expect(err).ToNot(BeZero())
			})
		})
	})
})
