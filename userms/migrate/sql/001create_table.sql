CREATE TABLE `user` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`login` VARCHAR(50) NOT NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
	`password` VARCHAR(32) NOT NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
	PRIMARY KEY (`id`),
	UNIQUE INDEX `login` (`login`),
	INDEX `login` (`login`)
)
COLLATE='utf8_unicode_ci'
ENGINE=InnoDB;