# Latest Changes
Added nano syntax highlighting  
Added Code syntax highlighting

The following functions & built-ins now modify the arrays and hashes in place. That is, they no longer return results from copies of the data.  
This has made entry/retrieve/update actions much much faster, especially when working with large arrays and hashes.   

Basically, with everything global and lots of built-ins now work directly on the hashes and arrays you need to be aware that you may/will need to add a copy() somewhere in your code if you want things to work as normal.  
 
I expect you'll work it out.  It now works like most other languages  

**copy()**  
**delete()**  
**extend()**  
**insert()**  
**push()**  
**reverse()**  
**rsort()**
**set()**  
**shift()**  
**sort()**  
**swap()**

**sec2time()** is moved to a function  
**sorted()** is moved to function  

**type()** now shows type null  

include files have been amended to cover changes with hash and array builtins.  
Files in the examples directory have been updated

It now means arrays/hashes passed as parameters to a function will be fully global, so any changes to the hash/array within the function will also happen outside the function.  

function sillysort(Array) {  
    return sort(Array)  
}  

function sillysort2(Array) {  
    arr=copy(Array)  
    return sort(arr)  
}  

a=[2,5,4,3,1]  
b=sillysort(a)  
println("a=",a)  
println("b=",a)  

a=[2,5,4,3,1]  
b=sillysort2(a)  
println("a=",a)  
println("b=",b)  
reverse(b)  
println("b=",b)  
