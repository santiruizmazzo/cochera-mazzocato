# language: es
Característica: Filtrar pagos
    Como administrador
    quiero obtener los pagos que cumplan con un filtro
    para únicamente ver información de pagos de ese grupo

    Escenario: Ningún filtro seleccionado
        Dado que existe un pago con id 1
        Y existe un pago con id 2
        Y no seleccioné ningún filtro
        Cuando intento obtener todos los pagos
        Entonces obtengo toda la información del pago 1
        Y toda la información del pago 2
    
    Escenario: No hay pagos
        Dado que no hay pagos creados
        Cuando intento obtener todos los pagos
        Entonces se me informa que no existen pagos creados
    
    Escenario: Filtro por id de inquilino
        Dado que existe un inquilino con id 1
        Y que existe un inquilino con id 2
        Y que existe un pago con id 1 que realizó el inquilino 1
        Y que existe un pago con id 2 que realizó el inquilino 1
        Y que existe un pago con id 3 que realizó el inquilino 2
        Y que seleccioné el filtro de inquilino con id 2
        Cuando intento obtener todos los pagos
        Entonces obtengo toda la información del pago con id 3