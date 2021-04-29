e1 = [-5, 2, 2, 2]
e2 = [-2, 4, 1, -1]
e3 = [1, -2, 6, 1]
erro = 10 ** -1

# e1 = [3, -2, 3, 5]
# e2 = [2, -1, 4, -3]
# e3 = [2, 2, 5, 4]
dr = []
x = [-1.0, 0.0, -0.5]
erro_relativo = 1
iteracao = 1

while erro < erro_relativo do
  x0 = [x[0], x[1], x[2]]
  x[0] = (1.0/e1[0]) * (e1[3]  - e1[1]*x0[1] - e1[2]*x0[2])
  x[1] = (1.0/e2[1]) * (e2[3]  - e2[0]*x0[0] - e2[2]*x0[2])
  x[2] = (1.0/e3[2]) * (e3[3]  - e3[0]*x0[0] - e3[1]*x0[1])

  dr[0] = x[0] - x0[0]
  dr[1] = x[1] - x0[1]
  dr[2] = x[2] - x0[2]

  temp1 = 0
  temp2 = 0
  for i in 0..2
    if temp1 < dr[i].abs
      temp1 = dr[i]
      # print "#{dr[i]}, "
    end
    if temp2 < x[i].abs
      temp2 = x[i].abs
    end
  end
  erro_relativo = temp1.abs / temp2

  puts "x#{iteracao}: [#{x[0]}, #{x[1]}, #{x[2]}]"
  puts "Erro relativo é #{erro_relativo}\n\n"
  iteracao += 1
end