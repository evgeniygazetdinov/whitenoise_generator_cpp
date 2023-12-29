<h1> for first up mysql container</h1>
docker run  --name some-mysql3 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql



dbeaver string connection
jdbc:mysql://localhost:3306/mysql?allowPublicKeyRetrieval=true&useSSL=false



init sql code 

CREATE TABLE users (
	id INT auto_increment not null primary key,
	name varchar(100) NULL,
	age INT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
