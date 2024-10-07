## Banco

<p>Este projeto em Go exemplifica conceitos básicos de orientação a objetos ao implementar um sistema bancário simples com contas correntes e poupança. O código utiliza interfaces para criar uma função genérica que realiza o pagamento de boletos, 
  permitindo que diferentes tipos de conta (corrente ou poupança) sejam tratadas de maneira uniforme.</p>

<p>No exemplo, a interface verificarConta define o comportamento esperado para contas que implementem a função Sacar. 
A função PagarBoleto recebe qualquer objeto que atenda a essa interface, garantindo a abstração do tipo de conta. As operações de depósito e saque são aplicadas a contas de poupança, e o saldo resultante é exibido.</p>

<p>Este projeto é um bom exemplo de como a orientação a objetos em Go pode ser usada para criar código flexível e reutilizável.</p>

Para executar o projeto em Go, siga os passos abaixo:

### Execute o programa: No terminal, rode o comando abaixo:
`go run main.go`
