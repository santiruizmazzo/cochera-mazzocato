# language: es
Característica: Filtrar inquilinos
    Como administrador
    quiero obtener los inquilinos que cumplan con un filtro
    para únicamente ver información de inquilinos de ese grupo

    Escenario: Ningún filtro seleccionado
        Dado que existe un inquilino con id 1
        Y existe un inquilno con id 2
        Y no seleccioné ningún filtro
        Cuando intento obtener todos los inquilinos
        Entonces obtengo toda la información del inquilino 1
        Y toda la información del inquilino 2
    
    Escenario: No hay inquilinos
        Dado que no hay inquilinos creados
        Cuando intento obtener todos los inquilinos
        Entonces se me informa que no existen inquilinos creados
    
    Escenario: Filtro por nombre que coincide por completo
        Dado que existe un inquilino llamado "Marcelo" con id 1
        Y que existe un inquilino llamado "Agustín" con id 2
        Y que existe un inquilino llamado "Marcelo" con id 3
        Y que seleccioné el filtro de nombre "Marcelo"
        Cuando intento obtener todos los inquilinos
        Entonces obtengo toda la información del inquilino 1
        Y obtengo toda la información del inquilino 3
    
    Escenario: Filtro por nombre que coincide parcialmente
        Dado que existe un inquilino llamado "Martín" con id 1
        Y que existe un inquilino llamado "Mario" con id 2
        Y que existe un inquilino llamado "Alonso" con id 3
        Y que existe un inquilino llamado "Lamar" con id 4
        Y que seleccioné el filtro de nombre "Mar"
        Cuando intento obtener todos los inquilinos
        Entonces obtengo toda la información del inquilino 1
        Y obtengo toda la información del inquilino 2
        Y obtengo toda la información del inquilino 4