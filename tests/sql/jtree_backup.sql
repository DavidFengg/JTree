-- MySQL dump 10.13  Distrib 5.7.21, for osx10.13 (x86_64)
--
-- Host: localhost    Database: TestJTree
-- ------------------------------------------------------
-- Server version	5.7.21

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
-- Table structure for table `experiments`
--

DROP TABLE IF EXISTS `experiments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `experiments` (
  `experiment_id` varchar(255) NOT NULL,
  `study_id` varchar(50) DEFAULT NULL,
  `panel_assay_screened` varchar(50) DEFAULT NULL,
  `test_date` date DEFAULT NULL,
  `chip_cartridge_barcode` varchar(50) DEFAULT NULL,
  `complete_date` date DEFAULT NULL,
  `pcr` varchar(50) DEFAULT NULL,
  `sample_id` varchar(50) DEFAULT NULL,
  `project_name` varchar(50) DEFAULT NULL,
  `priority` varchar(50) DEFAULT NULL,
  `opened_date` date DEFAULT NULL,
  `project_id` varchar(50) DEFAULT NULL,
  `has_project_files` tinyint(1) DEFAULT NULL,
  `procedure_order_datetime` datetime DEFAULT NULL,
  PRIMARY KEY (`experiment_id`),
  KEY `sample_id_idx` (`sample_id`),
  CONSTRAINT `sample_id_ex` FOREIGN KEY (`sample_id`) REFERENCES `samples` (`sample_id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `experiments`
--

