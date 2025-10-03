# 2. Leader election on a ring
## 1/
Le message doit contenir deux informations : l’ID du plus grand noeud rencontré et un booléen indiquant si le leader a été élu, ce qui permettra de propager cette information dans le système.

## 2/
Pour un anneau composé de 3 nœuds, il faut un channel entre chaque nœud ainsi qu’un channel pour chaque nœud pour donner la décision finale. Il faut donc 6 canaux différents au total.

A la fin du main, j'attend le messages de chacun des noeuds pour valider l'election du leader et mettre fin au script.
