# Tarea 6: Integración y Seguridad

- Integrar el backend desarrollado con el frontend
- Integrar un sistema de autenticación para los usuarios, permitiendo a los usuarios registrarse, iniciar sesión y gestionar sus entradas (6.4).
- La aplicación debe permitir a los usuarios crear, editar, eliminar y listar los pedidos.
- Implementar medidas de seguridad para proteger la información y las cuentas de los usuarios.

## Notes

En este entregable el objetivo es implementar seguiridad en la API y la gestion de credenciales de los usuarios:

- Para la Autentication de los usuarios se utilizo el metodo de basic con user and password. Este permite llamar al endpoint `/auth/login` que genera un par the JWTTokens, de los cuales se utilizan para poder acceder a las demas API. Los tokens contienen un token de accesso que da acceso necesario a los diferentes recursos de la API dependiendo si el usuario es Admin o customer, este tiene una duracion  de 24hr, y el segundo token es un token de refresco el cual solo tiene accesso al endpoint `/auht/refresh` que permite generar tokens.
- Para manejo de contraseña esta son guardadas en la base de datos depues de pasar por un hash function basado en el algorimot `hs256` y ademas es salted tambien para evitar rainbow table attacks.
- Ademas se agregaron varios endpoint tanto registro de usuarios como manejo de orders y direcciones.

## Demos

https://github.com/user-attachments/assets/2bfb81ba-74d7-4851-b6bf-8aac473decc6

https://github.com/user-attachments/assets/5e4a5af9-a6d1-406b-a3ab-a3c8cb68b817

https://github.com/user-attachments/assets/cb771619-fae5-4adc-ace1-d5acc8bf6069

https://github.com/user-attachments/assets/652349dd-7ebe-4ab5-8e88-083eb47faec9
