# What do three dots “./…” mean in Go command line invocations?

**source: https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations**

```
An import path is a pattern if it includes one or more "..." wildcards, each of which can match any string, including the empty string and strings containing slashes. Such a pattern expands to all package directories found in the GOPATH trees with names matching the patterns. As a special case, x/... matches x as well as x's subdirectories. For example, net/... expands to net and packages in its subdirectories.
```
