# language: es
Característica: Obtener inquilino
    Como administrador
    quiero obtener la información de un inquilino
    para ver sus datos de contacto

    Escenario: Inquilino existe
        Dado que existe el inquilino con id 1
        Cuando intento obtener su información
        Entonces se me brindan todos sus datos de contacto
    
    Escenario: Inquilino no existe
        Dado que no hay inquilinos creados
        Cuando intento obtener la información del inquilino con id 10
        Entonces se me informa que el inquilino buscado no existe