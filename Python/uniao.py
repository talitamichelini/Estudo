# união de conceitos
# quando efetuamos um programa, unimos todos os conceitos 

#Exemplo: saber se a palavra tem a letra Y
# palavra = "Python"

# for letra in palavra:
#     if letra == "y" or "Y":
#         print("Essa palavra tem a letra y")

#segundo a tabela ASCII letra maiúscula e minúscula tem valores diferentes
#portanto é importante colocar as duas

# método diferente
palavra = "Paralelepípedo"
letraProcurada = "a"

# for letra in palavra:
#     if letra == letraProcurada:
#         print("Essa palavra possue", letraProcurada)

#para dizer quantas letras tem na palavra
contador = 0
for letra in palavra:
    if letra == letraProcurada:
        contador += 1
print("A palavra possue", contador, "letras", letraProcurada)
# nesse exemplo, unimos laço de repetição com condição
