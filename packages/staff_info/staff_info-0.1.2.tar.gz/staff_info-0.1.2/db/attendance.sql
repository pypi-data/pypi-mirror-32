-- MySQL dump 10.13  Distrib 5.7.22, for Linux (x86_64)
--
-- Host: localhost    Database: attendance
-- ------------------------------------------------------
-- Server version	5.7.22-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `commission`
--

DROP TABLE IF EXISTS `commission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `commission` (
  `idEmp` int(11) NOT NULL,
  `comm_percent` decimal(2,1) DEFAULT NULL,
  `date_of_change` date DEFAULT NULL,
  `year` year(4) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `com_emp_idx` (`idEmp`),
  KEY `com_result_idx` (`year`),
  CONSTRAINT `com_emp` FOREIGN KEY (`idEmp`) REFERENCES `employee` (`idEmp`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `com_result` FOREIGN KEY (`year`) REFERENCES `result` (`year`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `employee`
--

DROP TABLE IF EXISTS `employee`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `employee` (
  `idEmp` int(11) NOT NULL AUTO_INCREMENT,
  `fullName` varchar(45) DEFAULT NULL,
  `dateOfEmp` date DEFAULT NULL,
  `dob` date DEFAULT NULL,
  PRIMARY KEY (`idEmp`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `informal_vacation`
--

DROP TABLE IF EXISTS `informal_vacation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `informal_vacation` (
  `idEmp` int(11) NOT NULL,
  `date_of_change` date NOT NULL,
  `description` varchar(40) DEFAULT NULL,
  `informal_vacation_available` int(11) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `infor_emp_idx` (`idEmp`),
  CONSTRAINT `infor_emp` FOREIGN KEY (`idEmp`) REFERENCES `employee` (`idEmp`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `position`
--

DROP TABLE IF EXISTS `position`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `position` (
  `idEmp` int(11) NOT NULL,
  `prev_pos` varchar(25) DEFAULT NULL,
  `date_of_change` date DEFAULT NULL,
  `current_pos` varchar(25) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `pos_to_emp_idx` (`idEmp`),
  CONSTRAINT `pos_to_emp` FOREIGN KEY (`idEmp`) REFERENCES `employee` (`idEmp`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `result`
--

DROP TABLE IF EXISTS `result`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `result` (
  `year` year(4) NOT NULL,
  `fact_revenue` decimal(4,3) DEFAULT NULL,
  `fact_margin` decimal(4,3) DEFAULT NULL,
  `plan_revenue` decimal(6,2) DEFAULT NULL,
  `plan_margin` decimal(6,2) DEFAULT NULL,
  PRIMARY KEY (`year`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `salary`
--

DROP TABLE IF EXISTS `salary`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `salary` (
  `idEmp` int(11) NOT NULL,
  `previous` int(11) DEFAULT NULL,
  `date_of_change` date DEFAULT NULL,
  `current` int(11) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `idEmp_idx` (`idEmp`),
  CONSTRAINT `idEmp` FOREIGN KEY (`idEmp`) REFERENCES `employee` (`idEmp`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `timeoff`
--

DROP TABLE IF EXISTS `timeoff`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `timeoff` (
  `idEmp` int(11) NOT NULL,
  `date_of_change` date DEFAULT NULL,
  `status` varchar(10) DEFAULT NULL,
  `description` varchar(40) DEFAULT NULL,
  `timeoff_available` decimal(3,1) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `idEmp_idx` (`idEmp`),
  CONSTRAINT `Empid` FOREIGN KEY (`idEmp`) REFERENCES `employee` (`idEmp`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `vacation`
--

DROP TABLE IF EXISTS `vacation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vacation` (
  `idEmp` int(11) NOT NULL,
  `date_of_change` date DEFAULT NULL,
  `description` varchar(45) DEFAULT NULL,
  `vacation_available` int(11) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `vac_emp_idx` (`idEmp`),
  CONSTRAINT `vac_emp` FOREIGN KEY (`idEmp`) REFERENCES `employee` (`idEmp`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-05-23 12:01:18
