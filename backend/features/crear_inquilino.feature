# language: es
Característica: Crear inquilino
    Como administrador
    quiero crear un nuevo inquilino
    para guardar sus datos de contacto

    Regla de negocio: DNI, nombre y apellido obligatorios
        Escenario: Falta algún dato
            Dado que no ingresé el DNI, nombre o apellido
            Cuando intento crear al inquilino
            Entonces se informa que faltan datos obligatorios

    Regla de negocio: DNI es único
        Escenario: DNI ya existente
            Dado que existe un inquilino cuyo DNI es 12345678
            Cuando intento crear uno nuevo con el mismo DNI
            Entonces se informa que ese DNI ya está ocupado

    Regla de negocio: DNI es entero mayor a 1
        Escenario: DNI negativo
            Dado que ingresé un DNI igual a -1
            Cuando intento crear al inquilino
            Entonces se informa que el DNI debe ser un entero mayor a 1

        Escenario: DNI igual a 0
            Dado que ingresé un DNI igual a 0
            Cuando intento crear al inquilino
            Entonces se informa que el DNI debe ser un entero mayor a 1

        Escenario: DNI no numérico
            Dado que ingresé un DNI igual a "hola"
            Cuando intento crear al inquilino
            Entonces se informa que el DNI debe ser un entero mayor a 1
    
    Regla de negocio: Celular es entero mayor a 1
        Escenario: Celular negativo
            Dado que ingresé un número de celular igual a -1
            Cuando intento crear al inquilino
            Entonces se informa que el número de celular debe ser un entero mayor a 1

        Escenario: Celular igual a 0
            Dado que ingresé un número de celular igual a 0
            Cuando intento crear al inquilino
            Entonces se informa que el número de celular debe ser un entero mayor a 1
    
    Regla de negocio: Email tiene formato válido
        Escenario: Email inválido
            Dado que ingresé "santiago" como email
            Cuando intento crear al inquilino
            Entonces se informa que el email debe tener formato válido

        Escenario: Email válido
            Dado que ingresé "santiago@gmail.com" como email
            Cuando intento crear al inquilino
            Entonces se crea correctamente