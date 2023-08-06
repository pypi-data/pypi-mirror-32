import re
from collections import defaultdict
from ifcschemareader import IfcSchema

IFCLINE_RE = re.compile("#(\d+)[ ]?=[ ]?(.*?)\((.*)\);[\\r]?$")

class IfcFile:
    """
    Parses an ifc file given by filename, entities can be retrieved by name and id
    The whole file is stored in a dictionary (in memory)
    """
    
    entsById = {}
    entsByName = {}

    def __init__(self, filename, schema):
        self.filename = filename
        self.schema = schema
        self.file = open(self.filename)
        self.entById, self.entsByName = self.read()
        print "Parsed from file %s: %s entities" % (self.filename, len(self.entById))
    
    def getEntityById(self, id):
        return self.entById.get(id, None)
    
    def getEntitiesByName(self, name):
        return self.entsByName.get(name, None)

    def read(self):
        """
        Returns 2 dictionaries, entById and entsByName
        """
        entById = {}
        entsByName = defaultdict(set)
        for line in self.file:
            e = self.parseLine(line)
            if e:
                entById[int(e["id"])] = e
                ids = e.get(e["name"],[])
                ids.append(e["id"])
                entsByName[e["name"]] = entsByName[e["name"]].union(set(ids))
                
        return [entById, entsByName]

    def parseLine(self, line):
        """
        Parse a line 
        """ 
        m = IFCLINE_RE.search(line)  # id,name,attrs
        if m:
            id, name, attrs = m.groups()
        else:
            return False
        
        return {"id": id, "name": name, "attributes": self.parseAttributes(name, attrs)}

    def parseAttributes(self, ent_name, attrs_str):
        """
        Parse the attributes of a line
        """
        parts = []
        lastpos = 0
        
        while lastpos < len(attrs_str):
            newpos = self.nextString(attrs_str, lastpos)
            parts.extend(self.parseAttribute(attrs_str[lastpos:newpos-1]))
            lastpos = newpos
        
        schema_attributes = self.schema.getAttributes(ent_name)

        assert len(schema_attributes) == len(parts), \
            "Expected %s attributes, got %s (entity: %s" % \
            (len(schema_attributes), len(parts), ent_name)
        
        attribute_names = [a[0] for a in schema_attributes]
        
        return dict(zip(attribute_names, parts))

    def parseAttribute(self, attr_str):
        """
        Map a single attribute to a python type (recursively)
        """
        parts = []
        lastpos = 0
        while lastpos < len(attr_str):
            newpos = self.nextString(attr_str, lastpos)
            s = attr_str[lastpos:newpos-1]
            if (s[0] == "(" and s[-1] == ")"): # list, recurse
                parts.append(self.parseAttribute(s[1:-1]))
            else:
                try:
                    parts.append(float(s)) # number, any kind
                except ValueError:
                    if s[0] == "'" and s[-1] == "'": # string
                        parts.append(s[1:-1])
                    elif s == "$":
                        parts.append(None)
                    else:
                        parts.append(s) # ref, enum or other

            lastpos = newpos
        
        return parts


    def nextString(self, s, start):
        """
        Parse the data part of a line
        """
        parens = 0
        quotes = 0

        for pos in range(start,len(s)):
            c = s[pos]
            if c == "," and parens == 0 and quotes == 0:
                return pos+1
            elif c == "(" and quotes == 0:
                parens += 1
            elif c == ")" and quotes == 0:
                parens -= 1
            elif c == "\'" and quotes == 0:
                quotes = 1
            elif c =="\'" and quotes == 1:
                quotes = 0
            
        return len(s)+1                  



if __name__ == "__main__":
    import time

    t1 = time.time()
    schema = IfcSchema("IFC2X3_TC1.exp")
    t2 = time.time()
    print "Loading schema took: %s s \n" % ((t2-t1))

    t1 = time.time()
    ifcfile = IfcFile("testdata/AC11-FZK-Haus-IFC.ifc", schema)
    t2 = time.time()
    print "Loading file took: %s s \n" % ((t2-t1))
    
    print ifcfile.getEntityById(233), "\n"
    print ifcfile.getEntityById(216), "\n"
