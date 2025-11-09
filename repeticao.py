#laços de repetição = fazer uma contagem numeral automaticamente sem a necessidade de escrever cada número manualmente.
#tipos de laços de repetição: for e while

#For = Quando sabemos quantas vezes queremos repetir.
#While = Quando não sabemos quantas vezes queremos repetir. Depende de uma condição.

#range = um intervalo. O último número não é incluso (de 1 a 11 = 1 a 10).
# = por exemplo: print(1), print(2) ... print(10). Com o FOR, podemos fazer isso automaticamente.


#FOR com lista de "números"
# for numero in range(1, 6):
#     print(numero)


#FOR com lista de "strings"
# amigos = ("Luis", "Paloma", "Aline")
# for escolher in amigos:
#     print(f"Olá, {escolher}! Tudo bem?")
    
#FOR percorrendo uma string, uma palavra.
# palavra = "Developer"
# for letra in palavra:
#     print(letra)


#WHILE contagem regressiva
contador = 3
while contador > 0: #ENQUANTO o contador for maior que 0
    print(contador)
    contador -= 1 #Negativo por ser contagem regrassiva
    #se colocar contador += 1, o laço nunca vai acabar, porque o contador sempre vai ser maior que 0.
print("GO!🏎️")
