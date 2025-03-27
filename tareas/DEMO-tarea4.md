
# Tarea 4: Convertir el frontend de la Plataforma en una SPA  con AngularJS

Convertir el frontend de la plataforma e-Commerce en una SPA (Single Page Application) utilizando AngularJS. Esta aplicación permitirá a los usuarios gestionar una lista de productos, incluyendo operaciones como agregar, editar, eliminar y buscar productos. El proyecto deberá demostrar competencias en la arquitectura de AngularJS, manejo de datos con AJAX, navegación con rutas, validación de formularios y animaciones.

## Requerimientos:

- Implementar una interfaz de usuario para gestionar productos (nombre, categoría, valor, etc.).
- Utilizar componentes de AngularJS para estructurar la aplicación (4.1).
- Crear servicios que utilicen AJAX para recuperar y almacenar datos (4.2), a través de una API RESTful. De momento Use JSON Server para crear una API propia o consuma la API del Sitio JSONPlaceholder
- Aplicar el sistema de rutas de AngularJS (4.3) para cambiar entre la visualización de la lista de contactos y el formulario de edición de contactos.
- Desarrollar formularios con validación (4.4) para la entrada de datos de contacto.
- Incorporar animaciones (4.5) para mejorar la interacción del usuario con la aplicación.
- La aplicación debe ser responsiva y funcional en diferentes dispositivos.
- Asegurar una experiencia de usuario fluida y dinámica.

## Notas

En este tarea decidi volver el repositorio a un mono-repo. Por lo tanto ahora en este repositorio viven varias aplicaciones.

1. [Ecommerce Site](/frontends/ecommerce/): Esta app tendra acceso limitado a la API y sera utilizado para los clientes.
2. [Admin Site](/frontends/admin/): Esta tendra full access a la API y podra realizar todas las operaciones CRUDs

El objetivo de esto es separar lo que tiene que ver con manejo del ecommerce y su content, de lo que es el ecommerce site. Ademas de brindar otro layer de seguridad.

Tambien decidi comenzar con la aplicacion de backend que sera escrita en Golang.

Tambien para la UI utilize un framework de React llamado [Refine](https://refine.dev/). Este es basicamente un conjunto de diferentes librerias, React Query, Material UI, React Router, etc. que estan unidas todas en el framework para realizar aplicaciones internal como Admin Panels, Dashboard, etc de una manera muy sencilla haciendo que la experiencia de desarrollo sea mas rapida y menos tediosa.

## Demo video

https://github.com/user-attachments/assets/ec1a33ae-fe86-4ee0-8ab8-ed93ab2dc5fe
