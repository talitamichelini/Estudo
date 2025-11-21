#if = se = confição é verdadeira
#elif = se + se não = a primeira é falsa, mas essa verdadeira
#else = senão = condição é falsa

#identar = dar espaço no começo da linha (tab) = indicar que aquele bloco de código pertence a uma estrutura condicional = que aquele código está dentro do outro.

#Você está em uma cafeteria e está com pouca grana. 
# O café custa R$5,00 e o chá R$3,00.

#Se tiver R$5,00 ou mais, você compra um café.
#Se tiver menos de R$5,00 mas R$3,00 ou mais, você compra um chá.
#Se tiver menos de R$3,00, você não compra nada.

dinheiro = float(input("Quanto tem na carteira?"))
#float = número com vírgula; input = entrada de dados
# (=) retorna um valor
if dinheiro >= 5:
    print("Pode comprar um café!")
elif dinheiro >= 3:
    print("Pode comprar um chá!")
else: 
    print("Não pode comprar nada!")
    # "else" não precisa de condição, por isso ele apresente ":" direto após a palavra.