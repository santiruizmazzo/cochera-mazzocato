# language: es
Característica: Registrar pago de inquilino
    Como administrador
    quiero registrar el pago de un inquilino
    para dejar asentado el pago mensual de su cochera

    Regla de negocio: Monto es entero mayor a 1
        Escenario: Monto vacío
            Dado que no ingresé ningún monto
            Cuando intento registrar el pago
            Entonces se informa que el monto es obligatorio
        
        Escenario: Monto negativo
            Dado que ingresé un monto igual a -1
            Cuando intento registrar el pago
            Entonces se informa que el monto debe ser un entero mayor a 1

        Escenario: Monto igual a 0
            Dado que ingresé un monto igual a 0
            Cuando intento registrar el pago
            Entonces se informa que el monto debe ser un entero mayor a 1

        Escenario: Monto no numérico
            Dado que ingresé un monto igual a "hola"
            Cuando intento registrar el pago
            Entonces se informa que el monto debe ser un entero mayor a 1
    
    Regla de negocio: Inquilino debe existir
        Escenario: ID de inquilino vacío
            Dado que no seleccioné el ID de ningún inquilino
            Cuando intento registrar el pago
            Entonces se informa que el ID del inquilino es obligatorio

        Escenario: ID de inquilino inexistente
            Dado que no existe ningún inquilino
            Y que seleccioné al inquilino con id 99
            Cuando intento registrar el pago
            Entonces se informa que el inquilino no existe

        Escenario: ID de inquilino existente
            Dado que existen inquilinos
            Y que seleccioné al inquilino con id 1
            Cuando intento registrar el pago
            Entonces se registra correctamente

    Regla de negocio: Método de pago es válido
        Escenario: Método de pago vacío
            Dado que no ingresé ningún método de pago
            Cuando intento registrar el pago
            Entonces se informa que el método de pago es obligatorio

        Escenario: Método de pago inexistente
            Dado que seleccioné el método de pago "cheque"
            Cuando intento registrar el pago
            Entonces se informa que el método de pago debe ser un texto dentro de las opciones válidas

        Escenario: Método de pago numérico
            Dado que seleccioné el método de pago 10
            Cuando intento registrar el pago
            Entonces se informa que el método de pago debe ser un texto dentro de las opciones válidas

        Escenario: Método de pago válido
            Dado que seleccioné el método de pago "efectivo" o "transferencia"
            Cuando intento registrar el pago
            Entonces se registra correctamente

    Regla de negocio: Período abonado es válido
        Escenario: Período abonado anterior a registro de inquilino
            Dado que seleccioné al inquilino con id 1 que fue registrado el 1/7/2025
            y que seleccioné el período abonado que incluye al 30/6/2025
            Cuando intento registrar el pago
            Entonces se informa que no se puede abonar un período anterior al de registro del inquilino