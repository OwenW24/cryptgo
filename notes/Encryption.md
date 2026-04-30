**Diffie Hellman Key Exchange**
based on discrete log problem $g^x \equiv h\mod{p}$
- public: $p, g$

Alice 
- secret $a$
- public $A \equiv g^a \mod{p}$
Bob
- secret $b$
- public $B \equiv g^b \mod{ p}$

Shared key
- $A^b \equiv g^{ab} \equiv B^a \mod{p}$

**RSA**
based on factoring problem pq
Alice 
- private: $p, q$ 
- public: $N = pq$, $e$ st $\gcd(e, (p-1)(q-1)) = 1$ 

Bob
- encrypt: $c \equiv m^e \mod{N}$

Alice
- decrypt $m \equiv c^{-e} \mod{N}$

**Goldwasser-Micali**
Alice
- private p, q
- a s.t. $\left( \frac{a}{p} \right) = \left( \frac{a}{q} \right) = -1$

Bob 
- $m \in \{ 0 ,1 \}$
- random $r$ s.t. $1<r<N$ and $\gcd(r, N) = 1$
- Encrypt
$$
c \equiv a^m r^2 \mod{N}
$$

Alice
- decrypt
$$
m \equiv \begin{cases}
0  & \text{if } \left( \frac{c}{p} \right) =1 \\
1  & \text{if } \left( \frac{c}{p} \right) =-1 \\
\end{cases}
$$



