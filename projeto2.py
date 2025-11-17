# Simulador de caixa eletrônico

# saldo inicial de R$500. Pode sacar até zerar.
# saque apenas de números inteiros.

# saldo é uma variável
# saldo = 500

# while saldo > 0: # foi definido que acaba chegar no  0. Loop.
#     # saque permitido
#     saque = int(input("Digite o valor do saque:"))

#     if saque <= saldo:
#         # o saque pode ser realizado
#         saldo -= saque # saldo = saldo - saque
#         # mostra o saldo atualizado
#         print("Saque realizado com sucesso.")
#         print(f"Saldo atual: R$ {saldo}") # f = formatação de string. .2f = duas casas decimais = R$0,00 

#     else: # limitado entre 500 e 0. Fora desse limite, não faz o saque.
#         print("Saldo insuficiente") # poderia detalhar mais para o usuário.

        # if = se
        # else = senão
        # elif = se + se não


# Usando em forma decimal (float) + menu

saldo = 500

while True: # loop infinito
    print("=== CAIXA ELETRÔNICO ===")
    print("1 - Ver saldo")
    print("2 - Fazer saque")
    print("3 - Sair")

    opcao = input("Escolha uma opção: ")
    
    if opcao == "1": # ver saldo
        print(f"Seu saldo é: R$ {saldo:.2f}")

    elif opcao == "2": # fazer saque
        saque = float(input("Digite o valor do saque:"))
        if saque <= saldo:
            saldo -= saque
            print("Saque realizado com sucesso.")
            print(f"Saldo atual: R$ {saldo:.2f}") # duas casas decimais
        else:
            print("Saldo insuficiente")
    elif opcao == "3": # sair
        print("Obrigado por usar o caixa eletrônico. Até logo!")
        break # sai do loop
    else: # opção inválida
        print("Opção inválida. Tente novamente.")

