package gocva

import "sort"

type Variant map[string]string

type Variants map[string]Variant

type CompoundVariant struct {
	Class string
	Attrs Variant
}

type CVA struct {
	Base             string
	Variants         Variants
	CompoundVariants []CompoundVariant
	DefaultVariants  Variant
}

func (cva *CVA) mergeDefaultVariants(variants Variant) Variant {
	mergedVariants := make(Variant, len(cva.DefaultVariants))

	for k, v := range cva.DefaultVariants {
		mergedVariants[k] = v
	}

	for k, v := range variants {
		if v != "" {
			mergedVariants[k] = v
		}
	}

	return mergedVariants
}

func (cva *CVA) applyVariants(cb *ClassBuilder, variants Variant) {
	sortedVariants := make([]string, 0, len(cva.Variants))

	for k := range cva.Variants {
		sortedVariants = append(sortedVariants, k)
	}

	sort.Strings(sortedVariants)

	for _, k := range sortedVariants {
		v := cva.Variants[k]

		variant, ok := variants[k]

		if !ok {
			continue
		}

		for k, class := range v {

			if k != variant {
				continue
			}

			cb.Add(class)

		}
	}
}

func (cva *CVA) applyCompoundVariant(cb *ClassBuilder, variants Variant) {
	for _, v := range cva.CompoundVariants {
		isValidCompound := true

		for k, v := range v.Attrs {
			variant, ok := variants[k]

			if !ok {
				isValidCompound = false
				break
			}

			if variant != v {
				isValidCompound = false
				break
			}
		}

		if isValidCompound {
			cb.Add(v.Class)
		}
	}
}

func (cva *CVA) Render(userVariants Variant) string {
	cb := ClassBuilder{}

	cb.Add(cva.Base)

	variants := cva.mergeDefaultVariants(userVariants)

	cva.applyVariants(&cb, variants)

	cva.applyCompoundVariant(&cb, variants)

	return cb.String()
}
