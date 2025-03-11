# PROGRAMMMING THE INTERNET

## Decripcion de la tarea

Crear un validador de formularios dinámico utilizando JavaScript ES6+ o TypeScript (puede usar frameworks moderno como React, Vue.js o Angular para mejorar la aplicación). Esta tarea tiene como fin aplicar los conocimientos avanzados en JavaScript, TypeScript y el uso de un framework para construir una aplicación web que realice validaciones de formularios de manera eficiente y dinámica.

### Requerimientos:

- Implementar un formulario web con varios campos como nombre, correo electrónico, contraseña, fecha nacimiento, cantidad numérica, y otros campos relevantes.

- Utilizar JavaScript ES6+ para agregar funcionalidades dinámicas al formulario.

- Integrar TypeScript para mejorar la estructura y tipado del código.

- Elegir y utilizar `React`, `Vue.js` o `Angular` para estructurar el proyecto y gestionar el estado del formulario.

### Funcionalidades de Validación:

- Validar los datos ingresados en tiempo real, proporcionando retroalimentación instantánea al usuario, por ejemplo: verificación de formato de correo electrónico, fortaleza de la contraseña, contraseñas iguales, fecha de nacimiento válida.

- Asegurar que todas las validaciones sean coherentes y proporcionen mensajes claros de error o confirmación.

- Validar que todos los datos estén ingresados y de manera correcta, considerar los campos obligatorios y opcionales

- Mostrar mensajes de alertas con la descripción de los mensajes

Crear un `README` detallado en el repositorio de GitHub, explicando el propósito del proyecto, las tecnologías usadas y cómo se implementaron las validaciones.

Incluir capturas de pantalla que muestren las diferentes etapas de validación del formulario.

Documentar los desafíos y soluciones encontrados durante el desarrollo.

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

## Demo video de los forms para Login y Sign Up

https://github.com/user-attachments/assets/ad119461-e6d6-44fc-98f9-733600ce36c6


