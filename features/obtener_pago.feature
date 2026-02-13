# language: es
Característica: Obtener pago de inquilino
    Como administrador
    quiero obtener el pago de un inquilino
    para consultar su información

    Escenario: Pago existe
        Dado que existe el pago con id 1
        Cuando intento obtener su información
        Entonces se me brindan todos sus datos
    
    Escenario: Pago no existe
        Dado que no hay pagos creados
        Cuando intento obtener la información del pago con id 10
        Entonces se me informa que el pago buscado no existe