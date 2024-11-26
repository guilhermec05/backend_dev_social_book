CREATE DATABASE  IF NOT EXISTS  devbook;
user devbook;

drop table if EXISTS users;

CREATE TABLE users(
    id int NOT NULL AUTO_INCREMENT,
    name  varchar(50) NOT NULL,
    nick varchar(50) NOT NULL unique,
    email  varchar(50) NOT NULL unique,
    password varchar(50) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    primary key (id)
 ) ENGINE=INNODB


 
CREATE TABLE follow( 
    user_id int NOT NULL foreign key(user_id) references  users(id) on delete cascade,
    follow_id int NOT NULL foreign key(follow_id) references  users(id) on delete cascade,
    primary key (user_id,follow_id)
 ) ENGINE=INNODB


CREATE TABLE publish(
    id int NOT NULL AUTO_INCREMENT,
    title  varchar(30) NOT NULL,
    content varchar(300) NOT NULL,
    author_id  int NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    primary key (id),
    foreign key(author_id) references  users(id) on delete cascade
 ) ENGINE=INNODB;
 

  CREATE TABLE like_publish(
    publish_id  int not null,
    user_id int not null
   	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    primary key (publish_id,user_id),
    foreign key(user_id) references  users(id) on delete cascade,
    foreign key(publish_id) references  publish(id) on delete cascade
 ) ENGINE=INNODB;
 