package gocva_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gungun974/gocva"
)

func cvaTests() {
	Describe("CVA", func() {
		It("should render empty with no parameters", func() {
			// arrange
			cva := gocva.CVA{}

			// act
			result := cva.Render(gocva.Variant{})

			// assert
			Expect(result).To(Equal(""))
		})

		It("should render base only when given a base", func() {
			// arrange
			cva := gocva.CVA{
				Base: "base-class",
			}

			// act
			result := cva.Render(gocva.Variant{})

			// assert
			Expect(result).To(Equal("base-class"))
		})

		It("should render basic variants when given variants", func() {
			// arrange
			cva := gocva.CVA{
				Variants: gocva.Variants{
					"variant": {
						"primary":   "variant-primary",
						"secondary": "variant-secondary",
						"danger":    "variant-danger",
					},
					"size": {
						"large": "text-large",
						"small": "text-small",
					},
				},
			}

			// act
			result := cva.Render(gocva.Variant{
				"variant": "secondary",
				"size":    "large",
			})

			// assert
			Expect(result).To(Equal("text-large variant-secondary"))
		})

		It("should render using default variants when no user variant is provided", func() {
			// arrange
			cva := gocva.CVA{
				Variants: gocva.Variants{
					"variant": {
						"primary":   "variant-primary",
						"secondary": "variant-secondary",
						"danger":    "variant-danger",
					},
					"size": {
						"large": "text-large",
						"small": "text-small",
					},
					"padding": {
						"with": "p-8",
						"none": "p-0",
					},
					"test": {
						"yes": "test",
					},
				},
				DefaultVariants: gocva.Variant{
					"variant": "danger",
					"padding": "with",
					"test":    "yes",
				},
			}

			// act
			result := cva.Render(gocva.Variant{
				"size":    "small",
				"padding": "none",
				"test":    "",
			})

			// assert
			Expect(result).To(Equal("p-0 text-small test variant-danger"))
		})

		It("should render empty when no default or user variant is provided", func() {
			// arrange
			cva := gocva.CVA{
				Variants: gocva.Variants{
					"variant": {
						"primary":   "variant-primary",
						"secondary": "variant-secondary",
						"danger":    "variant-danger",
					},
					"size": {
						"large": "text-large",
						"small": "text-small",
					},
					"padding": {
						"with": "p-8",
						"none": "p-0",
					},
				},
				DefaultVariants: gocva.Variant{},
			}

			// act
			result := cva.Render(gocva.Variant{})

			// assert
			Expect(result).To(Equal(""))
		})

		It("should render properly with compound variants", func() {
			// arrange
			cva := gocva.CVA{
				Variants: gocva.Variants{
					"variant": {
						"primary":   "variant-primary",
						"secondary": "variant-secondary",
						"danger":    "variant-danger",
					},
					"size": {
						"large": "text-large",
						"small": "text-small",
					},
					"padding": {
						"with": "p-8",
						"none": "p-0",
					},
				},
				DefaultVariants: gocva.Variant{},
			}

			// act
			result := cva.Render(gocva.Variant{})

			// assert
			Expect(result).To(Equal(""))
		})

		It("should render correctly in a real use case", func() {
			// arrange
			cva := gocva.CVA{
				Variants: gocva.Variants{
					"variant": {
						"primary":   "variant-primary",
						"secondary": "variant-secondary",
						"danger":    "variant-danger",
					},
					"size": {
						"large": "text-large",
						"small": "text-small",
					},
					"direction": {
						"left":   "",
						"center": "",
					},
				},
				CompoundVariants: []gocva.CompoundVariant{
					{
						Class: "md:ml-24",
						Attrs: gocva.Variant{
							"direction": "left",
							"size":      "large",
						},
					},
					{
						Class: "md:ml-[4.5rem]",
						Attrs: gocva.Variant{
							"direction": "left",
							"size":      "small",
						},
					},
				},
			}

			// act
			noCompoundResult := cva.Render(gocva.Variant{
				"variant":   "primary",
				"size":      "large",
				"direction": "center",
			})
			smallCompoundResult := cva.Render(gocva.Variant{
				"variant":   "primary",
				"size":      "small",
				"direction": "left",
			})
			largeCompoundResult := cva.Render(gocva.Variant{
				"variant":   "primary",
				"size":      "large",
				"direction": "left",
			})

			// assert
			Expect(noCompoundResult).To(Equal("text-large variant-primary"))
			Expect(smallCompoundResult).To(Equal("text-small variant-primary md:ml-[4.5rem]"))
			Expect(largeCompoundResult).To(Equal("text-large variant-primary md:ml-24"))
		})

		It("should not override defaults unless specified", func() {
			// arrange
			cva := gocva.CVA{
				Base: "relative text-center font-bold uppercase",
				Variants: gocva.Variants{
					"variant": {
						"white": "text-neutral-50",
						"black": "text-three",
						"blue":  "text-text-dark",
					},
					"direction": {
						"left":   "md:text-left",
						"center": "",
					},
					"size": {
						"large": "text-4xl",
						"small": "text-2xl",
					},
				},
				CompoundVariants: []gocva.CompoundVariant{
					{
						Class: "md:ml-24",
						Attrs: gocva.Variant{
							"direction": "left",
							"size":      "large",
						},
					},
					{
						Class: "md:ml-[4.5rem]",
						Attrs: gocva.Variant{
							"direction": "left",
							"size":      "small",
						},
					},
				},
				DefaultVariants: gocva.Variant{
					"direction": "left",
					"size":      "large",
				},
			}

			// act
			baseResult := cva.Render(gocva.Variant{
				"variant": "white",
			})
			centerResult := cva.Render(gocva.Variant{
				"variant":   "white",
				"direction": "center",
			})

			// assert
			Expect(
				baseResult,
			).To(Equal("relative text-center font-bold uppercase md:text-left text-4xl text-neutral-50 md:ml-24"))
			Expect(
				centerResult,
			).To(Equal("relative text-center font-bold uppercase text-4xl text-neutral-50"))
		})

		It("should override defaults when specified", func() {
			// arrange
			cva := gocva.CVA{
				Base: "mx-auto",
				Variants: gocva.Variants{
					"padding": {
						"default": "px-4 sm:px-8 md:px-16 lg:px-36",
						"none":    "",
					},
					"max": {
						"default": "max-w-screen-xl",
						"none":    "",
					},
				},
				DefaultVariants: gocva.Variant{
					"padding": "default",
					"max":     "default",
				},
			}

			// act
			withNoDefaultResult := cva.Render(gocva.Variant{
				"padding": "none",
				"max":     "none",
			})
			withDefaultResult := cva.Render(gocva.Variant{
				"padding": "",
			})

			// assert
			Expect(
				withNoDefaultResult,
			).To(Equal("mx-auto"))
			Expect(
				withDefaultResult,
			).To(Equal("mx-auto max-w-screen-xl px-4 sm:px-8 md:px-16 lg:px-36"))
		})
	})
}
