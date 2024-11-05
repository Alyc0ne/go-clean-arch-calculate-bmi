CREATE DATABASE IF NOT EXISTS `bmi` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `bmi`;

DROP TABLE IF EXISTS `bmi_condition`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
SET NAMES 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';
CREATE TABLE `bmi_condition` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `min` DECIMAL(4, 1),
  `max` DECIMAL(4, 1),
  `bmi_desc` TEXT COLLATE utf8mb4_unicode_ci NOT NULL,
  `bmi_advice` TEXT COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

-- LOCK TABLES `bmi_condition` WRITE;
-- /*!40000 ALTER TABLE `bmi_condition` DISABLE KEYS */;
-- INSERT INTO `bmi_condition` 
-- VALUES 
--     (1, 'อ้วนมาก', 30.0, NULL, 'ค่อนข้างอันตราย เสี่ยงต่อการเกิดโรคร้ายแรงที่แฝงมากับความอ้วน', 'ควรปรับพฤติกรรมการทานอาหาร และเริ่มออกกำลังกาย หาก BMI สูงกว่า 40.0 ควรไปตรวจสุขภาพและปรึกษาแพทย์');
-- /*!40000 ALTER TABLE `bmi_condition` ENABLE KEYS */;
-- UNLOCK TABLES;