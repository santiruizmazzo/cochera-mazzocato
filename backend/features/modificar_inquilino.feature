# language: es
Característica: Modificar inquilino
    Como administrador
    quiero modificar datos de un inquilino
    para corregir errores o actualizar información sobre él

    Escenario: Inquilino existe
        Dado que existe el inquilino con id 1
        Y que su DNI es 14571272
        Cuando intento modificar su DNI a 43295798
        Entonces se modifica correctamente
    
    Escenario: Inquilino no existe
        Dado que no existe el inquilino con id 4
        Cuando intento modificar su DNI a 43295798
        Entonces se informa que ese inquilino no existe

    Escenario: Modificación de id
        Dado que existe el inquilino con id 1
        Cuando intento modificar su id a 2
        Entonces se informa que el id no se puede modificar

    Escenario: Nuevo valor inválido
        Dado que existe el inquilino con id 1
        Y que su email es juanalberto@garcia.com
        Cuando intento modificar su email a illojuan.com
        Entonces se informa que el nuevo valor no es válido

    Escenario: Eliminar dato obligatorio
        Dado que existe el inquilino con id 1
        Y que su DNI es 777
        Cuando intento modificar su DNI a vacío
        Entonces se informa que el DNI no se puede borrar

    Escenario: Dato a modificar no existe
        Dado que existe el inquilino con id 1
        Cuando intento modificar su fecha de nacimiento a 11-09-2001
        Entonces se informa que el dato a modificar no existe
    
    Escenario: Múltiples datos a modificar
        Dado que existe el inquilino con id 1
        Y que su DNI es 41630284
        Y que su nombre es Nicolás
        Cuando intento modificar su DNI a 666 y su nombre a Lucifer
        Entonces se modifican correctamente
    
    Escenario: Nuevo valor ya en uso
        Dado que existe el inquilino con id 1
        Y que su DNI es 777
        Y que existe el inquilino con id 2
        Y que su DNI es 1234
        Cuando intento modificar el DNI del inquilino 1 a 1234
        Entonces se informa que ese DNI ya está en uso