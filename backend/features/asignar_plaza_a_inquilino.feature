# language: es
Característica: Asignar plaza a inquilino
    Como administrador
    quiero asignar una plaza a un inquilino
    para llevar registro de las plazas ocupadas por él

    Escenario: Plaza no ocupada
        Dado que existe el inquilino con id 1
        Y que existe la plaza con id 1
        Y que la plaza con id 1 no está ocupada
        Cuando intento asignar la plaza al inquilino
        Entonces se registra correctamente la asignación
    
    Escenario: Plaza ocupada
        Dado que existe el inquilino con id 1
        Y que existe la plaza con id 1
        Y que la plaza con id 1 está ocupada por el inquilino con id 2
        Cuando intento asignar la plaza al inquilino
        Entonces se informa que la plaza ya está ocupada

    Escenario: Plaza no existe
        Dado que existe el inquilino con id 1
        Y que no existe la plaza con id 3
        Cuando intento asignar la plaza al inquilino
        Entonces se informa que esa plaza no existe
    
    Escenario: Inquilino no existe
        Dado que no existe el inquilino con id 12
        Y que existe la plaza con id 3
        Cuando intento asignar la plaza al inquilino
        Entonces se informa que ese inquilino no existe