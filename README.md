# BigOrder

BigOrder is a simple tool that can be used to order objects in a manner of best to worst using user inputs. BigOrder has arisen from the innate difficulty of comparing and ranking a large set of items. With the rise of tier lists, ordering and sorting into ordered categories has become a huge part of internet culture, and in-person discussions as a whole. With that in mind, BigOrder was conceived as a way to remove the complexity from the ranking process.

## Features

Currently, BigOrder supports a StrictOrderer, where strict preferences between two items are required in order to sort elements properly. The current implementation is concurrency-friendly and also offers a reasonably low number of human comparisons, as it uses a binary selection sort.

### To-Do List (for v1.0.0)

-   [x] Sorting algorithm implementation
-   [x] StrictOrderer implementation and example usage
-   [ ] PartialOrderer implementation and example usage
-   [ ] JSON import of data to be sorted
-   [ ] Documentation
-   [ ] Tests

## How To Use

An orderer can be defined using the provided constructor which takes in an array of elements of the `Item` type found in package `items`.

The `sort()` method must be called in a goroutine. The method `GetNextComparison()` retrieves the next comparison to be made by the user, and `SendNextComparison(higher, lower item.Item)` sends the comparison results. The handling of human comparisons can be done as a separate goroutine (see [example](https://github.com/rohan-av/bigorder/blob/master/examples/main.go)).
