Solve for $x$, given $g^x \equiv h \mod{p}$

---

**Baby-Giant**
$N = \text{ord}_{p}(g)$, $b = \lfloor \sqrt{ N } \rfloor +1$
for $k = 0 \rightarrow b$ 

| baby: $g^k \mod{ p}$ | giant: $h(g^{-b})^k \mod{p}$ |
| -------------------- | ---------------------------- |
| $g^0 \mod{ p}$       | $h(g^{-b})^0 \mod{p}$        |
| $g^1 \mod{ p}$       | $h(g^{-b})^1 \mod{p}$        |
| $\vdots$             | $\vdots$                     |
| $g^b\mod{ p}$        | $h(g^{-b})^b \mod{p}$        |
find $i, j$ that satisfy $g^i \equiv h(g^{-b})^j \mod{ p}$ 
$x \equiv i +jb \mod{p-1}$

---

**Index calculus**
To solve DLP $g^x \equiv h \mod{ p}$, where $g$ is a primitive root
**1** Factor Base: Set smoothness bound $B$ and consider all primes $\leq B$ 
**2** Relation Building: Use random trials to find $m_{1}, m_{2}, \dots,m_{r}$ such that
$$
g^{m_{i}} \equiv p_{1}^{m_{i,1}}p_{2}^{m_{i,2}} \cdots p_{r}^{m_{i,r}} \mod{p} 
$$
is $B$-smooth. 

**3** Elimination: Solve $r\times r$ system to find $\log_g(p_i)$ 
$$
\begin{align}
m_{1} \equiv& m_{1,1}\log_{g}(p_{1}) +m_{1,2}\log_{g}(p_{2}) + \cdots + m_{1,r}\log_{g}(p_{r}) \\
m_{2} \equiv& m_{2,1}\log_{g}(p_{1}) +m_{2,2}\log_{g}(p_{2}) + \cdots + m_{2,r}\log_{g}(p_{r}) \\
\vdots& \\
m_{r} \equiv& m_{r,1}\log_{g}(p_{1}) +m_{r,2}\log_{g}(p_{2}) + \cdots + m_{r,r}\log_{g}(p_{r})
\end{align}
$$
**4** Log Computation: Use random trials to find $k$ such that $hg^{-k} \mod{p}$ is $B-$smooth

---

**Collision** randomly choose exponents $t \in 1 < n < N$ and find $g^i \equiv hg^j \mod{p}$, thus $x \equiv i -j \mod{p-1}$.

