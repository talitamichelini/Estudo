# Funções = Bloco de código reutilizável
# ao invés de escrever o mesmo código várias vezes,
# podemos escrever uma função e chamar ela quando necessário

# def = palavra paraa definir uma função
# tudo que estiver indentado (com espaço) abaixo do def
# faz parte da função

# def <nome_da_função>():
#     <bloco_da_função>


# def saudacao(nome):
#     print(f"Olá, {nome}! Bem-vindo(a)!")
# saudacao("Talita")
# saudacao("João")
# saudacao("Maria")


# Retorno de valores
# def somar(a, b, c):
#     return a + b + c
# # Chamando a função e armazenando o valor retornado
# resultado = somar(2, 4, 6)
# print(f"A soma é {resultado}")


# Calcular média
def calcular_media(n1, n2, n3):
    media = (n1 + n2 + n3) / 3
    return media
#chamando a função
resultado = calcular_media(8, 6, 7)
print(f"A média é igual a {resultado}")

if resultado <=5:
    print("Reprovado")
else:
    print("Aprovado")
