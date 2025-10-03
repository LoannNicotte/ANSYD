# 1. Message passing system
## 1/
Quand les deux sleep ont la même durée, cela ne change pas grand-chose, car l'action est gérée avec le channel qui donne la balle à celui qui attend.

Oui, le code est event-driven, car tant que le joueur ne reçoit pas de message, il n’avance pas et attend.

Oui, le code peut être simplifié en utilisant une seule fonction player, car le time.Sleep n’impacte pas l’ordre des actions.

## 2/
La struct Ball contient maintenant une vecteur d'entier. J'ai egalement cree une fonction inc qui prend en arguement une Ball et incremeent de k la k-ième composante du vecteur. Je reprend ensuite le même code qui gere le ping-pong entre les 2 joueurs.

## 3/
La probabilité est de 1/N_u car le tour d'avant, le token été sur le noeud u ayant N_u voisins et comme la distribution est uniforme on trouve 1/N_u

## 4/
La topologie de mon système est créée dans init_chan, qui crée les canaux entre chaque nœud. Ensuite, pour chaque joueur, on lui donne un nom ainsi que la liste des canaux auxquels il a accès. Le joueur attend un message de chacun de ses canaux ; dès qu’il reçoit un message, il l’annonce en affichant : "Reçu : Nom", puis choisit aléatoirement un canal vers lequel renvoyer le token. le token est envoyé toutes les 200 ms

Le programme s'arrete apres plus de 3 secondes et au moment ou un message est envoyer sur le chanel 4