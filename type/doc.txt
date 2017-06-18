** Structure:
- struct keeps data (memory layout)
- function with receiver becomes method for struct
- value_recv or pointer_recv
- value_rec, recv pass by value so modificaiton will not reflect
- pointer_rec, recv pass by reference so modificaiton will reflect

** Embedding:
- composition over inheritance
- methods/fields of inner_embedded types available to outer_type
- multiple inheritance is possible

** Alias:
- in-built type can given type name (alias)
- attach method on alias type name
- no boxing / unboxing done

** Orthogonal:
- there are minimum way to achieve given thing
- common way apply to multiple things

** Interface:
- set of methods [less methods, one,two,zero ;) ]
- any type having those methods implicitly implement interface
- type assertion
- type switch

** Retrofitted:
- PaidZip / MyZip - Zipper interface
- little modification (smart way)

==========================================================================================

** What next?
- diamond possible wrt multiple inheritance
- multiple inheritance with same field name / method name
- pointer variable for interface
- interface method with pointer receiver

** Links:
- Another Go at Language Design : http://web.stanford.edu/class/ee380/Abstracts/100428-pike-stanford.pdf
- CS university course by Rob pike : https://www.cs.cmu.edu/afs/cs.cmu.edu/academic/class/15440-f11/go/doc/docs.html
- When to use value / pointer receiver : https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
- Methods, Interface and Interface type : https://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html
