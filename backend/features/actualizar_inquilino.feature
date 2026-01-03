# language: es
Característica: Actualizar inquilino
    Como administrador
    quiero actualizar datos de un inquilino
    para corregir errores o actualizar información sobre él

    Escenario: Inquilino existe
        Dado que existe el inquilino con id 1
        Y que su DNI es 14571272
        Cuando intento actualizar su DNI a 43295798
        Entonces se modifica correctamente
    
    Escenario: Inquilino no existe
        Dado que no existe el inquilino con id 4
        Cuando intento actualizar su DNI a 43295798
        Entonces se informa que ese inquilino no existe

    Escenario: Modificación de id
        Dado que existe el inquilino con id 1
        Cuando intento actualizar su id a 2
        Entonces se informa que el id no se puede actualizar

    Escenario: Nuevo email inválido
        Dado que existe el inquilino con id 1
        Y que su email es juanalberto@garcia.com
        Cuando intento actualizar su email a illojuan.com
        Entonces se informa que el nuevo email no es válido

    Escenario: Eliminar dato obligatorio
        Dado que existe el inquilino con id 1
        Y que su DNI es 777
        Cuando intento actualizar su DNI a vacío
        Entonces se informa que el DNI no se puede borrar

    Escenario: Dato a actualizar no existe
        Dado que existe el inquilino con id 1
        Cuando intento actualizar su fecha de nacimiento a 11-09-2001
        Entonces se informa que el dato a actualizar no existe
    
    Escenario: Múltiples datos a actualizar
        Dado que existe el inquilino con id 1
        Y que su DNI es 41630284
        Y que su nombre es Nicolás
        Cuando intento actualizar su DNI a 666 y su nombre a Lucifer
        Entonces se modifican correctamente
    
    Escenario: Nuevo DNI ya en uso
        Dado que existe el inquilino con id 1
        Y que su DNI es 777
        Y que existe el inquilino con id 2
        Y que su DNI es 1234
        Cuando intento actualizar el DNI del inquilino 1 a 1234
        Entonces se informa que ese DNI ya está en uso
    
    Escenario: Nuevo email ya en uso
        Dado que existe el inquilino con id 1
        Y que su email es shigeru@nintendo.com
        Y que existe el inquilino con id 2
        Y que su email es shigeru@miyamoto.com
        Cuando intento actualizar el email del inquilino 1 a shigeru@miyamoto.com
        Entonces se informa que ese email ya está en uso