import sqlite3

print("Starting")
con = sqlite3.connect("/var/lib/portfolio/portfolio.db")
c = con.cursor()
print("Connected to db")
try:
    c.execute(
        """
    CREATE TABLE IF NOT EXISTS Blogs
    (
        Blog_id INTEGER PRIMARY KEY AUTOINCREMENT,
        Title TEXT UNIQUE NOT NULL,
        Date TEXT NOT NULL
    )
    """
    )
    print("Blog table successfully created")
except Exception as e:
    print(e)
    con.close()
    exit()
print("Created Blogs table")
try:
    c.execute(
        """
    CREATE TABLE IF NOT EXISTS Projects
    (
        Project_id INTEGER PRIMARY KEY AUTOINCREMENT,
        Title TEXT UNIQUE NOT NULL,
        Desc TEXT UNIQUE NOT NULL,
        Tech TEXT NOT NULL,
        GitLink TEXT DEFAULT NULL,
        WebLink TEXT DEFAULT NULL,
        Blog_id INTEGER DEFAULT NULL,
        FOREIGN KEY (Blog_id) references Blogs (Blog_id)
    )
    """
    )
except Exception as e:
    print(e)
    con.close()
    exit()

con.commit()
con.close()
print("Created projects table")
