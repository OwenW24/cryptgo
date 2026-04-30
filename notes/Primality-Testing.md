**Fermat Primality Test** 
*def 1*
$$
a^{n-1} \not \equiv 1 \mod{n} \implies n \text{ is composite}
$$
$a$ is a Fermat witness for compositeness of $n$ 

Algorithm
given $n \geq 2$ and potential witness $a \in \{ 2, \dots, n-2 \}$
- $a^{n-1} \not \equiv 1 \mod{n} \implies n \text{ is composite}$
- otherwise, test fails, try new $a$
*There exists some composite integers without Fermat Witnesses, Carmichael Numbers*

**Miller Rabin Test**
algorithm:
Given an odd integer $n$ and potential witness $a \in \{ 1,2,\dots, n-1 \}$
find $k, m$ such that $n-1 = 2^k m$ 

if $a^m \equiv \pm 1 \mod{ n}$ test fails (i=9)
from $i=1 \rightarrow k-1$: if $a^{2^i m } \equiv -1 \mod{ n}$ test fails
if no failure, return composite

test with multiple witnesses to determine primality. 75% of numbers are MR witnesses



