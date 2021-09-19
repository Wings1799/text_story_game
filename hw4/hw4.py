class Frog:


    def __init__(self, name, color, species):
        self.name = name
        self.color = color
        self.species = species
        self.age = 0
        self.partner = Frog

    def __str__(self):
        return (self.color + ' ' + self.species + 'frog age ' + str(self.age))

    def bday(self):
        self.age += 1
        return self.age

    def add_partner(self, frog):
        self.partner = frog
        return self.partner

james = Frog('James', 'blue', 'tree')
ted = Frog('Ted', 'green', 'tree')
print(james.__str__())
james.bday()
james.bday()
print(james)
james.add_partner(ted)
print(james.partner)