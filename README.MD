<h1> for first up mysql container</h1>
docker-compose -f docker-compose.yml build
<h1> for next attempts</h1>
docker-compose -f docker-compose.yml up -d db



dbeaver string connection
jdbc:mysql://localhost:3306/golang?allowPublicKeyRetrieval=true&useSSL=false



init sql code 

CREATE TABLE users (
	id INT auto_increment not null primary key,
	name varchar(100) NULL,
	age INT NULL
)