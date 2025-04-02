# Tarea 7: Desarrollo de un Chatbot)
 
## Requerimientos:
 
- Configurar la API de OpenAI:
- Registrar y obtener claves API de OpenAI.
- Configurar el entorno de desarrollo necesario para utilizar la API de OpenAI.
 
### Desarrollar el Chatbot:
 
- Implementar un chatbot utilizando la API de OpenAI.
- Diseñar el flujo de conversación para que el chatbot pueda responder preguntas basadas en los productos de la tienda virtual.
 
### Integrar el Chatbot en la Aplicación AngularJS:
 
- Conectar el chatbot con la aplicación web creada en AngularJS.
- Diseñar y desarrollar la interfaz de usuario del chatbot en la aplicación.
- Asegurar que el chatbot responda adecuadamente a las entradas del usuario y muestre las respuestas en la interfaz.

## Notas

En este entregable se realizo la implementacion de un chatbot consumiendo la API de OpenAI para Prompts:
1. Se implementaron nuevos endpoint en la backend escrita con go en el cual se encarga de hacer un wrapper sobre el SDK de OpenAI para go y ademas de las gestion de los mensajes y informacion que tanto el usuario como OpenAI necesitan para la comunicacion. Cabe destacar que este implementacion se baso en realizorlo con request HTTP pero una recomendacion seria utilizar Websocket para tener una comunicacion mas efectiva ya que se obtiene un stream donde el client y openai pueden interactuar sin tanto overhead que HTTP provee.
2. Se implemento el UI en la applicacion de front-end donde el usuario authenticado puede crear una session con el chat y realizar las consultas con el asistente AI.
3. Plus. Se crear nuevas tablas para poder almacenar los mensajes de los usuario y open ai para si en futuro se quiere reconstruir la conversacion completa se pueda realizar.

## Demo video

https://github.com/user-attachments/assets/7fe3e80a-6cf0-4fa8-8978-6f3d6833cf1b

