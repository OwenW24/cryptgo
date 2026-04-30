Factoring some integer $N$

**Pollard's $p-1$**
$a =2$ 
for $n = 2 \rightarrow \text{bound}$ 
compute $a^{n!} \mod{N}$, $d=\gcd(a^{n!} - 1, N)$
- $d = 1$ : n++
- $d=n$ : restart with new $a$
- $1 < d < N$ : return $d$ and terminate

| n        | $a^{n!} \mod{N}$       | $d=\gcd(a^{n!} - 1, N)$ |
| -------- | ---------------------- | ----------------------- |
| 2        | $a^{2!} \mod{N}$       |                         |
| 3        | $(a^{2!})^3 \mod{N}$   |                         |
| $\vdots$ | $(a^{n-1!})^n \mod{N}$ |                         |

---

**Factorization via Difference of Squares**
*def 1*
Let $N$ be an integer and suppose we have $a, b \in \mathbb{Z} / N \mathbb{Z}$ such that
$$
a^2 \equiv b^2 \mod{ N}, \quad \text{but } a \not \equiv \pm b \mod{ N}  
$$
$N$ is composite, and $\gcd(N, a\pm b)$ gives a nontrivial factor of $N$

*def 2*
An integer $n$ is called $B-$smooth if all of its prime factors are $\leq B$. 

Algorithm:
(1 - Factor Base) 
Choose a smoothness bound $B$. 
Assemble Factor Base = set of primes $\leq$ $B$ 

(2 - Relation Building) 
Find integers $a_{1},a_{2},\cdots, a_{r}$ such that $a_{i}^2 \equiv c_{i} \mod{N}$, where $c_i$ is $B-$smooth

(3 - Elimination) Take a product of $c_i$'s so that every prime appearing in the factorization has an even power, creating perfect square.

(4 - GCD Computation)
$$
a^2 \equiv (a_{i_{1}}a_{i_{2}}\cdots a_{i_{s}})^2 \equiv (a_{i_{1}}^2a_{i_{2}}^2\cdots a_{i_{s}}^2) \equiv (c_{i_{1}}c_{i_{2}}\cdots c_{i_{s}}) \equiv b^2 \mod{ N} 
$$
if $a \not \equiv b \mod{N}$, $\gcd(N, a-b)$ is a nontrivial factor of $N$

---

**The Quadratic Sieve**
*def 1*
The following are equivalent
$$
\begin{align}
m|F(\alpha) & \iff \\
m|F(\alpha + km) \forall k \in \mathbb{Z}  & \iff \\
\alpha^2 \equiv N \mod{m} 
\end{align}
$$



Choose some smoothness bound $B$, Find $B$-powersmooth Factor base. 
Using range $a = \lfloor \sqrt{ N } \rfloor + 1$, $[a, a + m]$
$F(x) = x^2 - N$

| $x$:    | $a$       | $a+1$         | $\dots$  | $a+m$         |
| ------- | --------- | ------------- | -------- | ------------- |
| $F(x)$: | $a^2 - N$ | $(a+1)^2 - N$ | $\dots$  | $(a+m)^2 - N$ |
Now sieve $F(x)$ to find values that are $B$-powersmooth 
- (numbers that sieved to 1)

Assemble perfect square using the list of powersmooth numbers and solve using factorization via difference of squares
$$
\begin{align}
a^2 \equiv& b^2 \mod{N}  \\
\gcd(N, a\pm b)& \text{ gives a nontrivial factor of $N$}
\end{align}
$$
