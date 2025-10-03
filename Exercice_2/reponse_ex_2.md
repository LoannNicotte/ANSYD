# 2. Leader election on a ring
## 1/
Le message doit contenir deux informations : l’ID du plus grand noeud rencontré et un booléen indiquant si le leader a été élu, ce qui permettra de propager cette information dans le système.

## 2/
Pour un anneau composé de 3 nœuds, il faut un channel entre chaque nœud ainsi qu’un channel pour chaque nœud pour donner la décision finale. Il faut donc 6 canaux différents au total.

A la fin du main, j'attend le messages de chacun des noeuds pour valider l'election du leader et mettre fin au script.

## 3/
Dans mon implémentation, j’ai choisi de mettre un buffer de taille 1 sur les canaux c1, c2 et c3 afin d’éviter un blocage immédiat lors de l’envoi du premier message par chaque nœud : cela permet de lancer la propagation du processus d’élection sans que les goroutines se retrouvent en attente mutuelle. 

Ensuite, j’ai structuré la fonction node avec des conditions (if/else if) pour traiter les différents cas possibles : réception d’un message de terminaison, réception d’un identifiant plus grand que le sien, ou réception du même identifiant qui signe la fin de l’élection. 

Enfin, dans le main, j’attends explicitement la décision de chaque nœud (m1, m2, m3) afin de m’assurer qu’ils ont tous convergé vers la même valeur maximale avant de conclure l’élection.

## 4/
Le code compile correctement et les trois nœuds convergent bien vers la même décision.

Une amélioration possible du code, en cas de crash d’un nœud, serait d’ajouter un mécanisme d’accusé de réception. Concrètement, à chaque envoi de message, le nœud expéditeur attendrait une confirmation de bonne réception. Si cette confirmation n’arrive pas dans un délai prédéfini, on considérerait que le nœud destinataire est tombé en panne. Dans ce cas, il ne serait plus possible de mener à bien l’élection, et une erreur serait renvoyée dans le canal Decision. De plus, il est important de veiller à ce qu’aucun nœud ne reste bloqué indéfiniment : chaque attente doit donc être limitée dans le temps. Au-delà de cette limite, le nœud en question doit être déclaré comme mort et l’algorithme d’élection interrompu proprement.
