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
  `use_password` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`user_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


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

INSERT INTO product (pro_name, pro_code, pro_price) VALUES ("Notebook", "Acer Nitro", 4502.5);
INSERT INTO product (pro_name, pro_code, pro_price) VALUES ("Ventilador", "Mondial", 200.0);
INSERT INTO product (pro_name, pro_code, pro_price) VALUES ("Faca", "Tramontina", 55.99);