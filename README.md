# Go Data Structures & Algorithms

> A collection of data structures and algorithms implemented in Go for personal reference and education.

## Data Structures

Bellow are a list of common data-structures created in Go with tests, implementations are grouped by abstraction and link to their associated wikipedia entry.

* ### List

  * [Linked List](https://en.wikipedia.org/wiki/Linked_list)
    * The field of each node that contains the address of the next node is called the 'next link' or 'next pointer'.
    * The 'head' of a list is its first node. The 'tail' of a list may refer either to the rest of the list after the head, or to the last node in the list.
    * Used in many list, queue, and stack implementations.
    * Great for creating circular lists.
    * Can easily model real world objects such as trains.
    * Used in separate chaining, which is present in certain hashable implementations to deal with hashing collisions.
    * Often used in the implementation of adjacency lists for graphs.
  
  * [Double-Linked List](https://en.wikipedia.org/wiki/Linked_list#Doubly_linked_list)
    * In a 'doubly linked list', each node contains, besides the next-node link, a second link field pointing to the 'previous' node in the sequence. The two links may be called 'forward('s') and 'backwards', or 'next' and 'prev'('previous').
    * Takes 2x the memory than a singly linked list.

* ### Queue

  * [Stack](https://en.wikipedia.org/wiki/Stack_(abstract_data_type)) - A stack can be easily implemented either through an array or a linked list. What identifies the data structure as a stack in either case is not the implementation but the interface: the user is only allowed to pop or push items onto the array or linked list, with few other helper operations.

* ### Map

  * Tree Map
  * Hash Map / Hash Table

## More information

[Data Structures Easy to Advanced Course [Youtube]](https://www.youtube.com/watch?v=RBSGKlAvoiM)
