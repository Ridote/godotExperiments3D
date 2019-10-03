# Proyecto RPG

## Servidor

Todo el procesamiento irá en la parte del servidor. El cliente mandará mensajes al servidor y el servidor responderá con la resolución del mensaje y si es válido o no.

Se mandarán los mensajes por WebSockets.

Los mensajes irán codificados con MSGPack. Se puede encontrar la librería para Godot en: https://github.com/xtpor/godot-msgpack

---

### Mensajes cliente a servidor

- Admin
    - `login` username, password
    - `register` username, password, email
    - `newPlayer` name, class?
- Mundo
    - `pos` ID, PX, PY, PZ, RX, RY
- Inventario
    - `equipItem` ID, itemID
    - `unequipItem` ID, itemID
    - `useItem` ID, itemID
    - `deleteItem` ID, itemID
- Acciones
    - `talk` ID, targetID
    - `answer` ID, targetID, option
---

### Mensajes servidor a cliente

- Admin
    - `auth` success, msg
    - `newPlayer` ID, name, model3D, owner (boolean)
- Eventos
    - 
- Info
    - `playerStats` ID, STR, AGI, INT
    - `playerState` ID, HP, MP
    - `pos` ID, PX, PY, PZ, RX, RZ
- Actiones
    - `inventory` ID
    - `useItem` ID, itemID
    - `newItem` ID, itemID
    - `deleteItem` ID, itemID
    - `equipItem` ID, itemID
    - `unequipItem` ID, itemID
    - `question` ID, targetID, options
    - `talk` ID, targetID, text

---

## Autenticación

- user / pass

---

## Estadísticas de personaje [stats]

- STR
    - HP
    - Attack
- AGI
    - Speed
    - Attack
- INT
    - MP
    - Magic attack
    
## Objetos
### Tipos
#### Armas
#### Armaduras
#### Accesorios
#### Consumibles
#### Objetos clave

## Inventario

### Estadísticas modificadas por objetos
- Armadura
- Defensa mágica
- Ataque elemental
- Defensa elemental

## Estado [state]

- HP
- MP

## Mecánicas

## Combate 

## Habilidades

## Personajes

## Historia 
