	query := datastore.NewQuery("Task").
		Filter("Category =", "Personal").
		Filter("Done =", false).
		Filter("Priority =", 4)  
