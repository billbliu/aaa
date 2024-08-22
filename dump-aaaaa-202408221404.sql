-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: 120.24.250.172    Database: aaaaa
-- ------------------------------------------------------
-- Server version	8.0.36-0ubuntu0.22.04.1

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
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=189 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (2,'p','888','/api/createApi','POST','','',''),(5,'p','888','/api/deleteApi','POST','','',''),(8,'p','888','/api/deleteApisByIds','DELETE','','',''),(11,'p','888','/api/enterSyncApi','POST','','',''),(7,'p','888','/api/getAllApis','POST','','',''),(4,'p','888','/api/getApiById','POST','','',''),(10,'p','888','/api/getApiGroups','GET','','',''),(3,'p','888','/api/getApiList','POST','','',''),(12,'p','888','/api/ignoreApi','POST','','',''),(9,'p','888','/api/syncApi','GET','','',''),(6,'p','888','/api/updateApi','POST','','',''),(13,'p','888','/authority/copyAuthority','POST','','',''),(15,'p','888','/authority/createAuthority','POST','','',''),(16,'p','888','/authority/deleteAuthority','POST','','',''),(17,'p','888','/authority/getAuthorityList','POST','','',''),(18,'p','888','/authority/setDataAuthority','POST','','',''),(14,'p','888','/authority/updateAuthority','PUT','','',''),(96,'p','888','/authorityBtn/canRemoveAuthorityBtn','POST','','',''),(95,'p','888','/authorityBtn/getAuthorityBtn','POST','','',''),(94,'p','888','/authorityBtn/setAuthorityBtn','POST','','',''),(72,'p','888','/autoCode/addFunc','POST','','',''),(65,'p','888','/autoCode/createPackage','POST','','',''),(69,'p','888','/autoCode/createPlug','POST','','',''),(62,'p','888','/autoCode/createTemp','POST','','',''),(68,'p','888','/autoCode/delPackage','POST','','',''),(63,'p','888','/autoCode/delSysHistory','POST','','',''),(60,'p','888','/autoCode/getColumn','GET','','',''),(56,'p','888','/autoCode/getDB','GET','','',''),(57,'p','888','/autoCode/getMeta','POST','','',''),(67,'p','888','/autoCode/getPackage','POST','','',''),(64,'p','888','/autoCode/getSysHistory','POST','','',''),(59,'p','888','/autoCode/getTables','GET','','',''),(66,'p','888','/autoCode/getTemplates','GET','','',''),(70,'p','888','/autoCode/installPlugin','POST','','',''),(58,'p','888','/autoCode/preview','POST','','',''),(71,'p','888','/autoCode/pubPlug','POST','','',''),(61,'p','888','/autoCode/rollback','POST','','',''),(46,'p','888','/casbin/getPolicyPathByAuthorityId','POST','','',''),(45,'p','888','/casbin/updateCasbin','POST','','',''),(54,'p','888','/customer/customer','DELETE','','',''),(51,'p','888','/customer/customer','GET','','',''),(53,'p','888','/customer/customer','POST','','',''),(52,'p','888','/customer/customer','PUT','','',''),(55,'p','888','/customer/customerList','GET','','',''),(89,'p','888','/email/emailTest','POST','','',''),(90,'p','888','/email/sendEmail','POST','','',''),(39,'p','888','/fileUploadAndDownload/breakpointContinue','POST','','',''),(38,'p','888','/fileUploadAndDownload/breakpointContinueFinish','POST','','',''),(42,'p','888','/fileUploadAndDownload/deleteFile','POST','','',''),(43,'p','888','/fileUploadAndDownload/editFileName','POST','','',''),(37,'p','888','/fileUploadAndDownload/findFile','GET','','',''),(44,'p','888','/fileUploadAndDownload/getFileList','POST','','',''),(40,'p','888','/fileUploadAndDownload/removeChunk','POST','','',''),(41,'p','888','/fileUploadAndDownload/upload','POST','','',''),(106,'p','888','/info/createInfo','POST','','',''),(107,'p','888','/info/deleteInfo','DELETE','','',''),(108,'p','888','/info/deleteInfoByIds','DELETE','','',''),(110,'p','888','/info/findInfo','GET','','',''),(111,'p','888','/info/getInfoList','GET','','',''),(109,'p','888','/info/updateInfo','PUT','','',''),(47,'p','888','/jwt/jsonInBlacklist','POST','','',''),(21,'p','888','/menu/addBaseMenu','POST','','',''),(23,'p','888','/menu/addMenuAuthority','POST','','',''),(25,'p','888','/menu/deleteBaseMenu','POST','','',''),(27,'p','888','/menu/getBaseMenuById','POST','','',''),(22,'p','888','/menu/getBaseMenuTree','POST','','',''),(19,'p','888','/menu/getMenu','POST','','',''),(24,'p','888','/menu/getMenuAuthority','POST','','',''),(20,'p','888','/menu/getMenuList','POST','','',''),(26,'p','888','/menu/updateBaseMenu','POST','','',''),(92,'p','888','/simpleUploader/checkFileMd5','GET','','',''),(93,'p','888','/simpleUploader/mergeFileMd5','GET','','',''),(91,'p','888','/simpleUploader/upload','POST','','',''),(81,'p','888','/sysDictionary/createSysDictionary','POST','','',''),(82,'p','888','/sysDictionary/deleteSysDictionary','DELETE','','',''),(78,'p','888','/sysDictionary/findSysDictionary','GET','','',''),(80,'p','888','/sysDictionary/getSysDictionaryList','GET','','',''),(79,'p','888','/sysDictionary/updateSysDictionary','PUT','','',''),(75,'p','888','/sysDictionaryDetail/createSysDictionaryDetail','POST','','',''),(77,'p','888','/sysDictionaryDetail/deleteSysDictionaryDetail','DELETE','','',''),(73,'p','888','/sysDictionaryDetail/findSysDictionaryDetail','GET','','',''),(76,'p','888','/sysDictionaryDetail/getSysDictionaryDetailList','GET','','',''),(74,'p','888','/sysDictionaryDetail/updateSysDictionaryDetail','PUT','','',''),(97,'p','888','/sysExportTemplate/createSysExportTemplate','POST','','',''),(98,'p','888','/sysExportTemplate/deleteSysExportTemplate','DELETE','','',''),(99,'p','888','/sysExportTemplate/deleteSysExportTemplateByIds','DELETE','','',''),(103,'p','888','/sysExportTemplate/exportExcel','GET','','',''),(104,'p','888','/sysExportTemplate/exportTemplate','GET','','',''),(101,'p','888','/sysExportTemplate/findSysExportTemplate','GET','','',''),(102,'p','888','/sysExportTemplate/getSysExportTemplateList','GET','','',''),(105,'p','888','/sysExportTemplate/importExcel','POST','','',''),(100,'p','888','/sysExportTemplate/updateSysExportTemplate','PUT','','',''),(85,'p','888','/sysOperationRecord/createSysOperationRecord','POST','','',''),(87,'p','888','/sysOperationRecord/deleteSysOperationRecord','DELETE','','',''),(88,'p','888','/sysOperationRecord/deleteSysOperationRecordByIds','DELETE','','',''),(83,'p','888','/sysOperationRecord/findSysOperationRecord','GET','','',''),(86,'p','888','/sysOperationRecord/getSysOperationRecordList','GET','','',''),(84,'p','888','/sysOperationRecord/updateSysOperationRecord','PUT','','',''),(50,'p','888','/system/getServerInfo','POST','','',''),(48,'p','888','/system/getSystemConfig','POST','','',''),(49,'p','888','/system/setSystemConfig','POST','','',''),(1,'p','888','/user/admin_register','POST','','',''),(33,'p','888','/user/changePassword','POST','','',''),(32,'p','888','/user/deleteUser','DELETE','','',''),(28,'p','888','/user/getUserInfo','GET','','',''),(31,'p','888','/user/getUserList','POST','','',''),(36,'p','888','/user/resetPassword','POST','','',''),(30,'p','888','/user/setSelfInfo','PUT','','',''),(35,'p','888','/user/setUserAuthorities','POST','','',''),(34,'p','888','/user/setUserAuthority','POST','','',''),(29,'p','888','/user/setUserInfo','PUT','','',''),(113,'p','8881','/api/createApi','POST','','',''),(116,'p','8881','/api/deleteApi','POST','','',''),(118,'p','8881','/api/getAllApis','POST','','',''),(115,'p','8881','/api/getApiById','POST','','',''),(114,'p','8881','/api/getApiList','POST','','',''),(117,'p','8881','/api/updateApi','POST','','',''),(119,'p','8881','/authority/createAuthority','POST','','',''),(120,'p','8881','/authority/deleteAuthority','POST','','',''),(121,'p','8881','/authority/getAuthorityList','POST','','',''),(122,'p','8881','/authority/setDataAuthority','POST','','',''),(140,'p','8881','/casbin/getPolicyPathByAuthorityId','POST','','',''),(139,'p','8881','/casbin/updateCasbin','POST','','',''),(146,'p','8881','/customer/customer','DELETE','','',''),(147,'p','8881','/customer/customer','GET','','',''),(144,'p','8881','/customer/customer','POST','','',''),(145,'p','8881','/customer/customer','PUT','','',''),(148,'p','8881','/customer/customerList','GET','','',''),(137,'p','8881','/fileUploadAndDownload/deleteFile','POST','','',''),(138,'p','8881','/fileUploadAndDownload/editFileName','POST','','',''),(136,'p','8881','/fileUploadAndDownload/getFileList','POST','','',''),(135,'p','8881','/fileUploadAndDownload/upload','POST','','',''),(141,'p','8881','/jwt/jsonInBlacklist','POST','','',''),(125,'p','8881','/menu/addBaseMenu','POST','','',''),(127,'p','8881','/menu/addMenuAuthority','POST','','',''),(129,'p','8881','/menu/deleteBaseMenu','POST','','',''),(131,'p','8881','/menu/getBaseMenuById','POST','','',''),(126,'p','8881','/menu/getBaseMenuTree','POST','','',''),(123,'p','8881','/menu/getMenu','POST','','',''),(128,'p','8881','/menu/getMenuAuthority','POST','','',''),(124,'p','8881','/menu/getMenuList','POST','','',''),(130,'p','8881','/menu/updateBaseMenu','POST','','',''),(142,'p','8881','/system/getSystemConfig','POST','','',''),(143,'p','8881','/system/setSystemConfig','POST','','',''),(112,'p','8881','/user/admin_register','POST','','',''),(132,'p','8881','/user/changePassword','POST','','',''),(149,'p','8881','/user/getUserInfo','GET','','',''),(133,'p','8881','/user/getUserList','POST','','',''),(134,'p','8881','/user/setUserAuthority','POST','','',''),(151,'p','9528','/api/createApi','POST','','',''),(154,'p','9528','/api/deleteApi','POST','','',''),(156,'p','9528','/api/getAllApis','POST','','',''),(153,'p','9528','/api/getApiById','POST','','',''),(152,'p','9528','/api/getApiList','POST','','',''),(155,'p','9528','/api/updateApi','POST','','',''),(157,'p','9528','/authority/createAuthority','POST','','',''),(158,'p','9528','/authority/deleteAuthority','POST','','',''),(159,'p','9528','/authority/getAuthorityList','POST','','',''),(160,'p','9528','/authority/setDataAuthority','POST','','',''),(187,'p','9528','/autoCode/createTemp','POST','','',''),(178,'p','9528','/casbin/getPolicyPathByAuthorityId','POST','','',''),(177,'p','9528','/casbin/updateCasbin','POST','','',''),(185,'p','9528','/customer/customer','DELETE','','',''),(183,'p','9528','/customer/customer','GET','','',''),(184,'p','9528','/customer/customer','POST','','',''),(182,'p','9528','/customer/customer','PUT','','',''),(186,'p','9528','/customer/customerList','GET','','',''),(175,'p','9528','/fileUploadAndDownload/deleteFile','POST','','',''),(176,'p','9528','/fileUploadAndDownload/editFileName','POST','','',''),(174,'p','9528','/fileUploadAndDownload/getFileList','POST','','',''),(173,'p','9528','/fileUploadAndDownload/upload','POST','','',''),(179,'p','9528','/jwt/jsonInBlacklist','POST','','',''),(163,'p','9528','/menu/addBaseMenu','POST','','',''),(165,'p','9528','/menu/addMenuAuthority','POST','','',''),(167,'p','9528','/menu/deleteBaseMenu','POST','','',''),(169,'p','9528','/menu/getBaseMenuById','POST','','',''),(164,'p','9528','/menu/getBaseMenuTree','POST','','',''),(161,'p','9528','/menu/getMenu','POST','','',''),(166,'p','9528','/menu/getMenuAuthority','POST','','',''),(162,'p','9528','/menu/getMenuList','POST','','',''),(168,'p','9528','/menu/updateBaseMenu','POST','','',''),(180,'p','9528','/system/getSystemConfig','POST','','',''),(181,'p','9528','/system/setSystemConfig','POST','','',''),(150,'p','9528','/user/admin_register','POST','','',''),(170,'p','9528','/user/changePassword','POST','','',''),(188,'p','9528','/user/getUserInfo','GET','','',''),(171,'p','9528','/user/getUserList','POST','','',''),(172,'p','9528','/user/setUserAuthority','POST','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_customers`
--

DROP TABLE IF EXISTS `exa_customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_customers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint unsigned DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint unsigned DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_customers`
--

LOCK TABLES `exa_customers` WRITE;
/*!40000 ALTER TABLE `exa_customers` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_file_chunks`
--

DROP TABLE IF EXISTS `exa_file_chunks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_file_chunks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `exa_file_id` bigint unsigned DEFAULT NULL,
  `file_chunk_number` bigint DEFAULT NULL,
  `file_chunk_path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_file_chunks`
--

LOCK TABLES `exa_file_chunks` WRITE;
/*!40000 ALTER TABLE `exa_file_chunks` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_file_chunks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_file_upload_and_downloads`
--

DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_file_upload_and_downloads` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_file_upload_and_downloads`
--

LOCK TABLES `exa_file_upload_and_downloads` WRITE;
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` DISABLE KEYS */;
INSERT INTO `exa_file_upload_and_downloads` VALUES (1,'2024-08-22 14:01:35.295','2024-08-22 14:01:35.295',NULL,'10.png','https://qmplusimg.henrongyi.top/gvalogo.png','png','158787308910.png'),(2,'2024-08-22 14:01:35.295','2024-08-22 14:01:35.295',NULL,'logo.png','https://qmplusimg.henrongyi.top/1576554439myAvatar.png','png','1587973709logo.png');
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_files`
--

DROP TABLE IF EXISTS `exa_files`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_files` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `file_md5` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `file_path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `chunk_total` bigint DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_files_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_files`
--

LOCK TABLES `exa_files` WRITE;
/*!40000 ALTER TABLE `exa_files` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_files` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jwt_blacklists`
--

DROP TABLE IF EXISTS `jwt_blacklists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jwt_blacklists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text COLLATE utf8mb4_general_ci COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jwt_blacklists`
--

LOCK TABLES `jwt_blacklists` WRITE;
/*!40000 ALTER TABLE `jwt_blacklists` DISABLE KEYS */;
/*!40000 ALTER TABLE `jwt_blacklists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_apis`
--

DROP TABLE IF EXISTS `sys_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=110 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_apis`
--

LOCK TABLES `sys_apis` WRITE;
/*!40000 ALTER TABLE `sys_apis` DISABLE KEYS */;
INSERT INTO `sys_apis` VALUES (1,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/jwt/jsonInBlacklist','jwt加入黑名单(退出，必选)','jwt','POST'),(2,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/user/deleteUser','删除用户','系统用户','DELETE'),(3,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/user/admin_register','用户注册','系统用户','POST'),(4,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/user/getUserList','获取用户列表','系统用户','POST'),(5,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/user/setUserInfo','设置用户信息','系统用户','PUT'),(6,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/user/setSelfInfo','设置自身信息(必选)','系统用户','PUT'),(7,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/user/getUserInfo','获取自身信息(必选)','系统用户','GET'),(8,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/user/setUserAuthorities','设置权限组','系统用户','POST'),(9,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/user/changePassword','修改密码（建议选择)','系统用户','POST'),(10,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/user/setUserAuthority','修改用户角色(必选)','系统用户','POST'),(11,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/user/resetPassword','重置用户密码','系统用户','POST'),(12,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/api/createApi','创建api','api','POST'),(13,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/api/deleteApi','删除Api','api','POST'),(14,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/api/updateApi','更新Api','api','POST'),(15,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/api/getApiList','获取api列表','api','POST'),(16,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/api/getAllApis','获取所有api','api','POST'),(17,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/api/getApiById','获取api详细信息','api','POST'),(18,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/api/deleteApisByIds','批量删除api','api','DELETE'),(19,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/api/syncApi','获取待同步API','api','GET'),(20,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/api/getApiGroups','获取路由组','api','GET'),(21,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/api/enterSyncApi','确认同步API','api','POST'),(22,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/api/ignoreApi','忽略API','api','POST'),(23,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/authority/copyAuthority','拷贝角色','角色','POST'),(24,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/authority/createAuthority','创建角色','角色','POST'),(25,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/authority/deleteAuthority','删除角色','角色','POST'),(26,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/authority/updateAuthority','更新角色信息','角色','PUT'),(27,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/authority/getAuthorityList','获取角色列表','角色','POST'),(28,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/authority/setDataAuthority','设置角色资源权限','角色','POST'),(29,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/casbin/updateCasbin','更改角色api权限','casbin','POST'),(30,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/casbin/getPolicyPathByAuthorityId','获取权限列表','casbin','POST'),(31,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/menu/addBaseMenu','新增菜单','菜单','POST'),(32,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/menu/getMenu','获取菜单树(必选)','菜单','POST'),(33,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/menu/deleteBaseMenu','删除菜单','菜单','POST'),(34,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/menu/updateBaseMenu','更新菜单','菜单','POST'),(35,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/menu/getBaseMenuById','根据id获取菜单','菜单','POST'),(36,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/menu/getMenuList','分页获取基础menu列表','菜单','POST'),(37,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/menu/getBaseMenuTree','获取用户动态路由','菜单','POST'),(38,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/menu/getMenuAuthority','获取指定角色menu','菜单','POST'),(39,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/menu/addMenuAuthority','增加menu和角色关联关系','菜单','POST'),(40,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/fileUploadAndDownload/findFile','寻找目标文件（秒传）','分片上传','GET'),(41,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/fileUploadAndDownload/breakpointContinue','断点续传','分片上传','POST'),(42,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/fileUploadAndDownload/breakpointContinueFinish','断点续传完成','分片上传','POST'),(43,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/fileUploadAndDownload/removeChunk','上传完成移除文件','分片上传','POST'),(44,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/fileUploadAndDownload/upload','文件上传示例','文件上传与下载','POST'),(45,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/fileUploadAndDownload/deleteFile','删除文件','文件上传与下载','POST'),(46,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/fileUploadAndDownload/editFileName','文件名或者备注编辑','文件上传与下载','POST'),(47,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/fileUploadAndDownload/getFileList','获取上传文件列表','文件上传与下载','POST'),(48,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/system/getServerInfo','获取服务器信息','系统服务','POST'),(49,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/system/getSystemConfig','获取配置文件内容','系统服务','POST'),(50,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/system/setSystemConfig','设置配置文件内容','系统服务','POST'),(51,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/customer/customer','更新客户','客户','PUT'),(52,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/customer/customer','创建客户','客户','POST'),(53,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/customer/customer','删除客户','客户','DELETE'),(54,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/customer/customer','获取单一客户','客户','GET'),(55,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/customer/customerList','获取客户列表','客户','GET'),(56,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/getDB','获取所有数据库','代码生成器','GET'),(57,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/getTables','获取数据库表','代码生成器','GET'),(58,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/createTemp','自动化代码','代码生成器','POST'),(59,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/preview','预览自动化代码','代码生成器','POST'),(60,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/getColumn','获取所选table的所有字段','代码生成器','GET'),(61,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/installPlugin','安装插件','代码生成器','POST'),(62,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/pubPlug','打包插件','代码生成器','POST'),(63,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/createPackage','配置模板','模板配置','POST'),(64,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/getTemplates','获取模板文件','模板配置','GET'),(65,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/getPackage','获取所有模板','模板配置','POST'),(66,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/delPackage','删除模板','模板配置','POST'),(67,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/getMeta','获取meta信息','代码生成器历史','POST'),(68,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/rollback','回滚自动生成代码','代码生成器历史','POST'),(69,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/getSysHistory','查询回滚记录','代码生成器历史','POST'),(70,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/delSysHistory','删除回滚记录','代码生成器历史','POST'),(71,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/autoCode/addFunc','增加模板方法','代码生成器历史','POST'),(72,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysDictionaryDetail/updateSysDictionaryDetail','更新字典内容','系统字典详情','PUT'),(73,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysDictionaryDetail/createSysDictionaryDetail','新增字典内容','系统字典详情','POST'),(74,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysDictionaryDetail/deleteSysDictionaryDetail','删除字典内容','系统字典详情','DELETE'),(75,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysDictionaryDetail/findSysDictionaryDetail','根据ID获取字典内容','系统字典详情','GET'),(76,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysDictionaryDetail/getSysDictionaryDetailList','获取字典内容列表','系统字典详情','GET'),(77,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysDictionary/createSysDictionary','新增字典','系统字典','POST'),(78,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysDictionary/deleteSysDictionary','删除字典','系统字典','DELETE'),(79,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysDictionary/updateSysDictionary','更新字典','系统字典','PUT'),(80,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysDictionary/findSysDictionary','根据ID获取字典','系统字典','GET'),(81,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysDictionary/getSysDictionaryList','获取字典列表','系统字典','GET'),(82,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysOperationRecord/createSysOperationRecord','新增操作记录','操作记录','POST'),(83,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysOperationRecord/findSysOperationRecord','根据ID获取操作记录','操作记录','GET'),(84,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysOperationRecord/getSysOperationRecordList','获取操作记录列表','操作记录','GET'),(85,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysOperationRecord/deleteSysOperationRecord','删除操作记录','操作记录','DELETE'),(86,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysOperationRecord/deleteSysOperationRecordByIds','批量删除操作历史','操作记录','DELETE'),(87,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/simpleUploader/upload','插件版分片上传','断点续传(插件版)','POST'),(88,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/simpleUploader/checkFileMd5','文件完整度验证','断点续传(插件版)','GET'),(89,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/simpleUploader/mergeFileMd5','上传完成合并文件','断点续传(插件版)','GET'),(90,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/email/emailTest','发送测试邮件','email','POST'),(91,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/email/sendEmail','发送邮件','email','POST'),(92,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/authorityBtn/setAuthorityBtn','设置按钮权限','按钮权限','POST'),(93,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/authorityBtn/getAuthorityBtn','获取已有按钮权限','按钮权限','POST'),(94,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/authorityBtn/canRemoveAuthorityBtn','删除按钮','按钮权限','POST'),(95,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysExportTemplate/createSysExportTemplate','新增导出模板','表格模板','POST'),(96,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysExportTemplate/deleteSysExportTemplate','删除导出模板','表格模板','DELETE'),(97,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysExportTemplate/deleteSysExportTemplateByIds','批量删除导出模板','表格模板','DELETE'),(98,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysExportTemplate/updateSysExportTemplate','更新导出模板','表格模板','PUT'),(99,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysExportTemplate/findSysExportTemplate','根据ID获取导出模板','表格模板','GET'),(100,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysExportTemplate/getSysExportTemplateList','获取导出模板列表','表格模板','GET'),(101,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysExportTemplate/exportExcel','导出Excel','表格模板','GET'),(102,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysExportTemplate/exportTemplate','下载模板','表格模板','GET'),(103,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/sysExportTemplate/importExcel','导入Excel','表格模板','POST'),(104,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/info/createInfo','新建公告','公告','POST'),(105,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/info/deleteInfo','删除公告','公告','DELETE'),(106,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/info/deleteInfoByIds','批量删除公告','公告','DELETE'),(107,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/info/updateInfo','更新公告','公告','PUT'),(108,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/info/findInfo','根据ID获取公告','公告','GET'),(109,'2024-08-22 14:01:25.593','2024-08-22 14:01:25.593',NULL,'/info/getInfoList','获取公告列表','公告','GET');
/*!40000 ALTER TABLE `sys_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authorities`
--

DROP TABLE IF EXISTS `sys_authorities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `uni_sys_authorities_authority_id` (`authority_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9529 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authorities`
--

LOCK TABLES `sys_authorities` WRITE;
/*!40000 ALTER TABLE `sys_authorities` DISABLE KEYS */;
INSERT INTO `sys_authorities` VALUES ('2024-08-22 14:01:26.523','2024-08-22 14:01:33.122',NULL,888,'普通用户',0,'dashboard'),('2024-08-22 14:01:26.523','2024-08-22 14:01:34.817',NULL,8881,'普通用户子角色',888,'dashboard'),('2024-08-22 14:01:26.523','2024-08-22 14:01:33.677',NULL,9528,'测试角色',0,'dashboard');
/*!40000 ALTER TABLE `sys_authorities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authority_btns`
--

DROP TABLE IF EXISTS `sys_authority_btns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authority_btns` (
  `authority_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authority_btns`
--

LOCK TABLES `sys_authority_btns` WRITE;
/*!40000 ALTER TABLE `sys_authority_btns` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_authority_btns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authority_menus`
--

DROP TABLE IF EXISTS `sys_authority_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authority_menus`
--

LOCK TABLES `sys_authority_menus` WRITE;
/*!40000 ALTER TABLE `sys_authority_menus` DISABLE KEYS */;
INSERT INTO `sys_authority_menus` VALUES (1,888),(1,8881),(1,9528),(2,888),(2,8881),(2,9528),(3,888),(4,888),(4,8881),(5,888),(5,8881),(6,888),(6,8881),(7,888),(7,8881),(8,888),(8,8881),(8,9528),(9,888),(9,8881),(10,888),(10,8881),(11,888),(11,8881),(12,888),(13,888),(13,8881),(14,888),(14,8881),(15,888),(15,8881),(16,888),(16,8881),(17,888),(17,8881),(18,888),(19,888),(20,888),(21,888),(22,888),(23,888),(24,888),(25,888),(26,888),(27,888),(28,888),(29,888),(30,888);
/*!40000 ALTER TABLE `sys_authority_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_auto_code_histories`
--

DROP TABLE IF EXISTS `sys_auto_code_histories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_auto_code_histories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `table_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '表名',
  `package` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模块名/插件名',
  `request` text COLLATE utf8mb4_general_ci COMMENT '前端传入的结构化信息',
  `struct_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '结构体名称',
  `business_db` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '业务库',
  `description` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'Struct中文名称',
  `templates` text COLLATE utf8mb4_general_ci COMMENT '模板信息',
  `Injections` text COLLATE utf8mb4_general_ci COMMENT '注入路径',
  `flag` bigint DEFAULT NULL COMMENT '[0:创建,1:回滚]',
  `api_ids` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api表注册内容',
  `menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `package_id` bigint unsigned DEFAULT NULL COMMENT '包ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_auto_code_histories`
--

LOCK TABLES `sys_auto_code_histories` WRITE;
/*!40000 ALTER TABLE `sys_auto_code_histories` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_auto_code_histories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_auto_code_packages`
--

DROP TABLE IF EXISTS `sys_auto_code_packages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_auto_code_packages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `desc` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  `label` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示名',
  `template` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模版',
  `package_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '包名',
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_packages_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_auto_code_packages`
--

LOCK TABLES `sys_auto_code_packages` WRITE;
/*!40000 ALTER TABLE `sys_auto_code_packages` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_auto_code_packages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menu_btns`
--

DROP TABLE IF EXISTS `sys_base_menu_btns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menu_btns` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menu_btns`
--

LOCK TABLES `sys_base_menu_btns` WRITE;
/*!40000 ALTER TABLE `sys_base_menu_btns` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_btns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menu_parameters`
--

DROP TABLE IF EXISTS `sys_base_menu_parameters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL,
  `type` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menu_parameters`
--

LOCK TABLES `sys_base_menu_parameters` WRITE;
/*!40000 ALTER TABLE `sys_base_menu_parameters` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_parameters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menus`
--

DROP TABLE IF EXISTS `sys_base_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint unsigned DEFAULT NULL,
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menus`
--

LOCK TABLES `sys_base_menus` WRITE;
/*!40000 ALTER TABLE `sys_base_menus` DISABLE KEYS */;
INSERT INTO `sys_base_menus` VALUES (1,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,0,'dashboard','dashboard',0,'view/dashboard/index.vue',1,'',0,0,'仪表盘','odometer',0),(2,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,0,'about','about',0,'view/about/index.vue',9,'',0,0,'关于我们','info-filled',0),(3,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,0,'admin','superAdmin',0,'view/superAdmin/index.vue',3,'',0,0,'超级管理员','user',0),(4,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,3,'authority','authority',0,'view/superAdmin/authority/authority.vue',1,'',0,0,'角色管理','avatar',0),(5,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,3,'menu','menu',0,'view/superAdmin/menu/menu.vue',2,'',1,0,'菜单管理','tickets',0),(6,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,3,'api','api',0,'view/superAdmin/api/api.vue',3,'',1,0,'api管理','platform',0),(7,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,3,'user','user',0,'view/superAdmin/user/user.vue',4,'',0,0,'用户管理','coordinate',0),(8,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,3,'dictionary','dictionary',0,'view/superAdmin/dictionary/sysDictionary.vue',5,'',0,0,'字典管理','notebook',0),(9,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,3,'operation','operation',0,'view/superAdmin/operation/sysOperationRecord.vue',6,'',0,0,'操作历史','pie-chart',0),(10,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,0,'person','person',1,'view/person/person.vue',4,'',0,0,'个人信息','message',0),(11,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,0,'example','example',0,'view/example/index.vue',7,'',0,0,'示例文件','management',0),(12,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,11,'upload','upload',0,'view/example/upload/upload.vue',5,'',0,0,'媒体库（上传下载）','upload',0),(13,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,11,'breakpoint','breakpoint',0,'view/example/breakpoint/breakpoint.vue',6,'',0,0,'断点续传','upload-filled',0),(14,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,11,'customer','customer',0,'view/example/customer/customer.vue',7,'',0,0,'客户列表（资源示例）','avatar',0),(15,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,0,'systemTools','systemTools',0,'view/systemTools/index.vue',5,'',0,0,'系统工具','tools',0),(16,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,15,'autoCode','autoCode',0,'view/systemTools/autoCode/index.vue',1,'',1,0,'代码生成器','cpu',0),(17,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,15,'formCreate','formCreate',0,'view/systemTools/formCreate/index.vue',3,'',1,0,'表单生成器','magic-stick',0),(18,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,15,'system','system',0,'view/systemTools/system/system.vue',4,'',0,0,'系统配置','operation',0),(19,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,15,'autoCodeAdmin','autoCodeAdmin',0,'view/systemTools/autoCodeAdmin/index.vue',2,'',0,0,'自动化代码管理','magic-stick',0),(20,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,15,'autoCodeEdit/:id','autoCodeEdit',1,'view/systemTools/autoCode/index.vue',0,'',0,0,'自动化代码-${id}','magic-stick',0),(21,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,15,'autoPkg','autoPkg',0,'view/systemTools/autoPkg/autoPkg.vue',0,'',0,0,'模板配置','folder',0),(22,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,0,'https://www.gin-vue-admin.com','https://www.gin-vue-admin.com',0,'/',0,'',0,0,'官方网站','customer-gva',0),(23,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,0,'state','state',0,'view/system/state.vue',8,'',0,0,'服务器状态','cloudy',0),(24,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,0,'plugin','plugin',0,'view/routerHolder.vue',6,'',0,0,'插件系统','cherry',0),(25,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,24,'https://plugin.gin-vue-admin.com/','https://plugin.gin-vue-admin.com/',0,'https://plugin.gin-vue-admin.com/',0,'',0,0,'插件市场','shop',0),(26,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,24,'installPlugin','installPlugin',0,'view/systemTools/installPlugin/index.vue',1,'',0,0,'插件安装','box',0),(27,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,24,'pubPlug','pubPlug',0,'view/systemTools/pubPlug/pubPlug.vue',3,'',0,0,'打包插件','files',0),(28,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,24,'plugin-email','plugin-email',0,'plugin/email/view/index.vue',4,'',0,0,'邮件插件','message',0),(29,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,15,'exportTemplate','exportTemplate',0,'view/systemTools/exportTemplate/exportTemplate.vue',5,'',0,0,'表格模板','reading',0),(30,'2024-08-22 14:01:31.031','2024-08-22 14:01:31.031',NULL,0,24,'anInfo','anInfo',0,'plugin/announcement/view/info.vue',5,'',0,0,'公告管理[示例]','scaleToOriginal',0);
/*!40000 ALTER TABLE `sys_base_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_data_authority_id`
--

DROP TABLE IF EXISTS `sys_data_authority_id`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_data_authority_id`
--

LOCK TABLES `sys_data_authority_id` WRITE;
/*!40000 ALTER TABLE `sys_data_authority_id` DISABLE KEYS */;
INSERT INTO `sys_data_authority_id` VALUES (888,888),(888,8881),(888,9528),(9528,8881),(9528,9528);
/*!40000 ALTER TABLE `sys_data_authority_id` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dictionaries`
--

DROP TABLE IF EXISTS `sys_dictionaries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dictionaries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dictionaries`
--

LOCK TABLES `sys_dictionaries` WRITE;
/*!40000 ALTER TABLE `sys_dictionaries` DISABLE KEYS */;
INSERT INTO `sys_dictionaries` VALUES (1,'2024-08-22 14:01:27.813','2024-08-22 14:01:28.294',NULL,'性别','gender',1,'性别字典'),(2,'2024-08-22 14:01:27.813','2024-08-22 14:01:28.805',NULL,'数据库int类型','int',1,'int类型对应的数据库类型'),(3,'2024-08-22 14:01:27.813','2024-08-22 14:01:29.201',NULL,'数据库时间日期类型','time.Time',1,'数据库时间日期类型'),(4,'2024-08-22 14:01:27.813','2024-08-22 14:01:29.640',NULL,'数据库浮点型','float64',1,'数据库浮点型'),(5,'2024-08-22 14:01:27.813','2024-08-22 14:01:30.039',NULL,'数据库字符串','string',1,'数据库字符串'),(6,'2024-08-22 14:01:27.813','2024-08-22 14:01:30.447',NULL,'数据库bool类型','bool',1,'数据库bool类型');
/*!40000 ALTER TABLE `sys_dictionaries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dictionary_details`
--

DROP TABLE IF EXISTS `sys_dictionary_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示值',
  `value` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dictionary_details`
--

LOCK TABLES `sys_dictionary_details` WRITE;
/*!40000 ALTER TABLE `sys_dictionary_details` DISABLE KEYS */;
INSERT INTO `sys_dictionary_details` VALUES (1,'2024-08-22 14:01:28.362','2024-08-22 14:01:28.362',NULL,'男','1','',1,1,1),(2,'2024-08-22 14:01:28.362','2024-08-22 14:01:28.362',NULL,'女','2','',1,2,1),(3,'2024-08-22 14:01:28.872','2024-08-22 14:01:28.872',NULL,'smallint','1','mysql',1,1,2),(4,'2024-08-22 14:01:28.872','2024-08-22 14:01:28.872',NULL,'mediumint','2','mysql',1,2,2),(5,'2024-08-22 14:01:28.872','2024-08-22 14:01:28.872',NULL,'int','3','mysql',1,3,2),(6,'2024-08-22 14:01:28.872','2024-08-22 14:01:28.872',NULL,'bigint','4','mysql',1,4,2),(7,'2024-08-22 14:01:28.872','2024-08-22 14:01:28.872',NULL,'int2','5','pgsql',1,5,2),(8,'2024-08-22 14:01:28.872','2024-08-22 14:01:28.872',NULL,'int4','6','pgsql',1,6,2),(9,'2024-08-22 14:01:28.872','2024-08-22 14:01:28.872',NULL,'int6','7','pgsql',1,7,2),(10,'2024-08-22 14:01:28.872','2024-08-22 14:01:28.872',NULL,'int8','8','pgsql',1,8,2),(11,'2024-08-22 14:01:29.270','2024-08-22 14:01:29.270',NULL,'date','','',1,0,3),(12,'2024-08-22 14:01:29.270','2024-08-22 14:01:29.270',NULL,'time','1','mysql',1,1,3),(13,'2024-08-22 14:01:29.270','2024-08-22 14:01:29.270',NULL,'year','2','mysql',1,2,3),(14,'2024-08-22 14:01:29.270','2024-08-22 14:01:29.270',NULL,'datetime','3','mysql',1,3,3),(15,'2024-08-22 14:01:29.270','2024-08-22 14:01:29.270',NULL,'timestamp','5','mysql',1,5,3),(16,'2024-08-22 14:01:29.270','2024-08-22 14:01:29.270',NULL,'timestamptz','6','pgsql',1,5,3),(17,'2024-08-22 14:01:29.713','2024-08-22 14:01:29.713',NULL,'float','','',1,0,4),(18,'2024-08-22 14:01:29.713','2024-08-22 14:01:29.713',NULL,'double','1','mysql',1,1,4),(19,'2024-08-22 14:01:29.713','2024-08-22 14:01:29.713',NULL,'decimal','2','mysql',1,2,4),(20,'2024-08-22 14:01:29.713','2024-08-22 14:01:29.713',NULL,'numeric','3','pgsql',1,3,4),(21,'2024-08-22 14:01:29.713','2024-08-22 14:01:29.713',NULL,'smallserial','4','pgsql',1,4,4),(22,'2024-08-22 14:01:30.124','2024-08-22 14:01:30.124',NULL,'char','','',1,0,5),(23,'2024-08-22 14:01:30.124','2024-08-22 14:01:30.124',NULL,'varchar','1','mysql',1,1,5),(24,'2024-08-22 14:01:30.124','2024-08-22 14:01:30.124',NULL,'tinyblob','2','mysql',1,2,5),(25,'2024-08-22 14:01:30.124','2024-08-22 14:01:30.124',NULL,'tinytext','3','mysql',1,3,5),(26,'2024-08-22 14:01:30.124','2024-08-22 14:01:30.124',NULL,'text','4','mysql',1,4,5),(27,'2024-08-22 14:01:30.124','2024-08-22 14:01:30.124',NULL,'blob','5','mysql',1,5,5),(28,'2024-08-22 14:01:30.124','2024-08-22 14:01:30.124',NULL,'mediumblob','6','mysql',1,6,5),(29,'2024-08-22 14:01:30.124','2024-08-22 14:01:30.124',NULL,'mediumtext','7','mysql',1,7,5),(30,'2024-08-22 14:01:30.124','2024-08-22 14:01:30.124',NULL,'longblob','8','mysql',1,8,5),(31,'2024-08-22 14:01:30.124','2024-08-22 14:01:30.124',NULL,'longtext','9','mysql',1,9,5),(32,'2024-08-22 14:01:30.535','2024-08-22 14:01:30.535',NULL,'tinyint','1','mysql',1,0,6),(33,'2024-08-22 14:01:30.535','2024-08-22 14:01:30.535',NULL,'bool','2','pgsql',1,0,6);
/*!40000 ALTER TABLE `sys_dictionary_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_export_template_condition`
--

DROP TABLE IF EXISTS `sys_export_template_condition`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_template_condition` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板标识',
  `from` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '条件取的key',
  `column` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '作为查询条件的字段',
  `operator` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作符',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_template_condition_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_export_template_condition`
--

LOCK TABLES `sys_export_template_condition` WRITE;
/*!40000 ALTER TABLE `sys_export_template_condition` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_export_template_condition` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_export_template_join`
--

DROP TABLE IF EXISTS `sys_export_template_join`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_template_join` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板标识',
  `joins` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联',
  `table` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联表',
  `on` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联条件',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_template_join_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_export_template_join`
--

LOCK TABLES `sys_export_template_join` WRITE;
/*!40000 ALTER TABLE `sys_export_template_join` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_export_template_join` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_export_templates`
--

DROP TABLE IF EXISTS `sys_export_templates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_templates` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `db_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '数据库名称',
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板名称',
  `table_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '表名称',
  `template_id` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板标识',
  `template_info` text COLLATE utf8mb4_general_ci,
  `limit` bigint DEFAULT NULL COMMENT '导出限制',
  `order` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_templates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_export_templates`
--

LOCK TABLES `sys_export_templates` WRITE;
/*!40000 ALTER TABLE `sys_export_templates` DISABLE KEYS */;
INSERT INTO `sys_export_templates` VALUES (1,'2024-08-22 14:01:32.812','2024-08-22 14:01:32.812',NULL,'','api','sys_apis','api','{\n\"path\":\"路径\",\n\"method\":\"方法（大写）\",\n\"description\":\"方法介绍\",\n\"api_group\":\"方法分组\"\n}',NULL,'');
/*!40000 ALTER TABLE `sys_export_templates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_ignore_apis`
--

DROP TABLE IF EXISTS `sys_ignore_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_ignore_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `method` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_ignore_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_ignore_apis`
--

LOCK TABLES `sys_ignore_apis` WRITE;
/*!40000 ALTER TABLE `sys_ignore_apis` DISABLE KEYS */;
INSERT INTO `sys_ignore_apis` VALUES (1,'2024-08-22 14:01:25.892','2024-08-22 14:01:25.892',NULL,'/swagger/*any','GET'),(2,'2024-08-22 14:01:25.892','2024-08-22 14:01:25.892',NULL,'/api/freshCasbin','GET'),(3,'2024-08-22 14:01:25.892','2024-08-22 14:01:25.892',NULL,'/uploads/file/*filepath','GET'),(4,'2024-08-22 14:01:25.892','2024-08-22 14:01:25.892',NULL,'/health','GET'),(5,'2024-08-22 14:01:25.892','2024-08-22 14:01:25.892',NULL,'/uploads/file/*filepath','HEAD'),(6,'2024-08-22 14:01:25.892','2024-08-22 14:01:25.892',NULL,'/autoCode/llmAuto','POST'),(7,'2024-08-22 14:01:25.892','2024-08-22 14:01:25.892',NULL,'/system/reloadSystem','POST'),(8,'2024-08-22 14:01:25.892','2024-08-22 14:01:25.892',NULL,'/base/login','POST'),(9,'2024-08-22 14:01:25.892','2024-08-22 14:01:25.892',NULL,'/base/captcha','POST'),(10,'2024-08-22 14:01:25.892','2024-08-22 14:01:25.892',NULL,'/init/initdb','POST'),(11,'2024-08-22 14:01:25.892','2024-08-22 14:01:25.892',NULL,'/init/checkdb','POST');
/*!40000 ALTER TABLE `sys_ignore_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_operation_records`
--

DROP TABLE IF EXISTS `sys_operation_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_operation_records` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求路径',
  `status` bigint DEFAULT NULL COMMENT '请求状态',
  `latency` bigint DEFAULT NULL COMMENT '延迟',
  `agent` text COLLATE utf8mb4_general_ci COMMENT '代理',
  `error_message` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '错误信息',
  `body` text COLLATE utf8mb4_general_ci COMMENT '请求Body',
  `resp` text COLLATE utf8mb4_general_ci COMMENT '响应Body',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_operation_records`
--

LOCK TABLES `sys_operation_records` WRITE;
/*!40000 ALTER TABLE `sys_operation_records` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_operation_records` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_authority`
--

DROP TABLE IF EXISTS `sys_user_authority`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_authority`
--

LOCK TABLES `sys_user_authority` WRITE;
/*!40000 ALTER TABLE `sys_user_authority` DISABLE KEYS */;
INSERT INTO `sys_user_authority` VALUES (1,888),(1,8881),(1,9528),(2,888);
/*!40000 ALTER TABLE `sys_user_authority` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_users`
--

DROP TABLE IF EXISTS `sys_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) COLLATE utf8mb4_general_ci DEFAULT '#fff' COMMENT '基础颜色',
  `authority_id` bigint unsigned DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`),
  KEY `idx_sys_users_uuid` (`uuid`),
  KEY `idx_sys_users_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_users`
--

LOCK TABLES `sys_users` WRITE;
/*!40000 ALTER TABLE `sys_users` DISABLE KEYS */;
INSERT INTO `sys_users` VALUES (1,'2024-08-22 14:01:31.471','2024-08-22 14:01:31.669',NULL,'fa0ff9a8-8f83-40a3-b900-bdf38f4dc9f1','admin','$2a$10$APPB.scxRsgJsrFGdUcBjO9TeUjrEPgws5lkpE9vg.dMDmkocP7..','Mr.奇淼','dark','https://qmplusimg.henrongyi.top/gva_header.jpg','#fff',888,'17611111111','333333333@qq.com',1),(2,'2024-08-22 14:01:31.471','2024-08-22 14:01:32.171',NULL,'b608dde7-b8bf-4088-a9d1-6aca5d9fe5b1','a303176530','$2a$10$Bhfsq0bG02z.YmHre/B0FujBi/JDm/reryeTisDMIM.juZkehN5fK','用户1','dark','https:///qmplusimg.henrongyi.top/1572075907logo.png','#fff',9528,'17611111111','333333333@qq.com',1);
/*!40000 ALTER TABLE `sys_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'aaaaa'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-08-22 14:04:38
