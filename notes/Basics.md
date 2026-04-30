*let $m$ be some integer* and $p$ be some prime integer


**Groups** $(G, *)$ Satisfies:
- identity: $\exists e\in G$ s.t. $e*a = a*e = a$
- inverse: $a * a^{-1} = a^{-1} * a = e$
- associative: $(a*b)*c = a*(b*c)$
If $G$ has finite elements, it is called a finite group. Number of elements in a finite group is called its order, denoted $|G|$ 
$G$ is abelian if it satisfies commutativity: $a*b = b*a$ 
Let $a \in G, k \in \mathbb{Z}^+$, positive powers of $a$ are denoted
$$
a^k = a*a* \cdots *a
$$$$
a^0 = e, \quad a^{-k} = (a^{-1})^k
$$
A group has finite order if $a^k = e$ for some positive integer $k$, else the group has infinite order

---

**1 Congruence Modulo** 
$$
m | a-b \iff a \equiv b \mod{m} 
$$
**2 Modular Ring** 
$$
\mathbb{Z} / m \mathbb{Z} = \{ 0, 1, 2, \dots, m - 1 \}
 $$
- abelian group 

**3 Invertibility**
$$
a, b \in \mathbb{Z}/m \mathbb{Z} \quad \text{ s.t. } \quad ab \equiv 1 \mod{m} 
$$
$$
b \equiv a^{-1} \mod{m} 
$$
$$
ab \equiv 1 \mod{m}  \iff \gcd(a, m) = 1
$$
**4 Modular Group**
$$
\mathbb{F}_{p}^*= (\mathbb{Z} / m \mathbb{Z})^* = \{ a \in \mathbb{Z} / m \mathbb{Z} | \gcd(a, m) = 1 \}
$$
$$
(\mathbb{Z} / p \mathbb{Z})^* = \{1, 2, \dots, p-1 \}
$$
$\mathbb{F}_{p}^*$ properties
$$
\begin{align}
\text{closure: } & ab \in \mathbb{F}_{p}^* \\
\text{identity: } & 1 \in \mathbb{F}_{p}^*, \quad 1\cdot a \equiv a  \mod{p} \\
\text{inverses: } &  a^{-1} \in \mathbb{F}_{p}^* \quad \text{s.t} \quad a\cdot a^{-1} \equiv 1 \mod{p} \\
\text{associativity: } &  (a)bc \equiv a(bc) \mod{p}   \\
\text{commutativity: } & ab \equiv ba \mod{p}  \\
\end{align}
$$

**5 Euler's Phi**
$$
\phi(m) = |(\mathbb{Z} / m \mathbb{Z})^* |
$$
$$
\phi(p) = p-1
$$

**6 Order** of $a$ modulo $p$ is the smallest exponent $k\geq 1$ such that
$$
a^k \equiv 1 \mod{ p} 
$$
Suppose $a$ has order $k$ modulo $p$ 
$$
a^n \equiv 1 \mod{p} \iff k | n 
$$
**7 Primitive Roots** $\exists g \in \mathbb{F}_{p}^*$ whose powers give every element of $\mathbb{F}_{p}^*$
$$
\mathbb{F}_{p}^* = \{ 1,g,g^2, g^3, \dots, g^{p-2} \}
$$

---

**8 Fermat's Little Theorem**$$
a \not \equiv 0 \mod{p}  \implies a ^{p-1} \equiv 1 \mod{ p} $$**9 Unique $e$th Root mod $p$** : given some $e$ such that $\gcd(e, p-1) = 1$, let $d \equiv e^{-1} \mod{ p-1}$ $$
x^e \equiv c \mod{p} \quad \text{has unique solution}\quad x\equiv c^d \mod{p} 
$$
**10 Euler's Formula for for $pq$**
$$
a^{(p-1)(q-1)} \equiv 1 \mod{pq} 
$$

**11 Unique $e$th Root mod $pq$** : given some $e$ such that $\gcd(e, (p-1)(q-1)) = 1$, let $d \equiv e^{-1} \mod{ (p-1)(q-1)}$ $$
x^e \equiv c \mod{pq} \quad \text{has unique solution}\quad x\equiv c^d \mod{pq} 
$$
---

**12 General Chinese Remainder**
$$
\begin{align}
x \equiv &  a_{1} \mod{m_{1}}  \\
x \equiv &  a_{2} \mod{m_{2}}  \\ 
\vdots &  \\
x \equiv &  a_{n} \mod{m_{n}}  \\
\end{align}
$$
The solution is unique modulo $m_{1}m_{2}\cdots m_{n}$


**Dividing/Simplifying Some Modular Equation**
If $d$ divides $a,b,n$ 
$$
ax \equiv b \mod{n} \quad \Rightarrow \quad \frac{a}{d} x \equiv \frac{b}{d} \mod{\frac{n}{d}} 
$$


---

**13 Square Roots modulo $p$**: Let $p$ be an odd prime and $a \not \equiv 0 \mod{p}$. Then $x^2 \equiv a \mod{p}$ either has 2 solutions or no solutions in $\mathbb{Z} / p \mathbb{Z}$.

**14 Quadratic Residue modulo $p$**: $a$ is a quadratic residue modulo $p$ 
$$
\exists c \quad \text{s.t.} \quad c^2 \equiv a \mod{p} 
$$
Properties
$$
\begin{align}
QR \cdot QR = &  QR  \\
QR \cdot NR = & NR \\
NR \cdot NR = & QR
\end{align}
$$
**15 Legendre Symbol of $a$ modulo $p$**
$$
\left( \frac{a}{p} \right) = 
\begin{cases}
1 & \text{if $a$ is QR mod $p$}  \\
-1 & \text{if $a$ is NR mod $p$} \\
0 & \text{if $p|a$}
\end{cases}
$$
$$
\begin{align}
\left( \frac{ab}{p} \right) = &  \left( \frac{a}{p} \right) \left( \frac{b}{p} \right)  \\
\left( \frac{a^m}{p} \right) = &  \left( \frac{a}{p} \right)^m \\
a \equiv b \mod{p} \implies &  \left( \frac{a}{p} \right) = \left( \frac{b}{p} \right)
\end{align}

