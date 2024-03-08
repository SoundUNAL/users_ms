-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema soundunal_users_db
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `soundunal_users_db` ;

-- -----------------------------------------------------
-- Schema soundunal_users_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `soundunal_users_db` DEFAULT CHARACTER SET utf8 ;
USE `soundunal_users_db` ;

-- -----------------------------------------------------
-- Table `soundunal_users_db`.`Role`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `soundunal_users_db`.`Role` ;

CREATE TABLE IF NOT EXISTS `soundunal_users_db`.`Role` (
  `idRol` INT AUTO_INCREMENT NOT NULL,
  `rol` VARCHAR(45) NOT NULL,
  `description` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`idRol`))
ENGINE = InnoDB
COMMENT = 'Rol';


-- -----------------------------------------------------
-- Table `soundunal_users_db`.`User`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `soundunal_users_db`.`User` ;

CREATE TABLE IF NOT EXISTS `soundunal_users_db`.`User` (
  `idUser` INT AUTO_INCREMENT NOT NULL, 
  `name` VARCHAR(45) NOT NULL,
  `lastname` VARCHAR(45) NOT NULL,
  `username` VARCHAR(45) NOT NULL,
  `password` VARCHAR(100) NOT NULL,
  `email` VARCHAR(45) NOT NULL,
  `phone` INT NOT NULL,
  `birthday` DATE NOT NULL,
  `lastconnection` DATETIME(6) NOT NULL,
  `idRol` INT NOT NULL,
  PRIMARY KEY (`idUser`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE,
  UNIQUE INDEX `username_UNIQUE` (`username` ASC) VISIBLE,
  INDEX `fk_User_Role_idx` (`idRol` ASC) VISIBLE,
  CONSTRAINT `fk_User_Role`
    FOREIGN KEY (`idRol`)
    REFERENCES `soundunal_users_db`.`Role` (`idRol`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `soundunal_users_db`.`Following`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `soundunal_users_db`.`Following` ;

CREATE TABLE IF NOT EXISTS `soundunal_users_db`.`Following` (
  `idUserFollower` INT NOT NULL,
  `idUserFollowed` INT NOT NULL,
  PRIMARY KEY (`idUserFollower`, `idUserFollowed`),
  INDEX `fk_User_has_User_User2_idx` (`idUserFollowed` ASC) VISIBLE,
  INDEX `fk_User_has_User_User1_idx` (`idUserFollower` ASC) VISIBLE,
  CONSTRAINT `fk_User_has_User_User1`
    FOREIGN KEY (`idUserFollower`)
    REFERENCES `soundunal_users_db`.`User` (`idUser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_User_has_User_User2`
    FOREIGN KEY (`idUserFollowed`)
    REFERENCES `soundunal_users_db`.`User` (`idUser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
