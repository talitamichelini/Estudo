# listas e tuplas
# listas  modificavel, tuplas não modificavel
# lista é representada por colchetes []
# tuplas é representada por parenteses ()

#dado a partir de posição 0
# posição:  0      1      2 
# ou posição: -3     -2     -1
#frutas = ["pera", "uva", "maçã", "banana"]
# numeros = [1, 2, 3, 4, 5]
# misturada = [1, "banana", 3.5, True] # int, string, float, booleano


# print(frutas[0])  # pera
# print(numeros[2])  # 3
# print(misturada[3])  # True
# print(frutas[-1])  # maçã  (posição negativa começa do final)


#inserindo elementos em uma lista

# append(): adiciona um item no final
#frutas.append("abacaxi")
#print(frutas)

# insert(): adiciona um item na posição desejada 
#frutas.insert(2, "amora")  # insere abacaxi na posição 1. Lembrando que a contagem começa do 0.
#print(frutas)

#removendo elementos de uma lista

# remove(): remove o item desejado
#frutas.remove("uva")
#print(frutas)

# pop(): remove o item da posição desejada
#frutas.pop(1)  # remove o item da posição 1
#print(frutas)

#SEMPRE O ÚLTIMO PRINT QUE VALE E SERÁ CONSIDERADO O CÓDIGO FINAL.


#tuplas
#lista de coisas fixas, que não podem ou eu não queira que sejam alteradas
#Exemplo dias da semana, meses do ano

# dias_semana = ("domingo", "segunda", "terça", "quarta", "quinta", "sexta", "sábado")
# #print(dias_semana)

# #print(dias_semana[0]) # domingo
# #print(dias_semana[-1]) # sábado

# #dias_semana[0] = "Domingo"  # Isso dará um erro, pois tuplas não podem ser modificadas
# #print(dias_semana)   #COMPROVADO
# #Para fazer essa alteração, seria necessário converter a tupla em lista,
# # fazer a alteração e depois converter de volta para tupla.

# # convertendo tupla para lista
# #exemplo
# dias_semana = list(dias_semana) 
# dias_semana[0] = "segunda"
# dias_semana[1] = "terça"
# dias_semana[2] = "quarta"
# dias_semana[3] = "quinta"
# dias_semana[4] = "sexta"
# dias_semana[5] = "sábado"
# dias_semana[6] = "domingo"
# # convertendo lista de volta para tupla
# dias_semana = tuple(dias_semana)
# print(dias_semana)
#com a conversão para lista, append, insert etc, podem ser usados normalmente.

#OUTRO EXEMPLO DE TUPLA
meses = ("J", "F", "M", "A", "M", "J", "A", "S", "O", "N")
print(meses)
#Está faltando Julho e Dezembro
meses = list(meses)
meses.insert(6, "J")  # Julho
meses.append("D")  # Dezembro (ou meses.insert(11, "D"))
meses = tuple(meses)
print(meses)