$$
Law of Quadratic Reciprocity. Let $p$ and $q$ be odd primes. 
$$
\left( \frac{p}{q} \right) = \begin{cases}
\left( \frac{q}{p} \right) & \text{if $p \equiv 1 \mod 4$ or $q \equiv 1 \mod 4$}  \\
-\left( \frac{q}{p} \right) & \text{if $p \equiv 3 \mod 4$ and $q \equiv 3 \mod 4$} 
\end{cases}
$$
$$
\begin{align}
\left(\frac{-1}{p} \right) &  = \begin{cases}
1  & \text{if $p \equiv 1 \mod 4$} \\
-1  & \text{if $p \equiv 3 \mod 4$}
\end{cases} \\
\left(\frac{2}{p} \right) & = \begin{cases}
1  & \text{if $p \equiv \pm1 \mod 8$} \\
-1  & \text{if $p \equiv \pm3\mod 8$} 
\end{cases}
\end{align}
$$

**16 Jacobi symbol of $a$ modulo $b$** ($b$ positive and odd)
$$
\begin{align}
b = & p_{1}^{e_{1}}p_{2}^{e_{2}} \cdots p_{r}^{e_{r}} \\
\left( \frac{a}{b} \right) = & \left( \frac{a}{p_{1}} \right)^{e_{1}} \left( \frac{a}{p_{2}} \right)^{e_{2}} \cdots \left( \frac{a}{p_{r}} \right)^{e_{r}}
\end{align}
$$
When $b$ is composite, a result of 1 does not necessarily mean $a$ is QR mod $b$
$$
\begin{align}
\left( \frac{a_{1} a_{2}}{b} \right) = &  \left( \frac{a_{1}}{b} \right) \left( \frac{a_{2}}{b} \right)  \\
\left( \frac{a}{b_{1}b_{2}} \right) = &  \left( \frac{a}{b_{1}} \right) \left( \frac{a}{b_{2}} \right)  \\
a_{1} \equiv a_{2} \mod{b} \implies &  \left( \frac{a_{1}}{b} \right) = \left( \frac{a_{2}}{b} \right)
\end{align}
$$
$$
\left( \frac{a}{b} \right) = \begin{cases}
\left( \frac{b}{a} \right) & \text{if $a \equiv 1 \mod 4$ or $b \equiv 1 \mod 4$}  \\
-\left( \frac{b}{a} \right) & \text{if $a \equiv 3 \mod 4$ and $b \equiv 3 \mod 4$} 
\end{cases}
$$
$$
\begin{align}
\left(\frac{-1}{b} \right) &  = \begin{cases}
1  & \text{if $b \equiv 1 \mod 4$} \\
-1  & \text{if $b \equiv 3 \mod 4$}
\end{cases} \\
\left(\frac{2}{b} \right) & = \begin{cases}
1  & \text{if $b \equiv \pm1 \mod 8$} \\
-1  & \text{if $b \equiv \pm3\mod 8$} 
\end{cases}
\end{align}
$$
---
**17 Elliptic Curve $E$** is the set of solutions to a Weierstrass equation
$$
E: y^2 = x^3 + ax + b, \quad \text{where there are not singularities } 4a^3 + 27b^2 \neq 0
$$
Together with extra point $\mathcal{O}$

Group Structure: Let $P,Q,R \in E$ 
$$
\begin{align}
\text{identity: } &  P + \mathcal{O} = \mathcal{O} + P= P \\
\text{inverse: } & P + (-P) = \mathcal{O} \\
\text{associative: } & (P+Q) + R = P + (Q + R) \\
\text{commutative: } & P+Q = Q + P
\end{align}
$$
**18 Elliptic Curve Addition**
$$
\lambda = 
\begin{cases}
\frac{y_{2} - y_{1}}{x_{2} - x_{1}} & \text{if } P \neq Q \\
\frac{3x_{1}^2 + a}{2y_{1}} & \text{if } P = Q
\end{cases}
$$
$$
\begin{align}
x_{3} = & \lambda^2 -x_{1}-x_{2} \\ \\
y_{3} = & \lambda(x_{1} - x_{3}) - y_{1}
\end{align}
$$

**19 Elliptic Curve Over Finite Fields**
$$
E(\mathbb{F}_{p}) = \{ (x,y) | x,y \in \mathbb{F}_{p} \text{ satisfying } y^2 = x^3 + ax + b \} \cup \{\mathcal{O}\}
$$

$$
\lambda \equiv
\begin{cases}
(y_{2} - y_{1})(x_{2} - x_{1} )^{-1} \mod{p} & \text{if } P \neq Q \\
(3x_{1}^2 + a)(2y_{1})^{-1}\mod{p} & \text{if } P = Q
\end{cases}
$$
$$
\begin{align}
x_{3} \equiv & \lambda^2 -x_{1}-x_{2} \mod{p} \\
y_{3} \equiv & \lambda(x_{1} - x_{3}) - y_{1} \mod{p} 
\end{align}
$$


**20 Elliptic Curve Discrete Log Problem***
Problem of finding some integer $n$ such that $nP = Q$, where $n = \log_P(Q)$ 

