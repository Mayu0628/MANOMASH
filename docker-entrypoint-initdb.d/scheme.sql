CREATE TABLE IF NOT EXISTS users (
  user_id       INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  user_name     VARCHAR(40)     NOT NULL,
  email         VARCHAR(254)     NOT NULL,
  password      VARCHAR(20)     NOT NULL,
  introduce		VARCHAR(100)	NOT NULL,
  created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE IF NOT EXISTS oshis (
  oshi_id		INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  oshi_name		VARCHAR(100)	NOT NULL,
  tag			VARCHAR(100)	NOT NULL,
  birthday		DATE 			NOT NULL DEFAULT,
  oshi_meet		VARCHAR(1024)	NOT NULL,
  like_point	VARCHAR(254)	NOT NULL,
  free_space	VARCHAR(1024)	NOT NULL,
  interest		VARCHAR(1024)	NOT NULL,
  reaction_num	INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  created_at	DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at	DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE oshi_tag (
  oshi_id		INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  tag			VARCHAR(100) NOT NULL,
  PRIMARY KEY (oshi_id,tag)
);

CREATE TABLE oshi_like (
  oshi_id		INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  like_point	VARCHAR(254) NOT NULL,
  PRIMARY KEY (oshi_id,like_point)
)


CREATE TABLE IF NOT EXISTS comments (
  oshi_id    	INTEGER UNSIGNED NOT NULL,
  nice_num		INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  comment       VARCHAR(255)     NOT NULL DEFAULT '',
  user_id       INTEGER UNSIGNED NOT NULL,
  created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
);
