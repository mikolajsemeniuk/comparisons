# Comparisons
the purpose of this repository is to check difference between different approaches like:
* monomorphized
* monomorphized with pointer
* interface
* interfaces with pointer
* generics
* generics with pointers 

```sh
go test ./... -bench=. -benchmem
-bench=. # show time
-benchmem # show memory
```