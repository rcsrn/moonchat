Modelado y Programación
=======================

Proyecto 1: Chat
-----------------


### DISEÑO:

Realizaremos un diseño orientado a objetos.
Comenzaremos con la parte del servidor. Necesitaremos los siguientes objetos: Server, ServerProcessor, MessageVerifier y MessageCreator y Room.

La clase Server representara a nuestro servidor; la misma se encargara de todo lo relacionado con el manejo de cuartos, además de estar todo el tiempo en espera de nuevas conexiones mientras nuestro programa se esté ejecutando. Una vez que alguna conexión sea aceptada por Server, esta será delegada a una instancia de la clase ServerProcessor, la cual se encargara de todo lo relacionado con la lectura y escritura de mensajes del cliente al servidor y del servidor hacia el cliente respectivamente. Cuando ServerProcessor reciba algún mensaje válido, dentro del protocolo de comunicación establecido, este realizara la petición recibida al servidor el cual, dependiendo de lo que se esté solicitando el cliente, responderá o bien con un error o con la información solicitada. Siempre que el servidor responda con un error significa que la petición recibida no puede ser realizada. Si el mensaje recibido no es válido entonces ServerProcessor responde con un mensaje de tipo error y se desconecta al cliente. Server servirá como el puente de comunicación entre ServerProcessor y Room con el fin de disminuir el acoplamiento entre clases. Esto implica que algunos métodos en Server serán solo para retornar los valores que regresen algunos métodos en la clase Room. Si el usuario no se ha identificado con el servidor entonces no podrá realizar ninguna acción más que la de tratar de identificarse, si trata de realizar alguna otra acción será desconectado.

ServerProcessor tendrá como propiedades la conexión del cliente establecida con el servidor, un MessageVerifier para verificar que los mensajes recibidos estén correctamente escritos, un MessageCreator para crear los diferentes tipos de mensajes, una referencia al servidor, el nombre de usuario -que por defecto al inicio será la cadena vacía-, el estatus del usuario -que por defecto al inicio será 'ACTIVE'-, una colección con los nombres de los cuartos a los que el usuario pertenece y además guardara el estado en una variable booleana para saber sí el usuario ya se ha identificado con el servidor. Dependiendo del tipo de mensaje recibido, que se traduce en la petición recibida de parte del cliente, ServerProcessor procederá de forma distinta, siempre respondiendo al cliente con algún mensaje, ya sea de error, advertencia o éxito.

Para la parte de mensajes tendremos un paquete Message, en la cual tendremos todos los tipos de mensajes que vamos a utilizar dado el protocolo. Cada mensaje tendrá un método con el cual obtener su formato JSON, que posteriormente será enviado a través de la conexión. Esta clase será utilizada tanto por el servidor como por el cliente.

La clase Room tendrá dos colecciones; una con los nombres de los usuarios dentro de la sala y otra de los usuarios invitados que potencialmente podrán unirse a la sala. Solo usuarios invitados podrán unirse a la sala y una vez que un usuario se una a la sala, será eliminado de los usuarios invitados y formará parte de los miembros de la sala. En la Room utilizaremos el patrón de diseño adaptador, adaptando la clase Set.
El comportamiento en Room se limita a solamente agregar y eliminar usuarios, verificar si un usuario en específico pertenece a la sala, regresar la lista de usuarios en la sala, remover usuarios de la lista de invitados y decirnos si la sala esta vacía o no.

Por parte del cliente tendremos una clase Client que se encargara de establecer la comunicación con el servidor a través de un enchufe y una clase ClientProcessor que será análoga a ServerProcessor, manejando todo lo relacionado con lectura y escritura de mensajes.

Al usuario se le solicitará la dirección IP y el puerto por el cual deseamos establecer una conexión con el servidor. Una vez establecida la conexión podrá comenzar a escribir mensajes, crear cuartos y todas las funcionalidades que ofrece el servidor. Los datos ingresados por el usuario serán procesados por ClientProcessor de forma que, dependiendo del caso, creara un Message el cual será enviado hacia el servidor.

Para la parte del cliente utilizaremos el patron de diseño modelo vista controlador (MVC).

Para la parte del modelo: Utilizaremos los siguientes objetos: Client, Client_Processor, MessageVerifier, MessageCreator y un archivo con las constantes que utilizaremos para lectura y escritura de mensajes además de utilizar, al igual que nuestro servidor, el paquete Message. Client se encargará de establecer la comunicación con el servidor a través de un enchufe, una vez establecida dicha conexión será delegada a un ClientProcessor el cual tendrá un ciclo de lectura para los mensajes enviados por el servidor. Estos mensajes serán recibidos por el controlador del modelo que a su vez mandara a imprimir los mensajes en consola
Para correr el programa al usuario deberá ingresar la dirección IP y el puerto por el cual deseamos establecer una conexión con el servidor. Una vez establecida la conexión podrá comenzar a escribir mensajes, crear cuartos y todas las funcionalidades que ofrece el servidor. Los datos ingresados por el usuarios serán procesados por ClientProcessor de forma que, dependiendo del caso, creara un Message el cual será enviado hacia el servidor.

Para la parte de la vista: Tendremos los siguientes objetos: Printer y ConsoleListener. Printer se encargará de todo lo relacionado a la impresión de mensajes por la consola, es decir todo lo que el usuario verá por la pantalla, tendrá como constantes las instrucciones para el uso del programa además de poder imprimir  errores de conexión o lectura. ConsoleListener solamente recibirá los mensajes que el usuario ingrese en una colección, la cual posteriormente será procesada por el Cliente a través de su ClientProcessor en el modelo.

Para la parte del controlador: Tendremos una clase ClientController (client_controller) la cual funcionara como intermediario entre la comunicación del modelo y la vista. Al correr el programa creara una instancia de Client y verificara que se haya ingresado la dirección IP y el puerto a donde se quiere establecer la conexión. Posteriormente creará un objeto Printer y un ConsoleListener para imprimir las instrucciones de uso del programa y leer los mensajes ingresados por la consola respectivamente. Lo primero que se le solicitara al usuario será un nombre de usuario con el cual identificarse con el servidor. Una vez ingresado el nombre de usuario se podrán a comenzar a mandar los diferentes tipos de mensajes.
 
### Compilación

Tanto para compilar el servidor y el cliente se tiene un archivo Makefile. Ejecutaremos los siguientes comandos

*Para el servidor: make buildServer
*Para el cliente: make buildClient

### Ejecución

*Para el servidor: make startServer
*Para el cliente ./client_controller

### Pruebas Unitarias
Entrar a test/ y ejecutar go test -v