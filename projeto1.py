# Jogo de Adivinhação

# No jogo, o usuário precisa adivinhar o número secreto
# Pode tentar várias vezes até acertar

# snake_case = underline (usado em Python)
# camelCase = letraMaiusculaNoMeio (usado em Go)

numero_secreto = 5
tentativa = 0

while tentativa != numero_secreto:
    # != diferente
    tentativa = int(input("Encontre o número secreto entre 0 e 10: "))

# para dar dicas ao usuário
    if tentativa < numero_secreto:
        print("Tente um número maior.")
    elif tentativa > numero_secreto:
        print("Tente um número menor.")
    else:
        print("Parabéns! Você acertou o número secreto.")
# Fim do jogo
