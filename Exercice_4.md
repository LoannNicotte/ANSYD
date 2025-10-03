# 4. Leader election

## 1/ For every $k$, what is the upper bound on the number of processes that are phase $k$ winners?
Un processus ne peut être gagnant à la phase $k$ que si son identifiant est le plus grand dans son $f(k-1)$-voisinage (à gauche et à droite).  
Cela implique que deux gagnants distincts ne peuvent pas être trop proches et donc être separé d'au moins $f(k-1)$

Par conséquent, le nombre total de gagnants est borné par le nombre de « blocs » de taille $f(k-1)$ dans l’anneau, soit :  

$$
\text{Nb gagnants} \;\leq\; \frac{n}{f(k-1)}
$$

## 2/ How many phases are needed to have only one leader?
On obtient un leader dès qu’un message `<probe>` fait **le tour complet de l’anneau**,  
c’est-à-dire lorsque :

$$
f(k) \;\geq\; n
$$

Comme la fonction $f$ est strictement croissante, il existe une fonction réciproque $g$ telle que :

$$
k \;\geq\; g(n)
$$

---

**Conclusion :**  
Le nombre minimal de phases nécessaires est donc :

$$
k_{\min} = g(n).
$$

## 3/ When $f (k) = 2^{k}$, what is that the total number of messages exchanged to elect a leader?

À la phase $k$ :  
- chaque gagnant envoie au plus $2 \cdot f(k) = 2 \cdot 2^k$ messages (un dans chaque direction),  
- le nombre de gagnants est au plus $\dfrac{n}{f(k-1)} \approx \dfrac{n}{2^{k-1}}$.  

Donc le nombre de messages à la phase $k$ est borné par :  

$$
M(k) \leq \frac{n}{2^{k-1}} \cdot 2 \cdot 2^k = 4n
$$

---

Le nombre total de messages échangés pendant toutes les phases est alors :  

$$
M_{\text{total}} \leq 4n \cdot \log_2(n)
$$

---

**Election :**  
Une fois le leader élu, il doit envoyer un message de terminaison qui fait le tour complet de l’anneau pour informer tous les processus.  
Cela nécessite **$n$ messages supplémentaires**.  

---

**Conclusion :**  
Le nombre total de messages échangés pour élire un leader est donc de l’ordre de :  

$$
O(n \log n + n) = O(n \log n).
$$