LOCK TABLES `experiments` WRITE;
/*!40000 ALTER TABLE `experiments` DISABLE KEYS */;
/*!40000 ALTER TABLE `experiments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `patients`
--

DROP TABLE IF EXISTS `patients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `patients` (
  `first_name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `initials` varchar(50) DEFAULT NULL,
  `gender` varchar(50) DEFAULT NULL,
  `mrn` varchar(50) DEFAULT NULL,
  `dob` date DEFAULT NULL,
  `on_hcn` varchar(50) DEFAULT NULL,
  `clinical_history` varchar(255) DEFAULT NULL,
  `patient_type` varchar(50) DEFAULT NULL,
  `se_num` varchar(50) DEFAULT NULL,
  `patient_id` varchar(50) NOT NULL,
  `date_received` date DEFAULT NULL,
  `referring_physican` varchar(150) DEFAULT NULL,
  `date_reported` date DEFAULT NULL,
  `surgical_date` date DEFAULT NULL,
  PRIMARY KEY (`patient_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `patients`
--

LOCK TABLES `patients` WRITE;
/*!40000 ALTER TABLE `patients` DISABLE KEYS */;
/*!40000 ALTER TABLE `patients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `resultdetails`
--

DROP TABLE IF EXISTS `resultdetails`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `resultdetails` (
  `VAF` float DEFAULT NULL,
  `c_nomenclature` varchar(255) DEFAULT NULL,
  `coverage` int(11) DEFAULT NULL,
  `exon` int(11) DEFAULT NULL,
  `gene` varchar(255) DEFAULT NULL,
  `p_nomenclature` varchar(255) DEFAULT NULL,
  `pcr` varchar(255) DEFAULT NULL,
  `quality_score` float DEFAULT NULL,
  `result` varchar(255) DEFAULT NULL,
  `results_details_id` varchar(255) NOT NULL,
  `results_id` varchar(255) DEFAULT NULL,
  `risk_score` float DEFAULT NULL,
  `uid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`results_details_id`),
  KEY `FK_results_id` (`results_id`),
  CONSTRAINT `FK_results_id` FOREIGN KEY (`results_id`) REFERENCES `results` (`results_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `resultdetails`
--

LOCK TABLES `resultdetails` WRITE;
/*!40000 ALTER TABLE `resultdetails` DISABLE KEYS */;
/*!40000 ALTER TABLE `resultdetails` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `results`
--

DROP TABLE IF EXISTS `results`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `results` (
  `failed_regions` varchar(255) DEFAULT NULL,
  `mean_depth_of_coveage` float DEFAULT NULL,
  `mlpa_pcr` varchar(255) DEFAULT NULL,
  `mutation` varchar(255) DEFAULT NULL,
  `overall_hotspots_threshold` float DEFAULT NULL,
  `overall_quality_threshold` float DEFAULT NULL,
  `results_id` varchar(255) NOT NULL,
  `uid` varchar(255) DEFAULT NULL,
  `verification_pcr` varchar(255) DEFAULT NULL,
  `experiment_id` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`results_id`),
  KEY `FK_experiment_id` (`experiment_id`),
  CONSTRAINT `FK_experiment_id` FOREIGN KEY (`experiment_id`) REFERENCES `experiments` (`experiment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `results`
--

LOCK TABLES `results` WRITE;
/*!40000 ALTER TABLE `results` DISABLE KEYS */;
/*!40000 ALTER TABLE `results` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `samples`
--

DROP TABLE IF EXISTS `samples`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `samples` (
  `sample_id` varchar(50) NOT NULL,
  `facility` varchar(255) DEFAULT NULL,
  `test_requested` varchar(50) DEFAULT NULL,
  `se_num` varchar(50) DEFAULT NULL,
  `date_collected` date DEFAULT NULL,
  `date_received` date DEFAULT NULL,
  `sample_type` varchar(50) DEFAULT NULL,
  `material_received` varchar(150) DEFAULT NULL,
  `material_received_num` varchar(150) DEFAULT NULL,
  `material_received_other` varchar(150) DEFAULT NULL,
  `volume_of_blood_marrow` float(5,1) DEFAULT NULL,
  `surgical_num` varchar(50) DEFAULT NULL,
  `tumor_site` varchar(255) DEFAULT NULL,
  `historical_diagnosis` varchar(255) DEFAULT NULL,
  `tumor_percnt_of_total` float(5,2) DEFAULT NULL,
  `tumor_percnt_of_circled` float(5,2) DEFAULT NULL,
  `reviewed_by` varchar(150) DEFAULT NULL,
  `h_e_slide_location` varchar(150) DEFAULT NULL,
  `non_uhn_id` varchar(50) DEFAULT NULL,
  `name_of_requestor` varchar(150) DEFAULT NULL,
  `dna_concentration` float(10,4) DEFAULT NULL,
  `dna_volume` float(5,1) DEFAULT NULL,
  `dna_location` varchar(255) DEFAULT NULL,
  `rna_concentration` float(10,4) DEFAULT NULL,
  `rna_volume` float(5,1) DEFAULT NULL,
  `rna_location` varchar(150) DEFAULT NULL,
  `wbc_location` varchar(50) DEFAULT NULL,
  `plasma_location` varchar(50) DEFAULT NULL,
  `cf_plasma_location` varchar(50) DEFAULT NULL,
  `pb_bm_location` varchar(50) DEFAULT NULL,
  `rna_lysate_location` varchar(50) DEFAULT NULL,
  `sample_size` varchar(50) DEFAULT NULL,
  `study_id` varchar(50) DEFAULT NULL,
  `sample_name` varchar(50) DEFAULT NULL,
  `date_submitted` date DEFAULT NULL,
  `container_type` varchar(50) DEFAULT NULL,
  `container_name` varchar(100) DEFAULT NULL,
  `container_id` varchar(100) DEFAULT NULL,
  `container_well` varchar(50) DEFAULT NULL,
  `copath_num` varchar(50) DEFAULT NULL,
  `other_identifier` varchar(50) DEFAULT NULL,
  `has_sample_files` tinyint(1) DEFAULT NULL,
  `dna_sample_barcode` varchar(50) DEFAULT NULL,
  `dna_extraction_date` date DEFAULT NULL,
  `dna_quality` varchar(255) DEFAULT NULL,
  `ffpe_qc_date` date DEFAULT NULL,
  `delta_ct_value` float(10,4) DEFAULT NULL,
  `comments` varchar(255) DEFAULT NULL,
  `rnase_p_date` date DEFAULT NULL,
  `dna_quality_by_rnase_p` float(10,4) DEFAULT NULL,
  `rna_quality` float(10,4) DEFAULT NULL,
  `rna_extraction_date` date DEFAULT NULL,
  `patient_id` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`sample_id`),
  KEY `patient_id` (`patient_id`),
  CONSTRAINT `patient_id` FOREIGN KEY (`patient_id`) REFERENCES `patients` (`patient_id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `samples`
--

LOCK TABLES `samples` WRITE;
/*!40000 ALTER TABLE `samples` DISABLE KEYS */;
/*!40000 ALTER TABLE `samples` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-04-05 13:01:33
