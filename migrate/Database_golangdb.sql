-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema golangdb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema golangdb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `golangdb` DEFAULT CHARACTER SET utf8mb3 ;
USE `golangdb` ;

-- -----------------------------------------------------
-- Table `golangdb`.`product`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `golangdb`.`product` (
  `pro_id` INT NOT NULL AUTO_INCREMENT,
  `pro_name` VARCHAR(120) NOT NULL,
  `pro_code` VARCHAR(20) NOT NULL,
  `pro_price` FLOAT NOT NULL,
  `pro_create_at` VARCHAR(50) NULL,
  `pro_update_at` VARCHAR(50) NULL,
  PRIMARY KEY (`pro_id`))
ENGINE = InnoDB
AUTO_INCREMENT = 4
DEFAULT CHARACTER SET = utf8mb3;

-- -----------------------------------------------------
-- Table `golangdb`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `golangdb`.`user` (
  `user_id` INT NOT NULL AUTO_INCREMENT,
  `user_username` VARCHAR(45) NOT NULL,
  `user_password` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`user_id`))
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `golangdb`.`log`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `golangdb`.`log` (
  `log_id` INT NOT NULL AUTO_INCREMENT,
  `log_method` VARCHAR(15) NOT NULL,
  `log_description` VARCHAR(150) NOT NULL,
  `log_data` VARCHAR(50) NULL,
  PRIMARY KEY (`log_id`))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

DELIMITER $$
	CREATE TRIGGER Tgr_UpdateAt BEFORE UPDATE
	ON product
	FOR EACH ROW
	BEGIN
		SET NEW.pro_update_at = NOW();
	END$$
DELIMITER ;

DELIMITER $$
	CREATE TRIGGER Tgr_CreateAt BEFORE INSERT
	ON product
	FOR EACH ROW
	BEGIN
		SET NEW.pro_create_at = NOW();
		SET NEW.pro_update_at = NOW();
	END$$
DELIMITER ;

DELIMITER $$
	CREATE TRIGGER Tgr_DataAt BEFORE INSERT
	ON log
	FOR EACH ROW
	BEGIN
		SET NEW.log_data = NOW();
	END$$
DELIMITER ;

DELIMITER $$
	CREATE TRIGGER Tgr_Log BEFORE INSERT
	ON product
	FOR EACH ROW
	BEGIN
		INSERT INTO log (log_method, log_description) VALUES ("GET", "product inserted directly into the database");
	END$$
DELIMITER ;

INSERT INTO product (pro_id, pro_name, pro_code, pro_price) VALUES (1, "Notebook", "Acer Nitro", 4502.5);
INSERT INTO product (pro_id, pro_name, pro_code, pro_price) VALUES (2, "Ventilador", "Mondial", 200.0);
INSERT INTO product (pro_id, pro_name, pro_code, pro_price) VALUES (3, "Faca", "Tramontina", 55.99);