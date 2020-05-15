# MemGraph
A memory graph data structure


How to use it:
g := NewMemGraph()
	type NodeType struct {
		name string
	}
	type NodeData struct {
		name string
		age  int
	}
	type relationship struct {
		name string
	}

	g.AddType(1, NodeType{"person"})
	g.AddType(2, NodeType{"laptop"})
	g.AddType(3, NodeType{"book"})

	g.Add(1, "Tom", NodeData{"Tom", 60})
	g.Add(1, "Mary", NodeData{"Mary", 50})
	g.Add(2, "Surface", NodeData{"Surface", 1})
	g.Add(2, "MacBook", NodeData{"MacBook", 2})
	g.Add(3, "Becoming", NodeData{"Becoming", 1})
	g.Add(3, "HungerGames", NodeData{"Hunger Games", 2})

	g.Connect(1, 2, "Tom", "Surface", relationship{"pc"})
	g.Connect(1, 3, "Tom", "HungerGames", relationship{"book"})

	g.Connect(1, 2, "Mary", "Surface", relationship{"pc"})
	g.Connect(1, 2, "Mary", "MacBook", relationship{"pc"})
	g.Connect(1, 3, "Mary", "HungerGames", relationship{"book"})
	g.Connect(1, 3, "Mary", "Becoming", relationship{"book"})
