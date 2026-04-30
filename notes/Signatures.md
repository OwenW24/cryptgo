**RSA**
Samantha
- secret: $p, q$
- verification exponent $e$ s.t. $\gcd(e, (p-1)(q-1))$
- signing exponent $d \equiv e^{-1} \mod{(p-1)(q-1)}$
- public: $N = pq, e$ 

Samantha
- Sign document $D$ where $1 < D < N$ 
$$
S \equiv D^d \mod{N} 
$$
Victor 
- Verify document
$$
S^e \equiv ? \quad D\mod{ N} 
$$

**Elgamal**
Samantha
- public: prime $p$, primitive root $g \mod{p}$
- secret signing key a s.t. $1<a<p$ 
- public: verification key $A \equiv g^a \mod{p}$ 

Samantha
- Sign document $D$ where $1 < D < p$ 
- Choose random k s.t. $1<k<p$ and $\gcd(k, p-1) = 1$
Compute signature
$$
S_{1} \equiv g^k \mod{p}  \quad S_{2} \equiv (D-aS_{1})k^{-1} \mod{ p-1} 
$$
- $S = (S_1, S_2)$

Victor
- Verify document
$$
A^{S_{1}} S_{1}^{S_{2}} \mod{p}  \equiv? \quad g^D \mod{p} 
$$

