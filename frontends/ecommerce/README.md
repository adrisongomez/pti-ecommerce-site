# PROGRAMMMING THE INTERNET

## Decripcion de la tarea

Convertir el frontend de la plataforma e-Commerce en una SPA (Single Page Application) utilizando AngularJS. Esta aplicación permitirá a los usuarios gestionar una lista de productos, incluyendo operaciones como agregar, editar, eliminar y buscar productos. El proyecto deberá demostrar competencias en la arquitectura de AngularJS, manejo de datos con AJAX, navegación con rutas, validación de formularios y animaciones.

### Requerimientos

- Implementar una interfaz de usuario para gestionar productos (nombre, categoría, valor, etc.).
- Utilizar componentes de AngularJS para estructurar la aplicación (4.1).
- Crear servicios que utilicen AJAX para recuperar y almacenar datos (4.2), a través de una API RESTful. De momento Use JSON Server para crear una API propia o consuma la API del Sitio JSONPlaceholder
- Aplicar el sistema de rutas de AngularJS (4.3) para cambiar entre la visualización de la lista de contactos y el formulario de edición de contactos.
- Desarrollar formularios con validación (4.4) para la entrada de datos de contacto.
- Incorporar animaciones (4.5) para mejorar la interacción del usuario con la aplicación.
- La aplicación debe ser responsiva y funcional en diferentes dispositivos.
- Asegurar una experiencia de usuario fluida y dinámica.

## Detalles del proyecto

### Herramientas requeridas

- `node` en su version `v22.14.0` o [`nvm`](https://github.com/nvm-sh/nvm)
- [`pnpm`](https://pnpm.io/installation)

### Como correr el proyecto

```BASH
nvm use # optional para que nvm utilize la version correcta de node
pnpm install # para instalar la dependencias
pnpm dev # para correr el proyecto con un servidor de desarrollo
pnpm build # para buildear la version de produccion
```

### Notas del proyecto

El proyecto se realizo con la libreria de `react` en su version `v19` con `vite`, `typescript` y `SWC`.

- `vite` es una herramienta de `bundling` muy utilizada en los ultimos años y se ha convertido en la herramienta de facto para la elaboracion de paginas web SPA sencillas.
- `typescript` es un superset de javascript que añade muchos features al lenguaje como son los tipo de datos, decoradores y otro tipo de sintaxis. Este al compilar se convierte a javascript
- `SWC` es una herramienta desarrollado por el equipo de Next.js que permite compilar codigo en typescript, tsx o jsx de una forma mucha raspida.

Otras herramienta de desarrollo utilizadas son `husky`, `prettier` y `eslint`.

- `prettier` es una herramienta que permite definir reglas de formato del codigo para asegurar un codigo homogeneo.
- `eslint` es una herramienta que permite definir reglas y estilo de codigo para asegurar la consistencia en el codigo.
- `husky` es una herramienta que permite ejecutar codigo arbitrario basado en eventos de git. En este caso, se utilizo en el event de `pre-commit` para asegurar que el codigo enviado a produccion, satisface las reglas y el formato establecidas por `eslint` y `prettier`.

Como CSS framework se utilizo `tailwindcss`.

Para routing en la aplicacion se utilizo [Tankstack Router](https://tanstack.com/router/latest), una alternativa a [React-Router](https://reactrouter.com/), este provee muchas ventajas como code split automatico lo que hace que el `page-load` de la pagina sea mas eficiente, File-base routes lo que permite organizar mejor el proyecto.

Para validacion se utilizo [Yup](https://www.npmjs.com/package/yup) que es una libreria que permite crear validationSchemas para validar los objectos que se van a construir en nuestros formulario.

Para form state management se utilizo [Formik](https://formik.org/) provee una API simple y intuitiva a la hora de manejar form state, ademas que presenta soporte para los validationSchema creados com **Yup**.

Tambien se utilizo como UI Sanbox and Testing Tool. [Storybook](https://storybook.js.org/). Una herramient que permite ir desarrollando los components de nuestra App de manera individual, ademas proveer un playground donde probar nuestros components.

### TODOs

- [ ] Implement [TankStack Query](https://tanstack.com/query/latest)
- [ ] Implement [OpenAPI Client generation]()
- [x] Implement [Mock Service Worker](https://mswjs.io/)
