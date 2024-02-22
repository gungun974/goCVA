# Overview

goCVA is a Go package inspired by [Class Variance Authority](https://cva.style/docs) for dynamic class name management, ideal for conditional styling.

# Getting Started

## Installation

Install goCVA with Go's package manager:

```bash
go get github.com/gungun974/gocva
```

Import goCVA into your project:

```go
import "github.com/gungun974/gocva"
```

# Usage

## Basic Usage

1. **Creating a CVA Instance**: Initialize `gocva.CVA` to start using the library.

```go
cva := gocva.CVA{}
```

2. **Rendering Without Parameters**: Without parameters, `Render` returns an empty string.

```go
result := cva.Render(gocva.Variant{})
// result is ""
```

3. **Adding a Base Class**: Specify a base class for consistent inclusion.

```go
cva := gocva.CVA{
    Base: "base-class",
}
result := cva.Render(gocva.Variant{})
// result is "base-class"
```

## Working with Variants

1. **Defining Variants**: Set up variants for conditional class names.

```go
cva := gocva.CVA{
    Variants: gocva.Variants{
        "variant": {
            "primary": "variant-primary",
            "secondary": "variant-secondary",
        },
        "size": {
            "large": "text-large",
            "small": "text-small",
        },
    },
}
```

2. **Rendering Variants**: Choose variants to apply during rendering.

```go
result := cva.Render(gocva.Variant{
    "variant": "secondary",
    "size":    "large",
})
// result is "text-large variant-secondary"
```

## Default Variants

Set default variants for automatic application when none are specified.

```go
cva := gocva.CVA{
    Variants: gocva.Variants{
        "variant": {
            "primary":   "variant-primary",
            "secondary": "variant-secondary",
        },
    },
    DefaultVariants: gocva.Variant{
        "variant": "primary",
    },
}
result := cva.Render(gocva.Variant{})
// result is "variant-primary"
```

## Compound Variants

Define classes based on multiple variant conditions.

```go
cva := gocva.CVA{
    CompoundVariants: []gocva.CompoundVariant{
        {
            Class: "md:ml-24",
            Attrs: gocva.Variant{
                "direction": "left",
                "size":      "large",
            },
        },
    },
}
result := cva.Render(gocva.Variant{
    "direction": "left",
    "size":      "large",
})
// result is "md:ml-24"
```

## Real-world Example

Manage button styles based on state and size with goCVA.

```go
cva := gocva.CVA{
    Base: "btn",
    Variants: gocva.Variants{
        "state": {
            "active": "btn-active",
            "disabled": "btn-disabled",
        },
        "size": {
            "large": "btn-large",
            "small": "btn-small",
        },
    },
}
result := cva.Render(gocva.Variant{
    "state": "active",
    "size": "large",
})
// result is "btn btn-active btn-large"
```

# How run test
Use `make test`
