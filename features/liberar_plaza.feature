# language: es
Característica: Liberar plaza
    Como administrador
    quiero liberar una plaza
    para registrar que no está ocupada por ningún inquilino

    Escenario: Plaza ocupada
        Dado que existe la plaza con id 1
        Y que está ocupada por el inquilino con id 1
        Cuando intento liberar la plaza
        Entonces se libera correctamente
    
    Escenario: Plaza no existe
        Dado que no existe la plaza con id 22
        Cuando intento liberar la plaza
        Entonces se informa que esa plaza no existe