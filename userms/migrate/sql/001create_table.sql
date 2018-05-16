CREATE TABLE 'user' (
	`id` INT NOT NULL AUTO_INCREMENT,
	`login` VARCHAR(50) NOT NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
	`password` VARCHAR(32) NOT NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
	PRIMARY KEY (`id`),
	UNIQUE INDEX `login` (`login`),
	INDEX `login` (`login`)
)

CREATE TABLE 'refresh_token' (
	'user_id' INT NOT NULL,
	'token'  VARCHAR(512) NOT NULL,
	'ttl' INT NOT NULL, 
	FOREIGN KEY (user_id) REFERENCES user(id)
) 

COLLATE='utf8_unicode_ci'
ENGINE=InnoDB;

