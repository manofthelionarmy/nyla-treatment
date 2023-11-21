-- MySQL dump 10.13  Distrib 8.0.33, for Linux (x86_64)
--
-- Host: mysql    Database: nyla_treatment
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `medicine`
--

DROP TABLE IF EXISTS `medicine`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `medicine` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `time_period_hr` smallint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `medicine`
--

LOCK TABLES `medicine` WRITE;
/*!40000 ALTER TABLE `medicine` DISABLE KEYS */;
INSERT INTO `medicine` VALUES (1,'Prednisolone','Anti-Inflammatory Steroid',12),(2,'Animax','Anti-infection Ointment',12),(3,'Clavamox','Antibiotic',12);
/*!40000 ALTER TABLE `medicine` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `treatment_time`
--

DROP TABLE IF EXISTS `treatment_time`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `treatment_time` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `recorded_time_taken` datetime DEFAULT NULL,
  `medicine_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_treatment_time_medicine` (`medicine_id`),
  CONSTRAINT `fk_treatment_time_medicine` FOREIGN KEY (`medicine_id`) REFERENCES `medicine` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `treatment_time`
--

LOCK TABLES `treatment_time` WRITE;
/*!40000 ALTER TABLE `treatment_time` DISABLE KEYS */;
INSERT INTO `treatment_time` VALUES (1,'2023-11-19 16:07:00',1),(2,'2023-11-19 16:07:00',3),(3,'2023-11-19 14:25:00',2),(4,'2023-11-20 09:36:00',2),(5,'2023-11-20 12:00:00',1),(6,'2023-11-20 12:00:00',3),(7,'2023-11-20 21:41:00',2),(8,'2023-11-21 10:15:00',3),(9,'2023-11-21 10:15:00',1),(10,'2023-11-21 10:46:00',2);
/*!40000 ALTER TABLE `treatment_time` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-11-21 23:44:49
