-- MySQL dump 10.13  Distrib 8.0.27, for macos11 (x86_64)
--
-- Host: localhost    Database: hackaton_db
-- ------------------------------------------------------
-- Server version	8.0.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lastname` varchar(45) DEFAULT NULL,
  `firstname` varchar(45) DEFAULT NULL,
  `condicion` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoices`
--

DROP TABLE IF EXISTS `invoices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `invoices` (
  `id` int NOT NULL AUTO_INCREMENT,
  `datetime` datetime DEFAULT NULL,
  `idcustomer` int DEFAULT NULL,
  `total` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=121 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoices`
--

LOCK TABLES `invoices` WRITE;
/*!40000 ALTER TABLE `invoices` DISABLE KEYS */;
INSERT INTO `invoices` VALUES (1,'2021-12-14 12:24:27',41,89697.7),(2,'2021-12-13 16:52:48',22,45227),(3,'2021-12-21 07:32:40',31,89738.9),(4,'2021-12-10 07:28:33',27,89618.5),(5,'2021-12-28 10:50:44',29,45126),(6,'2022-01-04 21:32:08',31,179277),(7,'2021-12-31 18:04:50',12,89500.7),(8,'2021-12-12 16:32:43',35,179314),(9,'2021-12-22 02:44:55',34,45273.2),(10,'2021-12-21 19:49:03',5,90337.6),(11,'2021-12-05 10:46:48',15,134679),(12,'2021-12-28 21:06:49',8,89583.3),(13,'2021-12-16 11:58:36',13,89894.1),(14,'2021-12-22 09:34:16',14,135305),(15,'2021-12-22 02:29:30',34,89999.7),(16,'2021-12-12 20:22:03',46,245.1),(17,'2021-12-24 09:19:32',24,89339.4),(18,'2021-12-09 19:27:17',3,179633),(19,'2022-01-01 10:26:45',34,89942.7),(20,'2021-12-06 12:15:30',47,179148),(21,'2021-12-27 13:49:51',22,134618),(22,'2022-01-03 22:05:35',24,409.5),(23,'2021-12-24 12:41:40',35,223700),(24,'2021-12-31 14:19:18',20,45148.6),(25,'2021-12-03 13:52:34',39,224230),(26,'2021-12-08 04:05:20',21,312912),(27,'2021-12-05 01:29:50',28,224178),(28,'2021-12-08 10:02:50',32,134562),(29,'2021-12-30 04:17:44',13,574.3),(30,'2021-12-06 01:09:35',8,89908.5),(31,'2021-12-03 05:38:31',10,134643),(32,'2021-12-04 20:24:57',19,45007.7),(33,'2021-12-25 13:49:56',39,45076.5),(34,'2021-12-10 09:08:00',9,134881),(35,'2021-12-18 08:57:54',8,556.2),(36,'2021-12-31 11:12:01',45,881.3),(37,'2021-12-28 07:34:25',4,89964.3),(38,'2021-12-14 22:17:26',47,179580),(39,'2021-12-26 19:56:58',15,134428),(40,'2021-12-26 07:46:58',19,45779.2),(41,'2021-12-14 20:40:50',45,45051.6),(42,'2022-01-02 02:44:43',30,44902.8),(43,'2022-01-02 12:56:27',8,224324),(44,'2021-12-14 18:29:51',35,134571),(45,'2021-12-07 00:17:37',3,45084.1),(46,'2021-12-18 13:02:32',1,89978.8),(47,'2021-12-11 12:51:47',50,45356.9),(48,'2021-12-20 16:55:08',35,90159.9),(49,'2022-01-04 07:37:10',50,89840.2),(50,'2021-12-01 04:47:05',26,90094),(51,'2021-12-09 23:00:09',13,134640),(52,'2021-12-29 17:22:26',50,90324.4),(53,'2021-12-19 08:14:31',26,134185),(54,'2021-12-10 00:01:30',12,537.2),(55,'2021-12-30 01:10:48',16,834.9),(56,'2021-12-01 05:18:55',42,134499),(57,'2021-12-18 18:24:34',9,179619),(58,'2021-12-20 08:11:05',6,89743.8),(59,'2021-12-07 13:36:02',40,44905.2),(60,'2021-12-30 16:56:32',3,179707),(61,'2021-12-22 05:42:50',24,89959.7),(62,'2022-01-01 21:20:47',14,134179),(63,'2021-12-07 09:59:28',9,89995.4),(64,'2021-12-30 03:47:20',46,134611),(65,'2021-12-29 20:20:28',43,179126),(66,'2021-12-17 20:38:16',22,569.8),(67,'2021-12-21 14:34:32',42,134431),(68,'2021-12-26 10:54:21',46,89766.9),(69,'2021-12-16 10:05:51',37,179541),(70,'2021-12-27 00:26:48',5,45248.5),(71,'2021-12-28 23:46:37',50,44962.6),(72,'2021-12-22 16:00:20',46,89734.7),(73,'2021-12-25 16:53:31',36,516.5),(74,'2021-12-27 23:44:30',38,89864),(75,'2021-12-31 01:46:44',33,90091.2),(76,'2021-12-13 00:45:18',1,378.3),(77,'2021-12-21 12:01:16',36,134632),(78,'2021-12-26 19:30:27',17,224181),(79,'2022-01-02 19:56:22',2,90198.2),(80,'2021-12-11 04:34:46',28,89901.4),(81,'2021-12-10 12:17:40',20,45297),(82,'2021-12-19 19:09:15',19,178940),(83,'2021-12-12 19:43:37',8,45517.5),(84,'2021-12-08 15:38:16',12,45378.4),(85,'2021-12-30 19:37:46',44,45039.1),(86,'2021-12-08 14:39:38',37,179036),(87,'2021-12-10 06:37:01',18,44944.5),(88,'2021-12-08 22:33:14',47,89804.2),(89,'2021-12-19 07:47:23',7,134910),(90,'2021-12-15 13:49:41',37,134685),(91,'2022-01-01 11:38:04',31,179008),(92,'2021-12-01 15:32:55',13,134248),(93,'2021-12-05 17:39:04',14,89767.2),(94,'2021-12-12 14:07:05',42,224272),(95,'2021-12-08 22:53:32',32,89939.1),(96,'2021-12-28 06:22:47',21,45012.7),(97,'2022-01-03 18:39:42',16,134919),(98,'2021-12-28 05:57:00',9,44991.1),(99,'2022-01-01 13:00:01',45,89477.1),(100,'2021-12-24 06:53:50',6,832.6);
/*!40000 ALTER TABLE `invoices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(45) DEFAULT NULL,
  `price` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'Mountain Dew',665.85),(2,'Dried Apple',663.3),(3,'Huck Towels White',510.63),(4,'Paper - Brown Paper Mini Cups',219),(5,'Mustard - Seed',493.5),(6,'Oyster - In Shell',373.61),(7,'Chocolate - Pistoles, Lactee, Milk',659.21),(8,'Sugar - Crumb',120.8),(9,'Bagels Poppyseed',438.23),(10,'Bread - Italian Sesame Poly',96.4),(11,'Juice - Propel Sport',121.43),(12,'Lotus Root',728),(13,'Isomalt',708.04),(14,'Muffin - Zero Transfat',268.13),(15,'Cake - Miini Cheesecake Cherry',463.95),(16,'Cherries - Bing, Canned',323.96),(17,'Pepper - Chilli Seeds Mild',509.24),(18,'Cod - Fillets',563.55),(19,'Table Cloth 120 Round White',145.88),(20,'Capon - Breast, Wing On',776.89),(21,'Yoplait Drink',423.47),(22,'Energy Drink - Redbull 355ml',203.68),(23,'Beef - Cooked, Corned',455.75),(24,'Syrup - Monin - Blue Curacao',498.67),(25,'Arctic Char - Fresh, Whole',457.65),(26,'Aspic - Light',365.3),(27,'Mace',116.26),(28,'Grenadillo',743.35),(29,'Sprouts Dikon',713.27),(30,'Drambuie',451.18),(31,'Dill Weed - Dry',617.86),(32,'Grouper - Fresh',306.36),(33,'Pork - Belly Fresh',647.03),(34,'Kellogs Raisan Bran Bars',557.09),(35,'Pasta - Angel Hair',510.81),(36,'Jack Daniels',159.24),(37,'Pepsi - 600ml',780.22),(38,'Calaloo',450.6),(39,'Skewers - Bamboo',651.33),(40,'Shrimp - 150 - 250',621.86),(41,'Tendrils - Baby Pea, Organic',576.64),(42,'Taro Leaves',557.54),(43,'Soy Protein',171.36),(44,'Bread - Multigrain Oval',228.1),(45,'Pastry - Chocolate Chip Muffin',117.43),(46,'Gatorade - Cool Blue Raspberry',269.65),(47,'Broom - Corn',710.26),(48,'Appetizer - Cheese Bites',103.11),(49,'Liquid Aminios Acid - Braggs',756.17),(50,'Syrup - Pancake',290.55),(51,'Chevril',113.34),(52,'Bread - Malt',163.51),(53,'Onions - Cooking',181.61),(54,'Muffin Batt - Blueberry Passion',384.98),(55,'Doilies - 10, Paper',416.7),(56,'Wine - Alsace Riesling Reserve',310.02),(57,'Chervil - Fresh',436.53),(58,'Glaze - Clear',742.33),(59,'Wine - Malbec Trapiche Reserve',599.65),(60,'Apples - Spartan',340.03),(61,'Pork - Butt, Boneless',617.44),(62,'Baking Powder',131.79),(63,'Syrup - Monin - Blue Curacao',97.66),(64,'Oysters - Smoked',598.53),(65,'Flounder - Fresh',194.23),(66,'Mince Meat - Filling',478.07),(67,'Longos - Chicken Cordon Bleu',325.22),(68,'Cheese - Asiago',118.44),(69,'Skirt - 24 Foot',437.71),(70,'Flour - Chickpea',844.96),(71,'Salmon - Fillets',638.34),(72,'Lumpfish Black',199.48),(73,'Sesame Seed Black',815.04),(74,'Gherkin',647.68),(75,'Mustard - Seed',273.88),(76,'Vinegar - Rice',757.31),(77,'Soup - Canadian Pea, Dry Mix',698.18),(78,'Wine - Baron De Rothschild',270.29),(79,'Nut - Pistachio, Shelled',611.54),(80,'Veal - Heart',722.97),(81,'Cookies - Englishbay Wht',538.09),(82,'Wine - Harrow Estates, Vidal',250.66),(83,'Pears - Fiorelle',522.65),(84,'Cheese - Bakers Cream Cheese',434.54),(85,'Crush - Grape, 355 Ml',801.12),(86,'Beer - True North Lager',312.08),(87,'Sardines',556.05),(88,'Tomato - Green',298.13),(89,'Tomatoes - Vine Ripe, Yellow',728.56),(90,'Bread Foccacia Whole',735.44),(91,'Water - Mineral, Carbonated',338.45),(92,'Syrup - Monin, Swiss Choclate',831.25),(93,'Nori Sea Weed',661.4),(94,'Lychee',212.56),(95,'Fenngreek Seed',236.74),(96,'Marsala - Sperone, Fine, D.o.c.',526.4),(97,'Vinegar - Balsamic',503.78),(98,'Clementine',612.18),(99,'Bread - White Epi Baguette',448.55),(100,'Chocolate - Feathers',646.12);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sales`
--

DROP TABLE IF EXISTS `sales`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sales` (
  `id` int NOT NULL AUTO_INCREMENT,
  `idinvoice` int DEFAULT NULL,
  `idproduct` int DEFAULT NULL,
  `quantity` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1001 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sales`
--

LOCK TABLES `sales` WRITE;
/*!40000 ALTER TABLE `sales` DISABLE KEYS */;
/*!40000 ALTER TABLE `sales` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-01-10 14:27:38
