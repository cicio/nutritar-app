import re
import os
from pathlib import Path
s = Path('src/anditable.html')
with open(s, "r") as cache:
    data = cache.read()
cache.close()

#print data
nutritable={}

def get_ingredient(s):
    start_ingred_link = s.find("<p>")
    if start_ingred_link == -1:
            return None, 0    
    start_ingredient = s.find(">", start_ingred_link)
    end_ingredient = s.find("<", start_ingredient +1)
    ingredient_name = s[start_ingredient +1: end_ingredient]
    start_ndi_link = s.find("<p>", end_ingredient + 4)
    if start_ndi_link == -1:
        return None, 0
    start_ndi = s.find(">", start_ndi_link)
    end_ndi = s.find("<", start_ndi + 1)
    ndi = int(s[start_ndi + 1: end_ndi])
#    return ingredient_name, ndi
    nutritable[ingredient_name] = ndi
    get_ingredient(s[end_ndi+4:])
    return nutritable    
        
#print get_ingredient(data)    

def max (l):
    maximum =0
    for e in l:
        if e[1] >= maximum:
            maximum = e[1]
    return maximum

def maxlist(table):
    max_lista=[]
    for ingredient in table:
        if ingredient[1] == max(table):
            max_lista.append(ingredient)
    return max_lista
            
nutritbl = get_ingredient(data)

#max_nutri = [ing for ing in nutritable if ing[1] == max(nutritbl)] 
#nutritbl.sort()
print(nutritbl)
import json
out = Path('src/foodandi.json')
with open(out, "w") as f:
    json.dump(nutritbl, f)
f.close()

            